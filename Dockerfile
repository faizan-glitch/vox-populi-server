FROM golang:1.17

COPY go.mod go.sum /go/src/github.com/faizan-glitch/vox-populi/

WORKDIR /go/src/github.com/faizan-glitch/vox-populi

RUN go mod download

COPY . .

EXPOSE 3000 3000

CMD [ "go", "run", "main.go" ]