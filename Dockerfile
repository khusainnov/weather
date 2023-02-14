#FROM postgres:latest
#
#EXPOSE 5432

FROM golang:1.19

WORKDIR /weather

COPY . .

RUN go mod download

#RUN GOOS=linux GOARCH=amd64 go build -o weather app/main.go
RUN go build -o weather app/main.go

EXPOSE 80

CMD ["./weather"]
