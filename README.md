# HTTP go client library
HTTP client in Go
## Installation
```bash
# Go Modules
module github.com/supalik/fthreeclient
require github.com/google/uuid v1.3.0
```
## Usage
import the corresponding HTTP package  to use the library, refer the example.go
```go
import "github.com/supalik/fthreeclient"
```
## Execute TCs
```
$ docker-compose up -d
$ go test ./goformclient/ -v -cover

