import os, json

# Looks through the repo, finds all json files and writes them to a CSV file to use as tempory DB for packages
Direc = "./repo"
packages = []
outputCSV = open("repo/packages.csv", 'w')
outputCSV.write("File, Name, Identifier, Type")
files = os.listdir(Direc)
files = [f for f in files if Direc+'/'+f.endswith(".json")]
for f in files:
        file = open(Direc+'/'+f, 'r')
        currentPackage = json.load(file)
        file.close()
        packages.append({
                "File": f, 
                "Name": currentPackage["Name"], 
                "Identifier": currentPackage["Identifier"], 
                "Type": currentPackage["Type"]
                })
        
for p in packages: outputCSV.write("\n" + p["File"] + ", " + p["Name"] + ", " + p["Identifier"] + ", " +p["Type"])
outputCSV.close()