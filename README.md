# tinygo-tls

## Setup

listen htttps://localhost:8443

```shell
> go run ./server
```

build on go

```shell
> go build -o tinygo-tls .
```

build on tinygo

```shell
> tinygo build -o tinygo-tls .
```

## Test TLS communication

```shell
> $ go run ./socat ./tinygo-tls
2020/10/19 10:46:12 exec: [./tinygo-tls]
HTTP/1.0 200 OK
Date: Mon, 19 Oct 2020 01:46:12 GMT
Content-Length: 6
Content-Type: text/plain; charset=utf-8

hello!
```
