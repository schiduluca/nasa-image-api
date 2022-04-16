package cache

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ImageCache interface {
	Put(key string, urls []string) error
	Get(key string) ([]string, error)
}

type FileCache struct {
}

func (f FileCache) Put(key string, urls []string) error {
	jsonString, _ := json.Marshal(urls)
	err := ioutil.WriteFile(fmt.Sprintf("internal/cache/files/%s.json", key), jsonString, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func (f FileCache) Get(key string) ([]string, error) {
	file, err := os.OpenFile(fmt.Sprintf("internal/cache/files/%s.json", key), os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)
	var result []string
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
