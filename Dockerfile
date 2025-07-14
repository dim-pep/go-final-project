# base
FROM golang:1.23.2-alpine3.20
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o todo
CMD [ "./todo" ]
