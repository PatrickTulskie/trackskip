# Use the official Go image as a base image
FROM golang:1.21

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod file to download dependencies
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the entire project inside the container
COPY . .

# The command below will default to running the CLI, but you can override it for testing
CMD ["go", "run", "./cli.go"]
