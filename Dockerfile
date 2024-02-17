FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o hgai

COPY . .

CMD [ "/app/hgai" ]

EXPOSE 3000