FROM golang:1.18.9-alpine3.17 as builder

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download

ADD . /main
WORKDIR /main

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build \
    -o /go/bin/main \
    -ldflags '-s -w'


FROM scratch as runner

COPY --from=builder /go/bin/main /app/main

ENTRYPOINT ["/app/main"]