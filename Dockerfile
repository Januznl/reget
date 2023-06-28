FROM golang:1.20-alpine as build

WORKDIR /app

ADD . /app

RUN apk add --update --no-cache ca-certificates
RUN CGO_ENABLED=0 go build .


FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app/reget /reget

ENTRYPOINT [ "/reget" ]