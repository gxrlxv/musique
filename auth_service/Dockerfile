FROM golang:1.18

COPY . /src
WORKDIR /src

RUN go mod download
RUN go build -o auth ./cmd/auth/main.go


CMD ["/src/auth"]