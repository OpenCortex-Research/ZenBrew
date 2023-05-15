#! /bin/sh
mkdir /opt/OpenCortex/
cd /opt/OpenCortex/
curl -L -s https://github.com/OpenCortex-Research/ZenBrew/releases/download/V0.1.3/ZenBrew.tar.gz -o ZenBrew.tar.gz
gunzip ZenBrew.tar.gz
tar -xf ZenBrew.tar
rm ZenBrew.tar
ln -s /opt/OpenCortex/ZenBrew/main.py /usr/bin/zenbrew
chmod +x /opt/OpenCortex/ZenBrew/main.py
chmod +x /usr/bin/zenbrew
