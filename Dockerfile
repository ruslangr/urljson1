FROM golang:alpine
RUN mkdir /app
WORKDIR /app
COPY ./src /app
CMD ["/app/urlToJson"]
