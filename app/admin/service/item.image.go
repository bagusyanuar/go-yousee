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
		var (
			checkPathGroup errgroup.Group
			uploadGroup    errgroup.Group
		)
		paths := []string{ImageOriginalPath, ImageThumbnailPath}
		fileSystem := common.FileSystem{
			File: request.Image,
		}
		for _, path := range paths {
			p := path
			checkPathGroup.Go(func() error {

				if err := fileSystem.CheckPath(p); err != nil {
					return err
				}
				return nil
			})
		}

		if err := checkPathGroup.Wait(); err != nil {
			return nil, err
		}

		ext := filepath.Ext(request.Image.Filename)
		image = fmt.Sprintf("%s%s", uuid.New().String(), ext)
		imageAddressOriginal := fmt.Sprintf("%s/%s", ImageOriginalPath, image)
		imageAddressThumbail := fmt.Sprintf("%s/%s", ImageThumbnailPath, image)

		//upload original image
		uploadGroup.Go(func() error {
			err := fileSystem.Upload(imageAddressOriginal)
			if err != nil {
				return err
			}
			return nil
		})

		//upload and resize thumbnail image
		uploadGroup.Go(func() error {
			err := fileSystem.UploadAndResize(imageAddressThumbail, 100, ext)
			if err != nil {
				return err
			}
			return nil
		})

		if err := uploadGroup.Wait(); err != nil {
			return nil, err
		}
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

func NewItemImage(itemImageRepo repositories.ItemImageRepository) ImageItemService {
	return &ItemImage{
		itemImageRepository: itemImageRepo,
	}
}
