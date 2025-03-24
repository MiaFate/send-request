#!/bin/zsh
dir=$(pwd)
#agregar la variable dir a la segunda linea del archivo finished.txt con sed
go build -o sendreq main.go
chmod +x scriptrunner.sh
#write out current crontab
crontab -l > mycron
#echo new cron into cron file
echo "*/1 * * * * $dir/scriptrunner.sh" >> mycron
#install new cron file
crontab mycron
rm mycron