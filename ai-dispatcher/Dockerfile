# Build the manager binary
FROM golang:1.13-stretch as builder

# Copy in the go src
WORKDIR /go/src/github.com/containers-ai/alameda
ADD . .

# Build
RUN ["/bin/bash", "-c", "CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags \"-X main.VERSION=`git rev-parse --abbrev-ref HEAD`-`git rev-parse --short HEAD``git diff --quiet || echo '-dirty'` -X 'main.BUILD_TIME=`date`' -X 'main.GO_VERSION=`go version`'\" -a -o ./ai-dispatcher/ai-dispatcher github.com/containers-ai/alameda/ai-dispatcher/cmd"]

# Copy the controller-manager into a thin image
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/github.com/containers-ai/alameda/ai-dispatcher/etc/ai-dispatcher.toml /etc/alameda/ai-dispatcher/ai-dispatcher.toml
COPY --from=builder /go/src/github.com/containers-ai/alameda/ai-dispatcher/ai-dispatcher .
EXPOSE 9091
ENTRYPOINT ["./ai-dispatcher"]
