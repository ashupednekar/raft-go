FROM golang

COPY . .

RUN go build -o server ./cmd

CMD ["./server"]

