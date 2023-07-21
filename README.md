# mahjong_sharer
[![codecov](https://codecov.io/gh/20jun01/mahjong_sharer/branch/main/graph/badge.svg?token=YF5I20WIWI)](https://codecov.io/gh/20jun01/mahjong_sharer)

## How to use
Using Docker
```
docker compose -f .\docker-compose.yaml up --build
```

If you want to up a container, use `docker compose up {container name}`

### front
1. up container
2. access to `localhost:3000`

### back
1. up container
2. application is running on `localhost:{port}` (port is defined in .env in root in go directory)
