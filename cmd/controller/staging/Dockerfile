FROM golang:1.13 as builder
WORKDIR /controller/
COPY . .
RUN  CGO_ENABLED=0 GOOS=linux go build -mod vendor -o staging_controller ./cmd/controller/staging/

FROM alpine
COPY --from=builder /controller/staging_controller /workspace/jobs/staging_controller

ENTRYPOINT [ "/workspace/jobs/staging_controller" ]
