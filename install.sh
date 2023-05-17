#! /bin/sh
mkdir /media/p4/OpenCortex/
cd /media/p4/OpenCortex/
curl -L -s https://github.com/OpenCortex-Research/ZenBrew/archive/refs/tags/V0.2.1.tar.gz -o ZenBrew.tar.gz
gunzip ZenBrew.tar.gz
tar -xf ZenBrew.tar
rm ZenBrew.tar
ln -s /media/p4/OpenCortex/ZenBrew/main.py /usr/bin/zenbrew
chmod +x /media/p4/OpenCortex/ZenBrew/main.py
chmod +x /usr/bin/zenbrew
