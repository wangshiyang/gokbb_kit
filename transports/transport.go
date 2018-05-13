package transports

import (
	"shawn/gokbb_kit/services"
	"github.com/go-kit/kit/endpoint"
	"shawn/gokbb_kit/common/param"
	"context"
	"net/http"
	"encoding/json"
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

func DecodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request param.UppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
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
