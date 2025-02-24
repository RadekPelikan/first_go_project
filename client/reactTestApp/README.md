# React Go Test App Dockerized

## Running

Build an docker image: <Br/>
Running without flags prompts you with options

```bash
$ ./build.sh -r
```

Run docker image

```bash
$ docker run -d -p 8080:80 --name my-container first-go-app:latest
```

Delete containers

```bash
docker ps -a -q --filter ancestor=first-go-app | xargs docker rm -f
```
