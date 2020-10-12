FROM golang:1.14.2-stretch AS builder

LABEL maintainer="ali.fareghbal@gmail.com"

WORKDIR /go/src/kube-pofc 

# Install git for `go get`
RUN apt-get update && apt-get install -y git && \
    apt-get autoremove -y && \
    apt-get clean

# Copy source files
COPY *.go ./

# Build commands
RUN go get ./... && go install ./... 
RUN go build -o kube-pofc

# Get debian base image 
FROM debian:stretch-slim 
# Set default application Envs
ENV VERSION 0.0.1
ENV NODE_NAME localhost
# Expose Application Port
EXPOSE 9000

# Get binrary from builder
COPY --from=builder /go/src/kube-pofc /usr/bin/

CMD [ "/usr/bin/kube-pofc" ]
