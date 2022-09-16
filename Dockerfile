#FROM golang:1.13.8
#
#RUN mkdir /weather
#WORKDIR /weather
#
#RUN export GO111MODULE=on
#RUN cd /weather && git clone https://github.com/khusainnov/weather.git
#RUN go mod download
#
#RUN cd /app && go build main.go
#
#EXPOSE 80
#
#CMD ["./main"]
FROM golang:1.13.8

RUN mkdir /weather
WORKDIR /weather

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY /app/main.go app/
#COPY /config/.env config/
#COPY /config/config.yml config/
COPY /config/* config/
COPY /frontend/*.gohtml frontend/
COPY /internal/entity/*.go internal/entity/
COPY /pkg/handler/*.go pkg/handler/
COPY /pkg/repository/*.go pkg/repository/
COPY /pkg/service/*.go pkg/service/
COPY /schema/*.sql schema/
COPY server.go .

#RUN GOOS=linux GOARCH=amd64 go build -o weather app/main.go
RUN go build -o weather app/main.go
#RUN docker pull postgres
#RUN docker run --name weather-db POSTGRES_PASSWORD='qwerty' -p 5434:5432 -d postgres

CMD ["./weather"]
