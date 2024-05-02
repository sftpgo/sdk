#!/bin/bash

# https://buf.build/docs/ecosystem/cli-overview
# https://buf.build/grpc/go?version=v1.3.0

for PLUGIN in auth eventsearcher ipfilter kms notifier
do
    cd ${PLUGIN}
    echo "generate code for plugin ${PLUGIN}"
    buf generate
    cd ..
done