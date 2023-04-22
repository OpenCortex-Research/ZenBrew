import json, requests
from fuzzywuzzy import fuzz
with open("/opt/OpenCortex/ZenBrew/settings.json") as jsonfile:
        settings = json.load(jsonfile)

def fuzzSort(fuzzList):
        return fuzzList[0]

class Repo:
        def __init__(self, url):
                self.url = url
                file = requests.get(self.url + "repo.json")
                self.json = json.loads(str(file.text))
                file.close()
                self.name = self.json["Name"]
                self.format = self.json["Format"]
                self.packageFile = self.json["Packages"]
                self.packages = []
                packageCSV = str(requests.get(self.url + self.packageFile).text)
                packageCSV = packageCSV.split('\n')
                packageCSV.pop(0)
                for i in range(len(packageCSV)):
                        if packageCSV[i] == '': break
                        else:
                                packageCSV[i] = packageCSV[i].split(', ')
                                self.packages.append(Package(self.url, packageCSV[i][0]))

        def searchPackages(self, searcTerm):
                searchFuzz = []
                for i in self.packages:
                        if fuzz.ratio(searcTerm.lower(), i.Identifier.lower()) >= 0.8:
                                searchFuzz.append([fuzz.ratio(searcTerm, i.Identifier), i.Identifier])
                searchFuzz.sort(key=fuzzSort, reverse=True)
                return (searchFuzz[0:10])

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
                self.url = repoURL + jsonFile
                self.downloaded = False
                file = requests.get(self.url)
                text = str(file.text)
                self.json = json.loads(text)
                file.close()
                self.Identifier = self.json["Identifier"]
                self.Script = self.json["Script"]
                self.FileType = self.json["FileType"]
                self.PackageLocation = self.json["Package Location"]
                self.Type = self.json["Type"]

        def download(self):
                file = requests.get(self.PackageLocation)
                save = open(settings["zenBrewDir"] + "cache/" + self.Identifier + "." + self.FileType, "wb")
                save.write(file.content)
                file.close()
                save.close()

testRepo = Repo("https://zen.judahfuller.com/repo/")
testRepo.searchPackages("Test py")