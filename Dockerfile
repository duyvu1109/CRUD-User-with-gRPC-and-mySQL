FROM golang:1.18-alpine

RUN mkdir -p /app

WORKDIR /app

COPY . /app/

RUN go mod download all

# COPY *.go ./

RUN cd server && go build -o /go-app

EXPOSE 8080

CMD [ "/go-app" ]