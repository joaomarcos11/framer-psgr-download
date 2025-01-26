package service

import (
	"errors"
	"fmt"

	"github.com/filipeandrade6/framer-psgr-download/domain/usecases"
)

func Generate(key string, psgrSvc usecases.Presigner) (string, error) {
	url, err := psgrSvc.GeneratePreSignedUrl("fiap44-framer-images", fmt.Sprintf("%s.zip", key))
	if err != nil {
		return "", errors.New(fmt.Sprintf("failed to generate pre-signed url to download: %s", err))
	}

	return url, nil
}
