package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type ImageClient interface {
	GetImagesURLFromDate(date time.Time, count int) ([]string, error)
}

type Nasa struct {
	url    string
	apiKey string
	client http.Client
}

func NewNasaImageClient(url, apiKey string) Nasa {
	return Nasa{
		url:    url,
		apiKey: apiKey,
		client: http.Client{},
	}
}

func (n Nasa) GetImagesURLFromDate(date time.Time, count int) ([]string, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf(n.url+"?earth_date=%s&camera=NAVCAM&api_key=%s", date.Format(time.RFC3339), n.apiKey), nil)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := n.client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	var result struct {
		Photos []struct {
			ImageSrc string `json:"img_src"`
		} `json:"photos"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return []string{}, err
	}

	res := make([]string, count)
	for index, photo := range result.Photos {
		if index == count {
			break
		}
		res[index] = photo.ImageSrc
	}

	return res, nil
}
