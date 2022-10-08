package interfaces

import (
	"github.com/abhishekr7/ganga/app/models"
)

type IAlbumRepository interface {
	getAlbumsByArtist(name string) ([]models.AlbumModel, error)
}