#! /bin/sh
echo "    ____                   ______           __             "
echo "   / __ \____  ___  ____  / ____/___  _____/ /____  _  __  "
echo "  / / / / __ \/ _ \/ __ \/ /   / __ \/ ___/ __/ _ \y| |/_/  "
echo " / /_/ / /_/ /  __/ / / / /___/ /_/ / /  / /_/  __/>  <    "
echo " \____/ .___/\___/_/ /_/\____/\____/_/   \__/\___/_/|_|    "
echo "     /_/                       ZenBrew Package Manager (Development)"
echo ""
echo "This installer is intented for use with the OpenCortex CorOS-dev-environment!"

echo "Do you want to install the ZenBrew Package Manager (y/n) "
read REPLY

if [ "$REPLY" = "y" ] || [ "$REPLY" = "Y" ]
then
    set -e # Exit immediately if a command exits with a non-zero status
    mkdir -p /media/p4/OpenCortex/ZenBrew/
    mount --bind ./ /media/p4/OpenCortex/ZenBrew/

    if [ -f "/usr/bin/zenbrew" ]; then
        rm /usr/bin/zenbrew
    fi
    
    ln -s /media/p4/OpenCortex/ZenBrew/main.py /usr/bin/zenbrew
    chmod +x /media/p4/OpenCortex/ZenBrew/main.py
    chmod +x /usr/bin/zenbrew
    echo "ZenBrew installed!"
    zenbrew --help
fi

