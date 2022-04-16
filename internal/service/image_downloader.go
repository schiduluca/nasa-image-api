package service

import (
	"time"

	"github.com/schiduluca/nasa-image-api/internal/cache"
	"github.com/schiduluca/nasa-image-api/internal/client"
)

type ImageDownloader struct {
	imageClient client.ImageClient
	cache       cache.ImageCache
}

func NewImageDownloader(imageClient client.ImageClient, cache cache.ImageCache) ImageDownloader {
	return ImageDownloader{
		imageClient: imageClient,
		cache:       cache,
	}
}

func (d ImageDownloader) GetImagesFromLastNDays(days int) (map[string][]string, error) {
	result := make(map[string][]string)
	now := time.Now()
	for i := 0; i < days; i++ {
		dateString := now.Format("2006-01-02")
		get, err := d.cache.Get(dateString)
		result[dateString] = get
		if err != nil {
			// file is not found or corrupted
			date, innerErr := d.imageClient.GetImagesURLFromDate(now, 3)
			if innerErr != nil {
				return nil, innerErr
			}
			result[dateString] = date

			innerErr = d.cache.Put(dateString, date)
			if innerErr != nil {
				return nil, innerErr
			}
		}
		now = now.AddDate(0, 0, -1)
	}

	return result, nil
}
