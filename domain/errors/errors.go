package errors

import "errors"

var (
	ErrGeneratePreSignedURL = errors.New("failed to generate pre-signed url")
)
