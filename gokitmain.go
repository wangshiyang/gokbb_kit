package main

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"shawn/gokbb_kit/transports"
	"net/http"
	"shawn/gokbb_kit/services"
)

func main() {
	//var svc services.StringService
	//svc = services.StringServiceStrut{}
	svc := services.StringServiceStrut{}

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
