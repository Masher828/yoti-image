FROM golang:1.20.4-alpine3.18

WORKDIR /go/src/yoti/

COPY . .

RUN go mod tidy

RUN go get

RUN go build main.go

CMD [ "./main"]