### Step 1: Build stage
FROM golang:1.21-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code and build the binary
COPY . .

# Copy the .env file to the build stage
COPY .env ./.env

RUN CGO_ENABLED=0 GOOS=linux go build -o myapp ./cmd

### 
## Step 2: Runtime stage
FROM alpine:latest


# Copy only the binary and .env file from the build stage to the final image
COPY --from=builder /app/myapp /
COPY --from=builder /app/.env / 
ENV HOSTNAME=build_dev_container

# Expose the port your application will run on
EXPOSE 8080

# Set the entry point for the container
ENTRYPOINT ["/myapp"]