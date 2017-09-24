# Sphinx Exporter for Prometheus [![CircleCI](https://circleci.com/gh/foxdalas/sphinx_exporter.svg?style=svg)](https://circleci.com/gh/foxdalas/sphinx_exporter)

[![Docker Repository on Quay](https://quay.io/repository/foxdalas/sphinx-exporter/status "Docker Repository on Quay")](https://quay.io/repository/foxdalas/sphinx-exporter)
[![Docker Pulls](https://img.shields.io/docker/pulls/foxdalas/sphinx-exporter.svg?maxAge=604800)](https://hub.docker.com/r/foxdalas/sphinx-exporter/)

A [sphinx](http://sphinxsearch.com) exporter for prometheus.


## Building

The sphinx exporter exports metrics from a sphinx server for
consumption by prometheus. The server is specified as `--sphinx.address` and `--sphinx.port` flag
to the program (default is `localhost:9306`).

By default the sphinx\_exporter serves on port `0.0.0.0:9247` at `/metrics`

```
make
./sphinx_exporter
```

Alternatively a Dockerfile is supplied

```
docker build -t sphinx_exporter .
docker run sphinx_exporter
```

## Collectors

The exporter collects a number of statistics from the server:

```

# HELP sphinx_agent_connect Number of agent connect since server start.
# TYPE sphinx_agent_connect counter
sphinx_agent_connect 0
# HELP sphinx_agent_retry Number of agent retry since server start.
# TYPE sphinx_agent_retry counter
sphinx_agent_retry 0
# HELP sphinx_avg_dist_local Number of avg_dist_local since server start.
# TYPE sphinx_avg_dist_local counter
sphinx_avg_dist_local 0
# HELP sphinx_avg_dist_wait Number of avg_dist_wait since server start.
# TYPE sphinx_avg_dist_wait counter
sphinx_avg_dist_wait 0
# HELP sphinx_avg_dist_wall Number of avg_dist_wall since server start.
# TYPE sphinx_avg_dist_wall counter
sphinx_avg_dist_wall 0
# HELP sphinx_avg_query_cpu Number of avg_query_cpu since server start.
# TYPE sphinx_avg_query_cpu gauge
sphinx_avg_query_cpu 0
# HELP sphinx_avg_query_readkb Number of avg_query_readkb since server start.
# TYPE sphinx_avg_query_readkb gauge
sphinx_avg_query_readkb 0
# HELP sphinx_avg_query_reads Number of avg_query_reads since server start.
# TYPE sphinx_avg_query_reads gauge
sphinx_avg_query_reads 0
# HELP sphinx_avg_query_readtime Number of avg_query_readtime since server start.
# TYPE sphinx_avg_query_readtime gauge
sphinx_avg_query_readtime 0
# HELP sphinx_avg_query_wall Number of avg_query_wall since server start.
# TYPE sphinx_avg_query_wall counter
sphinx_avg_query_wall 0
# HELP sphinx_command_delete Number of delete command since server start.
# TYPE sphinx_command_delete counter
sphinx_command_delete 0
# HELP sphinx_command_flushattrs Number of flushattrs command since server start.
# TYPE sphinx_command_flushattrs counter
sphinx_command_flushattrs 0
# HELP sphinx_command_keywords Number of keywords command since server start.
# TYPE sphinx_command_keywords counter
sphinx_command_keywords 672
# HELP sphinx_command_persist Number of persist command since server start.
# TYPE sphinx_command_persist counter
sphinx_command_persist 0
# HELP sphinx_command_search Number of search command since server start.
# TYPE sphinx_command_search counter
sphinx_command_search 802
# HELP sphinx_command_status Number of status command since server start.
# TYPE sphinx_command_status counter
sphinx_command_status 133
# HELP sphinx_command_update Number of update command since server start.
# TYPE sphinx_command_update counter
sphinx_command_update 0
# HELP sphinx_connections Number of connections since the server started.
# TYPE sphinx_connections counter
sphinx_connections 59691
# HELP sphinx_dist_local Number of dist_local since server start.
# TYPE sphinx_dist_local counter
sphinx_dist_local 0
# HELP sphinx_dist_queries Number of distributed queries since server start.
# TYPE sphinx_dist_queries counter
sphinx_dist_queries 0
# HELP sphinx_dist_wait Number of dist_wait since server start.
# TYPE sphinx_dist_wait counter
sphinx_dist_wait 0
# HELP sphinx_dist_wall Number of dist_wall since server start.
# TYPE sphinx_dist_wall counter
sphinx_dist_wall 0
# HELP sphinx_maxed_out Number of max children barier since the server started.
# TYPE sphinx_maxed_out counter
sphinx_maxed_out 0
# HELP sphinx_qcache_cached_queries Number of qcache_cached_queries since server start.
# TYPE sphinx_qcache_cached_queries counter
sphinx_qcache_cached_queries 0
# HELP sphinx_qcache_max_bytes Number of qcache_max_bytes since server start.
# TYPE sphinx_qcache_max_bytes counter
sphinx_qcache_max_bytes 1.6777216e+07
# HELP sphinx_qcache_thresh_msec Number of qcache_thresh_msec since server start.
# TYPE sphinx_qcache_thresh_msec counter
sphinx_qcache_thresh_msec 3000
# HELP sphinx_qcache_ttl_sec Number of qcache_ttl_sec since server start.
# TYPE sphinx_qcache_ttl_sec counter
sphinx_qcache_ttl_sec 60
# HELP sphinx_qcache_used_bytes Number of qcache_used_bytes since server start.
# TYPE sphinx_qcache_used_bytes counter
sphinx_qcache_used_bytes 0
# HELP sphinx_queries Number of queries since server start.
# TYPE sphinx_queries counter
sphinx_queries 802
# HELP sphinx_query_cpu Number of query_cpu since server start.
# TYPE sphinx_query_cpu gauge
sphinx_query_cpu 0
# HELP sphinx_query_readkb Number of query_readkb since server start.
# TYPE sphinx_query_readkb gauge
sphinx_query_readkb 0
# HELP sphinx_query_reads Number of query_reads since server start.
# TYPE sphinx_query_reads gauge
sphinx_query_reads 0
# HELP sphinx_query_readtime Number of query_readtime since server start.
# TYPE sphinx_query_readtime gauge
sphinx_query_readtime 0
# HELP sphinx_query_wall Number of query_wall since server start.
# TYPE sphinx_query_wall counter
sphinx_query_wall 0.291
# HELP sphinx_up Could the searchd server be reached.
# TYPE sphinx_up gauge
sphinx_up 1
# HELP sphinx_uptime Number of seconds since the server started.
# TYPE sphinx_uptime gauge
sphinx_uptime 92852
# HELP sphinx_index_count Number of indexes.
# TYPE sphinx_index_count counter
sphinx_index_count 10
```
