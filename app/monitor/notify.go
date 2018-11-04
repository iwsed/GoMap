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

  // SetMaxEvents to 1 to allow at most 1 event's to be received
  // on the Event channel per watching cycle.
  //
  // If SetMaxEvents is not set, the default is to send all events.
  w.SetMaxEvents(1)

  // Only notify rename and move events.
//  w.FilterOps(watcher.Rename, watcher.Move)
  w.FilterOps(watcher.Create, watcher.Write, watcher.Chmod)

  go func() {
    for {
      select {
      case event := <-w.Event:
        fmt.Println(event) // Print the event's info.
        fmt.Println("hello")
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

  // Watch this folder for changes.
  if err := w.Add("/srv/git/objects"); err != nil {
    log.Fatalln(err)
  }

  // Watch test_folder recursively for changes.
  if err := w.AddRecursive("/srv/git/objects"); err != nil {
    log.Fatalln(err)
  }

  // Print a list of all of the files and folders currently
  // being watched and their paths.
/*
  for path, f := range w.WatchedFiles() {
    fmt.Printf("%s: %s\n", path, f.Name())
  }

  fmt.Println()
*/

  // Trigger 2 events after watcher started.
  go func() {
    w.Wait()
    w.TriggerEvent(watcher.Create, nil)
    w.TriggerEvent(watcher.Write, nil)
    w.TriggerEvent(watcher.Chmod, nil)
  }()

  // Start the watching process - it'll check for changes every 100ms.
  if err := w.Start(time.Millisecond * 100); err != nil {
    log.Fatalln(err)
  }
}
