# build UI
FROM node:10.15.0-alpine AS dist
LABEL stage=intermediate

COPY . /app

WORKDIR /app/web

RUN npm install

RUN npm run build

# build service
FROM golang:1.11.1-alpine3.8 AS build
LABEL stage=intermediate

RUN apk add --no-cache git build-base

RUN go get github.com/golang/dep/cmd/dep
RUN go get github.com/GeertJohan/go.rice
RUN go get github.com/GeertJohan/go.rice/rice

COPY . /go/src/github.com/sgrzywna/passphrase-web

COPY --from=dist /app/web/dist /go/src/github.com/sgrzywna/passphrase-web/web/dist

WORKDIR /go/src/github.com/sgrzywna/passphrase-web

RUN make clean \
    && make dep \
    && make rice \
    && make build

# build release
FROM alpine:3.8

WORKDIR /app

COPY --from=build /go/src/github.com/sgrzywna/passphrase-web/passphrase-web /app/
COPY --from=build /go/src/github.com/sgrzywna/passphrase-web/dicts /app/dicts

ENTRYPOINT ["/app/passphrase-web", "-dicts", "/app/dicts"]
