FROM alpine:3.10

RUN apk add --no-cache ca-certificates

ADD ./cleanup-operator /cleanup-operator

ENTRYPOINT ["/cleanup-operator"]
