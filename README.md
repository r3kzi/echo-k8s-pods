# Echo Pods for Kubernetes

## tcp

```bash
$ docker run -p 9000:9000 rekzi/tcp-echo -prefix one
listening on port [::]:9000 with prefix: one
```
```bash
$ date | nc localhost 9000
one Sa 6. Apr 09:59:27 CEST 2019
```

## http

```bash
docker run -p 8080:8080 rekzi/http-echo -version v1                                                                      7.3s  So 07 Apr 2019 17:01:18 CEST
2019/04/07 15:04:45 Server is starting...
2019/04/07 15:04:45 Server is ready to handle requests at :8080 
```

```bash
$ curl localhost:8080
v1
```