import json
import os
import os.path
import subprocess
with open("/media/p4/OpenCortex/ZenBrew/settings.json") as jsonfile:
    settings = json.load(jsonfile)

def getFile(url, file = ""):
    if file == "":
        file = url.split("/")[-1]
        url = url.removesuffix(file)
    destination = settings["OpenCortexDir"] + "cache/"
    if os.path.exists(destination) is False:
        os.mkdir(destination)
    subprocess.call(["curl", "-s", "-L", "-o", os.path.join(destination, file), url + file])
    return open(os.path.join(destination, file), 'r').read()