#!/bin/sh
echo "Installing ZenBrew..."
mkdir -p /media/p4/OpenCortex
mkdir -p /media/p4/OpenCortex/bin
mkdir -p /media/p4/OpenCortex/ZenBrew
cd /media/p4/OpenCortex/bin
curl -O https://github.com/OpenCortex-Research/ZenBrew/releases/download/V1.0.0/zenbrew
chmod +x /media/p4/OpenCortex/bin/zenbrew
cd /media/p4/OpenCortex/ZenBrew
curl -O https://github.com/OpenCortex-Research/ZenBrew/releases/download/V1.0.0/settings.json
cd /media/p4/OpenCortex/bin
./zenbrew install zenbrew
