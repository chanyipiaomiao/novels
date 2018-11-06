package dtype

import (
	"bytes"
	"text/template"
)

// ParseResult 解析结果
type ParseResult struct {
	Items []interface{}
}

// Site Site
type Site struct {
	URL             string // 小说目录的URL
	SiteName        string // 小说网站的名称
	Parser          string // 解析器 就是网站的域名
	PreviousChapter string // 前一个的章节
	LastChapter     string // 最新的章节
	LastChapterLink string // 最新章节链接
}

// Novel 小说
type Novel struct {
	Name    string // 名称
	Site    []Site // URL地址
	Crontab string // 计划任务
}

var novelPrintTemplate = `
名称: {{.Name}} 计划任务: {{.Crontab}}
	{{range .Site}}
	{{.SiteName}}
    	========
        解析器: {{.Parser}}
        目录URL: {{.URL}}
        前一个章节: {{.PreviousChapter}}
        最新章节: {{.LastChapter}}
        最新章节URL: {{.LastChapterLink}}
	{{end}}
`

func (n *Novel) String() string {
	buff := new(bytes.Buffer)
	tpl, _ := template.New("test").Parse(novelPrintTemplate)
	tpl.Execute(buff, n)
	return buff.String()
}
