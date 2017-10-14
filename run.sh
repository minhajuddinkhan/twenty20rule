#!/bin/bash
OS=`uname`;
chmod +x twenty20rule   
if [ $OS == "Linux" ]; then 
        sudo cp twenty20rule /usr/bin   
        sudo echo "/usr/bin/twenty20rule &" >> /etc/rc.local  
        twenty20rule &
fi     

if [$OS == "Darwin"]; then

    echo "do something for osx"

fi
