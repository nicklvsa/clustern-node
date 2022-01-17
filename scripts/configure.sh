#!/usr/bin/env bash

wget -O /usr/local/bin/pacapt https://github.com/icy/pacapt/raw/ng/pacapt
chmod 755 /usr/local/bin/pacapt
ln -sv /usr/local/bin/pacapt /usr/local/bin/pacman || true 
pacapt update
pacapt install git
wget https://golang.org/dl/go1.15.7.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.15.7.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
mkdir clustern 
cd clustern
echo "hello world"
