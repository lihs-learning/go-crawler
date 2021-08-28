这是一个在慕课网的实战课学习项目

感谢 ccmouse 老师

## 针对原课程项目有一些重要改动

- [x] 去重采用 redis 的 Cuckoo Filter
- [ ] 分布式调用采用 grpc

## 已知问题

- [ ] Cuckoo Filter 最大存放量仅 10,000 左右
    - 后续按 id 分桶存入

## 运行其他依赖服务

Docker 运行 带有 Cuckoo Filter 的 Redis

```bash
docker run -d \
    --name redis-rebloom \
    -v ~/DockerData/redis-rebloom:/data \
    -p 6379:6379 \
    redislabs/rebloom \
    redis-server \
    --loadmodule "/usr/lib/redis/modules/redisbloom.so" \
    --appendonly yes --requirepass "Redis12345679#"
```

Docker 运行 Elastic Search

```bash
docker run -d --name=single-elasticsearch \
  -p 9200:9200 \
  -p 9300:9300 \
  -e "discovery.type=single-node" \
  -v elasticsearch_data:/usr/share/elasticsearch/data \
  docker.elastic.co/elasticsearch/elasticsearch:7.14.0
```
