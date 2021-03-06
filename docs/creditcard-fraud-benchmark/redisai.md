# aibench Supplemental Guide: RedisAI


### Benchmarking inference performance -- RedisAI Benchmark Go program

To measure inference performance in aibench, you first need to load
the data using the instructions in the overall [Reference data Loading section](https://github.com/RedisAI/aibench#reference-data-loading). 

Once the data is loaded,
just use the corresponding `aibench_run_inference_redisai` binary for the DL Solution
being tested:

```bash
# make sure you're on the root project folder
cd $GOPATH/src/github.com/RedisAI/aibench
./scripts/run_inference_redisai.sh
```

#### Sequence diagram - RedisAI Solution

The following diagram illustrates the sequence of requests made for each inference.


![Sequence diagram - RedisAI Solution][aibench_client_redisai]

[aibench_client_redisai]: aibench_client_redisai.png



 
## Installation 

### Local Installation -- Download the RedisAI Docker image

```bash
docker pull redisai/redisai

# Start RedisAI container 
docker run -t --rm -p 6379:6379 
```

### Production Installation -- Install RedisAI on production VM

Follow [building and running from source](https://oss.redislabs.com/redisearch/Quick_Start.html#building_and_running_from_source)
while ensuring one of the high availability methods, both [OSS](https://redis.io/topics/cluster-tutorial#redis-cluster-master-slave-model) or [Enterprise](https://redislabs.com/redis-enterprise/technology/highly-available-redis/). 