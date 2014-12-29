FROM ubuntu

# File Author / Maintainer
MAINTAINER Corey McClelland

RUN apt-get update -y

#install python
RUN apt-get install -y software-properties-common
RUN apt-get install -y python-software-properties

RUN add-apt-repository ppa:webupd8team/java
RUN gpg --keyserver pgp.mit.edu --recv-keys F758CE318D77295D
RUN gpg --export --armor F758CE318D77295D | apt-key add -
RUN gpg --keyserver pgp.mit.edu --recv-keys 2B5C1B00
RUN gpg --export --armor 2B5C1B00 | apt-key add -
RUN echo "deb http://www.apache.org/dist/cassandra/debian 20x main" | tee -a /etc/apt/sources.list.d/cassandra.list


# Update the repository
RUN apt-get update

# java
RUN echo oracle-java8-installer shared/accepted-oracle-license-v1-1 select true | /usr/bin/debconf-set-selections
RUN apt-get install -y oracle-java8-installer ca-certificates

# gradle
RUN apt-get install -y gradle
