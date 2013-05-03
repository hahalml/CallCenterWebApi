package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"
	"visoline/ini"
	"visoline/mahonia"
	"vistech/odbc"
)

type AppConfig struct {
	Port       string
	ViewPath   string
	StaticPath string
	Server     string
	UserName   string
	Password   string
	DbName     string
}

var config = new(AppConfig)
var (
	view      *template.Template
	viewFuncs = template.FuncMap{
		"fs": func(t time.Time) string {
			return t.Format("2006-01-02 15:04")
		},
		"fd": func(t time.Time) string {
			return t.Format("2006-01-02")
		},
	}
)

func main() {
	http.HandleFunc(config.StaticPath, Static)
	http.HandleFunc("/", Index)
	http.HandleFunc("/CallCenterWebApi", CallCenterWebApi)
	log.Fatal(http.ListenAndServe(config.Port, nil))
}

//绘制等值线数据服务 demo
func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		view.ExecuteTemplate(w, "index", nil)
	}
}

//绘制等值线服务
func CallCenterWebApi(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		content := r.FormValue("content")
		telephones := r.FormValue("telephones")
		tels := strings.Split(telephones, ",")
		var data []byte
		buf := bytes.NewBuffer(data)
		buf.WriteString(fmt.Sprintf("%s,%s,%s,%s,%s,%s", "序号", "电话号码", "允许外呼通道", "月", "日", "积分"))
		//回车换行符
		buf.WriteByte('\r')
		buf.WriteByte('\n')
		for k, v := range tels {
			buf.WriteString(fmt.Sprintf("%d,%s,%s,%s,%s,%s", k+1, v, "", "1", "1", "100"))
			buf.WriteByte('\r')
			buf.WriteByte('\n')
		}
		//编码转换，utf-8转换到GBK
		encode := mahonia.NewEncoder("GBK")
		encodeData := encode.ConvertString(string(buf.Bytes()))
		result := []byte(encodeData)

		t := time.Now().Add(60 * time.Second)
		id := fmt.Sprintf("2-%s-%s", t.Format("20060102"), t.Format("150405.000"))
		id = strings.Replace(id, ".", "", 1)
		ioutil.WriteFile(fmt.Sprintf(`CallTask\%s.csv`, id), result, 0600)
		conn, err := odbc.Connect(fmt.Sprintf("DSN=%s;UID=%s;PWD=%s", config.Server, config.UserName, config.Password))
		if err != nil {
			fmt.Println(err)
			fmt.Fprint(w, err)
		}
		stmt, err := conn.ExecDirect(fmt.Sprintf("INSERT INTO YZ_GroupCall(Name,TheDate,TheTime,CallType,Content,Status1,Status2,TelAtt) VALUES('%s','%s','%s','%s','%s','%s','%s','%s.csv')", t.Format("200601021504.000"), t.Format("2006-01-02"), t.Format("15:04"), "2", content, "1", "0", id))
		if err != nil {
			fmt.Println(err)
			fmt.Fprint(w, err)
		}
		stmt.Close()
		conn.Close()
		fmt.Fprint(w, "ok")
	}
}
func init() {
	cfg, err := ini.Load("CallCenterWebApi.ini", false)
	if err != nil {
		panic("CallCenterWebApi.ini not find")
	}
	appSetting, ok := cfg.Sections["程序设置"]
	if !ok {
		panic("CallCenterWebApi.ini setting error")
	}
	config.Port = appSetting.Pairs["端口"]
	config.ViewPath = appSetting.Pairs["模板文件"]
	config.StaticPath = appSetting.Pairs["静态文件"]
	config.Server = appSetting.Pairs["DSN名称"]
	config.UserName = appSetting.Pairs["数据库用户名"]
	config.Password = appSetting.Pairs["数据库密码"]
	absViewPath, err := filepath.Abs(config.ViewPath)
	if err != nil {
		panic(err)
	}
	view, err = template.New("view").Funcs(viewFuncs).ParseGlob(absViewPath)
	if err != nil {
		panic(err)
	}
}

//静态文件服务
func Static(w http.ResponseWriter, r *http.Request) {
	absPath, err := filepath.Abs(r.URL.Path)
	log.Println(absPath, r.URL.Path)
	if err != nil {
		http.NotFound(w, r)
	}
	http.ServeFile(w, r, absPath)
}
