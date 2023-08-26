package service

import (
	"fmt"
	"path/filepath"

	"github.com/bagusyanuar/go-yousee/app/admin/repositories"
	"github.com/bagusyanuar/go-yousee/app/admin/request"
	"github.com/bagusyanuar/go-yousee/common"
	"github.com/bagusyanuar/go-yousee/model"
	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

type (
	ImageItemService interface {
		Create(request request.ItemImageRequest) (*model.ItemImage, error)
		GetDataByItemID(id string) ([]model.ItemImage, error)
	}

	ItemImage struct {
		itemImageRepository repositories.ItemImageRepository
	}
)

const (
	ImageThumbnailPath = "assets/item/thumbnails"
	ImageOriginalPath  = "assets/item/originals"
)

// Create implements ImageItemService.
func (svc *ItemImage) Create(request request.ItemImageRequest) (*model.ItemImage, error) {
	itemID, _ := uuid.Parse(request.ItemID)
	var image string
	if request.Image != nil {
		paths := []string{ImageOriginalPath, ImageThumbnailPath}
		fileSystem := common.FileSystem{
			File: request.Image,
		}
		img, err := svc.uploadImages(fileSystem, paths...)
		if err != nil {
			return nil, err
		}
		image = img
	}

	entity := model.ItemImage{
		ItemID: itemID,
		Image:  image,
		Type:   request.Type,
	}
	return svc.itemImageRepository.Create(entity)
}

// GetDataByItemID implements ImageItemService.
func (svc *ItemImage) GetDataByItemID(id string) ([]model.ItemImage, error) {
	panic("unimplemented")
}

func (svc *ItemImage) uploadImages(fs common.FileSystem, paths ...string) (string, error) {
	jobGroup := errgroup.Group{}

	//job group for checking image paths
	for _, path := range paths {
		p := path
		jobGroup.Go(func() error {
			if err := fs.CheckPath(p); err != nil {
				return err
			}
			return nil
		})
	}

	if err := jobGroup.Wait(); err != nil {
		return "", err
	}

	ext := filepath.Ext(fs.File.Filename)
	image := fmt.Sprintf("%s%s", uuid.New().String(), ext)
	imageAddressOriginal := fmt.Sprintf("%s/%s", ImageOriginalPath, image)
	imageAddressThumbail := fmt.Sprintf("%s/%s", ImageThumbnailPath, image)

	//upload original image
	jobGroup.Go(func() error {
		err := fs.Upload(imageAddressOriginal)
		if err != nil {
			return err
		}
		return nil
	})

	//upload and resize thumbnail image
	jobGroup.Go(func() error {
		err := fs.UploadAndResize(imageAddressThumbail, 100, ext)
		if err != nil {
			return err
		}
		return nil
	})

	if err := jobGroup.Wait(); err != nil {
		return "", err
	}
	return image, nil
}
func NewItemImage(itemImageRepo repositories.ItemImageRepository) ImageItemService {
	return &ItemImage{
		itemImageRepository: itemImageRepo,
	}
}
