# transcribe
Simple markdown parser as Service

Running
---

 - Go
 - Godep (opt)

```
go get
go run transcribe.go
```
Using Godep
```
godep get
go run transcribe.go
```

Endpoint
---

`PUT /markdown`: Convert the given body into HTML


Outputs are sanitized.

The server listens on port `8080` by default. If you want it to listen on another port you can set on `PORT` environment variable.

Consul
---

We can connect to Consul using the provided environment variables from API
 - `CONSUL_HTTP_ADDR`
 - `CONSUL_HTTP_TOKEN`
 - `CONSUL_HTTP_AUTH`
 - `CONSUL_HTTP_SSL`
 - `CONSUL_HTTP_SSL_VERIFY`

 Any changes need from default can be set on environment variables, so the client can pick easily without code modification.
 In case of failure the application will still run. So we don't get strings attached on Consul.
