#!/bin/bash

echo "Installing Openresty............."

echo "Import GPG key"
wget -qO - https://openresty.org/package/pubkey.gpg | sudo apt-key add -

echo "Add Openresty repository"
sudo apt-get -y install software-properties-common

echo "Install Openresty"
sudo apt-get update
sudo apt-get install openresty -y
