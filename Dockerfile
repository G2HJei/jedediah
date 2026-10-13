# syntax=docker/dockerfile:1

# ---- build: a static binary for the target platform -----------------------
# Runs on the builder's own platform and cross-compiles, so the multi-arch
# image in CI needs no emulation. Generated *_templ.go files are committed,
# so templ is not needed here.
FROM --platform=$BUILDPLATFORM golang:1.27-alpine AS build
WORKDIR /src
COPY . .
ARG TARGETOS TARGETARCH
ARG VERSION=dev
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH \
    go build -trimpath -ldflags "-s -w -X main.version=$VERSION" -o /out/jedediah ./cmd/jedediah
# An empty data folder owned by the runtime user, so a fresh named volume is writable.
RUN mkdir /out/data

# ---- runtime: the binary alone, as an unprivileged user -------------------
# distroless/static has tzdata, so TZ works ("today" and overdue dates use it).
# UTC unless the user sets their own.
FROM gcr.io/distroless/static-debian13:nonroot
COPY --from=build /out/jedediah /jedediah
COPY --from=build --chown=65532:65532 /out/data /data
ENV JEDEDIAH_DATA_DIR=/data TZ=UTC
VOLUME /data
EXPOSE 8765
# All interfaces inside the container; publish the port on 127.0.0.1 only
# (see SPECS.md). No browser to open in here. Extra flags (--demo) append.
ENTRYPOINT ["/jedediah", "--addr", "0.0.0.0:8765", "--open=false"]
