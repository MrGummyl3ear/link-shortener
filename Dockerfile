FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go mod download
RUN go mod tidy
RUN go build ./...
ARG DATABASE="in-memory"
ENV database_env=$DATABASE
ENTRYPOINT /app/url_shortener -database=$database_env