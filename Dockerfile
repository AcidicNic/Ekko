FROM golang:1.14

ADD . go/src/ekko
WORKDIR go/src/ekko/src/

EXPOSE 80

CMD ["go", "run", "main.go"]