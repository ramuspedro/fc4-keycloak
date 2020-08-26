FROM golang:1.14
COPY ./main.go .
RUN go build main.go
EXPOSE 8080
ENTRYPOINT [ "./main" ]