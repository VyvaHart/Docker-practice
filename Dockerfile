FROM golang

WORKDIR /app
COPY . .

RUN go build -o main main.go

EXPOSE 8080
CMD ["go", "run", "main.go"]
