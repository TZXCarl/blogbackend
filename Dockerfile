FROM golang:alpine AS build

RUN apk add --update bash
RUN apk add git openssh-client
# Install tools required for project
# Run `docker build --no-cache .` to update dependencies
RUN apk add --no-cache git

RUN go get github.com/golang/dep/cmd/dep

# List project dependencies with Gopkg.toml and Gopkg.lock
# These layers are only re-built when Gopkg files are updated
#COPY Gopkg.lock Gopkg.toml /go/src/file/
RUN mkdir -p /go/src/file
COPY . /go/src/file/
WORKDIR /go/src/file/
# Install library dependencies
RUN dep ensure -v

# Copy the entire project and build it
# This layer is rebuilt when a file changes in the project directory
# COPY . /go/src/file/
RUN go build -o app

# This results in a single layer image
FROM alpine
COPY --from=build /go/src/file/app /go/bin/app

EXPOSE 9010

ENTRYPOINT ["/go/bin/app"]
