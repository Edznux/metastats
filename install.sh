#!/bin/bash

# systemd
sudo cp metastats.service /etc/systemd/system/metastats.service
sudo systemctl enable metastats

# copy config file
sudo mkdir -p /etc/metastats/
sudo cp config.toml /etc/metastats/

# copy binary
sudo cp metastats /usr/local/bin/metastats
