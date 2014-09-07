# Set the base image to Ubuntu
FROM ubuntu

# File Author / Maintainer
MAINTAINER Corey McClelland <corey@wordfinder.io>

RUN apt-get update
RUN apt-get install -y software-properties-common python-software-properties
RUN add-apt-repository ppa:webupd8team/java
RUN add-apt-repository ppa:cwchien/gradle
RUN apt-get update

# java
RUN echo oracle-java8-installer shared/accepted-oracle-license-v1-1 select true | /usr/bin/debconf-set-selections
RUN apt-get install -y oracle-java8-installer ca-certificates

# gradle
RUN apt-get install -y gradle

# git
RUN apt-get install -y git git-core
RUN mkdir -p /root/.ssh
ADD wordfinder_rsa /root/.ssh/id_rsa
RUN chmod 700 /root/.ssh/id_rsa
RUN echo "Host github.com\n\tStrictHostKeyChecking no\n" >> /root/.ssh/config

# clone the autocomplete service
RUN git clone git@github.com:sandalon/starlight.git /home/starlight

# add the lookup data
ADD words2entries.txt .
RUN mv words2entries.txt /home/starlight/src/words2Entries.txt

EXPOSE 8000

WORKDIR /home/starlight
CMD ["run"]
ENTRYPOINT ["gradle"]
