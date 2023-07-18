#!/usr/bin/python2
import argparse, json, sys, logging, subprocess, os, time
from src.repo import Repo
from src.funcs import getFile
with open("/media/p4/OpenCortex/ZenBrew/settings.json") as jsonfile:
        settings = json.load(jsonfile)

parser = argparse.ArgumentParser(formatter_class=argparse.ArgumentDefaultsHelpFormatter,)

parser.add_argument("Operation", choices=["install", "update", "upgrade", "remove", "packages"])
parser.add_argument("-q", "--quiet", default=False, action="store_true", help="Reduce output text")
parser.add_argument("-r", "--repo", default=settings['defaultRepo'], action="store", help="Set a custom repo")
parser.add_argument("-d", "--detailed", action="store_true", help="Shows more detail during search or selection")
parser.add_argument("Package", nargs='?')

args = parser.parse_args()
operation = args.Operation
if args.Operation in ["install", "upgrade", "remove"] and args.Package is None:
    parser.error("<Package> required")
if args.Package: 
    package = args.Package.split('@')
    if len(package) == 1:
        package.append(False)

repo = Repo(args.repo)
"""
searchFuzz=repo.searchPackages(package)
def Search():
    if len(searchFuzz) == 0:
        print("No Packages Found")
        sys.exit()
    print("Found Packages: ")
    for i in range(len(searchFuzz)):
        print(" "+ str(i) + ". " + searchFuzz[i][1] + ":")
        if args.detailed: 
            print("     Author: " + repo.getPackageInfo(searchFuzz[i][1], "Author"))
            print("     Description: " + repo.getPackageInfo(searchFuzz[i][1], "Description"))
            print("     Location: " + repo.getPackageInfo(searchFuzz[i][1], "Package Location"))
        else: print("     " + repo.getPackageInfo(searchFuzz[i][1], "Description"))

if operation == "search":
    Search()"""

def update():
    print("Downloading Repos:")
    destination = settings["/media/p4/OpenCortex/ZenBrew/"] + "repo/"
    if os.path.exists(destination) is False:
        os.mkdir(destination)
    for repo in settings["repos"]:
        print("Downloading: " + repo["name"])
        file = open(destination + repo["name"] + ".json", 'w')
        file.write(json.loads(getFile(repo["url"])))
        file.close()
        print("Downloaded: " + repo["name"])
    print("Saving Update Time")
    timeFile = open(destination + "updateTime.txt", 'w')
    timeFile.write(str(int(time.time())))
    timeFile.close()
    print("Done!")

def install(package):
    isMatch = repo.packageMatch(package[0])
    if package[1]: print("Installing a custom version will overwrite any other installed versions.")
    if not isMatch:
        print("Package does not currently exist in repo")
        sys.exit()
        toInstall = input("Which package would you like to install: [0] ")
        if toInstall == "": toInstall = 0
        else: toInstall = int(toInstall)
        isMatch = repo.packageMatch(searchFuzz[toInstall][1])
    isMatch.install(package[1])

if operation == "update":
    isMatch = repo.packageMatch(package[0])
    if package[1]: print("Installing a custom version will overwrite any other installed versions.")
    if not isMatch:
        print("Package does not currently exist in repo")
        sys.exit()
        toInstall = input("Which package would you like to install: [0] ")
        if toInstall == "": toInstall = 0
        else: toInstall = int(toInstall)
        isMatch = repo.packageMatch(searchFuzz[toInstall][1])
    isMatch.update(package[1])

if operation == "uninstall":
    isMatch = repo.packageMatch(package[0])
    if not isMatch:
        print("Package does not currently exiist in repo")
        sys.exit()
        toInstall = input("Which package would you like to install: [0] ")
        if toInstall == "": toInstall = 0
        else: toInstall = int(toInstall)
        isMatch = repo.packageMatch(searchFuzz[toInstall][1])
    isMatch.uninstall()

if operation == "packages":
    for i in repo.packages:
        print(i.Identifier + ":")
        if args.detailed: 
            print("     Author: " + repo.getPackageInfo(i.Identifier, "Author"))
            print("     Description: " + repo.getPackageInfo(i.Identifier, "Description"))
            print("     Location: " + repo.getPackageInfo(i.Identifier, "Package Location"))
        else: print("     " + repo.getPackageInfo(i.Identifier, "Description"))


timeFile = open("/media/p4/OpenCortex/ZenBrew/updateTime.txt", 'r')
timeSinceUpdate = int(timeFile.read())
timeFile.close()
if int(time.time()) - 86400 > timeSinceUpdate:
    update()



if settings["clearCache"] == True: subprocess.call(["rm", "-r", settings["OpenCortexDir"] + "cache/"])