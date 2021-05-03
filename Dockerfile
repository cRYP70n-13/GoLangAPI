FROM golang:latest

# Add maintainer information
LABEL maintainer="cRYP70n-13 otmane.kimdil@gmail.com"

# Set a working directory inside of the container
WORKDIR /app

# Copy the go.mod file to container for the dependencies
COPY go.mod .

# Copy the GoModules expected hash Files
COPY go.sum .

# Download the necessary dependencies
RUN go mod download

# Copy the source code to the workstation
COPY . .

# Build the app
RUN go build -o golangApi

# Remove all the source file
RUN find . -name "*.go" -type f -delete

# Make the PORT available outside the docker contaner as an env varaible
EXPOSE $PORT

# Run the App
CMD ["./golangApi"]
