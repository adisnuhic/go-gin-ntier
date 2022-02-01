# Accept the Go version for the image to be set as a build argument.
# Default to Go 1.17
ARG GO_VERSION=1.17

FROM golang:${GO_VERSION}-alpine AS builder

# Install the Certificate-Authority certificates for the app to be able to make
# calls to HTTPS endpoints.
# Git is required for fetching the dependencies.
RUN apk add --no-cache ca-certificates git 

# add base packages to alpine
RUN apk add build-base

# Unit tests
#RUN CGO_ENABLED=0 go test -v

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /tmp/app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
#COPY go.sum .

# Download modules
RUN go mod download

# Import the code
COPY . .

# Run tests
RUN go test ./...

# Build the app
RUN go build -o ./out/api .


# Stage 2 built from stage 1 
# Running container
FROM alpine

WORKDIR /app

# Import the Certificate-Authority certificates for enabling HTTPS.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy migrations from first stage
COPY --from=builder /tmp/app/migrations ./migrations

# Copy executable file
COPY --from=builder /tmp/app/out/api .

# Perform any further action as an unprivileged user.
USER nobody:nobody


# Run the binary program
CMD ["/app/api"]