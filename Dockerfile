FROM        quay.io/prometheus/busybox:latest
MAINTAINER  The Prometheus Authors <prometheus-developers@googlegroups.com>

COPY sphinx_exporter /bin/sphinx_exporter

ENTRYPOINT ["/bin/sphinx_exporter"]
EXPOSE     9247
