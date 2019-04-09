# First stage: build the executable.
FROM golang:1.11 as builder
WORKDIR /go/src/github.com/rekzi/http-echo/
COPY . .
# CGO_ENABLED=0 to build a statically-linked executable
# GOFLAGS=-mod=vendor to force `go build` to look into the `/vendor` folder.
ENV CGO_ENABLED=0 GOFLAGS=-mod=vendor
RUN go build -installsuffix 'static' -o http-echo main.go

# Final stage: the running container.
FROM alpine:3.9 AS final
RUN apk add --update --no-cache ca-certificates
WORKDIR /bin/
# Import the compiled executable from the first stage.
COPY --from=builder /go/src/github.com/rekzi/http-echo/http-echo .
EXPOSE 9000
ENTRYPOINT [ "/bin/http-echo" ]
