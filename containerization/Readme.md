## we can use dockerfile concept to convert 

### golang scode into container image 

### Dockerfile

```
FROM golang:1.18-alpine as codeBuilder 
# taking sample docker image from docker hub 
LABEL name=ashutoshh
LABEL email=ashutoshh@linux.com
# just information about image designer 
RUN mkdir /myapp 
WORKDIR /myapp
# changing directory 
COPY  .  .  
# we have .dockerignore to avoide unwanted data 
RUN go build -o ashuapp
# so build binary will be save in current working directory by the name 
# of ashuapp 
FROM golang:1.18-alpine
LABEL name=dummyashu
LABEL email=dummyashutohh@linux.com 
RUN mkdir /deploy
COPY --from=codeBuilder  /myapp/ashuapp  /deploy/ 
# from above stage i took only bin file 
WORKDIR /deploy
CMD ["./ashuapp"]

```

### building image 

```
docker  build -t  ashugolang:appv2  .
```

### creating container 

```
docker run -itd  --name ashuc2   ashugolang:appv2  
```

### checking output 

```
docker  logs  ashuc2
```

### checking binary code inside running container 

```
docker  exec -it ashuc2  sh
/deploy # ls
ashuapp
/deploy # exit
```

## GOlang for docker 

### understanding docker socket concept for golang connection 

<img src="socket.png">


### By default golang client support packages are not present for docker 

### we can get golang packages from github 

```
go get  github.com/docker/docker/client
```

### list of images using go code 

```
go run . 
Docker client is now ready !!
sha256:9448a048f115d9c198ce8e6b25ce20769b461cdaedd09be3f713695eb76d08c6
```

### GO SDK docker docs 

[URLforGoSDK](https://docs.docker.com/engine/api/sdk/examples/)

