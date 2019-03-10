#!/bin/bash

if [ "$1" = "" ]; then
	echo "In order to continue, you must specify either latest or the release version."
	exit 1
elif [ "$1" = "latest" ]; then
	deb=$(/usr/bin/curl -s https://api.github.com/repos/EOSIO/eosio.cdt/releases/latest | grep "browser_download_url.*rpm" | grep "centos" | cut -d '"' -f 4)
	filename=$(/usr/bin/curl -s https://api.github.com/repos/EOSIO/eosio.cdt/releases/latest | grep "name.*rpm" | cut -d '"' -f 4)
else
	deb=$(/usr/bin/curl -s https://api.github.com/repos/EOSIO/eosio.cdt/releases/tags/$1 | grep "browser_download_url.*rpm" | cut -d '"' -f 4)
	filename=$(/usr/bin/curl -s https://api.github.com/repos/EOSIO/eosio.cdt/releases/tags/$1 | grep "name.*rpm" | cut -d '"' -f 4)
fi

echo "package:" $deb
if [ "$deb" = "" ]; then
	echo "Either $1 is not a valid release, or there is not a published .rpm package for the release."
	exit 1
fi

/usr/bin/wget $deb && /usr/bin/rpm -ihv "$filename" && rm -f "$filename"
