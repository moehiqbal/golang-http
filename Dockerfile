# Build
FROM golang:1.17-buster AS build

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GOARCH=wasm
ENV GOOS=js

WORKDIR /app
COPY go,mod .

RUN go mod tidy

COPY *.go .

RUN go build -o ./docker-golang-http

# Deploy
FROM gcr.io/distroless/base-debian10
WORKDIR /app
COPY --from=build /docker-golang-http /docker-golang-http
EXPOSE 3000
USER nonroot:nonroot
ENTRYPOINT ["/docker-golang-http"]