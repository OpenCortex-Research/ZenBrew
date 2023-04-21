from repo import Package
from settings import Settings
import os, sys, subprocess
def addToInstallList(package, location):
        installList = open("installQueue.txt", 'a')
        installList.append(package, location)
        installList.close()

def install(package: Package, settings=Settings()):
        package.download()
        print("downloaded")
        match package.FileType:
                case "py": exec(open("cache/" + package.Identifier + "." + package.FileType).read())
                case "sh":subprocess.call(["sh", "cache/" + package.Identifier + "." + package.FileType])
        if settings.clearCache == True: 
                subprocess.call(["rm", "-r", "cache/"])
                subprocess.call(["mkdir", "cache/"])

install(Package("https://zen.judahfuller.com/repo/", "test.json"))