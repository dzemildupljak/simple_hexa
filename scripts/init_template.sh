# Create the cmd folder and main.go file
mkdir cmd
touch cmd/main.go

# Create the internal folder structure
mkdir -p internal/app/application
mkdir -p internal/app/domain
mkdir -p internal/app/ports/inbound
mkdir -p internal/app/ports/outbound
mkdir -p internal/infrastructure/persistence

# Create the application services files
touch internal/app/application/service.go
touch internal/app/application/service_impl.go

# Create the domain files
touch internal/app/domain/model.go
touch internal/app/domain/repository.go

# Create the inbound ports file
touch internal/app/ports/inbound/handler.go

# Create the outbound ports file
touch internal/app/ports/outbound/repository.go

# Create the infrastructure files
touch internal/infrastructure/persistence/repository_impl.go

# Create the pkg, scripts, config, and tests folders
mkdir pkg
mkdir scripts
mkdir config
mkdir tests

# Create the go.mod and go.sum files
touch go.mod
touch go.sum

# Create the README.md file
touch README.md
