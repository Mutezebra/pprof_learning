## 主要用于记录自己的pprof学习

**系统且详细的学习pprof可见[here](https://github.com/wolfogre/go-pprof-practice)**

常用的几个shell命令

go tool pprof -seconds 10 "http://localhost:6060/debug/pprof/profile"

go tool pprof -http=:8080 "http://localhost:6060/debug/pprof/heap"

top list 


**可以进入 localhost:6060/debug/pprof 查看性能详细**
