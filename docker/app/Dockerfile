FROM golang:1.22.1-alpine
COPY . .
RUN mkdir /app
RUN go build -o /app/x2tg1 cmd/main.go
CMD ["/app/x2tg1"]