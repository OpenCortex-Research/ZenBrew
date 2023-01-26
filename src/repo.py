import urllib, json

class Repo:
        def __init__(self, url):
                self.url = url
                self.downloaded = False
                file = urllib.urlopen(self.url + "repo.json")
                self.json = json.load(file)
                file.close()
                self.name = self.json["Name"]
                self.format = self.json["Format"]
                self.packageFile = self.json["Packages"]
        
        def downloadPackagesList(self):
                pass

class localRepo(Repo):
        def __init__(self, path):
                self.downloaded = True
                file = open(self.path + "repo.json")
                self.json = json.load(file)
                file.close()
                self.name = self.json["Name"]
                self.format = self.json["Format"]
                self.packageFile = self.json["Packages"]
        
        def downloadPackagesList(self):
                pass