package utils

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/RaulCatalinas/StylelintBC/internal/enums"
	"github.com/skratchdot/open-golang/open"
)

func Exists(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}

func CreateFolder(name string) {
	err := os.Mkdir(name, 0755)

	if err != nil {
		errorMessage := GetErrorMessage("CreateFolder")

		WriteMessage(WriteMessageProps{
			Type:    enums.MessageTypeError,
			Message: strings.Replace(string(errorMessage), "{folderName}", name, -1),
		})

		os.Exit(1)
	}
}

func CreateEmptyJsonFile(name string) {
	emptyObject := make(map[string]interface{})

	file, err := os.Create(name + ".json")

	if err != nil {
		WriteMessage(WriteMessageProps{
			Type:    enums.MessageTypeError,
			Message: string(GetErrorMessage("CreateEmptyFile")),
		})

		os.Exit(1)
	}

	defer file.Close()

	encoder := json.NewEncoder(file)

	encodeErr := encoder.Encode(emptyObject)

	if encodeErr != nil {
		WriteMessage(WriteMessageProps{
			Type:    enums.MessageTypeError,
			Message: string(GetErrorMessage("CreateEmptyFile")),
		})

		os.Exit(1)
	}
}

func CreateEmptyFile(fileName string) {
	file, err := os.Create(fileName)

	if err != nil {
		WriteMessage(WriteMessageProps{
			Type:    enums.MessageTypeError,
			Message: string(GetErrorMessage("CreateEmptyFile")),
		})

		os.Exit(1)
	}

	defer file.Close()
}

func FilterAndTrim(files []string) []string {
	var result []string

	for _, file := range files {
		file = strings.TrimSpace(file)

		if file != "" {
			result = append(result, file)
		}
	}
	return result
}

func OpenURL(url string) {
	err := open.Run(url)

	if err != nil {
		WriteMessage(WriteMessageProps{
			Type:    enums.MessageTypeError,
			Message: string(GetErrorMessage("GitHubRepoOpen")),
		})

		os.Exit(1)
	}
}
