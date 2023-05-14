FROM golang:1.20.3-bullseye
WORKDIR /workspace
RUN useradd -m nonroot && chown nonroot:nonroot -R /workspace && chmod 777 /workspace
USER nonroot
