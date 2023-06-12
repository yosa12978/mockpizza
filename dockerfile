FROM golang:1.20-alpine3.17

WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o mockPizza ./cmd/mockPizza/main.go
EXPOSE 5000
CMD [ "./mockPizza" ]