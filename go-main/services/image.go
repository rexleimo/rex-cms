package services

import (
	"rexai.com/helpers"
	"rexai.com/models"
)

type Image struct {
}

func (i *Image) List() ([]models.Image, error) {
	dbHelper := helpers.GetDbInstance()
	var images []models.Image
	dbHelper.GetDb().Find(&images)
	return images, nil
}

func (i *Image) Get(id int) (*models.Image, error) {
	dbHelper := helpers.GetDbInstance()
	var image models.Image
	dbHelper.GetDb().First(&image, id)
	return &image, nil
}

func (i *Image) Create(image *models.Image) error {
	dbHelper := helpers.GetDbInstance()
	tx := dbHelper.GetDb().Create(&image)
	return tx.Error
}

func (i *Image) Delete(id int) error {
	dbHelper := helpers.GetDbInstance()
	var image models.Image
	dbHelper.GetDb().Delete(&image, id)
	return nil
}

func (i *Image) Update(id int, image models.Image) (*models.Image, error) {
	dbHelper := helpers.GetDbInstance()
	var imageUpdate models.Image
	dbHelper.GetDb().First(&imageUpdate, id)
	dbHelper.GetDb().Model(&imageUpdate).Updates(image)
	return &imageUpdate, nil
}
