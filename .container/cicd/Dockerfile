FROM golang:1.20.6-bookworm

RUN apt-get update \
 && apt-get install -y --no-install-recommends \
    zip \
    binutils-mingw-w64-x86-64 \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/*

WORKDIR /workspace

RUN useradd -u 1001 -m nonroot && chown nonroot:nonroot -R /workspace
# WA: https://github.com/actions/runner/issues/2411

USER nonroot
