package shared

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Serializer interface {
	SerializeToFile(data interface{}, path string) error
	DeserializeFromFile(path string, out interface{}) error
}

var DefaultSerializer = &JsonSerializer{}

type JsonSerializer struct{}

func (ser *JsonSerializer) SerializeToFile(data interface{}, path string) error {
	bytes, err := json.Marshal(data)

	if err != nil {
		return err
	}

	destFile, _ := os.Create(path)
	destFile.Close()

	destFile, err = os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
	defer destFile.Close()

	destFile.Write(bytes)

	return nil
}

func (ser *JsonSerializer) DeserializeFromFile(path string, v interface{}) error {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	err = json.Unmarshal(data, v)

	if err != nil {
		return err
	}

	return nil
}
