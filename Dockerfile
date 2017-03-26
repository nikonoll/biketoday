# Base this docker container off the official golang docker image.
# Docker containers inherit everything from their base.
FROM golang:1.8

# Create a directory inside the container to store all our application and then make it the working directory.
RUN mkdir -p /go/src/biketoday
WORKDIR /go/src/biketoday

# Copy the example-app directory (where the Dockerfile lives) into the container.
COPY . /go/src/biketoday

# Download and install any required third party dependencies into the container.
RUN go get github.com/urfave/cli

# Now tell Docker what command to run when the container starts
#CMD ["go-wrapper", "get", "run"]
