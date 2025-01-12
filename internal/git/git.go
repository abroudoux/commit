package git

import (
	"os"
	"os/exec"
)

func AddAllFiles() error {
	cmd := exec.Command("git", "add", ".")
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func WriteCommitMessage() error {
	cmd := exec.Command("git", "commit", ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func PushCode() error {
	cmd := exec.Command("git", "push", "-u", "origin", "HEAD")
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}


func CheckIfUpstreamBranchExists() error {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "--symbolic-full-name", "@{u}")
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}