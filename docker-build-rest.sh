#!/usr/bin/env bash
# build docker image
docker build -f dockerfile-rest . -t einkaufsliste-rest:latest
