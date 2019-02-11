#!/bin/bash

# systemd
sudo systemctl stop metastats # stop it in case it already exist, so we can update.
sudo cp metastats.service /etc/systemd/system/metastats.service
sudo systemctl enable metastats

# copy config file
sudo mkdir -p /etc/metastats/
sudo cp config.toml /etc/metastats/

# copy binary
sudo cp metastats /usr/local/bin/metastats

# Starts
sudo systemctl start metastats

# And check status
sudo systemctl status metastats

# Install bash completion
echo "source <(metastats completion)" >> ~/.bashrc

# TODO : add rclone saving.
# crontab -e
# @daily rclone sync /var/log/metastats/ metastats:/
