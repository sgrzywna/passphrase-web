# minimize stage
FROM node:8.12.0-alpine AS minimize
LABEL stage=intermediate

RUN npm install html-minifier -g
RUN npm install uglify-js -g

COPY . /go/src/github.com/sgrzywna/passphrase-web

RUN uglifyjs \
    --compress \
    --mangle \
    --toplevel \
    --output /tmp/index.js \
    /go/src/github.com/sgrzywna/passphrase-web/web/static/index.js

RUN html-minifier \
    --collapse-whitespace \
    --remove-comments \
    --remove-optional-tags \
    --remove-redundant-attributes \
    --remove-script-type-attributes \
    --remove-tag-whitespace \
    --use-short-doctype \
    --minify-css true \
    --minify-js true \
    --output /tmp/index.tmpl \
    /go/src/github.com/sgrzywna/passphrase-web/web/template/index.tmpl

# build stage
FROM golang:1.11.1-alpine3.8 AS build
LABEL stage=intermediate

RUN apk add --no-cache git build-base

RUN go get github.com/golang/dep/cmd/dep
RUN go get github.com/GeertJohan/go.rice
RUN go get github.com/GeertJohan/go.rice/rice

WORKDIR /go/src/github.com/sgrzywna/passphrase-web

COPY . /go/src/github.com/sgrzywna/passphrase-web

COPY --from=minimize /tmp/index.js /go/src/github.com/sgrzywna/passphrase-web/web/static/index.js
COPY --from=minimize /tmp/index.tmpl /go/src/github.com/sgrzywna/passphrase-web/web/template/index.tmpl

RUN make clean \
    && make dep \
    && make rice \
    && make build

# release stage
FROM alpine:3.8

WORKDIR /app

COPY --from=build /go/src/github.com/sgrzywna/passphrase-web/passphrase-web /app/
COPY --from=build /go/src/github.com/sgrzywna/passphrase-web/dicts /app/dicts

ENTRYPOINT ["/app/passphrase-web", "-dicts", "/app/dicts"]
