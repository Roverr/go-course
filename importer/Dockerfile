## Build stage
FROM golang:1.10-alpine AS build-env
ADD . /go/src/github.com/Roverr/go-course/importer
WORKDIR /go/src/github.com/Roverr/go-course/importer
RUN apk add --update --no-cache git
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN go build -o server

## Creating potential production image
FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /app
COPY --from=build-env /go/src/github.com/Roverr/go-course/importer/server /app/
ENTRYPOINT [ "/app/server" ]