FROM postgres:16-alpine3.19
RUN mkdir /var/lib/postgresql/csv
WORKDIR /var/lib/postgresql/csv
COPY ./ueba.csv /var/lib/postgresql/csv/
COPY ./script.sql /docker-entrypoint-initdb.d/
