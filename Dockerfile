
FROM ubuntu:latest

EXPOSE 8000

WORKDIR /app

ENV DB_HOST=localhost
ENV DB_PORT=5432
ENV DB_USER=root
ENV DB_PW=root
ENV DB_NAME=root
ENV DB_SSL=disable



COPY ./main main

ENTRYPOINT ["./main"]