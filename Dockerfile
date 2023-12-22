# Use an official Golang runtime as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy the Go application files to the working directory
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 80 for HTTP (Cloudflare Flexible SSL)
EXPOSE 80

# Run the Go application
CMD ["./main"]
