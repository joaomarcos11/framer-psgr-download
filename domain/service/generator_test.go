package service

import (
	"testing"

	"github.com/filipeandrade6/framer-psgr-download/domain/errors"
	"github.com/stretchr/testify/assert"
)

type mockPresignerSuccess struct{}

func NewMockPresignerSuccess() *mockPresignerSuccess {
	return &mockPresignerSuccess{}
}

func (mockPresignerSuccess) GeneratePreSignedUrl(bucket, filename string) (string, error) {
	return "http://success.com", nil
}

type mockPresignerFail struct{}

func NewMockPresignerFail() *mockPresignerFail {
	return &mockPresignerFail{}
}

func (mockPresignerFail) GeneratePreSignedUrl(bucket, filename string) (string, error) {
	return "", errors.ErrGeneratePreSignedURL
}

func TestGenerate(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		psgr := NewMockPresignerSuccess()
		_, err := Generate("email@email.com", psgr)

		assert.NoError(t, err)
	})

	t.Run("fail to generate pre-signed url", func(t *testing.T) {
		psgr := NewMockPresignerFail()
		_, err := Generate("email@email.com", psgr)

		assert.Error(t, err)
		assert.ErrorAs(t, err, &errors.ErrGeneratePreSignedURL)
	})
}
