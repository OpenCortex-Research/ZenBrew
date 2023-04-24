import argparse, json, sys, logging
from repo import Repo
with open("../settings.json") as jsonfile:
        settings = json.load(jsonfile)

parser = argparse.ArgumentParser(formatter_class=argparse.ArgumentDefaultsHelpFormatter,)

parser.add_argument("Operation", choices=["install", "update", "uninstall"])
parser.add_argument("-q", "--quiet", default=False, action="store_true", help="Reduce output text")
parser.add_argument("-r", "--repo", default=settings['defaultRepo'], action="store", help="Set a custom repo")
parser.add_argument("-d", "--detailed", action="store_true", help="Shows more detail during search or selection")
parser.add_argument("Package")

args = parser.parse_args()
operation = args.Operation
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

if operation == "install":
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
