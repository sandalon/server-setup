#!/bin/bash

cqlsh db 9160 -f /create.cql
cd /home/server-setup/cassandra/dataloader
go run main.go db wordfinder /entries.txt /words2entries.txt
