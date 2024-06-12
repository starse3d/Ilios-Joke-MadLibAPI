FROM golang:1.22.2

WORKDIR /Ilios-Joke-MadlibAPI

COPY go.mod /Ilios-Joke-MadlibAPI/

COPY . .

RUN go build -o main

CMD ["./main"]