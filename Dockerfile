FROM ubuntu:14.04

RUN apt-get update && apt-get install --no-install-recommends -y ca-certificates

COPY bin/releasifier /usr/bin/
RUN mkdir /usr/bin/temp
RUN mkdir /usr/bin/bundle

EXPOSE 7331

CMD ["releasifier", "-config=/etc/releasifier-prod.conf"]
