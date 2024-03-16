FROM golang:1.22.1-alpine3.19
WORKDIR /workspace
COPY . /workspace/
EXPOSE 8080
RUN apk update && apk upgrade && apk add --no-cache make
RUN make build
ENTRYPOINT [ "./main" ]