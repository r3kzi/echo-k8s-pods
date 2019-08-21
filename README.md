# Echo Pods for Kubernetes

[![Build Status](https://travis-ci.org/rekz1/echo-k8s-pods.svg?branch=master)](https://travis-ci.org/rekz1/echo-k8s-pods)

## tcp

```bash
$ docker run -p 9000:9000 rekzi/tcp-echo -version v1
2019/04/09 17:05:03 Server is starting...
2019/04/09 17:05:03 listening on port [::]:9000 with prefix: v1
```

```bash
$ echo "" | nc localhost 9000                                                                                                         Di 09 Apr 2019 19:03:32 CEST
v1
```

## http

```bash
$ docker run -p 8080:8080 rekzi/http-echo -version v1
2019/04/07 15:04:45 Server is starting...
2019/04/07 15:04:45 Server is ready to handle requests at :8080 
```

```bash
$ curl localhost:8080
v1

$ curl localhost:8080/showRequest -X POST -H "TestHeader: test1111111111111" -H "Content-Type: application/json" --data '{"value": "my_password"}' 
POST /showRequest HTTP/1.1
Host: localhost:8080
Accept: */*
Content-Length: 24
Content-Type: application/json
Testheader: test1111111111111
User-Agent: curl/7.61.0
{"value": "my_password"}

```

### prometheus endpoint

```bash
$ curl -s localhost:8080/metrics | grep http_echo                                                                                     Di 09 Apr 2019 19:07:01 CEST
# HELP http_echo_requests_processed_total The total number of processed events
# TYPE http_echo_requests_processed_total counter
http_echo_requests_processed_total 42
```