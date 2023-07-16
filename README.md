# ZenBrew

Homebrew Package Manager for the Quad Cortex

Install using:

```bash
curl -s -L https://github.com/OpenCortex-Research/ZenBrew/raw/main/install.sh | sh
```

Update using:

```bash
curl -s -L https://github.com/OpenCortex-Research/ZenBrew/raw/main/update.sh | sh
```

Uninstall using:

```bash
curl -s -L https://github.com/OpenCortex-Research/ZenBrew/raw/main/uninstall.sh | sh
```

Or use ZenBrew to update itself

## Available commands

ZenBrew [install, update, uninstall] package_name - runs install option on a package
ZenBrew packages - lists all packages

## Maintaining ZenBrew

To develop and test, you can use the [CorOS emulation environment Docker container](https://github.com/VanIseghemThomas/OpenCortex/tree/main/CorOS-dev-environment). This has out of the box support for getting started with developing for ZenBrew. Anything inside the `mount` folder, will be available in the emulated filesystem under `/mnt/host`. Cloning the repo in this folder, you can use the `dev_install` script to install ZenBrew as a developer. Any changes to ZenBrew will be updated live.

Since the QC relies on Python 2.7, we need to develop for that target. This is quite the challenge since a lot of support for Python 2 has dropped over time.

### VSCode

For development in VSCode, download the **v2021.9.1246542782** Python extension. This version has support for Python 2 interpreters, codelens,...

### File Structure

#### ZenBrew files

ZenBrew files are installed to
/media/p4/OpenCortex/ZenBrew/

The main ZB files are:

<!-- prettier-ignore -->
- repo - Repository cache files
  - Contains each repos .json file and a file with the timestamp of the last repo update
- packageCache - Package cache files. This folder is regularly cleared during package install/updates/uninstall
- settings.json - Settings file
- installedPackages.json - Installed packages and there current versions

#### Package files

Package files are installed to  
/media/p4/OpenCortex/{package name}/

Each package folder should contain:

<!-- prettier-ignore -->
- install.sh
- update.sh
- uninstall.sh
- {package name}.json
