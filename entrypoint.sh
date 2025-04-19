#!/bin/sh
set -e

echo "Waiting for MinIO to be ready..."
until curl -s http://chat.minio:9000/minio/health/live; do
  sleep 2
done

echo "Retrieving MinIO credentials..."
export MINIO_ACCESS_KEY="minioadmin"
export MINIO_SECRET_KEY="minioadmin123"

echo "Starting application..."
exec "/app"

