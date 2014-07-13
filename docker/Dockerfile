# Set the base image to Ubuntu
FROM ubuntu:precise

# File Author / Maintainer
MAINTAINER Corey McClelland

# Add application repository URL to the default sources
RUN echo "deb http://archive.ubuntu.com/ubuntu/ raring main universe" >> /etc/apt/sources.list
RUN apt-get install -y software-properties-common python-software-properties
RUN add-apt-repository ppa:webupd8team/java
#RUN apt-key adv --keyserver pgp.mit.edu --recv-keys F758CE318D77295D
RUN gpg --keyserver pgp.mit.edu --recv-keys F758CE318D77295D
RUN gpg --export --armor F758CE318D77295D | apt-key add -
RUN gpg --keyserver pgp.mit.edu --recv-keys 2B5C1B00
RUN gpg --export --armor 2B5C1B00 | apt-key add -
RUN echo "deb http://www.apache.org/dist/cassandra/debian 20x main" | tee -a /etc/apt/sources.list.d/cassandra.list
RUN echo "deb http://archive.ubuntu.com/ubuntu precise main universe" >> /etc/apt/sources.list
RUN add-apt-repository ppa:cwchien/gradle
RUN add-apt-repository ppa:chris-lea/node.js

# Update the repository
RUN apt-get update

# Install necessary tools
RUN apt-get install -y nano wget dialog net-tools

#install python
RUN apt-get install -y python-software-properties

# java
RUN echo oracle-java8-installer shared/accepted-oracle-license-v1-1 select true | /usr/bin/debconf-set-selections
RUN apt-get install -y oracle-java8-installer ca-certificates

# gradle
RUN apt-get install -y gradle

#install go
RUN apt-get install -y golang

# download and install nodejs and git
RUN apt-get install -y nodejs
RUN apt-get install -y git git-core
#RUN npm config set registry http://registry.npmjs.org/
#RUN npm install -g npm
#RUN ln -s /usr/bin/nodejs /usr/bin/node
RUN npm install pm2 -g

# install cassandra
VOLUME ["/var/lib/cassandra"]
RUN apt-get -y --force-yes install cassandra

# install cql client

# begin adding files
ADD create.cql .
ADD drop.cql .
ADD cassandra-env.sh /etc/cassandra/cassandra-env.sh
#ADD data /var/lib/cassandra

# git
RUN mkdir -p /root/.ssh
ADD wordfinder_rsa /root/.ssh/id_rsa
RUN chmod 700 /root/.ssh/id_rsa
#RUN ssh-agent /bin/bash
#RUN ssh-add /root/.ssh/id_rsa
RUN echo "Host github.com\n\tStrictHostKeyChecking no\n" >> /root/.ssh/config
#RUN ssh -v git@github.com

RUN git clone git@github.com:sandalon/wordfinder.git /home/wordfinder
RUN git clone git@github.com:sandalon/starlight.git /home/starlight

# Expose ports
EXPOSE 3000
EXPOSE 9160
EXPOSE 9042

#ADD data /var/lib/cassandra
ADD cassandra.yaml /etc/cassandra/cassandra.yaml
ADD words2entries.txt .
RUN mv words2entries.txt /home/starlight/src/words2Entries.txt
ADD run.sh .
ENTRYPOINT ["./run.sh"]




# Set the default command to execute
# when creating a new container
#WORKDIR /home/starlight
#RUN gradle run


#WORKDIR /home/wordfinder
#RUN npm install

#pm2 start app.js -i max
#ENTRYPOINT ["pm2", "start", "server.js", "-i", "max", "--no-daemon"]
