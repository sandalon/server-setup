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

# git
RUN mkdir -p /root/.ssh
ADD wordfinder_rsa /root/.ssh/id_rsa
RUN chmod 700 /root/.ssh/id_rsa
RUN echo "Host github.com\n\tStrictHostKeyChecking no\n" >> /root/.ssh/config

# wordfinder site
RUN git clone git@github.com:sandalon/wordfinder.git /home/wordfinder

# install libs
WORKDIR /home/wordfinder
RUN npm install

# Expose ports
EXPOSE 3000

# environment variables
ENV NODE_ENV production

ENTRYPOINT ["pm2", "start", "server.js", "-i", "max", "--no-daemon"]
