package main

import (
	_ "embed"
	"os"

	git "github.com/abroudoux/commit/internal/git"
	repository "github.com/abroudoux/commit/internal/repository"
	utils "github.com/abroudoux/commit/internal/utils"
)

func main() {
	err := utils.IsGitInstalled()
	if err != nil {
		utils.PrintErrorAndExit(err)
	}

	err = utils.IsInGitRepository()
	if err != nil {
		utils.PrintErrorAndExit(err)
	}

	if len(os.Args) > 1 {
		err := repository.FlagMode()
		if err != nil {
			println(err)
		}
		os.Exit(0)
	}

	err = git.AddAllFiles()
	if err != nil {
		utils.PrintErrorAndExit(err)
	}

	err = git.WriteCommitMessage()
	if err != nil {
		utils.PrintErrorAndExit(err)
	}

	err = git.CheckIfUpstreamBranchExists()
	if err == nil {
		err = git.PushCode()
		if err != nil {
			utils.PrintErrorAndExit(err)
		}
	} else {
		createUpstreamBranch, err := utils.AskUser("Upstream branch does not exist. Would you like to create it?")
		if err != nil {
			utils.PrintErrorAndExit(err)
		}

		if createUpstreamBranch {
			err := git.PushCode()
			if err != nil {
				utils.PrintErrorAndExit(err)
			}

			println("Upstream branch created successfully.")
			os.Exit(0)
		}

		println("Upstream branch not created. Exiting...")
	}

	os.Exit(0)
}
