FROM golang:1.10.1 AS build

# Install tools required for project
# Run `docker build --no-cache .` to update dependencies
RUN apk add --no-cache git
RUN go get github.com/golang/dep/cmd/dep

# List project dependencies with Gopkg.toml and Gopkg.lock
# These layers are only re-built when Gopkg files are updated
COPY Gopkg.lock Gopkg.toml /go/src/file/
RUN mkdir -p /go/src/file
WORKDIR /go/src/file/
# Install library dependencies
RUN dep ensure -vendor-only

# Copy the entire project and build it
# This layer is rebuilt when a file changes in the project directory
COPY . /go/src/file/
RUN go build -o app

# This results in a single layer image
FROM scratch
COPY --from=build /go/src/file/app /go/bin/app
ENTRYPOINT ["/go/bin/app"]
CMD ["--help"]