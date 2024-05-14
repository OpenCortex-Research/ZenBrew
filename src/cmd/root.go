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
			fmt.Printf("Already installed at version: %s", installedVer)
			fmt.Println("To upgrade the package, use the upgrade command")
			fmt.Println("To reinstall the package, use the reinstall command")
			os.Exit(0)
		}

		log.Info("Downloading repos")
		for _, repo_url := range repos_links {
			repo.DownloadRepoJson(repo_url)
			repos = append(repos, repo.DownloadRepoJson(repo_url))
		}

		log.Info("Checking if the package is available in the repos")
		var package_link pkg.PackageLink
		var usedRepo repo.Repo
		for _, repo := range repos {
			if repo.CheckPackage(pkg_to_install) {
				usedRepo = repo
				package_link_url := repo.URL + "/packages/" + pkg_to_install
				package_link_file := utils.DownloadFile(package_link_url)
				err := json.Unmarshal(package_link_file, &package_link)
				if err != nil {
					log.Error("Failed to unmarshal JSON:", err)
					panic("Failed to unmarshal JSON")
				}
				break
			}
		}
		log.Info("Installing package")
		selected_package := pkg.DownloadPackageMetadata(package_link)
		selected_package.Download(version)
		selected_package.Install()
		utils.AddInstalledPackage(pkg_to_install, selected_package.Latest, "installed", usedRepo.Name)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	fmt.Println(`    ____                   ______           __             `)
	fmt.Println(`   / __ \____  ___  ____  / ____/___  _____/ /____  _  __  `)
	fmt.Println(`  / / / / __ \/ _ \/ __ \/ /   / __ \/ ___/ __/ _ \| |/_/  `)
	fmt.Println(` / /_/ / /_/ /  __/ / / / /___/ /_/ / /  / /_/  __/>  <    `)
	fmt.Println(` \____/ .___/\___/_/ /_/\____/\____/_/   \__/\___/_/|_|    `)
	fmt.Println(`     /_/                       ZenBrew Package Manager `)
	fmt.Println(` `)
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
