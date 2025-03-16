# Step 1: Use an official Go image as the base image
FROM golang:1.24-alpine as build

# Step 2: Set the working directory in the container
WORKDIR /app

# Step 3: Copy the Go module files into the container
COPY go.mod go.sum ./

# Step 4: Download the Go dependencies
RUN go mod tidy

# Step 5: Copy the entire application code into the container
COPY . .

# Step 6: Build the Go binary
RUN GOOS=linux GOARCH=amd64 go build -o task-service ./cmd

# Step 7: Create a new image based on a smaller, optimized base image
FROM alpine:latest

# Step 8: Install dependencies for running the Go binary (e.g., glibc)
RUN apk --no-cache add ca-certificates

# Step 9: Copy the binary from the build stage
COPY --from=build /app/task-service /usr/local/bin/task-service

# Step 10: Expose the port that the service will listen on
EXPOSE 8080

# Step 11: Define the command to run the service
ENTRYPOINT ["/usr/local/bin/task-service"]
