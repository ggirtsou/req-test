# requests-test

```bash
go run server/main.go

# in a new terminal
go run client/main.go
```

`io.Copy(ioutil.Discard, resp.Body)` is needed to release the sockets.
This is what it looks like when you don't discard the body:

```bash
$ netstat
...
tcp        0      0 localhost:47176         localhost:http-alt      TIME_WAIT
tcp        0      0 localhost:42310         localhost:http-alt      TIME_WAIT
tcp        0      0 localhost:34137         localhost:http-alt      TIME_WAIT
tcp        0      0 localhost:59068         localhost:http-alt      TIME_WAIT
tcp        0      0 localhost:46366         localhost:http-alt      TIME_WAIT
...
```
