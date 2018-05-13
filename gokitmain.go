package main

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"shawn/gokbb_kit/transports"
	"net/http"
	"shawn/gokbb_kit/services"
	"github.com/go-kit/kit/log"
	"os"
	"shawn/gokbb_kit/common/util"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)
	// go语言的interface没有明确的数据类型 只要定义好strut 并实现interface的方法就可以同种类型转换
	var svc services.StringService
	svc = services.StringServiceStrut{}
	//svc := services.StringServiceStrut{}
	svc = util.LoggingMiddleware{logger, svc}

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

	http.ListenAndServe(":8080", nil)
}
