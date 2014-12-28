FROM ubuntu

# File Author / Maintainer
MAINTAINER Corey McClelland <corey@wnd.io>

RUN apt-get update

#install cqlsh
RUN apt-get install -y python-setuptools
RUN easy_install pip
RUN pip install cqlsh

#install git
RUN apt-get install -y git git-core

RUN apt-get install tar

#install mercurial
RUN apt-get install -y mercurial
RUN apt-get install -y wget

#install go
RUN apt-get install -y python-software-properties
RUN mkdir /go
ENV GOPATH /go
RUN wget -P /usr/local https://storage.googleapis.com/golang/go1.3.3.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf /usr/local/go1.3.3.linux-amd64.tar.gz
ENV PATH $PATH:/usr/local/go/bin
RUN go get github.com/gocql/gocql

# clone repo
RUN mkdir -p /root/.ssh
ADD wordfinder_rsa /root/.ssh/id_rsa
RUN chmod 700 /root/.ssh/id_rsa
RUN echo "Host github.com\n\tStrictHostKeyChecking no\n" >> /root/.ssh/config
RUN git clone git@github.com:sandalon/server-setup.git /home/server-setup

ADD create.cql .
ADD cqlsh .
RUN mv cqlsh /usr/local/bin

ADD entries.txt .
ADD entry2meta.txt .
ADD entry2title.txt .
ADD words2entries.txt .

ADD run.sh .
RUN chmod +x run.sh

ENTRYPOINT ["/run.sh"]
#ENTRYPOINT ["/bin/bash"]
