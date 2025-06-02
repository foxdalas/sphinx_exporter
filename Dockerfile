FROM       alpine:3.22
MAINTAINER Maxim Pogozhiy <foxdalas@gmail.com>

ARG TARGETARCH

RUN apk add --no-cache libc6-compat
COPY sphinx-exporter /bin/sphinx_exporter

ENTRYPOINT ["/bin/sphinx_exporter"]
EXPOSE     9247
