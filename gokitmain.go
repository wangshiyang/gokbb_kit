package main

import (
	httptransport "github.com/go-kit/kit/transport/http"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"shawn/gokbb_kit/transports"
	"net/http"
	"shawn/gokbb_kit/services"
	log "github.com/go-kit/kit/log"
	"os"
	"shawn/gokbb_kit/common/util"
	"shawn/gokbb_kit/middleware/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/nats-io/go-nats"
	"flag"
	log2 "log"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	natsURL := flag.String("nats-url", nats.DefaultURL, "URL for connection to NATS")
	flag.Parse()

	nc, err := nats.Connect(*natsURL)

	if err != nil {
		log2.Fatal(err)
	}

	defer nc.Close()

	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_count",
		Help:      "Number of request recevied",
	}, fieldKeys)

	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)

	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{})

	// go语言的interface没有明确的数据类型 只要定义好strut 并实现interface的方法就可以同种类型转换
	var svc services.StringService
	svc = services.StringServiceStrut{}
	//svc := services.StringServiceStrut{}
	svc = util.LoggingMiddleware{logger, svc}
	svc = metrics.InstrumentingMiddleware{requestCount, requestLatency, countResult, svc}

	uppercaseHandler := httptransport.NewServer(
		transports.MakeUppercaseEndpoint(svc),
		transports.DecodeUppercaseRequest,
		transports.EncodeResponse,
	)

	countHandler := httptransport.NewServer(
		transports.MakeCountEndpoint(svc),
		transports.DecodeCountRequest,
		transports.EncodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", nil)
}
