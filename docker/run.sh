#!/bin/bash

cat /etc/cassandra/cassandra.yaml

cassandra
sleep 10s
cqlsh -f create.cql
cd /home/starlight
gradle run &
cd /home/wordfinder
npm install
pm2 start server.js -i max --no-daemon
