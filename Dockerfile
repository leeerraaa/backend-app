FROM golang:alpine

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN apk update
RUN apk add --no-cache postgresql-client
RUN apk add --no-cache git && apk add --no-cach bash && apk add build-base

# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

# build go app
RUN go mod download
RUN go build -o backend-app ./cmd/main.go

EXPOSE 8000

CMD ["./backend-app"]