## Golang code organization

cat ~/export-http-proxy.sh
```sh
#!/bin/bash
export http_proxy=socks5://127.0.0.1:1080
export https_proxy=$http_proxy
export | grep http
```

```sh
go install github.com/lixit/hello_go
export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))
hello_go
```


```sh
go run helloworld.go

go build helloworld.go
./helloworld

go get gopl.io/ch1/helloworld

go fmt

```
