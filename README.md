# docker-env
Dockerでwebの環境構築を行う(完了)

### Docker

- Docker Desktop
  - https://docs.docker.com/desktop/
- docker-compose 3.8に必要な
  - Docker Engine: [20.10.1+](https://docs.docker.com/compose/compose-file/)
  - docker-compose : [1.25.0以上](https://docs.docker.com/compose/release-notes/#1255)

```
docker -v
Docker version 20.10.12, build 20.10.12-0ubuntu2~20.04.1
docker-compose --version
docker-compose version 1.25.0, build unknown
```

# Hello World

### docker-compose
```
> make compose/build
> make compose/up
```

### db (mysql)

データベースの新規作成。
```
> make mysql/init
> make mysql/client

mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| ...                |
+--------------------+
5 rows in set (0.01 sec)

```

### migration (flyway)

flyway: https://flywaydb.org/documentation/

#### baseline
```
// 初期化コマンド
> make flyway/baseline
```

# Bye World
```
make docker-compose down
```

### Bye db
```
make __drop/mysql
```

# Delete World
### 全コンテナ削除
```
make rm-container
```

### 全DockerImage削除
```
make rm-images
```

### Dockerログ確認
```
make compose/logs
```
