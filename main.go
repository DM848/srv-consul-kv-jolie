package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type ConsulEntry struct {
	Key string `json:"Key"`
	Val string `json:"Value"`
}

type ConsulJolieEntry struct {
	Key string `json:"key"`
	Val string `json:"val"`
	Err string `json:"err"`
}

func (jr *ConsulJolieEntry) data() ([]byte, error) {
	return json.Marshal(jr)
}

func main() {
	port := os.Getenv("WEB_SERVER_PORT")
	if port == "" {
		panic("missing environment variable WEB_SERVER_PORT")
	}

	client := &http.Client{
		//CheckRedirect: redirectPolicyFunc,
	}

	router := httprouter.New()
	router.GET("/health", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Fprintf(w, `{"status":"ok"}`)
	})

	router.GET("/get", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		jr := &ConsulJolieEntry{}
		defer func(w http.ResponseWriter, jr *ConsulJolieEntry) {
			data, err := jr.data()
			w.Header().Set("Content-Type", "application/json")
			if err != nil {
				_, _ = w.Write([]byte(`{"err":"unable to unmarshal kv entry"}`))
				return
			}

			_, _ = w.Write(data)
		}(w, jr)

		key := r.URL.Query().Get("key")
		if key == "" {
			w.WriteHeader(404)
			return
		}
		jr.Key = key
		if key == "test.data" {
			jr.Val = "success"
			return
		}

		url := "http://consul-node:8500/v1/kv/" + key + "?raw=true"
		resp, err := client.Get(url)
		if err != nil {
			w.WriteHeader(400)
			jr.Err = "key was not found"
			return
		}
		defer resp.Body.Close()

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			w.WriteHeader(500)
			jr.Err = "Error parsing JSON"
			return
		}

		var entries []*ConsulEntry
		err = json.Unmarshal(data, &entries)
		if err != nil {
			w.WriteHeader(500)
			jr.Err = "Error unmarshaling JSON"
			return
		}

		if len(entries) == 0 || entries[0].Val == "" {
			w.WriteHeader(404)
			jr.Err = "value in entry was empty. key = " + key
			return
		}

		jr.Val = entries[0].Val
	})

	log.Fatal(http.ListenAndServe(":" + port, router))
}
