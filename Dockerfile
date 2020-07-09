FROM golang:latest

RUN mkdir /myapp

COPY src/ /myapp/

WORKDIR /myapp

ENTRYPOINT ["go", "run", "/myapp/main.go"]

