FROM ubuntu

# File Author / Maintainer
MAINTAINER Corey McClelland <corey@wnd.io>

RUN apt-get update

#install git
RUN apt-get install -y git git-core

#install mercurial
RUN apt-get install -y mercurial

#install go
RUN apt-get install -y python-software-properties
ENV GOPATH /go
RUN apt-get install -y golang
RUN go get github.com/gocql/gocql

# clone repo
RUN mkdir -p /root/.ssh
ADD wordfinder_rsa /root/.ssh/id_rsa
RUN chmod 700 /root/.ssh/id_rsa
RUN echo "Host github.com\n\tStrictHostKeyChecking no\n" >> /root/.ssh/config
RUN git clone git@github.com:sandalon/wordfinder.git /home/server-setup

ADD entries.txt .
ADD words2entries.txt .
ADD create.cql .

ADD run.sh .
RUN chmod +x run.sh

ENTRYPOINT ["/run.sh"]
