FROM golang:1.18

COPY . /src
WORKDIR /src

RUN go mod download
RUN go build -o music ./cmd/music/main.go


CMD ["/src/music"]