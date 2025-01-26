package awslambda

import (
	"github.com/filipeandrade6/framer-psgr-download/domain/usecases"

	"github.com/aws/aws-lambda-go/lambda"
)

func Start(psgr usecases.Presigner) {
	lambdaHndlr := New(psgr)
	lambda.Start(lambdaHndlr.Handler)
}
