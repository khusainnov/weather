#FROM golang:latest
#
#RUN mkdir /build
#WORKDIR /build
#
#RUN export GO111MODULE=on
#RUN go get github.com/khusainnov/weather/app/
#RUN cd /build && git clone https://github.com/khusainnov/weather.git
#
#RUN cd /build/weather/app/ && go build
#
#EXPOSE 80
#
#ENTRYPOINT ["/build/khusainnov/app/main"]