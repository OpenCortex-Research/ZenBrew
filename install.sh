mkdir /opt/OpenCortex/
cd /opt/OpenCortex/
curl "https://github.com/JudahZF/ZenBrew/releases/download/V0.1.0/ZenBrew-0_1_0.tar.gz" -o ZenBrew-0_1_0.tar.gz
tar -xf ZenBrew-0_1_0.tar.gz
rm ZenBrew-0_1_0.tar.gz
ln -s ZenBrew/main.py /bin/ZenBrew