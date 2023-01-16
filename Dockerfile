FROM ubuntu:latest

WORKDIR /app

EXPOSE 8000

ENV HOST=localhost DB_PORT=5432
ENV USER=root PASSWORD=postgres DB_NAME=cicdgolang 

COPY ./main main

ENTRYPOINT [ "./main" ]