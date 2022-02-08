# Dockerfile for Sphinx SE
# https://hub.docker.com/_/alpine/
FROM alpine:3.14

# https://sphinxsearch.com/blog/
ENV SPHINX_VERSION 3.4.1-efbcc65

# install dependencies
RUN apk add --no-cache mariadb-connector-c-dev \
	postgresql-dev \
	wget

# set up and expose directories
RUN mkdir -pv /opt/sphinx/log /opt/sphinx/index
VOLUME /opt/sphinx/index

# http://sphinxsearch.com/downloads/sphinx-3.3.1-b72d67b-linux-amd64-musl.tar.gz
RUN wget http://sphinxsearch.com/files/sphinx-${SPHINX_VERSION}-linux-amd64-musl.tar.gz -O /tmp/sphinxsearch.tar.gz \
	&& cd /opt/sphinx && tar -xf /tmp/sphinxsearch.tar.gz \
	&& rm /tmp/sphinxsearch.tar.gz

# point to sphinx binaries
ENV PATH "${PATH}:/opt/sphinx/sphinx-3.4.1/bin"
RUN indexer -v

# redirect logs to stdout
RUN ln -sv /dev/stdout /opt/sphinx/log/query.log \
    	&& ln -sv /dev/stdout /opt/sphinx/log/searchd.log

# expose TCP port
EXPOSE 3306

RUN mkdir -p /opt/sphinx/conf
RUN mkdir -p /opt/sphinx/tests

COPY docs.xml /opt/sphinx/tests/
COPY sphinx.conf /opt/sphinx/conf/

# allow custom config file to be passed
ARG SPHINX_CONFIG_FILE=/opt/sphinx/conf/sphinx.conf
ENV SPHINX_CONFIG_FILE ${SPHINX_CONFIG_FILE}

# prepare a start script
RUN echo "exec searchd --nodetach --config \${SPHINX_CONFIG_FILE}" > /opt/sphinx/start.sh

RUN /opt/sphinx/sphinx-3.4.1/bin/indexer --all --config "/opt/sphinx/conf/sphinx.conf"

CMD sh /opt/sphinx/start.sh
