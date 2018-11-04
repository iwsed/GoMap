# VS Code的快捷键使用过程

# VSCode了解过程

Github 位置  Microsoft/vscode  [官方地址](https://code.visualstudio.com/)

shortcut - keyboard [官方地址](https://code.visualstudio.com/docs/getstarted/tips-and-tricks)

key-map [macOS地址](https://code.visualstudio.com/shortcuts/keyboard-shortcuts-macos.pdf)


- 查看快捷键  CTRL+Shift+P,  Macbook上可以直接  code -help执行查看 `Visual Studio Code 1.27.2`
- 当日志提示具体的报错位置，想快速定位程序位置， 可以使用  code -r -g filename:128,  其中 128为文件行数；
- 比较两个文件，可以使用 code -r -d file1 file2, 目的是比较文件， 如果能用diff -y file1 file2 也可以完成类似功能[available at macOs]
- 还有一个其它人推荐的方法  ls |code -r - 可以将文件名放到code

### 键盘操作

- 光标移动，按词进行移动 option + left, right； 行首，行尾  cmd + left, right.
- 代码块移动， cmd+shift+\
- 代码块纵向连接， option + cmd + up, down, esc 取消， 功能同ultraEdit.
- 代码块上下移动，  optiopn + up, dowm;  如果加上shift就是快速复制本行内容；     //可以快速实现代码移动到代码块中
- 词选择，  option + left, right + shift 选择， 直到空格为止；
- 换行快速写入， cmd+enter
- 添加注释  cmd + /
- 代码合并 cmd + j; 大小写转换  cmd + shift + p;; transform
- 调整字符前后位置 ctrl+t
- 调整代码内部分内容排序， cmd+shift+p, sort...
- 撤销光标的移动， cmd+u
- 拷贝一行, alt+shift+up, down


### 跳转

- function定义跳转  option + click, 
- function跳转回来  ctrl + -

### 多光标操作
- 键盘方式是 cmd + option + up, down;  
- 可以通过 cmd + left, right 快速移动整个光标；
- 创建一个光标 click new place + cmd + d

### 变量定义

- 在变量的使用部分， 使用option+光标click，可以快速定位到文件定义位置；
- 在变量中按住option + cmd + click, 可以打开函数的定义位置；