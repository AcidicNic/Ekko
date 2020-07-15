[![Go Report Card](https://goreportcard.com/badge/github.com/acidicnic/ekko)](https://goreportcard.com/report/github.com/acidicnic/ekko)
# Ekko
Encrypted Logless Messaging Platform that enables messaging in countries where social platforms are filtered, or for people that need anonymity during times of political turmoil.

## Chat Page

To use the chat app:
[Click Here to View](http://ekko-chat.dev.ekko.cc)

## To Run
```bash
$ git clone https://www.github.com/AcidicNic/Ekko
$ cd Ekko
$ cd src
$ go build main.go
$ go run main.go
```
navigate to localhost:8080

## To Run with Docker
```bash
$ git clone https://www.github.com/AcidicNic/Ekko
$ docker build -t ekko-image .
$ docker run -p 8080:8080 --rm --name ekko-container ekko-image
```
navigate to localhost:8080

## To deploy to caprover
```bash
$ git clone https://www.github.com/AcidicNic/Ekko
$ docker build -t ekko-image .
$ caprover deploy
```