FROM golang:1.14

ADD . go/src/ekko
WORKDIR go/src/ekko/src/

RUN go build . && go install .

ENTRYPOINT [ "/go/bin/ekko" ]

EXPOSE 443

CMD ["go", "run", "main.go"]