FROM golang:1.20-alpine

# Install bash (useful for running commands like 'go run')
RUN apk add --no-cache bash

# Set the working directory inside the container
WORKDIR /code

# Copy the Go code into the container (this line is crucial)
COPY . /code

# Default command (Go specific)
CMD ["go", "run", "main.go"]