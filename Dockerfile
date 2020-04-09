FROM golang:1.13-alpine as builder
WORKDIR /go/theykk
COPY nodeinfo.go .
RUN CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o nodeinfo

# Create a minimal container to run a Golang static binary
FROM scratch
COPY --from=builder /go/theykk/nodeinfo .
ENTRYPOINT ["/nodeinfo"]
