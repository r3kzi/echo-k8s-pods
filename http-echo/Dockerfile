FROM golang:1.11 as builder
WORKDIR /go/src/github.com/rekzi/http-echo/
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o http-echo main.go

FROM alpine:3.9
RUN apk add --update --no-cache ca-certificates
WORKDIR /bin/
COPY --from=builder /go/src/github.com/rekzi/http-echo/http-echo .
EXPOSE 9000
ENTRYPOINT [ "/bin/http-echo" ]
