package upload

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/google/uuid"
)

func CheckImageExt(fileName string) bool {
	ext := strings.ToUpper(path.Ext(fileName))

	return ext == ".PNG"
}

func GetTempFilePath() string {
	filePath := ""
	for {
		uuid := uuid.New()
		fileName := uuid.String()

		filePath = fmt.Sprint("images/", fileName, ".png")

		if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
			break
		}
	}
	return filePath
}
