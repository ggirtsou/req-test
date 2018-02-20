# requests-test

```bash
go run server/main.go

# in a new terminal
go run client/main.go
```

with `io.Copy(ioutil.Discard, resp.Body)` in `makeReq`:

```text
made 922 requests
made 923 requests
panic: Get http://localhost:8080: dial tcp: lookup localhost: too many open files

goroutine 4548 [running]:
```

without discarding the body:

```text
made 28209 requests
made 28211 requests
panic: Get http://localhost:8080: dial tcp 127.0.0.1:8080: connect: cannot assign requested address
```
