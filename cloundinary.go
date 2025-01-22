package main

import (
	"context"
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

type CldnryStore struct {
	Cloudinary   *cloudinary.Cloudinary
	Context      context.Context
	UploadPreset string
}

func NewCloudinaryStore(cloudinaryURL string, ctx context.Context, uploadPreset string) (*CldnryStore, error) {
	logInfo("Running: Cloudinary INIT - NewCloudinaryStore")
	cld, err := cloudinary.New()
	if err != nil {
		return nil, err
	}
	return &CldnryStore{
		Cloudinary:   cld,
		Context:      ctx,
		UploadPreset: uploadPreset,
	}, nil
}

func (s *CldnryStore) GetImageURLByPublicId(publicId string) (string, error) {
	logInfo("Running: Cloudinary - GetImageURLByPublicId")
	img, err := s.Cloudinary.Image(publicId)
	if err != nil {
		return "", err
	}
	url, err := img.String()
	if err != nil {
		return "", err
	}
	return url, nil
}

func (s *CldnryStore) UploadImage(imageID uuid.UUID, stringbase64URI string) (string, error) {
	logInfo("Running: Cloudinary - UploadImage")
	res, err := s.Cloudinary.Upload.UnsignedUpload(s.Context, stringbase64URI, s.UploadPreset, uploader.UploadParams{
		FilenameOverride: imageID.String(),
		PublicID:         imageID.String(),
	})
	if err != nil {
		return "", err
	}
	if res.Error.Message != "" {
		return "", fmt.Errorf("%v", res.Error.Message)
	}
	return res.SecureURL, nil
}
