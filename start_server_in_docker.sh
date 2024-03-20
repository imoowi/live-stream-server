#!/bin/sh
docker build -f Dockerfile --progress=plain -t live-stream-server --cpu-shares 2 --no-cache .
docker rm -f live-stream-server || true
docker run -d -p 8000:8000 -v ./configs:/configs -v ./runtime:/runtime --name live-stream-server live-stream-server