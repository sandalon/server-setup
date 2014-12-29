# Set the base image to Ubuntu
FROM sandalon/java8

# File Author / Maintainer
MAINTAINER Corey McClelland <corey@wordfinder.io>

RUN apt-get update

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
