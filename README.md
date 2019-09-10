## Go worker

### Dockerfile

Used multistage build to decreasing size docker image


### Start Worker

```
docker pull docker.isuvorov.com:5005/buzzguru/go
docker run -p 8000:8000 docker.isuvorov.com:5005/buzzguru/go
```