FROM ubuntu:14.04

RUN apt-get update && apt-get install --no-install-recommends -y ca-certificates

COPY ./bin/releasifier-linux-amd64 /usr/bin/releasifier
COPY ./etc/releasifier-prod.conf /etc/releasifier.conf
RUN mkdir /usr/bin/temp
RUN mkdir /usr/bin/bundle

EXPOSE 7331

CMD ["releasifier", "-config=/etc/releasifier.conf"]
