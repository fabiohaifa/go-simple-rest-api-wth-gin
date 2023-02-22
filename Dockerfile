FROM golang:1.20-alpine

WORKDIR /app

RUN go mod init api-wth-gin
RUN go mod tidy

# COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go get github.com/gin-gonic/gin
RUN go build -o /api-wth-gin

EXPOSE 80

CMD [ "/api-wth-gin" ]