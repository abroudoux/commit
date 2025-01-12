package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func PrintErrorAndExit(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func IsGitInstalled() error {
	cmd := exec.Command("git", "version")
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func IsInGitRepository() error {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func AskUser(question string) (bool, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s (y/n) [yes]: ", question)
	input, err := reader.ReadString('\n')
	if err != nil {
		return false, fmt.Errorf("error reading user input: %v", err)
	}

	response := strings.TrimSpace(input)
	if response == "y" || response == "yes" || response == "" || response == "Y" || response == "YES" {
		return true, nil
	}

	return false, nil
}