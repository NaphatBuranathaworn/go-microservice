# Build the cipher service binary
FROM golang:1.15.10-stretch as base

# add the working directory for the project
WORKDIR /go/src/watermarksvc

# Copy the service code
# COPY api api
COPY cmd/database cmd/database
COPY internal internal
COPY pkg/database pkg/database
COPY vendor vendor
COPY go.mod go.mod
COPY go.sum go.sum

# building service binary at path discovergy/www
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOFLAGS=-mod=vendor go build -o database ./cmd/database

## Using the multi-stage image to just run the binary
FROM alpine:latest as final

# Init working directory to root /
WORKDIR /

# Copy just the binary from the base image
COPY --from=base /go/src/watermarksvc/database .

# just an indication that this port will be exposed by this container
EXPOSE 3334

# command to run at the immediate start of the container
ENTRYPOINT ["./database"]

# docker build -t database_image . -f images/database/Dockerfile
# docker run -d  -p 3334:8081 --name database_container database_image