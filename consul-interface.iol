
type ConsulKey: void {
    .key: string
}

type ConsulResponse: void {
    .key: string
    .val: string
    .err: string
}

interface ConsulGetter {
    RequestResponse:
        get( ConsulKey )( ConsulResponse )
}