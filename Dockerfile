FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go mod download
RUN go mod tidy
RUN go build ./cmd/link-shortener
ARG DATABASE="in-memory"
ENV database_env=$DATABASE
ENTRYPOINT /app/link-shortener -database=$database_env