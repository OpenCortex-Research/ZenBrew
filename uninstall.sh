#! /bin/sh
echo "    ____                   ______           __             "
echo "   / __ \____  ___  ____  / ____/___  _____/ /____  _  __  "
echo "  / / / / __ \/ _ \/ __ \/ /   / __ \/ ___/ __/ _ \| |/_/  "
echo " / /_/ / /_/ /  __/ / / / /___/ /_/ / /  / /_/  __/>  <    "
echo " \____/ .___/\___/_/ /_/\____/\____/_/   \__/\___/_/|_|    "
echo "     /_/                       ZenBrew Package Manager "
echo " "

echo "Do you want to uninstall the ZenBrew Package Manager (y/n) "
read REPLY

if [ "$REPLY" = "y" ] || [ "$REPLY" = "Y" ]
then
    set -e # Exit immediately if a command exits with a non-zero status
    echo "[!] Uninstalling ZenBrew..."
    rm -r /media/p4/OpenCortex/ZenBrew
    rm /usr/bin/zenbrew
    rm -r /media/p4/OpenCortex/cache/
    exit 0
fi
exit 1