#!/bin/bash

SCRIPT_PATH=$(readlink -f $(dirname $0))

printf "In\n\n\n\n\n\n\n" | openssl req -newkey rsa:2048 -nodes -keyout /etc/ssh/ssh_pos_key.pem -x509 -days 365 -out /etc/ssh/ssh_pos_cert.pem

#Permission to those files for grafana
sudo chmod 777 /etc/ssh/ssh_pos_cert.pem
sudo chmod 777 /etc/ssh/ssh_pos_key.pem