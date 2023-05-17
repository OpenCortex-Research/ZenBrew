#! /bin/sh
mkdir /media/p4/OpenCortex/cache
cd /media/p4/OpenCortex/cache
curl -L https://github.com/OpenCortex-Research/ZenBrew/archive/refs/tags/V0.2.0.tar.gz -o ZenBrew.tar.gz
gunzip ZenBrew.tar.gz
tar -xf ZenBrew.tar
rm ZenBrew.tar
cp -R -p -P -f /media/p4/OpenCortex/cache/ZenBrew/ /media/p4/OpenCortex/
rm -r /media/p4/OpenCortex/cache/
rm /usr/bin/zenbrew
ln -s /media/p4/OpenCortex/ZenBrew/main.py /usr/bin/zenbrew
chmod +x /media/p4/OpenCortex/ZenBrew/main.py
chmod +x /usr/bin/zenbrew
