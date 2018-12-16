include "consul-interface.iol"

outputPort Consul {
    Location: "socket://consul-kv-jolie:8888/"
    Interfaces: ConsulGetter
    Protocol: http { .method = "get" }
}