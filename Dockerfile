# Now copy it into our base image.
#FROM golang:1.15
FROM golang:1.17
RUN mkdir /app
# copy everything in the root directory into our /app directory
ADD . /app
# execute further commands inside our /app directory
WORKDIR /app

RUN go build -o loadgen .

ENTRYPOINT ["/app/loadgen"]
