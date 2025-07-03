# Use the official Go image as the base image
FROM golang:1.24

# Set the working directory inside the container
WORKDIR /app

# Copy the Go program into the container
COPY main.go ./

# Download dependencies and build the Go program
RUN go mod init cvss-app && go mod tidy && go build -o cvss-app

# Specify the command to run the program
ENTRYPOINT ["./cvss-app"]
