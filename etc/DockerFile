FROM golang:latest

MAINTAINER ReaTion "reation_11@163.com"

WORKDIR /home_order_service

COPY . /home_order_service

RUN go build order.go

EXPOSE 8080

RUN chmod +x /home_order_service/order

CMD ["/home_order_service/order -f /home_order_service/etc/produce.yaml"]
