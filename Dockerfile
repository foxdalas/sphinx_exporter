FROM       alpine:3.15
MAINTAINER Maxim Pogozhiy <foxdalas@gmail.com>

COPY sphinx_exporter /bin/sphinx_exporter

ENTRYPOINT ["/bin/sphinx_exporter"]
EXPOSE     9247
