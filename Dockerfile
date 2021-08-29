FROM golang:1.17-alpine as builder

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download

COPY ./cmd/server.go  ./
COPY ./graph ./graph

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=arm
ARG GOARM=7
RUN go build \
    -o /go/bin/server \
    -ldflags '-s -w'


FROM arm32v7/debian:bullseye-slim as runner

COPY --from=builder /go/bin/server /app/server

CMD ["/app/server"]