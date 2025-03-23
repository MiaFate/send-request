#!/bin/zsh

# This script is used to run the script
# It will run the script and pass the arguments to the script
# It will also log the output of the script to a file
# The script will be run in the background
# dir=$(pwd)
cd /Users/miafate/programacion/golang/src/github.com/miafate/goscripts/send-request
echo $dir


#hace un cat de finished.txt y lo guarda en una variable text
text=$(cat finished.txt)
#si text es igual a "" entonces ejecuta el script
if [ "$text" = "" ]; then
    # ejecuta el script
    ./sendreq # si no quedan productos por procesar guarda "exit" en el archivo finished.txt
elif [ "$text" = "exit" ]; then
    echo $text
    crontab -r # elimina el crontab
fi

#./sendreq
