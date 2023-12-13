FROM golang:1.18

WORKDIR /fridge_app
COPY ./ ./

RUN go mod download
RUN go build -o ./cmd/app ./cmd/main.go

WORKDIR /fridge_app/cmd
CMD ["./app"]