FROM golang:latest

WORKDIR $GOPATH/src/dreamon
ADD . $GOPATH/src/dreamon
RUN go build .

EXPOSE 8080
ENTRYPOINT ["./dreamon"]

