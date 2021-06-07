#FROM alpine:latest
FROM golang:alpine
#FROM umputun/baseimage:app-latest
RUN mkdir /app
WORKDIR /app
COPY ./src /app
CMD ["/app/urlToJson"]
