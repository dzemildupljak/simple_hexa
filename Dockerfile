### Step 1: Build stage
FROM golang:1.21-alpine as builder

WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code and build the binary
COPY . .
# Load environment variables from .env file
COPY .env .

RUN CGO_ENABLED=0 GOOS=linux go build -o myapp ./cmd

### 
## Step 2: Runtime stage
FROM scratch

# Copy only the binary from the build stage to the final image
COPY --from=builder /app/myapp /
COPY --from=builder /app/.env /


# Expose the port your application will run on
EXPOSE 8080

# Set the entry point for the container
ENTRYPOINT ["/myapp"]