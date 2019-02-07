# Metastats 

Metastats is a simple "monitoring" tool for your life.
It is monitoring some fun statistic from your computer (Linux only) like the number of keystrokes and click.

It saves its data inside simple CSV file. Thoose can be easily loaded from spreadsheets or any other stats tools.

## Usage

To run properly, Metastats needs root privileges : it must listen on all keystrokes and click events from /dev/inputs/

Metastats will create multiple file once you execute the `./install.sh` file :
- `/etc/metastats/config.toml` : The global config file.
- `/var/log/metastats/data/*` : All the events are stored here. By default, you will see `mice.dat` and `keyboard.dat`.
- `/var/log/metastats/metastats.log` : The log file.
- `/etc/systemd/system/metastats.service` : The Systemd service file.
- `/usr/local/bin/metastats` : The binary itself.

```
# install metastats
git clone https://github.com/Edznux/metastats
cd metastats
./install.sh

# check if everything is ok
sudo systemctl status metastats
```
