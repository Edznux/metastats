# Metastats 

Metastats is a *simple* monitoring tool for your life.
It is monitoring some fun statistic from your computer (Linux only) like the number of keystrokes, mouse click and uptime. (might expan later)

It saves its data inside simple CSV file. Thoose can be easily loaded from spreadsheets or any other stats tools.

It also provided non automatic monitoring for everything in your life.
For example you can follow your workout progress, weight gain/loss, books read, mountain climbed...

## Install

To run properly, Metastats needs root privileges : it must listen on all keystrokes and click events from /dev/inputs/

Metastats will create multiple file once you execute the `./install.sh` file :
- `/etc/metastats/config.toml` : The global config file.
- `/var/log/metastats/data/*` : All the events are stored here. By default, you will see `mice.csv` and `keyboard.csv`.
- `/var/log/metastats/metastats.log` : The log file.
- `/etc/systemd/system/metastats.service` : The Systemd service file.
- `/usr/local/bin/metastats` : The binary itself.

```
# install metastats
git clone https://github.com/Edznux/metastats
cd metastats
./install.sh
```

## Usage

Let it run as a daemon, it will save some stats inside the data folder (see installation).
You can add custom stats with `metastats add some-task someValue1 someValue2`. For exemple, monitor your pushups with : 
`metastats add pushups 25` will add a line with 25, at the current timestamp to the file `pushups.csv` in the data folder.

If you forgot to add something, or were AFK, you can "fake" the time with the `--at` flag.
You can provide pretty standards english sentences.

```
sudo metastats add books --at "yesterday 2pm"
```

Don't forget the quotes or it will break up in multiple arguments and only take the first word (and put the rest inside the `add` command)

## Goal

This projects is mainly developed to get better understanding about myself, usage of my computers, reading books counts, sports achievements etc.
And provide an easy to digest files to do some stats visualisation.

I (am/will be) using it for :
- Usage of my computer. [When booted on linux, so most of the time]
  - keystrokes 
  - clicks
  - uptime
  - network consumption
    - rx bytes
    - tx bytes
- Number of books read.
- Fitness.
- Maybe more, ¯\_(ツ)_/¯

## Requirement

This software will run out of the box on any *Linux* machine. It will need systemd for automatic startup.
I do not plan to extend this software to be compatible for windows or other OS (PR welcome tho).
