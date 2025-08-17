package domain_test

import (
	"testing"

	"encoder/domain"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()

	require.Error(t, err)
}

func TestVideoIdIsNotAUUID(t *testing.T) {
	video := domain.NewVideo()
	video.ID = "invalid-uuid"
	err := video.Validate()

	require.Error(t, err)
}

func TestVideoValidation(t *testing.T) {
	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.ResourceID = "123"
	video.FilePath = "teste"
	err := video.Validate()

	require.NoError(t, err)
}
