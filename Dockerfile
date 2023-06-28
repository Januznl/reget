FROM golang:1.20-alpine as build

ADD . /app

RUN apk add --update --no-cache ca-certificates


FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app/reget /reget

