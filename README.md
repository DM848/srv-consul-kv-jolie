# srv-consul-kv-jolie
Service that allows jolie interfaces to retrieve values from the Consul key/value database


How to use this in your Jolie service:
```jolie
include "consul.iol" // I suggest using git submodules (version control) => 
                     // include "vendor/srv-consul-kv-jolie/consul.iol"
include "console.iol"

main
{
    entry.key = "GITHUB_ACCESS_TOKEN";
    get@Consul(entry)( out );
    println@Console( out.val )() // token is printed to terminal
}
```


NB! This only support the retrieval of values.
