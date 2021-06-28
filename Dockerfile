FROM golang:alpine
RUN mkdir /app
WORKDIR /app
COPY ./src /app
ENV TZ=Asia/Yekaterinburg
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone  
CMD ["/app/urlToJson"]
