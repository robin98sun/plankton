#!/usr/bin/env bash

# build app
go build -o reducer

# run with env variables
JADE_TASKID=abcdefg \
JADE_PROTOCOL=http \
JADE_SELFNODE_ADDR=0.0.0.0 \
JADE_SELFNODE_PORT=8081 \
JADE_SUBTASKS=1234567 \
JADE_TTL=120 \
./reducer &