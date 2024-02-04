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

var ping_cmd = &cobra.Command{
	Use:   "ping",
	Short: "ping test",
	Long:  `ping test`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("pong")
	},
}

var install_cmd = &cobra.Command{
	Use:   "install package_name",
	Short: "Install a package",
	Long:  `Install a package`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		utils.GetSettings()
		repos_links := utils.Preferences.Repos
		var repos []repo.Repo
		for _, repo_url := range repos_links {
			repo.DownloadRepoJson(repo_url)
			repos = append(repos, repo.DownloadRepoJson(repo_url))
		}
		var possible_packages []pkg.PackageLink
		for _, repo := range repos {
			if repo.CheckPackage(args[0]) {
				package_link_url := repo.URL + "/packages/" + args[0]
				package_link_file := utils.DownloadFile(package_link_url)
				var package_link pkg.PackageLink
				err := json.Unmarshal(package_link_file, &package_link)
				if err != nil {
					log.Error("Failed to unmarshal JSON:", err)
					panic("Failed to unmarshal JSON")
				}
				possible_packages = append(possible_packages, package_link)
			}
		}
		var selection int = 0
		if len(possible_packages) >= 0 {
			log.Info("Multiple repos contain the package")
			i := 0
			for _, temp_package := range possible_packages {
				log.Info(fmt.Sprintf("%d: %s - %s", i, temp_package.Name, temp_package.URL))
			}
			log.Info("Please select a package to install: ")
			fmt.Scanln(&selection)
		}
		log.Info("Installing package")
		selected_package := pkg.DownloadPackageMetadata(possible_packages[selection])
		selected_package.Download()
		selected_package.Install()
	},
}

var update_cmd = &cobra.Command{
	Use:   "update package_name",
	Short: "Update a package",
	Long:  `Update a package`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		utils.GetSettings()
		repos_links := utils.Preferences.Repos
		var repos []repo.Repo
		for _, repo_url := range repos_links {
			repo.DownloadRepoJson(repo_url)
			repos = append(repos, repo.DownloadRepoJson(repo_url))
		}
		var possible_packages []pkg.PackageLink
		for _, repo := range repos {
			if repo.CheckPackage(args[0]) {
				package_link_url := repo.URL + "/packages/" + args[0]
				package_link_file := utils.DownloadFile(package_link_url)
				var package_link pkg.PackageLink
				err := json.Unmarshal(package_link_file, &package_link)
				if err != nil {
					log.Error("Failed to unmarshal JSON:", err)
					panic("Failed to unmarshal JSON")
				}
				possible_packages = append(possible_packages, package_link)
			}
		}
		var selection int = 0
		if len(possible_packages) >= 0 {
			log.Info("Multiple repos contain the package")
			i := 0
			for _, temp_package := range possible_packages {
				log.Info(fmt.Sprintf("%d: %s - %s", i, temp_package.Name, temp_package.URL))
			}
			log.Info("Please select a package to install: ")
			fmt.Scanln(&selection)
		}
		log.Info("Updating package")
		selected_package := pkg.DownloadPackageMetadata(possible_packages[selection])
		selected_package.Download()
		selected_package.Update()
	},
}

var uninstall_cmd = &cobra.Command{
	Use:   "install package_name",
	Short: "Install a package",
	Long:  `Install a package`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		utils.GetSettings()
		repos_links := utils.Preferences.Repos
		var repos []repo.Repo
		for _, repo_url := range repos_links {
			repo.DownloadRepoJson(repo_url)
			repos = append(repos, repo.DownloadRepoJson(repo_url))
		}
		var possible_packages []pkg.PackageLink
		for _, repo := range repos {
			if repo.CheckPackage(args[0]) {
				package_link_url := repo.URL + "/packages/" + args[0]
				package_link_file := utils.DownloadFile(package_link_url)
				var package_link pkg.PackageLink
				err := json.Unmarshal(package_link_file, &package_link)
				if err != nil {
					log.Error("Failed to unmarshal JSON:", err)
					panic("Failed to unmarshal JSON")
				}
				possible_packages = append(possible_packages, package_link)
			}
		}
		var selection int = 0
		if len(possible_packages) >= 0 {
			log.Info("Multiple repos contain the package")
			i := 0
			for _, temp_package := range possible_packages {
				log.Info(fmt.Sprintf("%d: %s - %s", i, temp_package.Name, temp_package.URL))
			}
			log.Info("Please select a package to install: ")
			fmt.Scanln(&selection)
		}
		log.Info("Installing package")
		selected_package := pkg.DownloadPackageMetadata(possible_packages[selection])
		selected_package.Download()
		selected_package.Install()
	},
}

var patch_cmd = &cobra.Command{
	Use:   "patch",
	Short: "patch the system",
	Long:  `link the binaries to the system`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		utils.GetSettings()
		repos_links := utils.Preferences.Repos
		var repos []repo.Repo
		for _, repo_url := range repos_links {
			repo.DownloadRepoJson(repo_url)
			repos = append(repos, repo.DownloadRepoJson(repo_url))
		}
		var possible_packages []pkg.PackageLink
		for _, repo := range repos {
			if repo.CheckPackage(args[0]) {
				package_link_url := repo.URL + "/packages/" + args[0]
				package_link_file := utils.DownloadFile(package_link_url)
				var package_link pkg.PackageLink
				err := json.Unmarshal(package_link_file, &package_link)
				if err != nil {
					log.Error("Failed to unmarshal JSON:", err)
					panic("Failed to unmarshal JSON")
				}
				possible_packages = append(possible_packages, package_link)
			}
		}
		var selection int = 0
		if len(possible_packages) >= 0 {
			log.Info("Multiple repos contain the package")
			i := 0
			for _, temp_package := range possible_packages {
				log.Info(fmt.Sprintf("%d: %s - %s", i, temp_package.Name, temp_package.URL))
			}
			log.Info("Please select a package to install: ")
			fmt.Scanln(&selection)
		}
		log.Info("Installing package")
		selected_package := pkg.DownloadPackageMetadata(possible_packages[selection])
		selected_package.Download()
		selected_package.Install()
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ZenBrew.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	root_cmd.AddCommand(ping_cmd)
}
