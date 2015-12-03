FROM postgres

ADD ./init_db.sh /docker-entrypoint-initdb.d

RUN mkdir -p /data
COPY data /data
