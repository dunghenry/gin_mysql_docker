FROM golang:1.19.1-bullseye

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /gin_mysql

EXPOSE 3000

CMD [ "/gin_mysql" ]