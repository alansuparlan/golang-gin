package service

import (
	"testing"

	"github.com/alansuparlan/golang-gin/entity"
	"github.com/alansuparlan/golang-gin/repository"
	"github.com/stretchr/testify/assert"
)

const (
	TITLE       = "Video Title"
	DESCRIPTION = "Video Description"
	URL         = "https://youtu.be/JgW-i2QjgHQ"
)

func getVideo() entity.Video {
	return entity.Video{
		Title:       TITLE,
		Description: DESCRIPTION,
		URL:         URL,
	}
}

func TestFindAll(t *testing.T) {
	videoRepository := repository.NewVideoRepository()
	defer videoRepository.CloseDB()

	videoService := New(videoRepository)

	videoService.Save(getVideo())

	videos := videoService.FindAll()

	firstVideo := videos[0]
	assert.NotNil(t, videos)
	assert.Equal(t, TITLE, firstVideo.Title)
	assert.Equal(t, DESCRIPTION, firstVideo.Description)
	assert.Equal(t, URL, firstVideo.URL)
}
