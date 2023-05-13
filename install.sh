mkdir /opt/OpenCortex/
cd /opt/OpenCortex/
curl "https://github.com/OpenCortex-Research/ZenBrew/releases/download/V0.1.0/ZenBrew-0_1_1.tar.gz" -o ZenBrew-0_1_1.tar.gz
gunzip ZenBrew-0_1_1.tar.gz
tar -xf ZenBrew-0_1_1.tar
rm ZenBrew-0_1_1.tar
ln -s ZenBrew/main.py /bin/ZenBrew