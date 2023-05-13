mkdir /opt/OpenCortex/cache
cd /opt/OpenCortex/cache
curl -L https://github.com/OpenCortex-Research/ZenBrew/releases/download/V0.1.3/ZenBrew.tar.gz -o ZenBrew.tar.gz
gunzip ZenBrew.tar.gz
tar -xf ZenBrew.tar
rm ZenBrew.tar
cp --recursive --preserve --update /opt/OpenCortex/cache/ZenBrew/ /opt/OpenCortex/
rm -r /opt/OpenCortex/cache/
rm /bin/zenbrew
ln -s /opt/OpenCortex/ZenBrew/main.py /bin/zenbrew
chmod +x /opt/OpenCortex/ZenBrew/main.py
chmod +x /bin/zenbrew