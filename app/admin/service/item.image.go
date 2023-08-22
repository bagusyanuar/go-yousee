package service

import (
	"errors"
	"fmt"
	"sync"

	"github.com/bagusyanuar/go-yousee/app/admin/repositories"
	"github.com/bagusyanuar/go-yousee/app/admin/request"
	"github.com/bagusyanuar/go-yousee/common"
	"github.com/bagusyanuar/go-yousee/model"
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

// Create implements ImageItemService.
func (svc *ItemImage) Create(request request.ItemImageRequest) (*model.ItemImage, error) {
	errorCheckPath := make(chan error)
	wgCheckpathDone := make(chan bool)

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go svc.checkPath(i, &wg, errorCheckPath)
	}

	go func() {
		wg.Wait()
		close(wgCheckpathDone)
	}()

	select {
	case <-wgCheckpathDone:
		break
	case err := <-errorCheckPath:
		close(errorCheckPath)
		fmt.Println("error check path" + err.Error())
		return nil, err
	}

	return nil, nil
}

const (
	ImageThumbnailPath = "assets/item/thumbnails"
)

func (svc *ItemImage) checkPath(i int, wg *sync.WaitGroup, e chan error) {
	defer wg.Done()
	text := fmt.Sprintf("do job %d", i)
	fmt.Println(text)

	if i == 1 {
		e <- errors.New("error on 3")
	} else {
		fileSystem := common.FileSystem{}
		if err := fileSystem.CheckPath(ImageThumbnailPath); err != nil {
			e <- err
		}
	}

	// return nil
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
