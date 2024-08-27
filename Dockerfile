FROM golang:1.22.2

WORKDIR /TaskTracker

COPY . .

RUN go mod download

WORKDIR /TaskTracker/cmd/application

RUN go build -o /TaskTracker/build/app

CMD ["/TaskTracker/build/app"]

WORKDIR /TaskTracker

