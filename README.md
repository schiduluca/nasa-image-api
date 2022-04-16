# nasa-image-api

## Build
```go build cmd/main.go ```


## Run
```
./main {lastNDays}

```

Example:
```./main 5```

`{lastNDays}` defaults to 10



## Things to be improved
- Add unit tests
- Use a database for the storage
- Make requests in paralel to fetch images from NASA API
- Add a config file with env variables for URL, API KEY etc.
