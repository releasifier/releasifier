FROM ubuntu:14.04

RUN apt-get update && apt-get install --no-install-recommends -y ca-certificates

COPY bin/releasifier /usr/bin/

EXPOSE 7331

CMD ["releasifier", "-config=/etc/releasifier.conf"]
