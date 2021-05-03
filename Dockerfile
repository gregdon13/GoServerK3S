FROM alpine

WORKDIR /home/

ADD go-server /home/

CMD ./go-server
