
#!/usr/bin/env bash

# cores
normal=$'\e[0m'  
C=$(printf '\033')                                                 
green="${C}[1;32m"
yellow="${C}[1;33m"
RED="${C}[1;31m"
# fim das cores

instalacao(){
	echo "${green}[OK] ${normal}starting the installation."
	sleep 2
	echo "${yellow}[!!] ${normal}looking for the go path."

	if [ -x "$(command -v go)" ]; then
		echo "${green}[OK] ${normal}SUCCESS."
		go get -v "github.com/fatih/color"
		go build confusion.go
		go build check.go
		echo "${green}[OK] ${normal}to use this tool, type ${yellow}./confusion${normal} or ${yellow}./check"

	else
		echo "${RED}plz install golang lol"
	fi
}

instalacao
