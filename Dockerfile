FROM golang:1.17-alpine3.14 as builder

COPY go.mod go.sum /go/src/github.com/faizan-glitch/vox-populi/

WORKDIR /go/src/github.com/faizan-glitch/vox-populi

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/vox-populi github.com/faizan-glitch/vox-populi

FROM alpine

RUN apk add --no-cache ca-certificates && update-ca-certificates

COPY --from=builder /go/src/github.com/faizan-glitch/vox-populi/build/vox-populi /usr/bin/vox-populi

# EXPOSE 3000 3000

ENTRYPOINT ["/usr/bin/vox-populi"]