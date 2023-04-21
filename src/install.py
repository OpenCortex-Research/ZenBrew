from repo import Package
import subprocess, json
with open("/OpenCortex/ZenBrew/settings.json") as jsonfile:
        settings = json.load(jsonfile)

def addToInstallList(package, location):
        installList = open("installQueue.txt", 'a')
        installList.append(package, location)
        installList.close()

def installPackahe(package: Package):
        package.download()
        print("Downloaded")
        packagesList = open(settings["zenBrewDir"] + "installedPackages.csv", 'a')
        packagesList.write(package.Identifier + "\n")
        packagesList.close()
        if package.Type == "DirectRun":
                match package.FileType:
                        case "py": exec(open(settings["zenBrewDir"] + "cache/" + package.Identifier + "." + package.FileType).read())
                        case "sh":subprocess.call(["sh", settings["zenBrewDir"] + "cache/" + package.Identifier + "." + package.FileType])
        elif package.Type == "executable":
                subprocess.call(["mv", settings["zenBrewDir"] + "cache/" + package.Identifier + "." + package.FileType, settings["zenBrewDir"] + "bin/"])
                subprocess.call(["chmod", "755", settings["zenBrewDir"] + "bin/" + package.Identifier + "." + package.FileType])
        if settings["clearCache"] == True: 
                subprocess.call(["rm", "-r", settings["zenBrewDir"] + "cache/"])
                subprocess.call(["mkdir", settings["zenBrewDir"] + "cache/"])
        print("Installed!")