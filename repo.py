import json, os, os.path, urllib, subprocess, logging
with open("/media/p4/OpenCortex/ZenBrew/settings.json") as jsonfile:
        settings = json.load(jsonfile)

def fuzzSort(fuzzList):
        return fuzzList[0]

def getFile(url, file):
        destination = settings["OpenCortexDir"] +"cache/"
        if os.path.exists(destination) is False:
                os.mkdir(destination)
        urllib.urlretrieve(url + file, os.path.join(destination, file))
        json = open(destination + file, 'r')
        output = str(json.read())
        json.close()
        return output

def updateInstallLog(id, version):
        with open(settings["zenBrewDir"] + "installedPackages.json", 'r') as file:
                installedLog = json.load(file)
        if version == -1: installedLog.pop(id)
        else: installedLog[id] = version
        with open(settings["zenBrewDir"] + "installedPackages.json", 'w') as file:
                file.write(json.dumps(installedLog))

def isInstalled(id):
        with open(settings["zenBrewDir"] + "installedPackages.json", 'r') as file:
                installedLog = json.load(file)
        if id in installedLog: return True
        else: return False

class Repo:
        def __init__(self, url):
                self.url = url
                file = getFile(self.url, "repo.json")
                self.json = json.loads(str(file))
                self.name = self.json["Name"]
                self.format = self.json["Format"]
                self.packageFile = self.json["Packages"]
                self.packages = []
                packageCSV = str(getFile(self.url, self.packageFile))
                packageCSV = packageCSV.split('\n')
                packageCSV.pop(0)
                for i in range(len(packageCSV)):
                        if packageCSV[i] == '': break
                        else:
                                packageCSV[i] = packageCSV[i].split(', ')
                                self.packages.append(Package(self.url, packageCSV[i][0]))

        """def searchPackages(self, searcTerm):
                searchFuzz = []
                for i in self.packages:
                        if fuzz.ratio(searcTerm.lower(), i.Identifier.lower()) >= 0.8:
                                searchFuzz.append([fuzz.ratio(searcTerm, i.Identifier), i.Identifier])
                searchFuzz.sort(key=fuzzSort, reverse=True)
                return (searchFuzz[0:10])"""

        def getPackageInfo(self, package, info):
                for i in self.packages:
                        if i.Identifier == package:
                                return i.json[info]
        
        def packageMatch(self, package):
                for i in self.packages:
                        if i.Identifier == package:
                                return i
                return False

class Package:
        def __init__(self, repoURL, jsonFile):
                self.url = repoURL
                self.downloaded = False
                text = getFile(repoURL, jsonFile)
                self.json = json.loads(text)
                self.Identifier = self.json["Identifier"]
                self.FileType = self.json["FileType"]
                self.versions = self.json["versions"]
                self.newestVer = len(self.versions)-1
                self.Location = settings["OpenCortexDir"] + self.Identifier + "/"

        def download(self, version=False):
                if version == False: version = self.newestVer
                subprocess.call(["curl", "-L", "-s", (self.versions[version]["Location"]+ self.versions[version]["FileName"]), "-o", (settings["OpenCortexDir"] + "cache/" + self.versions[version]["FileName"])])
                subprocess.call(["cp", "-R", "-p", "-P", "-f", settings["OpenCortexDir"] + "cache/" + self.versions[version]["FileName"], settings["OpenCortexDir"]])
                subprocess.call(["gunzip", settings["OpenCortexDir"] + self.versions[version]["FileName"]])
                subprocess.call(["tar", "-xf", settings["OpenCortexDir"] + self.versions[version]["FileName"][:-3], "-C", settings["OpenCortexDir"]])
                #subprocess.call(["cp", "--recursive", "--preserve", "--update", settings["OpenCortexDir"] + self.Identifier, settings["OpenCortexDir"]])
                subprocess.call(["rm", settings["OpenCortexDir"] + self.versions[version]["FileName"][:-3]])
                #subprocess.call(["rm", "-r", settings["zenBrewDir"] + self.Identifier])
                return True
        
        def install(self, version=False):
                if version == False: version = self.newestVer
                if self.download(version):
                        subprocess.call(["bash", self.Location + "install.sh"])
                        updateInstallLog(self.Identifier, self.versions[version]["id"])
                else: print("Error")
        
        def update(self, version=False):
                if isInstalled(self.Identifier):
                        if version == False: version = self.newestVer
                        if self.download(version):
                                subprocess.call(["bash", self.Location + "update.sh"])
                                updateInstallLog(self.Identifier, self.versions[version]["id"])
                        else: print("Error")
                else: print("Package Not Installed")
        
        def uninstall(self):
                subprocess.call(["bash", self.Location + "uninstall.sh"])
                subprocess.call(["rm", "-r", self.Location])
                updateInstallLog(self.Identifier, -1)
