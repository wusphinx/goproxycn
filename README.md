# goproxycn
private goproxycn docker server

内网私有代码仓库可通过传入`GONOPROXY`环境变量过掉**GOPROXY**代码

# use
```
$ docker build -t goproxycn .
$ docker run  --name goproxy -e GONOPROXY=your.gitlab.cn  -p 8080:8080  -d goproxycn
$ export GOPROXY=http://localhost:8080
```