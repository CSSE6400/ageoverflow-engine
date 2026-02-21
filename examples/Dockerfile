FROM ubuntu:latest

ARG ENGINE_VERSION=v0.9.0

RUN apt-get update && apt-get install -y curl

RUN ARCH=$(arch | sed s/aarch64/arm64/ | sed s/x86_64/amd64/) && \
     curl -o engine -L https://github.com/CSSE6400/ageoverflow-engine/releases/download/${ENGINE_VERSION}/engine-${ENGINE_VERSION}-linux-${ARCH} && chmod +x engine

CMD ["./engine", "--help"]