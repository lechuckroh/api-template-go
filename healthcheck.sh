#!/bin/sh

server_port=8080
curl --fail http://localhost:${server_port}/api/healthcheck || exit 1
exit 0
