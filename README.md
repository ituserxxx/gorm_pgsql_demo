# gorm_pgsql_demo

启动 pgsql 数据库
```
docker run  -it -d  \
--name pgsql \
-e POSTGRES_USER=root \
-e POSTGRES_PASSWORD=root \
-e POSTGRES_DB=demo \
-e ALLOW_IP_RANGE=0.0.0.0/0 \
-p 5432:5432 \
-v ./pgdata:/var/lib/postgresql/data \
postgres

docker run -it -d \
--name pgadmin4 \
-e PGADMIN_DEFAULT_EMAIL=root@qq.com \
-p 5433:80 \
-e PGADMIN_DEFAULT_PASSWORD=123456 dpage/pgadmin4



```

