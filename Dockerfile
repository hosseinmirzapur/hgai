FROM golang:latest

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -v -o hgai

COPY . .

CMD [ "./hgai" ]
