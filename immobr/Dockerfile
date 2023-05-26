FROM golang:alpine3.18 as BUILDER
WORKDIR /app
COPY . /app
RUN go build -o immobr ./main.go

FROM alpine
WORKDIR /app
COPY --from=BUILDER /app/immobr .
ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz
