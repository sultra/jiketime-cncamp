FROM golang:1.18.1-alpine3.15 as build
WORKDIR $GOPATH/src
COPY go.mod go.sum homework/module2/*.go ./
RUN go mod download && \ 
  go build -o $GOPATH/bin/app ./main.go


FROM alpine:3.15.4
COPY --from=build /go/bin/app /app
CMD [ "/app" ]