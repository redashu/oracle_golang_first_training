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
