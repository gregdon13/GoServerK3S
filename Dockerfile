FROM alpine

WORKDIR /home/

ADD go-server /home/
EXPOSE 8080
CMD ["./go-server"]
