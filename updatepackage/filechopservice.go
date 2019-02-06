package updatepackage

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"os"

	"nova/shared"
)

var Megabyte int64 = 1 * (1 << 20)

type FileChopService interface {
	ChopFile(file string, destDir string, chunkSize int64) ([]string, error)
	JoinFile(files []string, destFile string) error
}

type StraightFileChopService struct{}

func NewStraightFileChopService() *StraightFileChopService {
	return &StraightFileChopService{}
}

// https://socketloop.com/tutorials/golang-how-to-split-or-chunking-a-file-to-smaller-pieces
func (service *StraightFileChopService) ChopFile(filename string, destDir string, chunkSize int64) ([]string, error) {
	fileToBeChunked := filename

	file, err := os.Open(fileToBeChunked)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	fileInfo, _ := file.Stat()

	var fileSize int64 = fileInfo.Size()

	const fileChunk = 1 * (1 << 20) // 1 MB, change this to your requirement

	// calculate total number of parts the file will be chunked into

	totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

	result := []string{}

	for i := uint64(0); i < totalPartsNum; i++ {

		partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
		partBuffer := make([]byte, partSize)

		file.Read(partBuffer)

		// write to disk
		fileName := fmt.Sprintf("%v/%v.dat", destDir, i)
		_, err := os.Create(fileName)

		if err != nil {
			return nil, err
		}

		// write/save buffer to disk
		ioutil.WriteFile(fileName, partBuffer, os.ModeAppend)
		result = append(result, fileName)
	}

	return result, nil
}

func (service *StraightFileChopService) JoinFile(sourceFiles []string, destination string) error {

	if shared.FileOrFolderExists(destination) {
		//Try delete file
		if err := os.Remove(destination); err != nil {
			return err
		}
	}

	destFile, _ := os.Create(destination)
	destFile.Close()

	destFile, err := os.OpenFile(destination, os.O_APPEND|os.O_WRONLY, 0600)
	defer destFile.Close()

	if err != nil {
		return errors.New("Could not create destination file")
	}

	for _, file := range sourceFiles {
		data, err := ioutil.ReadFile(file)

		_, err = destFile.Write(data)

		if err != nil {
			return err
		}
	}

	return nil
}
