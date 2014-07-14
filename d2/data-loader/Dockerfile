FROM ubuntu

# File Author / Maintainer
MAINTAINER Corey McClelland <corey@wnd.io>

RUN apt-get update

#install cqlsh
RUN apt-get install -y python-setuptools
RUN easy_install pip
RUN pip install cqlsh

ADD create.cql .
ADD word.cql .
ADD cqlsh .
RUN mv cqlsh /usr/local/bin

ADD word.tsv .
ADD lookup.tsv .

ADD run.sh .
RUN chmod +x run.sh

ENTRYPOINT ["/run.sh"]
