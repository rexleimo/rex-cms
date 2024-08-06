package services_test

import (
	"testing"

	"rexai.com/helpers"
	"rexai.com/models"
	"rexai.com/services"
)

func TestMain(m *testing.M) {
	helpers.GetDbInstance()
	m.Run()
}

func TestImageInsert(t *testing.T) {
	service := services.Image{}
	entity := &models.Image{
		URL:      "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png",
		Width:    200,
		Height:   200,
		Size:     100,
		MimeType: "image/png",
	}
	err := service.Create(entity)
	if err != nil {
		t.Errorf("Error inserting image: %v", err)
	}
}

func TestImageGet(t *testing.T) {
	service := services.Image{}
	entity, err := service.Get(1)
	if err != nil {
		t.Errorf("Error getting image: %v", err)
	}
	if entity.Id != 1 {
		t.Errorf("Expected image ID to be 1, got %v", entity.Id)
	}
	t.Logf("Image: %v", entity)
}

func TestImageList(t *testing.T) {
	service := services.Image{}
	entities, err := service.List()
	if err != nil {
		t.Errorf("Error listing images: %v", err)
	}
	if len(entities) == 0 {
		t.Errorf("Expected at least one image, got %v", len(entities))
	}
	t.Logf("Images: %v", entities)
}

func TestImageDelete(t *testing.T) {
	service := services.Image{}
	err := service.Delete(1)
	if err != nil {
		t.Errorf("Error deleting image: %v", err)
	}
	lists, err := service.List()
	if err != nil {
		t.Errorf("Error listing images: %v", err)
	}
	t.Log("lists", lists)
}
