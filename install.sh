#!/bin/bash

# systemd
sudo systemctl stop metastats # stop it in case it already exist, so we can update.
sudo cp metastats.service /etc/systemd/system/metastats.service
sudo systemctl enable metastats

# copy config file
sudo mkdir -p /etc/metastats/
sudo cp config.example.toml /etc/metastats/config.toml

# copy binary
sudo cp metastats /usr/local/bin/metastats

# Starts
sudo systemctl start metastats

# And check status
sudo systemctl status metastats

# Install bash completion
METASTATS_BASHRC="source <(metastats completion)"
if grep --quiet "$METASTATS_BASHRC" ~/.bashrc; then
    echo "Already in your ~/.bashrc"
else
    echo $METASTATS_BASHRC >> ~/.bashrc
fi

# TODO : add rclone saving.
# crontab -e
# @daily rclone sync /var/log/metastats/ metastats:/
