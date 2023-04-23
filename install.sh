sudo apt install wget python3 python3-pip -y -qq
pip3 install python-Levenshtein fuzzywuzzy fuzzywuzzy[speedup]
mkdir /opt/OpenCortex/
cd /opt/OpenCortex/
wget ""
unzip ZenBrew.zip
cd ZenBrew
sudo ln -s src/main.py /bin/ZenBrew