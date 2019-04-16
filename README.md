# pex-awardwallet

## Start server dev
```sh
docker-compose up
```

## Build server prod
```sh
docker build -t "pex-awardwallet" .
docker -it -p 8080:8080 pex-awardwallet
```

## Refresh dependencies
```sh
godep save
```
