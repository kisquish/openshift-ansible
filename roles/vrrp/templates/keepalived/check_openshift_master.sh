#!/bin/bash

PORT=8443

netstat -tan|grep LISTEN|grep ":$PORT" &> /dev/null

rc=$?

if [ $rc -ne 0 ]; then
    echo "Openshift master is not listening on port $PORT"
    exit 1
else
    echo "Openshift master is listening on port $PORT"
    exit 0
fi
