include "consul-interface.iol"
include "console.iol"

execution { single }

outputPort Consul {
    Location: "socket://localhost:80/"
    Interfaces: ConsulGetter
    Protocol: http { .method = "get" }
}

// run tests
main
{
    entry.key = "test.data";
    get@Consul(entry)( out );
    println@Console( out.val )()
}