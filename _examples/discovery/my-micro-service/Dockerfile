FROM mysql:5.7 as god-jay-mysql
ENV MYSQL_DATABASE=god-jay-test-db \
    MYSQL_ROOT_PASSWORD=god-jay-pw
ADD schema.sql /docker-entrypoint-initdb.d
EXPOSE 3306


FROM alpine as god-jay-server
RUN mkdir -p /service/etc
COPY ./db-service /service
COPY cmd/db-service/etc /service/etc
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
WORKDIR /service/
COPY ./wait-for /service
RUN chmod +x ./wait-for
#CMD ["./db-service"]
EXPOSE 50051


FROM alpine as god-jay-client
RUN mkdir -p /service/etc
COPY ./api-service /service
COPY cmd/api-service/etc /service/etc
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
WORKDIR /service/
#CMD ["./api-service"]