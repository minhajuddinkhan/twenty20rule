#!/bin/bash

APP=twenty20rule

withoutGo() {
   exec git clone https://github.com/minhajuddinkhan/twenty20rule
   cd twenty20rule && sudo chmod +x twenty20rule && sudo cp twenty20rule /usr/local/bin   
   echo /usr/bin/$APP > /etc/rc.local & 
}

withGo(){
    echo "with go."
}


echo $1

if [ "$1" == "-s" ]; then
    withoutGo
else
    withGo    
fi

