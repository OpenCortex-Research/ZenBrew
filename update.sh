mkdir /opt/OpenCortex/cache
cd /opt/OpenCortex/cache
curl -L https://github.com/OpenCortex-Research/ZenBrew/releases/download/V0.1.3/ZenBrew-0_1_3.tar.gz -o ZenBrew-0_1_3.tar.gz
gunzip ZenBrew-0_1_3.tar.gz
tar -xf ZenBrew-0_1_3.tar
rm ZenBrew-0_1_3.tar
cp --recursive --preserve --update /opt/OpenCortex/cache/ZenBrew/ /opt/OpenCortex/
rm -r /opt/OpenCortex/cache/
rm /bin/zenbrew
ln -s /opt/OpenCortex/ZenBrew/main.py /bin/zenbrew
chmod +x /opt/OpenCortex/ZenBrew/main.py
chmod +x /bin/zenbrew