FROM golang:1.11.2 as builder

WORKDIR /app
COPY . /app

RUN go test ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o webserver .

FROM dm848/cs-jolie:v1

WORKDIR /server

# test jolie interface
COPY . /server
COPY --from=builder /app/webserver .
RUN chmod +x test-jolie-interface.sh
RUN chmod +x webserver
RUN ./test-jolie-interface.sh


FROM dm848/consul-service:v1

WORKDIR /server
COPY --from=builder /app/webserver .

# COPY ContainerPilot configuration
COPY containerpilot.json5 /etc/containerpilot.json5
ENV CONTAINERPILOT=/etc/containerpilot.json5

ENV WEB_SERVER_PORT 8888
EXPOSE 8888
CMD ["/bin/containerpilot"]
