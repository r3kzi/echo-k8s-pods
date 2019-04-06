# Example Pods for Kubernetes

## TCP Server

```bash
$ docker run -p 9000:9000 rekzi/tcp-echo -prefix one
listening on port [::]:9000 with prefix: one
```

```bash
$ docker run -p 9001:9000 rekzi/tcp-echo -prefix two
listening on port [::]:9000 with prefix: two
```

```bash
$ date | nc localhost 9000
one Sa 6. Apr 09:59:27 CEST 2019

$ date | nc localhost 9001
two Sa 6. Apr 10:00:42 CEST 2019
```