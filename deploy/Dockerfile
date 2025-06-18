FROM golang:1.24 AS builder

WORKDIR /app

ENV CGO_ENABLED=1

COPY go.mod go.sum ./
RUN go mod download -x && go mod verify

COPY cmd cmd
COPY internal internal
COPY docs docs
COPY pkg pkg
RUN mkdir ./bin && go build -v -o ./bin ./...

FROM ubuntu:noble AS final

ARG RELEASE
ENV RELEASE=${RELEASE}
WORKDIR /app

COPY --from=builder /app/bin/ .
