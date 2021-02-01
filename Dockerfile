FROM golang

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

# this will ideally be built by the ONBUILD below ;)
CMD ["go-micro", "run"]

ONBUILD COPY . /go/src/app
ONBUILD RUN go-micro download
ONBUILD RUN go-micro install