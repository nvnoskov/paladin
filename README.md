## Go worker
Caching service for mobile application


### Dockerfile

Used multistage build to decreasing size docker image


### Start Worker

```
docker pull registry.gitlab.com/kurs-kz/paladin:v4.5
docker run -p 8000:8000 registry.gitlab.com/kurs-kz/paladin:v4.5
```