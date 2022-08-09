FROM golang:1.18

COPY . /src
WORKDIR /src

RUN go mod download
RUN go build -o userPlaylist ./cmd/userPlaylist/main.go


CMD ["/src/playlist"]