package main

import (
	"log"

	"github.com/filipeandrade6/framer-psgr-download/adapters/presigner/awss3"
	"github.com/filipeandrade6/framer-psgr-download/controllers/awslambda"
)

func main() {
	presigner, err := awss3.New()
	if err != nil {
		log.Fatalf("failed to configure pre-signer: %s", err)
	}

	awslambda.Start(presigner)
}
