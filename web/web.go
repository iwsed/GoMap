package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
)

const form = `
	<html>
	<head>
	  <style>
	  table, th, td {
    	border: 1px solid green;
	  </style>
	</head>
	<body>
        <form action="#" method="post" name="bar">
            <input type="text" name="in" />
            <input type="submit" value="submit"/>
        </form>
	</body>
	
	</html>
`

const tpl = `
<!DOCTYPE html>
<html>
    <head>
        <meta charset="UTF-8">
		<title>{{.Title}}</title>
		<style>    
			table, th, td {
			border: 1px solid green;
		</style>
    </head>
    <body>
		{{range .Items}}<div>{{ . }}</div>{{else}}
		<div><strong>no rows</strong></div>{{end}}
		<table>
		<tr><th>Inst</th><th>count</th><th>rate</th><th>flag</th></tr>
		<tr><td>hello</td><td>200</td><td>70</td><td>true</td></tr>
		</table>
		<!-- 在这里定义表格，然后显示出机构成功率的状态 -->
    </body>
</html>`

type Inst struct {
	name string
	cnt  int
	rate int
	flag bool
}

type InstRate struct {
	Title string
	Items [5]Inst
}

func SimpleServer1(w http.ResponseWriter, request *http.Request) {
	check := func(err error) {
		if err != nil {
			fmt.Println(err)
		}
	}

	t, err := template.New("webpage").Parse(tpl)
	check(err)

	var data InstRate

	data.Items[0].name = "140141"
	data.Items[0].cnt = 200
	data.Items[0].rate = 70
	data.Items[0].flag = true

	data.Items[1].name = "165001"
	data.Items[1].cnt = 500
	data.Items[1].rate = 80
	data.Items[1].flag = true

	err = t.Execute(w, data)
	check(err)
}

/* handle a simple get request */
func SimpleServer(w http.ResponseWriter, request *http.Request) {

	check := func(err error) {
		if err != nil {
			fmt.Println(err)
		}
	}

	t, err := template.New("webpage").Parse(tpl)
	check(err)

	data := struct { // 这里可以定义机构数据， 然后循环显示出机构的状态；
		Title string
		Items []string
	}{
		Title: "机构成功率",
		Items: []string{
			"165001",
			"140141",
		},
	}

	fmt.Println(data)

	err = t.Execute(w, data)
	//err = t.Execute(w, "<script>alert('you have been pwned')</script>")
	check(err)
	//io.WriteString(w, handleData())
}

func FormServer(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch request.Method {
	case "GET":
		/* display the form to the user */
		io.WriteString(w, form)
	case "POST":
		/* handle the form data, note that ParseForm must
		   be called before we can extract form data */
		//request.ParseForm();
		//io.WriteString(w, request.Form["in"][0])
		io.WriteString(w, request.FormValue("in"))
	}
}

func main() {
	http.HandleFunc("/test1", SimpleServer1)
	http.HandleFunc("/test2", FormServer)
	if err := http.ListenAndServe(":8088", nil); err != nil {
		panic(err)
	}
}
