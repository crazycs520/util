# Note Manager 

自己用markdown 记笔记，以前每次都要打开文件夹找到 .md 笔记文件，然后用 typora 打开，然后就写下这个小工具。

## 前提

* 用Golang 写的，需要安装Go

* 我的 note 笔记目录在 `/Users/cs/note` 下，可更改

* 默认打开 .md 的软件是 typora , 可以在源码里面更改

  ```go
  var fileURL = "/Users/cs/note"				//note文件夹路径
  var cmdtypora = "/Applications/Typora 2.app/Contents/MacOS/Typora"	//typora 
  ```

* `/Users/cs/note` 下的大致目录结构如下

  ```shell
  ds											#文件夹
  	MIT6.824.md								#.md文件
  go
  	go_learn_note.md
  linux
  	command.md
  macos
  	mac_os_skill.md
  mongoDB
  	mongoDB.md
  	mongoose.md
  mysql
  	mysql.md
  ```

* 将源文件 `note.go` 用 `go build note.go` 编译后生成的`note`文件复制到可执行目录`PATH`下，我复制在`/Users/cs/bin`下




## 使用

1. `note ls`查看 ``/Users/cs/note`下的所有目录，`note ll` 目录下查看所有目录及文件

   ```shell
   ▶ note ls
   ds
   go
   linux
   macos
   mongoDB
   mysql
   #note ll 查看目录下的所有目录和文件
   ▶ note ll
   ds
   	MIT6.824.md
   go
   	go_learn_note.md
   linux
   	command.md
   macos
   	mac_os_skill.md
   mongoDB
   	mongoDB.md
   	mongoose.md
   mysql
   	mysql.md
   ```

2、添加

* 添加一个新笔记目录

```shell
note add test
```

* 添加一个新笔记（要指定笔记目录）,默认添加后会用 typora 打开

```
note add test note.md
```

3、删除

* 删除一个笔记文件

```
note rm test note.md
```

* 删除一个空笔记目录

```
note rm test
```

