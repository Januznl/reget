FROM golang:1.20-alpine as build

RUN apk add --update --no-cache ca-certificates

FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY reget /reget

ENTRYPOINT [ "/reget" ]