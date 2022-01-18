#!/usr/bin/env bash

if [[ -f "/clustern-node" ]]
then
    exit 0
fi

wget -O /usr/local/bin/pacapt https://github.com/icy/pacapt/raw/ng/pacapt
chmod 755 /usr/local/bin/pacapt
ln -sv /usr/local/bin/pacapt /usr/local/bin/pacman || true 
pacapt update
pacapt install git
pacapt install screen
wget https://golang.org/dl/go1.15.7.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.15.7.linux-amd64.tar.gz
rm -rf go1.15.7.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
mkdir app
git clone https://github.com/nicklvsa/clustern-node
cd clustern-node
screen go run main.go
