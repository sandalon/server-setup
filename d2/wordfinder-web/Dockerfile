# Set the base image to Ubuntu
FROM sandalon/nodejs

# File Author / Maintainer
MAINTAINER Corey McClelland <corey@wordfinder.io>

RUN apt-get update -y

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
