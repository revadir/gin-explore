FROM golang:1.18-bullseye

ENV APP_HOME /app
RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"

COPY ./go.mod/ ./
COPY ./go.sum ./
RUN go env -w GOPROXY=direct
RUN go mod download
COPY ./src ./

RUN go build -o /ginTest

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

EXPOSE 8080
ENTRYPOINT ["/ginTest"]