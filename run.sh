#!/bin/bash
OS=`uname`;
echo $OS


if [ "$1" == "-s" ]; then
   cd $HOME 
   git clone https://github.com/minhajuddinkhan/twenty20rule
   cd twenty20rule 
   chmod +x twenty20rule
   
   if [ $OS == "Linux" ]; then 
        echo "INSIDo!"
        sudo cp twenty20rule /usr/bin   
        sudo echo "/usr/bin/twenty20rule &" >> /etc/rc.local  
        twenty20rule &
   fi     
else
   go get github.com/minhajuddinkhan/twenty20rule 
   if [ $OS == "Linux" ]; then 
      twenty20rule &
      sudo echo "$GOPATH/bin/twenty20rule" >> /etc/rc.local    
   fi         
fi

