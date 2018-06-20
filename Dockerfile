# Build Geth in a stock Go builder container
FROM golang:1.10-alpine as build
RUN apk add --no-cache git

ENV SOURCEPATH $GOPATH/src/github.com/drkaka/dockerclean

ADD . $SOURCEPATH
WORKDIR $SOURCEPATH

# Use Glide to install dependencies 
RUN go get -v github.com/Masterminds/glide
RUN cd $GOPATH/src/github.com/Masterminds/glide && go install
RUN glide install

# Test
RUN go test ./...

# Build
RUN mkdir /dockerclean
RUN go build -o /dockerclean/dockerclean main.go

# Pull Geth into a second stage deploy alpine container
FROM alpine:3.7 as runtime
RUN apk add --no-cache ca-certificates

# Install to PATH
COPY --from=build /dockerclean/dockerclean /usr/local/bin/

ENTRYPOINT ["dockerclean"]