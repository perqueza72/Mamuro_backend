FROM golang:1.19
RUN mkdir /server
COPY . /server
WORKDIR /server
RUN go build -o main main.go
RUN chmod +x main
CMD ["./main"]

EXPOSE 8080
