FROM golang:1.11 as builder
WORKDIR /go/src/github.com/rekzi/tcp-echo/
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o tcp-echo main.go

FROM alpine:3.9
RUN apk add --update --no-cache ca-certificates
WORKDIR /bin/
COPY --from=builder /go/src/github.com/rekzi/tcp-echo/tcp-echo .
EXPOSE 9000
ENTRYPOINT [ "/bin/tcp-echo" ]
