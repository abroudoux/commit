package repository

import (
	"fmt"
	"net/http"
	"os"
)

func FlagMode() error {
	flag := os.Args[1]

	if flag == "--version" || flag == "-v" {
		latestRealease, err := getLatestRelease()
		if err != nil {
			println("Latest version not avaible")
			os.Exit(1)
		}

		println(latestRealease)
	} else if flag == "--help" || flag == "-h" {
		PrintHelpManual()
	}

	return nil
}

func PrintHelpManual() {
	fmt.Println("Usage: commit [options]")
	fmt.Printf("  %-20s %s\n", "commit", "Commits all changes and pushes to the current branch")
	fmt.Printf("  %-20s %s\n", "commit [--help | -h]", "Show this help message")
}

func getLatestRelease() (string, error) {
	url := "https://api.github.com/repos/abroudoux/commit/releases/latest"
	res, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error while fetching latest release: %v", err)
	}

	latestVersion := res.Header.Get("tag_name")
	return latestVersion, nil
}