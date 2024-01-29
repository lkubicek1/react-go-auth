@echo off
REM Building API Docker image
docker build -t my-api-image -f api/Dockerfile .

REM Building UI Docker image
docker build -t my-ui-image -f ui/Dockerfile .
