package config

// App configs
var App = struct {
	Domain       string
	HTTPPort     string
	SSLPort      string
	UseTLS       bool
	TLSCrtPath   string
	TLSKeyPath   string
	Title        string
	UserTimeZone string
	AdminEmails  []string
}{
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

// Mail configs
var Mail = struct {
	MailAccount string
	Password    string
	SMTPServer  string
	SMTPPort    string
}{
	MailAccount: "you@domain.com",
	Password:    "password",
	SMTPServer:  "smtp.somedomain.com",
	SMTPPort:    "465",
}

// Database configs
var Database = struct {
	Path string
}{
	Path: "./database.db",
}

// Regex configs
var Regex = struct {
	PatternQQ    string
	PatternStuID string
}{
	PatternQQ:    "^[0-9]{5,15}$",
	PatternStuID: "^20182410[0-9]{4}$",
}

// Log configs
var Log = struct {
	Filepath string
}{
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
