FROM golang:1.22-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go mod download github.com/gin-gonic/gin
RUN go mod download gorm.io/driver/postgres
RUN go mod download gorm.io/gorm
CMD ["go", "run", "main.go"]