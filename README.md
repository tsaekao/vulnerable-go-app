Vulnerability-goapp
====

Overview

This Web application was build for learning Web application vulnerability (ex. SQL injection, XSS, CSRF)


## Description
Vulnerability Web application build Golang
This Web application was build for learning Web application vulnerability (ex. SQL injection, XSS, CSRF)
It can be used local for learning by yourself 

### Requirement

- docker
- docker-compose

### install

```
$ git clone https://github.com/Snow-HardWolf/Vulnerability-goapp.git

$ cd Vulnerability-goapp

$ docker-compose -f runenv/docker-compose-local.yml up -d
Starting vulnapp-mysql    ... done
Starting vulnapp-csrftrap ... done
Starting vulnapp-goapp    ... done
```

**confirm**

```
$ docker-compose -f runenv/docker-compose-local.yml ps
      Name                    Command               State           Ports
----------------------------------------------------------------------------------
vulnapp-csrftrap   sh -c apk add git && go ge ...   Up      0.0.0.0:3030->3030/tcp
vulnapp-goapp      sh -c apk add git mysql-cl ...   Up      0.0.0.0:9090->9090/tcp
vulnapp-mysql      docker-entrypoint.sh mysql ...   Up      0.0.0.0:3306->3306/tcp
```





