# Set the base image to Ubuntu
FROM ubuntu

# File Author / Maintainer
MAINTAINER Corey McClelland <corey@wordfinder.io>

RUN apt-get update

RUN apt-get install -y software-properties-common python-software-properties
RUN add-apt-repository ppa:chris-lea/node.js

RUN apt-get update

# download and install nodejs and git
RUN apt-get install -y nodejs
RUN apt-get install -y git git-core
RUN npm install pm2 -g
