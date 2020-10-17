# tinygo-tls

listen htttps://localhost:8443

```shell
> cd server
> go run .
```

Test TLS communication

```
> cd ..
> go run .
HTTP/1.0 200 OK
Date: Sat, 17 Oct 2020 13:47:23 GMT
Content-Length: 6
Content-Type: text/plain; charset=utf-8

hello!
```
