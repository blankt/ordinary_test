#!/bin/sh
docker pull registry.cn-heyuan.aliyuncs.com/omar-docker/drone
docker rm -fv drone
docker run -d --name=drone --restart=always -p 8080:8000 \
        -v /data/drone/data/session:/session \
        -v /data/drone/data:/data:rw \
        -v /data/drone/config:/config:rw \
        -v /usr/share/zoneinfo:/usr/share/zoneinfo:ro \
        -v /etc/ssl/certs:/etc/ssl/certs:ro \
        -e TZ=Asia/Shanghai \
        -e DEBUG= \
        -e LOG_LEVEL=2 \
        registry.cn-heyuan.aliyuncs.com/omar-docker/drone