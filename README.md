这是一个在慕课网的实战课学习项目

感谢 ccmouse 老师

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
