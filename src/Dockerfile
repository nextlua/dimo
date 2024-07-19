# Start from the latest golang base image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Add files to app folder
ADD . /app

# Build the Go app
RUN go build -o main cmd/main.go

# Expose ports to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]