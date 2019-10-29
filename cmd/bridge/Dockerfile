FROM golang:1.13 as builder
WORKDIR /bridge/
COPY . .
RUN  CGO_ENABLED=0 GOOS=linux go build -mod vendor -o bridge ./cmd/bridge/

FROM alpine
COPY --from=builder /bridge/bridge /workspace/jobs/bridge


ENTRYPOINT [ "/workspace/jobs/bridge" ]
