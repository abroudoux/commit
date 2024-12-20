package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

//go:embed assets/ascii.txt
var asciiArt string

func main() {
	err := isGitInstalled()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = isInGitRepository()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(os.Args) > 1 {
		err := flagMode()
		if err != nil {
			println(err)
		}
		os.Exit(0)
	}

	err = addAllFiles()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = writeCommitMessage()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = checkIfUpstreamBranchExists()
	if err == nil {
		err = pushCode()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		createUpstreamBranch, err := askUser("Upstream branch does not exist. Would you like to create it?")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if createUpstreamBranch {
			err := pushCode()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			println("Upstream branch created successfully.")
			os.Exit(0)
		}

		println("Upstream branch not created. Exiting...")
	}

	os.Exit(0)
}

func isGitInstalled() error {
	cmd := exec.Command("git", "version")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("git is not installed: %v", err)
	}

	return nil
}

func isInGitRepository() error {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error checking if in git repository: %v", err)
	}

	return nil
}

func addAllFiles() error {
	cmd := exec.Command("git", "add", ".")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error adding all files: %v", err)
	}

	return nil
}

func writeCommitMessage() error {
	cmd := exec.Command("git", "commit", ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error committing code: %v", err)
	}

	return nil
}

func checkIfUpstreamBranchExists() error {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "--symbolic-full-name", "@{u}")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("upstream branch does not exists: %v", err)
	}

	return nil
}

func askUser(question string) (bool, error) {
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

func pushCode() error {
	cmd := exec.Command("git", "push", "-u", "origin", "HEAD")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error pushing code: %v", err)
	}

	return nil
}

func flagMode() error {
	flag := os.Args[1]

	if flag == "--add" || flag == "-a" {
		filesModified, err := getAllFiles()
		if err != nil {
			return err
		}

		for _, file := range filesModified {
			println(file)
		}
	}
	if flag == "--version" || flag == "-v" {
		fmt.Println(asciiArt)
		latestRealease, err := getLatestRelease()
		if err != nil {
			println("Latest version not avaible")
			os.Exit(1)
		}

		println(latestRealease)
	} else if flag == "--help" || flag == "-h" {
		printHelpManual()
	}

	return nil
}

func printHelpManual() {
	fmt.Println("Usage: commit [options]")
	fmt.Printf("  %-20s %s\n", "commit", "Commits all changes and pushes to the current branch")
	fmt.Printf("  %-20s %s\n", "commit [--help | -h]", "Show this help message")
}

type File struct {
	Name string
	Status string
}

const (
	Modified string = "modified"
	Deleted string = "deleted"
	Renamed string = "renamed"
	Created string = "created"
)

func getAllFiles() ([]File, error) {
	var allFiles []File = []File{}

	filesModified, err := getfilesModified()
	if err != nil {
		return nil, fmt.Errorf("error getting modified files: %v", err)
	}

	// println("Files Modified:")
	for _, file := range filesModified {
		fileToAdd := File{Name: file, Status: Modified}
		allFiles = append(allFiles, fileToAdd)
	}

	filesDeleted, err := getFilesDeleted()
	if err != nil {
		return nil, fmt.Errorf("error getting deleted files: %v", err)
	}

	// println("Files Deleted:")
	for _, file := range filesDeleted {
		fileToAdd := File{Name: file, Status: Deleted}
		allFiles = append(allFiles, fileToAdd)
	}

	filesRenamed, err := getFilesRenamed()
	if err != nil {
		return nil, fmt.Errorf("error getting renamed files: %v", err)
	}

	// println("Files Renamed:")
	for _, file := range filesRenamed {
		fileToAdd := File{Name: file, Status: Renamed}
		allFiles = append(allFiles, fileToAdd)
	}

	filesCreated, err := getFilesCreated()
	if err != nil {
		return nil, fmt.Errorf("error getting created files: %v", err)
	}

	// println("Files Created:")
	for _, file := range filesCreated {
		fileToAdd := File{Name: file, Status: Created}
		allFiles = append(allFiles, fileToAdd)
	}

	return nil, nil
}

func getfilesModified() ([]string, error) {
	cmd := exec.Command("git", "ls-files", "--modified")
	filesModified, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("error getting modified files: %v", err)
	}

	return strings.Split(string(filesModified), "\n"), nil
}

func getFilesDeleted() ([]string, error) {
	cmd := exec.Command("git", "ls-files", "--deleted")
	filesDeleted, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("error getting deleted files: %v", err)
	}

	return strings.Split(string(filesDeleted), "\n"), nil
}

func getFilesRenamed() ([]string, error) {
	cmd := exec.Command("git", "ls-files", "--renames")
	filesRenames, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("error getting renamed files: %v", err)
	}

	return strings.Split(string(filesRenames), "\n"), nil
}

func getFilesCreated() ([]string, error) {
	cmd := exec.Command("git", "ls-files", "--others")
	filesCreated, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("error getting created files: %v", err)
	}

	return strings.Split(string(filesCreated), "\n"), nil
}
