#!/bin/bash

# Fetch aibench and its dependencies
cd $GOPATH/src/github.com/RedisAI/aibench/cmd
go get ./...

# Install desired binaries. At a minimum this includes aibench_generate_data, aibench_load_data, and one aibench_run_inference_*
# binaries:
cd $GOPATH/src/github.com/RedisAI/aibench/cmd
cd aibench_generate_data && go install
cd ../aibench_load_data && go install
cd ../aibench_run_inference_redisai && go install
cd ../aibench_run_inference_tensorflow_serving && go install
cd ../aibench_run_inference_flask_tensorflow && go install
