FROM golang:1.22.5

RUN mkdir app

WORKDIR /app

ADD . /app

RUN go build -o server .

CMD [ "./server" ]