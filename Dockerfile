FROM golang:1.15

COPY . .

EXPOSE 80

RUN go build -o server .

CMD ["./server"]