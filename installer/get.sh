#!/bin/sh
echo "Installing ZenBrew..."
mkdir -p /media/p4/OpenCortex
mkdir -p /media/p4/OpenCortex/bin
mkdir -p /media/p4/OpenCortex/ZenBrew
cd /media/p4/OpenCortex/bin
curl -O https://github.com/OpenCortex-Research/ZenBrew/archive/
chmod +x /media/p4/OpenCortex/Bin/zenbrew
cd /media/p4/OpenCortex/ZenBrew
curl -O https://github.com/OpenCortex-Research/ZenBrew/archive/
cd /media/p4/OpenCortex/bin
zenbrew install zenbrew