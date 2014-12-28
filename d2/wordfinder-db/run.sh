#!/bin/bash

IP='hostname --ip-address'

sed -i -e "s/# JVM_OPTS=\"$JVM_OPTS -Djava.rmi.server.hostname=<public name>\"/ JVM_OPTS=\"$JVM_OPTS -Djava.rmi.server.hostname=$IP\"/" /etc/cassandra/cassandra-env.sh



cassandra -f
