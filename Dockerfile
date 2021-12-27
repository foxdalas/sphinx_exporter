FROM       alpine:3.15
MAINTAINER Maxim Pogozhiy <foxdalas@gmail.com>

ARG TARGETARCH

COPY sphinx-exporter /bin/sphinx_exporter

ENTRYPOINT ["/bin/sphinx_exporter"]
EXPOSE     9247
