FROM golang

RUN apt-get update && apt-get -y install vim 
ENV APP  /home/dev/app
RUN mkdir -p $APP

ADD .  $APP

WORKDIR $APP

RUN go build -o main .
  
EXPOSE 8080
CMD go run main.go
