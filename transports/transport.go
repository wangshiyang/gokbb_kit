package transports

import (
	"shawn/gokbb_kit/services"
	"github.com/go-kit/kit/endpoint"
	"shawn/gokbb_kit/common/param"
	"context"
	"net/http"
	"encoding/json"
	"github.com/nats-io/go-nats"
	natstransport "github.com/go-kit/kit/transport/nats"
)

func MakeUppercaseEndpoint(svc services.StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(param.UppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return param.UppercaseResponse{v, err.Error()}, nil
		}

		return param.UppercaseResponse{v, ""}, nil
	}
}

func MakeCountEndpoint(svc services.StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(param.CountRequest)

		v := svc.Count(req.S)

		return param.CountResponse{v}, nil
	}
}

func MakeUppercaseHTTPEndpoint(nc *nats.Conn) endpoint.Endpoint {
	return natstransport.NewPublisher(nc,
		"stringsvc.uppercase",
		natstransport.EncodeJSONRequest,
		DecodeUppercaseResponse,
	).Endpoint()
}

func MakeCountHTTPEndpoint(nc *nats.Conn) endpoint.Endpoint {
	return natstransport.NewPublisher(nc,
		"stringsvc.count",
		natstransport.EncodeJSONRequest,
		DecodeCountResponse,
	).Endpoint()
}

func DecodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request param.UppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func DecodeUppercaseResponse(_ context.Context, msg *nats.Msg) (interface{}, error) {
	var response param.UppercaseResponse

	if err := json.Unmarshal(msg.Data, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func DecodeCountResponse(_ context.Context, msg *nats.Msg) (interface{}, error) {
	var response param.CountResponse

	if err := json.Unmarshal(msg.Data, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func DecodeUppercaseNatsRequest(_ context.Context, msg *nats.Msg) (interface{}, error) {
	var request param.UppercaseRequest

	if err := json.Unmarshal(msg.Data, &request); err != nil {
		return nil, err
	}

	return request, nil
}

func DecodeCountNatsRequest(_ context.Context, msg *nats.Msg) (interface{}, error) {
	var request param.CountRequest

	if err := json.Unmarshal(msg.Data, &request); err != nil {
		return nil, err
	}

	return request, nil
}

func DecodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request param.CountRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
