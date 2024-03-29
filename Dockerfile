# Copyright Contributors to the Open Cluster Management project

# Build the manager binary
FROM golang:1.21 as builder

WORKDIR /workspace

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY . .

# Build
RUN go build -a -o manager main.go

# Use ubi-minimal as minimal base image to package the manager binary
FROM registry.access.redhat.com/ubi8/ubi-minimal:latest
WORKDIR /
COPY --from=builder /workspace/manager .

ENTRYPOINT ["/manager"]
