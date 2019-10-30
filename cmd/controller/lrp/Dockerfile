FROM golang:1.13 as builder
WORKDIR /controller/
COPY . .
RUN  CGO_ENABLED=0 GOOS=linux go build -mod vendor -o lrp_controller ./cmd/controller/lrp/

FROM alpine
COPY --from=builder /controller/lrp_controller /workspace/jobs/lrp_controller

ENTRYPOINT [ "/workspace/jobs/lrp_controller" ]
