#!/bin/bash

# Building API Docker image
docker build -t my-api-image -f api/Dockerfile .

# Building UI Docker image
docker build -t my-ui-image -f ui/Dockerfile .
