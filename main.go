package main

import (
	"database/sql"
	"math"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"github.com/prometheus/common/version"
	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	namespace = "sphinx"
)

var (
	labels = []string{"index"}
)

// Exporter collects metrics from a searchd server.
type Exporter struct {
	sphinx string

	up                    *prometheus.Desc
	uptime                *prometheus.Desc
	connections           *prometheus.Desc
	maxed_out             *prometheus.Desc
	command_search        *prometheus.Desc
	command_update        *prometheus.Desc
	command_delete        *prometheus.Desc
	command_keywords      *prometheus.Desc
	command_persist       *prometheus.Desc
	command_status        *prometheus.Desc
	command_flushattrs    *prometheus.Desc
	agent_connect         *prometheus.Desc
	agent_retry           *prometheus.Desc
	queries               *prometheus.Desc
	dist_queries          *prometheus.Desc
	query_wall            *prometheus.Desc
	query_cpu             *prometheus.Desc
	dist_wall             *prometheus.Desc
	dist_local            *prometheus.Desc
	dist_wait             *prometheus.Desc
	query_reads           *prometheus.Desc
	query_readkb          *prometheus.Desc
	query_readtime        *prometheus.Desc
	avg_query_wall        *prometheus.Desc
	avg_query_cpu         *prometheus.Desc
	avg_dist_wall         *prometheus.Desc
	avg_dist_local        *prometheus.Desc
	avg_dist_wait         *prometheus.Desc
	avg_query_reads       *prometheus.Desc
	avg_query_readkb      *prometheus.Desc
	avg_query_readtime    *prometheus.Desc
	qcache_max_bytes      *prometheus.Desc
	qcache_thresh_msec    *prometheus.Desc
	qcache_ttl_sec        *prometheus.Desc
	qcache_cached_queries *prometheus.Desc
	qcache_used_bytes     *prometheus.Desc
	qcache_hits           *prometheus.Desc
	index_count           *prometheus.Desc
	indexed_documents     *prometheus.Desc
	indexed_bytes         *prometheus.Desc
	field_tokens_title    *prometheus.Desc
	field_tokens_body     *prometheus.Desc
	total_tokens          *prometheus.Desc
	ram_bytes             *prometheus.Desc
	disk_bytes            *prometheus.Desc
	mem_limit             *prometheus.Desc
	threads_count         *prometheus.Desc
}

func NewExporter(server string, port string, timeout time.Duration) *Exporter {
	c := "@tcp(" + server + ":" + port + ")/"

	return &Exporter{
		sphinx: c,
		up: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "up"),
			"Could the searchd server be reached.",
			nil,
			nil,
		),
		uptime: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "uptime"),
			"Number of seconds since the server started.",
			nil,
			nil,
		),
		connections: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "connections"),
			"Number of connections since the server started.",
			nil,
			nil,
		),
		maxed_out: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "maxed_out"),
			"Number of max children barier since the server started.",
			nil,
			nil,
		),
		command_search: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "command_search"),
			"Number of search command since server start.",
			nil,
			nil,
		),
		command_update: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "command_update"),
			"Number of update command since server start.",
			nil,
			nil,
		),
		command_delete: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "command_delete"),
			"Number of delete command since server start.",
			nil,
			nil,
		),
		command_keywords: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "command_keywords"),
			"Number of keywords command since server start.",
			nil,
			nil,
		),
		command_persist: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "command_persist"),
			"Number of persist command since server start.",
			nil,
			nil,
		),
		command_status: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "command_status"),
			"Number of status command since server start.",
			nil,
			nil,
		),
		command_flushattrs: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "command_flushattrs"),
			"Number of flushattrs command since server start.",
			nil,
			nil,
		),
		agent_connect: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "agent_connect"),
			"Number of agent connect since server start.",
			nil,
			nil,
		),
		agent_retry: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "agent_retry"),
			"Number of agent retry since server start.",
			nil,
			nil,
		),
		queries: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "queries"),
			"Number of queries since server start.",
			nil,
			nil,
		),
		dist_queries: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "dist_queries"),
			"Number of distributed queries since server start.",
			nil,
			nil,
		),
		query_wall: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "query_wall"),
			"Number of query_wall since server start.",
			nil,
			nil,
		),
		query_cpu: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "query_cpu"),
			"Number of query_cpu since server start.",
			nil,
			nil,
		),
		dist_wall: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "dist_wall"),
			"Number of dist_wall since server start.",
			nil,
			nil,
		),
		dist_local: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "dist_local"),
			"Number of dist_local since server start.",
			nil,
			nil,
		),
		dist_wait: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "dist_wait"),
			"Number of dist_wait since server start.",
			nil,
			nil,
		),
		query_reads: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "query_reads"),
			"Total number of read operations since server start.",
			nil,
			nil,
		),
		query_readkb: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "query_readkb"),
			"Total number of KB read since server start.",
			nil,
			nil,
		),
		query_readtime: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "query_readtime"),
			"Total time spent doing read operations (in seconds) since server start.",
			nil,
			nil,
		),
		avg_query_wall: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "avg_query_wall"),
			"Number of avg_query_wall since server start.",
			nil,
			nil,
		),
		avg_query_cpu: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "avg_query_cpu"),
			"Number of avg_query_cpu since server start.",
			nil,
			nil,
		),
		avg_dist_wall: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "avg_dist_wall"),
			"Number of avg_dist_wall since server start.",
			nil,
			nil,
		),
		avg_dist_local: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "avg_dist_local"),
			"Number of avg_dist_local since server start.",
			nil,
			nil,
		),
		avg_dist_wait: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "avg_dist_wait"),
			"Number of avg_dist_wait since server start.",
			nil,
			nil,
		),
		avg_query_reads: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "avg_query_reads"),
			"Number of avg_query_reads since server start.",
			nil,
			nil,
		),
		avg_query_readkb: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "avg_query_readkb"),
			"Number of avg_query_readkb since server start.",
			nil,
			nil,
		),
		avg_query_readtime: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "avg_query_readtime"),
			"Number of avg_query_readtime since server start.",
			nil,
			nil,
		),
		qcache_max_bytes: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "qcache_max_bytes"),
			"Number of qcache_max_bytes since server start.",
			nil,
			nil,
		),
		qcache_thresh_msec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "qcache_thresh_msec"),
			"Number of qcache_thresh_msec since server start.",
			nil,
			nil,
		),
		qcache_ttl_sec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "qcache_ttl_sec"),
			"Number of qcache_ttl_sec since server start.",
			nil,
			nil,
		),
		qcache_cached_queries: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "qcache_cached_queries"),
			"Number of qcache_cached_queries since server start.",
			nil,
			nil,
		),
		qcache_used_bytes: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "qcache_used_bytes"),
			"Number of qcache_used_bytes since server start.",
			nil,
			nil,
		),
		qcache_hits: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "qcache_hits"),
			"Number of qcache_hits since server start.",
			nil,
			nil,
		),
		index_count: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "index_count"),
			"Number of indexes.",
			nil,
			nil,
		),
		indexed_documents: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "indexed_documents"),
			"Number of documents indexed",
			labels,
			nil,
		),
		indexed_bytes: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "indexed_bytes"),
			"Size (in bytes) of indexed documents",
			labels,
			nil,
		),
		field_tokens_title: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "field_tokens_title"),
			"Sums of per-field length titles over the entire index",
			labels,
			nil,
		),
		field_tokens_body: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "field_tokens_body"),
			"Sums of per-field length bodies over the entire index",
			labels,
			nil,
		),
		total_tokens: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "total_tokens"),
			"Total tokens",
			labels,
			nil,
		),
		ram_bytes: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "ram_bytes"),
			"Total size (in bytes) of the RAM-resident index portion",
			labels,
			nil,
		),
		disk_bytes: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "disk_bytes"),
			"Total size (in bytes) of indexes on disk",
			labels,
			nil,
		),
		mem_limit: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "mem_limit"),
			"Memory limit",
			labels,
			nil,
		),
		threads_count: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "threads_count"),
			"Number of threads",
			[]string{"state"},
			nil,
		),
	}
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.up
	ch <- e.uptime
	ch <- e.connections
	ch <- e.maxed_out
	ch <- e.command_search
	ch <- e.command_update
	ch <- e.command_delete
	ch <- e.command_keywords
	ch <- e.command_persist
	ch <- e.command_status
	ch <- e.command_flushattrs
	ch <- e.agent_connect
	ch <- e.agent_retry
	ch <- e.queries
	ch <- e.dist_queries
	ch <- e.query_wall
	ch <- e.query_cpu
	ch <- e.dist_wall
	ch <- e.dist_local
	ch <- e.dist_wait
	ch <- e.query_reads
	ch <- e.query_readkb
	ch <- e.query_readtime
	ch <- e.avg_query_wall
	ch <- e.avg_query_cpu
	ch <- e.avg_dist_wall
	ch <- e.avg_dist_local
	ch <- e.avg_dist_wait
	ch <- e.avg_query_reads
	ch <- e.avg_query_readkb
	ch <- e.avg_query_readtime
	ch <- e.qcache_max_bytes
	ch <- e.qcache_thresh_msec
	ch <- e.qcache_ttl_sec
	ch <- e.qcache_cached_queries
	ch <- e.qcache_used_bytes
	ch <- e.qcache_hits
	ch <- e.index_count
	ch <- e.indexed_documents
	ch <- e.indexed_bytes
	ch <- e.field_tokens_title
	ch <- e.field_tokens_body
	ch <- e.total_tokens
	ch <- e.ram_bytes
	ch <- e.disk_bytes
	ch <- e.mem_limit
	ch <- e.threads_count
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	status := 1

	db, err := sql.Open("mysql", e.sphinx)
	if err != nil {
		printerror(e, err, ch)
		return
	}
	// will close DB connection while return
	defer db.Close()

	rows, err := db.Query("SHOW STATUS")
	if err != nil {
		printerror(e, err, ch)
		return
	}
	variables := make(map[string]string)

	for rows.Next() {
		var metric string
		var counter string
		err = rows.Scan(&metric, &counter)
		if err != nil {
			printerror(e, err, ch)
			return
		}
		variables[metric] = counter
	}

	for k, v := range variables {
		switch {
		case k == "uptime":
			ch <- prometheus.MustNewConstMetric(e.uptime, prometheus.GaugeValue, parse(v))
		case k == "connections":
			ch <- prometheus.MustNewConstMetric(e.connections, prometheus.CounterValue, parse(v))
		case k == "maxed_out":
			ch <- prometheus.MustNewConstMetric(e.maxed_out, prometheus.CounterValue, parse(v))
		case k == "command_search":
			ch <- prometheus.MustNewConstMetric(e.command_search, prometheus.CounterValue, parse(v))
		case k == "command_update":
			ch <- prometheus.MustNewConstMetric(e.command_update, prometheus.CounterValue, parse(v))
		case k == "command_delete":
			ch <- prometheus.MustNewConstMetric(e.command_delete, prometheus.CounterValue, parse(v))
		case k == "command_keywords":
			ch <- prometheus.MustNewConstMetric(e.command_keywords, prometheus.CounterValue, parse(v))
		case k == "command_persist":
			ch <- prometheus.MustNewConstMetric(e.command_persist, prometheus.CounterValue, parse(v))
		case k == "command_status":
			ch <- prometheus.MustNewConstMetric(e.command_status, prometheus.CounterValue, parse(v))
		case k == "command_flushattrs":
			ch <- prometheus.MustNewConstMetric(e.command_flushattrs, prometheus.CounterValue, parse(v))
		case k == "agent_connect":
			ch <- prometheus.MustNewConstMetric(e.agent_connect, prometheus.CounterValue, parse(v))
		case k == "agent_retry":
			ch <- prometheus.MustNewConstMetric(e.agent_retry, prometheus.CounterValue, parse(v))
		case k == "queries":
			ch <- prometheus.MustNewConstMetric(e.queries, prometheus.CounterValue, parse(v))
		case k == "dist_queries":
			ch <- prometheus.MustNewConstMetric(e.dist_queries, prometheus.CounterValue, parse(v))
		case k == "query_wall":
			ch <- prometheus.MustNewConstMetric(e.query_wall, prometheus.CounterValue, parse(v))
		case k == "query_cpu":
			ch <- prometheus.MustNewConstMetric(e.query_cpu, prometheus.GaugeValue, parse(v))
		case k == "dist_wall":
			ch <- prometheus.MustNewConstMetric(e.dist_wall, prometheus.CounterValue, parse(v))
		case k == "dist_local":
			ch <- prometheus.MustNewConstMetric(e.dist_local, prometheus.CounterValue, parse(v))
		case k == "dist_wait":
			ch <- prometheus.MustNewConstMetric(e.dist_wait, prometheus.CounterValue, parse(v))
		case k == "query_reads":
			ch <- prometheus.MustNewConstMetric(e.query_reads, prometheus.CounterValue, parse(v))
		case k == "query_readkb":
			ch <- prometheus.MustNewConstMetric(e.query_readkb, prometheus.CounterValue, parse(v))
		case k == "query_readtime":
			ch <- prometheus.MustNewConstMetric(e.query_readtime, prometheus.CounterValue, parse(v))
		case k == "avg_query_wall":
			ch <- prometheus.MustNewConstMetric(e.avg_query_wall, prometheus.CounterValue, parse(v))
		case k == "avg_query_cpu":
			ch <- prometheus.MustNewConstMetric(e.avg_query_cpu, prometheus.GaugeValue, parse(v))
		case k == "avg_dist_wall":
			ch <- prometheus.MustNewConstMetric(e.avg_dist_wall, prometheus.CounterValue, parse(v))
		case k == "avg_dist_local":
			ch <- prometheus.MustNewConstMetric(e.avg_dist_local, prometheus.CounterValue, parse(v))
		case k == "avg_dist_wait":
			ch <- prometheus.MustNewConstMetric(e.avg_dist_wait, prometheus.CounterValue, parse(v))
		case k == "avg_query_reads":
			ch <- prometheus.MustNewConstMetric(e.avg_query_reads, prometheus.GaugeValue, parse(v))
		case k == "avg_query_readkb":
			ch <- prometheus.MustNewConstMetric(e.avg_query_readkb, prometheus.GaugeValue, parse(v))
		case k == "avg_query_readtime":
			ch <- prometheus.MustNewConstMetric(e.avg_query_readtime, prometheus.GaugeValue, parse(v))
		case k == "qcache_max_bytes":
			ch <- prometheus.MustNewConstMetric(e.qcache_max_bytes, prometheus.CounterValue, parse(v))
		case k == "qcache_thresh_msec":
			ch <- prometheus.MustNewConstMetric(e.qcache_thresh_msec, prometheus.CounterValue, parse(v))
		case k == "qcache_ttl_sec":
			ch <- prometheus.MustNewConstMetric(e.qcache_ttl_sec, prometheus.CounterValue, parse(v))
		case k == "qcache_cached_queries":
			ch <- prometheus.MustNewConstMetric(e.qcache_cached_queries, prometheus.CounterValue, parse(v))
		case k == "qcache_used_bytes":
			ch <- prometheus.MustNewConstMetric(e.qcache_used_bytes, prometheus.CounterValue, parse(v))
		case k == "qcache_used_bytes":
			ch <- prometheus.MustNewConstMetric(e.qcache_used_bytes, prometheus.CounterValue, parse(v))
		}
	}

	//Collect Indexes
	databases := make(map[string]string)

	indexes, err := db.Query("SHOW TABLES")
	if err != nil {
		printerror(e, err, ch)
		return
	}

	for indexes.Next() {
		var index string
		var index_type string
		err = indexes.Scan(&index, &index_type)
		if err != nil {
			printerror(e, err, ch)
			return
		}
		databases[index] = index_type
	}
	ch <- prometheus.MustNewConstMetric(e.index_count, prometheus.GaugeValue, float64(len(databases)))

	//Collect metrics per index
	for index, _ := range databases {

		// Distributed indexes has no index status

		if databases[index] == "distributed" {
			continue
		}
		metrics, err := db.Query("SHOW INDEX " + index + " STATUS")
		if err != nil {
			printerror(e, err, ch)
			return
		}

		for metrics.Next() {
			var metric string
			var value string
			err := metrics.Scan(&metric, &value)
			if err != nil {
				printerror(e, err, ch)
				return
			}
			switch {
			case metric == "indexed_documents":
				ch <- prometheus.MustNewConstMetric(e.indexed_documents, prometheus.CounterValue, parse(value), index)
			case metric == "indexed_bytes":
				ch <- prometheus.MustNewConstMetric(e.indexed_bytes, prometheus.CounterValue, parse(value), index)
			case metric == "field_tokens_title":
				ch <- prometheus.MustNewConstMetric(e.field_tokens_title, prometheus.CounterValue, parse(value), index)
			case metric == "field_tokens_body":
				ch <- prometheus.MustNewConstMetric(e.field_tokens_body, prometheus.CounterValue, parse(value), index)
			case metric == "total_tokens":
				ch <- prometheus.MustNewConstMetric(e.total_tokens, prometheus.CounterValue, parse(value), index)
			case metric == "ram_bytes":
				ch <- prometheus.MustNewConstMetric(e.ram_bytes, prometheus.CounterValue, parse(value), index)
			case metric == "disk_bytes":
				ch <- prometheus.MustNewConstMetric(e.disk_bytes, prometheus.CounterValue, parse(value), index)
			case metric == "mem_limit":
				ch <- prometheus.MustNewConstMetric(e.mem_limit, prometheus.CounterValue, parse(value), index)
			}
		}
	}

	threads_rows, err := db.Query("SHOW THREADS")
	if err != nil {
		printerror(e, err, ch)
		return
	}

	threads := make(map[string][]string)

	for threads_rows.Next() {
		var tid string
		var proto string
		var state string
		var time string
		var info string
		err := threads_rows.Scan(&tid, &proto, &state, &time, &info)
		if err != nil {
			printerror(e, err, ch)
			return
		}
		threads[state] = append(threads[state], tid)
	}

	for threads_state, threads_times := range threads {
		ch <- prometheus.MustNewConstMetric(e.threads_count, prometheus.CounterValue, float64(len(threads_times)), threads_state)
	}

	ch <- prometheus.MustNewConstMetric(e.up, prometheus.GaugeValue, float64(status))
}

func parse(stat string) float64 {
	switch {
	case stat == "ON":
		return 1.0
	case stat == "OFF":
		return 0
	}

	v, err := strconv.ParseFloat(stat, 64)
	if v == 0 {
		return 0.0
	}

	if err != nil {
		log.Errorf("Failed to parse %s: %s", stat, err)
		v = math.NaN()
	}
	return v
}

func printerror(e *Exporter, err error, ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(e.up, prometheus.GaugeValue, 0)
	return
}

func main() {
	var (
		address       = kingpin.Flag("sphinx.address", "Sphinx server address.").Default("127.0.0.1").String()
		port          = kingpin.Flag("sphinx.port", "Sphinx server port.").Default("9306").String()
		timeout       = kingpin.Flag("sphinx.timeout", "memcached connect timeout.").Default("1s").Duration()
		listenAddress = kingpin.Flag("web.listen-address", "Address to listen on for web interface and telemetry.").Default(":9247").String()
		metricsPath   = kingpin.Flag("web.telemetry-path", "Path under which to expose metrics.").Default("/metrics").String()
	)

	log.AddFlags(kingpin.CommandLine)
	kingpin.Version(version.Print("sphinx_exporter"))
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	log.Infoln("Starting sphinx_exporter", version.Info())
	log.Infoln("Build context", version.BuildContext())

	prometheus.MustRegister(NewExporter(*address, *port, *timeout))

	http.Handle(*metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
      <head><title>Sphinx Exporter</title></head>
      <body>
      <h1>Sphinx Exporter</h1>
      <p><a href='` + *metricsPath + `'>Metrics</a></p>
      </body>
      </html>`))
	})
	log.Infoln("Starting HTTP server on", *listenAddress)
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}
