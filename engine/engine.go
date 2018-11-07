package engine

import (
	"bytes"
	"fmt"
	"github.com/robfig/cron"
	"log"
	"novels/db"
	"novels/dtype"
	"novels/fetcher"
	"novels/parser"
	"novels/util"
	"text/template"
)

var emailTemplate = `
	{{range .Site}}
	<h4>{{.SiteName}}</h4>
        <p>最新章节: {{.LastChapter}}</p>
        <p>最新章节地址: <a href="{{.LastChapterLink}}">{{.LastChapterLink}}</a></p>
	{{end}}
`

// Run 运行
func Run() {

	novels, err := db.GetAllNovel()
	if err != nil {
		log.Fatalf("get all novels error: %s\n", err)
	}

	c := cron.New()

	for _, novel := range novels {
		n := novel
		schedule, _ := cron.ParseStandard(novel.Crontab)
		job := cron.FuncJob(func() {
			run(n)
		})
		c.Schedule(schedule, job)
	}

	c.Start()

}

func run(novel *dtype.Novel) {

	isUpdate := false

	for index, site := range novel.Site {
		siteParser := parser.GetParser(site.Parser)
		content, err := fetcher.Fetch(site.URL)
		if err != nil {
			log.Printf("名称: %s 目录: %s 获取目录错误: %s\n", novel.Name, site.URL, err)
			continue
		}
		parseResult := siteParser.ParserFunc(content, &site)
		if len(parseResult.Items) < 2 {
			continue
		}

		lastChapter := parseResult.Items[0].(string)
		if lastChapter == "" || lastChapter == site.PreviousChapter {
			continue
		}

		lastChapterLink := parseResult.Items[1].(string)
		if site.LastChapter != lastChapter {
			novel.Site[index].PreviousChapter = novel.Site[index].LastChapter
			novel.Site[index].LastChapter = lastChapter
			novel.Site[index].LastChapterLink = lastChapterLink
			isUpdate = true
		}

	}

	if isUpdate {
		d := db.NewNovelCURD()
		err := d.Update(novel)
		if err != nil {
			log.Printf("名称: %s 更新数据库失败: %s\n", novel.Name, err)
			return
		}
		log.Printf("%s 更新数据库完毕.\n", novel.Name)

		buf := new(bytes.Buffer)
		tpl, _ := template.New("test").Parse(emailTemplate)
		err = tpl.Execute(buf, novel)
		if err != nil {
			log.Printf("解析发送信息模板失败: %s\n", err)
			return
		}

		err = util.Send(fmt.Sprintf("你关注的 [ %s ] 更新了", novel.Name), buf.String())
		if err != nil {
			log.Printf("名称: %s 发送更新通知消息失败: %s\n", novel.Name, err)
			return
		}
		log.Printf("%s 发送更新通知成功\n", novel.Name)
	} else {
		log.Printf("%s 没有更新\n", novel.Name)
	}

}
