FROM sandalon/cassandra

# File Author / Maintainer
MAINTAINER Corey McClelland

RUN apt-get update -y

# add config files
ADD cassandra-env.sh /etc/cassandra/cassandra-env.sh
ADD cassandra.yaml /etc/cassandra/cassandra.yaml
#ADD run.sh /etc/cassandra/start.sh
#RUN chmod +x /etc/cassandra/start.sh
#RUN /etc/cassandra/start.sh

EXPOSE 9160
EXPOSE 9042
EXPOSE 7199

ENTRYPOINT ["cassandra","-f"]
#ENTRYPOINT ["/etc/cassandra/start.sh"]
