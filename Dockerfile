FROM golang:1.14
MAINTAINER weeknd.su@gmail.com
WORKDIR $GOPATH/src/godocker
ADD . $GOPATH/src/godocker
RUN go build
EXPOSE 9090
CMD ["/bin/bash","cronweb"]