package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// TypeApp struct
type TypeApp struct {
	Domain       string   `json:"domain"`
	HTTPPort     string   `json:"httpPort"`
	SSLPort      string   `json:"sslPort"`
	UseTLS       bool     `json:"useTLS"`
	TLSCrtPath   string   `json:"crtFilepath"`
	TLSKeyPath   string   `json:"keyFilepath"`
	Title        string   `json:"appTitle"`
	UserTimeZone string   `json:"userTimeZone"`
	AdminEmails  []string `json:"adminEmails"`
}

// App configs
var App = TypeApp{
	Domain:       "localhost",
	HTTPPort:     ":80",
	SSLPort:      ":443",
	UseTLS:       true,
	TLSCrtPath:   "./domain.crt",
	TLSKeyPath:   "./domain.key",
	Title:        "作业布告栏",
	UserTimeZone: "Asia/Shanghai",
	AdminEmails: []string{
		"you@domain.com",
	},
}

// TypePath struct
type TypePath struct {
	StaticFoler    string `json:"staticFolder"`
	UploadFolder   string `json:"uploadFolder"`
	ViewsFolder    string `json:"viewsFolder"`
	NameListFolder string `json:"namelistFolder"`
}

// Path configs
var Path = TypePath{
	StaticFoler:    "./app/static",
	UploadFolder:   "./upload",
	ViewsFolder:    "./app/views",
	NameListFolder: "./app/names",
}

// TypeMail struct
type TypeMail struct {
	MailAccount string `json:"mailAccount"`
	Password    string `json:"mailPassword"`
	SMTPServer  string `json:"smtpServer"`
	SMTPPort    string `json:"smtpPort"`
}

// Mail configs
var Mail = TypeMail{
	MailAccount: "you@domain.com",
	Password:    "password",
	SMTPServer:  "smtp.somedomain.com",
	SMTPPort:    "465",
}

// TypeDatabase struct
type TypeDatabase struct {
	Path string `json:"dbPath"`
}

// Database configs
var Database = TypeDatabase{
	Path: "./database.db",
}

// TypeRegex struct
type TypeRegex struct {
	PatternQQ    string `json:"patternQQ"`
	PatternStuID string `json:"patternStudentID"`
}

// Regex configs
var Regex = TypeRegex{
	PatternQQ:    "^[0-9]{5,15}$",
	PatternStuID: "^2018[0-9]{8}$",
}

// TypeLog struct
type TypeLog struct {
	Filepath string `json:"logPath"`
}

// Log configs
var Log = TypeLog{
	Filepath: "./logs",
}

// WebConstance configs
var WebConstance = map[string]string{
	"Title":                 "作业布告栏",
	"TaskList":              "作业列表",
	"SubmitHistory":         "提交记录",
	"AdminLogin":            "管理员登录",
	"UserLogin":             "登录",
	"UserLogout":            "登出",
	"Submit":                "提交",
	"PleaseLogin":           "未登录，请登录",
	"Available":             "可提交",
	"Unavailable":           "不可用",
	"SubmitPlaceHolder":     "选择文件(最大30MB)",
	"SubmitWarning":         "仅保存最后一次提交",
	"Email":                 "邮箱",
	"EmailConstr":           "qq.com",
	"EmailPlaceHolder":      "QQ号",
	"StuID":                 "学号",
	"StuIDPlaceHolder":      "201824100000",
	"LoginWarning":          "将发送邮件验证码，首次登录后学号将与邮箱绑定",
	"ConfirmKey":            "验证",
	"ConfirmKeyPlaceHolder": "邮件中的验证码",
	"SendConfirmKey":        "发送验证码",
	"MultiReg":              "重复注册, 请联系管理员",
	"Error":                 "错误",
	"UnknownError":          "未知错误",
	"MailError":             "邮件发送失败, 可能是提交过于频繁",
	"TimeValue":             "提交时间",
	"Subject":               "科目名称",
	"SubjectPlaceHolder":    "如：软件工程导论",
	"Subtitle":              "作业内容",
	"SubtitlePlaceHolder":   "如：第一章第1~12题",
	"FileFormat":            "作业格式",
	"FileFormatPlaceHolder": "如：pdf",
	"Date":                  "收取时间",
	"AdminSubmit":           "添加作业",
	"AdminWarning":          "添加后无法直接更改。如需更改请联系运维更改数据库",
	"SubmitTo":              "作业ID",
	"WantIDError":           "请求ID无效",
	"QueryIllegal":          "请求的ID不存在",
	"TaskNotEnd":            "请求的作业还未结束",
	"QueryFileSuccess":      "请求成功, 正在下载...",
	"QueryFileUnknownError": "请求失败",
	"NoAvailableFile":       "没有可用文件",
}

func init() {
	TryLoadConfig()
}

// TryLoadConfig while init
func TryLoadConfig() {
	filename := "./config.json"

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		ExportDefault()
		return
	}

	pool := make(map[string]interface{})

	err = json.Unmarshal(bytes, &pool)
	if err != nil {
		ExportDefault()
		return
	}

	defer func() {
		if x := recover(); x != nil {
			ExportDefault()
			panic("json error: " + fmt.Sprint(x))
		}
	}()

	appMap := pool["App"].(map[string]interface{})
	app := TypeApp{}
	app.Domain = appMap["domain"].(string)
	app.HTTPPort = appMap["httpPort"].(string)
	app.SSLPort = appMap["sslPort"].(string)
	app.UseTLS = appMap["useTLS"].(bool)
	app.TLSCrtPath = appMap["crtFilepath"].(string)
	app.TLSKeyPath = appMap["keyFilepath"].(string)
	app.Title = appMap["appTitle"].(string)
	app.UserTimeZone = appMap["userTimeZone"].(string)

	adminemailSlice := appMap["adminEmails"].([]interface{})
	var adminemails []string
	for _, i := range adminemailSlice {
		adminemails = append(adminemails, i.(string))
	}
	app.AdminEmails = adminemails
	App = app

	pathMap := pool["Path"].(map[string]interface{})
	path := TypePath{}
	path.StaticFoler = pathMap["staticFolder"].(string)
	path.UploadFolder = pathMap["uploadFolder"].(string)
	path.ViewsFolder = pathMap["viewsFolder"].(string)
	path.NameListFolder = pathMap["namelistFolder"].(string)
	Path = path

	mailMap := pool["Mail"].(map[string]interface{})
	mail := TypeMail{}
	mail.MailAccount = mailMap["mailAccount"].(string)
	mail.Password = mailMap["mailPassword"].(string)
	mail.SMTPServer = mailMap["smtpServer"].(string)
	mail.SMTPPort = mailMap["smtpPort"].(string)
	Mail = mail

	dbMap := pool["Database"].(map[string]interface{})
	db := TypeDatabase{}
	db.Path = dbMap["dbPath"].(string)
	Database = db

	regexMap := pool["Regex"].(map[string]interface{})
	reg := TypeRegex{}
	reg.PatternQQ = regexMap["patternQQ"].(string)
	reg.PatternStuID = regexMap["patternStudentID"].(string)
	Regex = reg

	logMap := pool["Log"].(map[string]interface{})
	log := TypeLog{}
	log.Filepath = logMap["logPath"].(string)
	Log = log

	webconstances := pool["WebConstance"].(map[string]interface{})
	for k, v := range webconstances {
		WebConstance[k] = v.(string)
	}
}

// ExportDefault when need
func ExportDefault() {
	filename := "./config.json"

	pool := make(map[string]interface{})

	pool["App"] = App
	pool["Path"] = Path
	pool["Mail"] = Mail
	pool["Database"] = Database
	pool["Regex"] = Regex
	pool["Log"] = Log
	pool["WebConstance"] = WebConstance

	bytes, err := json.MarshalIndent(pool, "", "    ")
	if err != nil {
		panic("config export error !?")
	}

	err = ioutil.WriteFile(filename, bytes, 0666)
	if err != nil {
		panic("config export error: " + err.Error())
	}
}
