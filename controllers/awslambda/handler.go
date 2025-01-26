package awslambda

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"

	"github.com/filipeandrade6/framer-psgr-download/domain/service"
	"github.com/filipeandrade6/framer-psgr-download/domain/usecases"
)

type Handler struct {
	psgr usecases.Presigner
}

func New(psgr usecases.Presigner) *Handler {
	return &Handler{psgr: psgr}
}

func (hdlr *Handler) Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	key := request.Body
	if key == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Bad Request: no filename was provided",
		}, nil
	}

	url, err := service.Generate(key, hdlr.psgr)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Internal Server Error: Generating pre-signed url for download",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       url,
	}, nil
}
