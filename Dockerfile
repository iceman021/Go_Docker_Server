FROM golang:1.18.0-alpine3.15
RUN mkdir /godocker_server
ADD .  /godocker_server
WORKDIR  /godocker_server
RUN go build -o main .
CMD ["/godocker_server/main"]
