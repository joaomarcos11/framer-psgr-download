package service

import (
	"fmt"
	"log/slog"

	"github.com/filipeandrade6/framer-psgr-download/domain/errors"
	"github.com/filipeandrade6/framer-psgr-download/domain/usecases"
)

func Generate(key string, psgrSvc usecases.Presigner) (string, error) {
	url, err := psgrSvc.GeneratePreSignedUrl("fiap44-framer-images", fmt.Sprintf("%s.zip", key))
	if err != nil {
		err = fmt.Errorf("%w: %w", errors.ErrGeneratePreSignedURL, err)
		slog.Error("generate", "err", err)
		return "", err
	}

	return url, nil
}
