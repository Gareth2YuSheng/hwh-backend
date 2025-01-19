package main

import (
	"context"
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

type CldnryStore struct {
	Cloudinary *cloudinary.Cloudinary
	Context    context.Context
}

func NewCloudinaryStore(cloudinaryURL string, ctx context.Context) (*CldnryStore, error) {
	logInfo("Running NewCloudinaryStore")
	cld, err := cloudinary.New()
	if err != nil {
		return nil, err
	}
	return &CldnryStore{
		Cloudinary: cld,
		Context:    ctx,
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

// func (s *CldnryStore) UploadImageTest() {
// 	logInfo("Running GetImageURLUploadImageTestByPublicId")
// 	res, err := s.Cloudinary.Upload.UnsignedUpload(s.Context, "test cloudinary images/hq720.jpg", "HomeworkHelp", uploader.UploadParams{
// 		PublicID: uuid.New().String(),
// 	})
// 	if err != nil {
// 		logError("Image Upload Test Failed", err)
// 		return
// 	}
// 	fmt.Printf("Cloundinary Image Test Response:\n%v\n", res.SecureURL)
// }

func (s *CldnryStore) UploadImage(imageID uuid.UUID, stringbase64URI string) (string, error) {
	logInfo("Running: Cloudinary - UploadImage")
	res, err := s.Cloudinary.Upload.UnsignedUpload(s.Context, "test cloudinary images/hq720.jpg", "HomeworkHelp", uploader.UploadParams{
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
