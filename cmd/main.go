package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/schiduluca/nasa-image-api/internal/cache"
	"github.com/schiduluca/nasa-image-api/internal/client"
	"github.com/schiduluca/nasa-image-api/internal/service"
)

const (
	NasaApiUrl = "https://api.nasa.gov/mars-photos/api/v1/rovers/curiosity/photos"
	NasaApiKey = "VaMUmd9BOyakdsdHQDGrlagkSHRjdhWseCd6oZOE"
)

func main() {
	daysArg := "10"
	if len(os.Args) == 2 {
		daysArg = os.Args[1]
	}
	defaultDays := 10
	daysInt, err := strconv.Atoi(daysArg)
	if err == nil {
		defaultDays = daysInt
	}

	cl := client.NewNasaImageClient(NasaApiUrl, NasaApiKey)
	imageCache := cache.FileCache{}

	imageService := service.NewImageDownloader(cl, imageCache)
	days, err := imageService.GetImagesFromLastNDays(defaultDays)
	if err != nil {
		return
	}
	bytes, err := json.Marshal(days)

	fmt.Println(string(bytes))
}
