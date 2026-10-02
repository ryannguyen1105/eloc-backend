package service

import (
	"context"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type CloudinaryService struct {
	cld *cloudinary.Cloudinary
}

func NewCloudinaryService() (*CloudinaryService, error) {
	cld, err := cloudinary.New()
	if err != nil {
		return nil, err
	}
	return &CloudinaryService{cld: cld}, nil
}

func (cloudinaryService *CloudinaryService) UploadImage(ctx context.Context, file multipart.File, folderName string) (string, error) {
    uploadResult, err := cloudinaryService.cld.Upload.Upload(
        ctx,
        file,
        uploader.UploadParams{
            Folder: folderName,
        },
    )
    if err != nil {
        return "", err
    }
    return uploadResult.SecureURL, nil
}