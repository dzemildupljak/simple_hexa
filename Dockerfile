### Step 1: Build stage
FROM golang:1.21-alpine as builder

WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code and build the binary
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o myapp ./cmd

### 
## Step 2: Runtime stage
FROM scratch

# ARG APP_PORT=8080
ENV APP_PORT=$APP_PORT
ENV NEW_RELIC_LICENCE=$NEW_RELIC_LICENCE

# Copy only the binary from the build stage to the final image
COPY --from=builder /app/myapp /

# Expose the port your application will run on
EXPOSE 8080

# Set the entry point for the container
ENTRYPOINT ["/myapp"]