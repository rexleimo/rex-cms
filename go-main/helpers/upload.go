package helpers

import (
	"time"

	"rexai.com/utils"
)

type Upload struct{}

func (Upload) UploadFile(fileContent []byte, fileName string) (string, error) {
	uploader := utils.FileToGithub{}
	cig := GetConfigInstance()
	uploader.Owner = cig.Config.Github.Owner
	uploader.Repo = cig.Config.Github.Repo
	uploader.Token = cig.Config.Github.Token

	time := time.Now().Format("2006-01-02")
	uploader.FilePath = time + "/" + fileName
	uri, err := uploader.Upload(fileContent)
	if err != nil {
		return "", err
	}
	return uri, nil
}
