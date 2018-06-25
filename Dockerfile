FROM golang:1.9-alpine

ENV HOME=/src
ENV GOPATH=/src
WORKDIR /src

COPY *.go ./

HEALTHCHECK --interval=30s --timeout=3s \
CMD curl -f http://localhost:$port/ || exit 1

USER nobody
CMD ["go", "run", "main.go"]