# GoLang实现目录文件监控


## golang收获

## 应用分析

项目组随着时间积累，知识在传承方面有些遗漏，知识分散在各自的电脑上；  

为了可以集中项目组的知识，构建一个项目组书籍是很有必要的， 经过分析采用GitBook，展示与显示都比较清晰，也简单

- markdown 文件编辑;
- html格式展示;

在本地可以参考官网的安装手册进行安装，但是在服务器的时候存在一个问题

1. 提交markdown文件到Server；
2. 服务器部署一个gitbook目录；

在服务器上的gitbook目录需要定时判断git server是否有文件更新，如果有更新则pull到服务器， 这样gitbook就会有内容更新；

如下是代码的实现部分，借用的是 watcher的包；

```go
package main

import (
  "fmt"
  "log"
  "time"
  "os/exec"

  "github.com/radovskyb/watcher"
)

func main() {
  w := watcher.New()

  w.SetMaxEvents(1)

//  w.FilterOps(watcher.Rename, watcher.Move)
  w.FilterOps(watcher.Create, watcher.Write, watcher.Chmod)

  go func() {
    for {
      select {
      case event := <-w.Event:
        fmt.Println(event) // Print the event's info.
        c := "pull.book.sh"
        cmd := exec.Command(c, "")
        err := cmd.Run()
        fmt.Printf("Command finished with error: %v\n", err)
      case err := <-w.Error:
        log.Fatalln(err)
      case <-w.Closed:
        return
      }
    }
  }()

  // Watch git/objects folder for changes.
  if err := w.Add("/srv/git/objects"); err != nil {
    log.Fatalln(err)
  }

  // Watch git/objects recursively for changes.
  if err := w.AddRecursive("/srv/git/objects"); err != nil {
    log.Fatalln(err)
  }


  // Trigger 2 events after watcher started.
  go func() {
    w.Wait()
    w.TriggerEvent(watcher.Create, nil)
    w.TriggerEvent(watcher.Write, nil)
    w.TriggerEvent(watcher.Chmod, nil)
  }()

  // Start the watching process - it'll check for changes every 1000ms.
  if err := w.Start(time.Millisecond * 1000); err != nil {
    log.Fatalln(err)
  }
}

```