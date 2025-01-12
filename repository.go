package main

import (
	"fmt"
	"net/http"
)

func getLatestRelease() (string, error) {
	url := "https://api.github.com/repos/abroudoux/commit/releases/latest"
	res, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error while fetching latest release: %v", err)
	}

	latestVersion := res.Header.Get("tag_name")
	return latestVersion, nil
}