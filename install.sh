mkdir /opt/OpenCortex/
cd /opt/OpenCortex/
curl "https://github.com/JudahZF/ZenBrew/releases/download/V0.1.0/ZenBrew.tar.gz" -o ZenBrew.tar.gz
tar -xf ZenBrew.tar.gz
rm ZenBrew.tar.gz
ln -s ZenBrew/main.py /bin/ZenBrew
