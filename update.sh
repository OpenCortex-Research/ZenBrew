#! /bin/sh
echo "    ____                   ______           __             "
echo "   / __ \____  ___  ____  / ____/___  _____/ /____  _  __  "
echo "  / / / / __ \/ _ \/ __ \/ /   / __ \/ ___/ __/ _ \| |/_/  "
echo " / /_/ / /_/ /  __/ / / / /___/ /_/ / /  / /_/  __/>  <    "
echo " \____/ .___/\___/_/ /_/\____/\____/_/   \__/\___/_/|_|    "
echo "     /_/                       ZenBrew Package Manager "
echo " "

echo "Do you want to update the ZenBrew Package Manager (y/n) "
read REPLY

if [ "$REPLY" = "y" ] || [ "$REPLY" = "Y" ]
then
    set -e # Exit immediately if a command exits with a non-zero status
    echo "[!] Updating ZenBrew..."

    # Create the directory if it doesn't exist
    if [ ! -d "/media/p4/OpenCortex/cache" ]; then
        mkdir /media/p4/OpenCortex/cache
    fi

    cd /media/p4/OpenCortex/cache
    echo "Update ZenBrew..."
    curl -L -s https://raw.githubusercontent.com/OpenCortex-Research/ZenBrew/main/ZenBrew.tar.gz -o ZenBrew.tar.gz
    
    echo "Extracting..."
    gunzip ZenBrew.tar.gz
    tar -xf ZenBrew.tar
    rm ZenBrew.tar

    echo "Copying Update"
    cp -R -p -P -f /media/p4/OpenCortex/cache/ZenBrew/ /media/p4/OpenCortex/
    rm -r /media/p4/OpenCortex/cache/

    echo "Creating links and setting permissions..."
    # Check if zenbrew is already installed
    if [ -f "/usr/bin/zenbrew" ]; then
        rm /usr/bin/zenbrew
    fi
    ln -s /media/p4/OpenCortex/ZenBrew/main.py /usr/bin/zenbrew
    chmod +x /media/p4/OpenCortex/ZenBrew/main.py
    chmod +x /usr/bin/zenbrew
    
    echo "ZenBrew Updated!"
    echo ""
    zenbrew --help
    exit 0
fi
exit 1