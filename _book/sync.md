# golang中的sync的

## WaitGroup
等待goroutine完成， wg.Wait()

其中， 可以使用wg.Add(1)， 来增加一个goroutine的等待标记， 通过wg.Done()来表示完成一个goroutine();

一定要先Add，然后Done， 否则会引发panic；

## Mutex
互斥过程， 表现形式是 

```go
type Inst struct {
    Name string
    Count int
    lock *sync.Locker
}
```

## RWMutex

- RWMutex 是单写多读锁，该锁可以加多个读锁或者一个写锁
- 读锁占用的情况下会阻止写，不会阻止读，多个 goroutine 可以同时获取读锁
- 写锁会阻止其他 goroutine（无论读和写）进来，整个锁由该 goroutine 独占
- 适用于读多写少的场景

### Lock() 和 Unlock()
- Lock() 加写锁，Unlock() 解写锁
- 如果在加写锁之前已经有其他的读锁和写锁，则 Lock() 会阻塞直到该锁可用，为确保该锁可用，已经阻塞的 Lock() 调用会从获得的锁中排除新的读取器，即写锁权限高于读锁，有写锁时优先进行写锁定
- 在 Lock() 之前使用 Unlock() 会导致 panic 异常
### RLock() 和 RUnlock()
- RLock() 加读锁，RUnlock() 解读锁
- RLock() 加读锁时，如果存在写锁，则无法加读锁；当只有读锁或者没有锁时，可以加读锁，读锁可以加载多个
- RUnlock() 解读锁，RUnlock() 撤销单词 RLock() 调用，对于其他同时存在的读锁则没有效果
- 在没有读锁的情况下调用 RUnlock() 会导致 panic 错误
- RUnlock() 的个数不得多余 RLock()，否则会导致 panic 错误

## Cond 
字面理解上应该是goroutine等待一个event；

```go
type Cond struct {
    L Locker    //Mutex or RWMutex
}
```

## Pool

比较经典的fmt.Println 的实现方式

```go
fmt.Println("xx")

var ppFree = sync.Pool{
    New: func() interface{} { return new(pp) },
}

func Println(a ...interface{}) (n int, err error) {
    return Fprintln(os.Stdout, a...)
}

func Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
    p := newPrinter()
    p.doPrintln(a)
    n, err = w.Write(p.buf)
    p.free()
    return
}
// newPrinter allocates a new pp struct or grabs a cached one.
func newPrinter() *pp {
    p := ppFree.Get().(*pp)
    p.panicking = false
    p.erroring = false
    p.fmt.init(&p.buf)
    return p
}
// doPrintln is like doPrint but always adds a space between arguments and a newline after the last argument.
func (p *pp) doPrintln(a []interface{}) {
    for argNum, arg := range a {
        if argNum > 0 {
            p.buf.WriteByte(' ')
        }
        p.printArg(arg, 'v')
    }
    p.buf.WriteByte('\n')
}
// free saves used pp structs in ppFree; avoids an allocation per invocation.
func (p *pp) free() {
    if cap(p.buf) > 64<<10 {
        return
    }
    p.buf = p.buf[:0]
    p.arg = nil
    p.value = reflect.Value{}
    ppFree.Put(p)
}
```