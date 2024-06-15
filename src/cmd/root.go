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

var debug bool

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
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if debug {
			log.SetLogLoggerLevel(log.LevelDebug)
		}

		// Setup install
		pkg_to_install := args[0]
		version := ""

		log.Info("Getting the current state of ZenBrew")
		utils.GetSettings("/media/p4/OpenCortex/ZenBrew/")
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
			} else {
				usedRepo = repo
				package_link_file := utils.DownloadFile(package_link_url)
				json_err := json.Unmarshal(package_link_file, &selected_package)
				if json_err != nil {
					log.Error("Failed to unmarshal JSON:", json_err)
					panic("Failed to unmarshal JSON")
				}
				break
			}
		}
		log.Info("Installing package")
		ver_index := selected_package.Download(version)
		selected_package.Install()
		var installed_package utils.Package
		installed_package.Name = selected_package.Name
		installed_package.Format = selected_package.Format
		installed_package.Maintainer = selected_package.Maintainer
		installed_package.Versions = selected_package.Versions
		installed_package.Latest = selected_package.Latest
		utils.AddInstalledPackage(installed_package, "installed", usedRepo.Name, ver_index)
	},
}

var uninstall_cmd = &cobra.Command{
	Use:   "uninstall package_name",
	Short: "Uninstall a package",
	Long:  `Uninstall a package`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if debug {
			log.SetLogLoggerLevel(log.LevelDebug)
		}
		// Setup install
		pkg_to_uninstall := args[0]

		log.Info("Getting the current state of ZenBrew")
		utils.GetSettings("/media/p4/OpenCortex/ZenBrew/")

		// Check if the package is already installed
		installed, _ := utils.CheckIfPackageInstalled(pkg_to_uninstall)
		if !installed {
			log.Info("Package is not already installed")
			log.Info("To install the package, use the install command")
			os.Exit(0)
		}

		installed_packages := utils.GetInstalledPackages()
		var selected_package pkg.Package
		var uninstall_package pkg.InstalledPackage
		for _, installed_package := range installed_packages {
			if installed_package.Name == pkg_to_uninstall {
				uninstall_package.Name = installed_package.Name
				uninstall_package.Format = installed_package.Format
				uninstall_package.Maintainer = installed_package.Maintainer
				uninstall_package.Version = installed_package.Version
				uninstall_package.Status = installed_package.Status
				uninstall_package.Repository = installed_package.Repository
				selected_package = pkg.FromInstalled(uninstall_package)
			}
		}

		log.Info("Uninstalling package")
		selected_package.Uninstall()
		utils.RemoveInstalledPackage(uninstall_package.Name, uninstall_package.Repository)
	},
}

var update_cmd = &cobra.Command{
	Use:   "update package_name",
	Short: "Update a package",
	Long:  `Update a package`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if debug {
			log.SetLogLoggerLevel(log.LevelDebug)
		}
		// Setup update
		pkg_to_update := args[0]
		version := ""

		log.Info("Getting the current state of ZenBrew")
		utils.GetSettings("/media/p4/OpenCortex/ZenBrew/")
		repos_links := utils.Preferences.Repos
		var repos []repo.Repo

		// Check if the package is already installed
		installed, installedVer := utils.CheckIfPackageInstalled(pkg_to_update)
		if !installed {
			log.Info("Package is not already installed")
			log.Info("To install the package, use the install command")
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
			package_link_url, err := repo.CheckPackage(pkg_to_update)
			if err != nil {
				log.Error("Failed to check package in repo:", err)
			} else {
				usedRepo = repo
				package_link_file := utils.DownloadFile(package_link_url)
				json_err := json.Unmarshal(package_link_file, &selected_package)
				if json_err != nil {
					log.Error("Failed to unmarshal JSON:", json_err)
					panic("Failed to unmarshal JSON")
				}
				break
			}
		}

		if selected_package.Latest == installedVer {
			log.Info("Package is already up to date")
			os.Exit(0)
		}

		log.Info("Updating package")
		ver_index := selected_package.Download(version)
		selected_package.Install()
		var update_package utils.Package
		update_package.Name = selected_package.Name
		update_package.Format = selected_package.Format
		update_package.Maintainer = selected_package.Maintainer
		update_package.Versions = selected_package.Versions
		update_package.Latest = selected_package.Latest
		utils.SetPackageStatus(update_package, "installed", ver_index, usedRepo.Name)
	},
}

var reinstall_cmd = &cobra.Command{
	Use:   "reinstall package_name",
	Short: "Reinstall a package",
	Long:  `Reinstall a package`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if debug {
			log.SetLogLoggerLevel(log.LevelDebug)
		}
		log.Error("Not implemented.")
		log.Error("Please use the uninstall command to uninstall the package, then the install command to reinstall it.")
		os.Exit(1)
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
	root_cmd.AddCommand(uninstall_cmd)
	root_cmd.AddCommand(reinstall_cmd)
	root_cmd.AddCommand(update_cmd)
	root_cmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Enable debug mode")
}
