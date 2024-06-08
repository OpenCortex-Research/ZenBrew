/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	log "log/slog"
	"os"

	"OpenCortex/ZenBrew/pkg"
	"OpenCortex/ZenBrew/repo"
	"OpenCortex/ZenBrew/utils"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var root_cmd = &cobra.Command{
	Use:   "ZenBrew",
	Short: "Homebrew Package Manager for the Quad Cortex",
	Long:  `Homebrew Package Manager for the Quad Cortex`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
}

var install_cmd = &cobra.Command{
	Use:   "install package_name",
	Short: "Install a package",
	Long:  `Install a package`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		// Setup install
		pkg_to_install := args[0]
		version := ""

		log.Info("Getting the current state of ZenBrew")
		utils.GetSettings()
		repos_links := utils.Preferences.Repos
		var repos []repo.Repo

		// Check if the package is already installed
		installed, installedVer := utils.CheckIfPackageInstalled(pkg_to_install)
		if installed {
			log.Info("Package already installed")
			log.Info(fmt.Sprintf("Already installed at version: %s", installedVer))
			log.Info("To upgrade the package, use the upgrade command")
			log.Info("To reinstall the package, use the reinstall command")
			os.Exit(0)
		}

		log.Info("Downloading repos")
		for _, repo_url := range repos_links {
			repos = append(repos, repo.DownloadRepoJson(repo_url))
		}

		log.Info("Checking if the package is available in the repos")
		var selected_package pkg.Package
		var usedRepo repo.Repo
		for _, repo := range repos {
			package_link_url, err := repo.CheckPackage(pkg_to_install)
			if err != nil {
				log.Error("Failed to check package in repo:", err)
			}
			usedRepo = repo
			package_link_file := utils.DownloadFile(package_link_url)
			json_err := json.Unmarshal(package_link_file, &selected_package)
			if json_err != nil {
				log.Error("Failed to unmarshal JSON:", json_err)
				panic("Failed to unmarshal JSON")
			}
			break
		}
		log.Info("Installing package")
		selected_package.Download(version)
		selected_package.Install()
		utils.AddInstalledPackage(pkg_to_install, selected_package.Latest, "installed", usedRepo.Name)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	fmt.Println(`  ______          ____                    `)
	fmt.Println(` |___  /         |  _ \                   `)
	fmt.Println(`    / / ___ _ __ | |_) |_ __ _____      __`)
	fmt.Println(`   / / / _ \ '_ \|  _ <| '__/ _ \ \ /\ / /`)
	fmt.Println(`  / /_|  __/ | | | |_) | | |  __/\ V  V / `)
	fmt.Println(` /_____\___|_| |_|____/|_|  \___| \_/\_/  `)
	fmt.Println(`               QuadCortex Package Manager `)
	fmt.Println(`                          From OpenCortex `)
	utils.Lock()
	err := root_cmd.Execute()
	utils.Unlock()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	root_cmd.AddCommand(install_cmd)
}
