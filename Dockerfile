FROM golang:1.22.2

WORKDIR /Ilios-Joke-MadlibAPI

COPY go.mod go.sum

RUN go mod downlaod

COPY . .

RUN go build -o main

CMD ["go run main.go jokes.go madLibs.go"]