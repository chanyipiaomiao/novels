package cli

import (
	"fmt"
	"log"
	"net/url"
	"novel-update-notice/engine"
	"novel-update-notice/util"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
	nparser "novel-update-notice/parser"
)

var (
	app = kingpin.New("novels", "novels is a novel update notice program")

	novelCURD    = app.Command("novel", "novel增删改查")
	action       = novelCURD.Flag("action", "action: update | delete | run | debug").Short('a').Required().String()
	name         = novelCURD.Flag("name", "小说/影视名称").Short('n').String()
	tableURL     = novelCURD.Flag("tableURL", "小说/影视目录URL").Short('u').String()
	siteName     = novelCURD.Flag("sitename", "网站的名称").Short('s').String()
	novelCronTab = novelCURD.Flag("cron", "设置计划任务,默认是每隔5分钟,支持Linux标准的cron表达式").Short('c').Default("*/5 * * * *").String()

	parserSub = app.Command("parser", "解析器")
	_         = parserSub.Flag("list", "列出支持的parser").Short('l').Default("all").String()

	cronTab    = app.Command("cron", "cron表达式")
	cronString = cronTab.Flag("expr", "检查cron表达式是否正确").Short('e').Required().String()
)

// InitCli 初始化cli
func InitCli() {
	c, err := app.Parse(os.Args[1:])
	if err != nil {
		log.Fatalf("parse cli args error: %s\n", err)
	}

	switch c {

	case "novel":
		switch *action {
		case "update":

			if _, err = util.CheckCron(*novelCronTab); err != nil {
				log.Fatalf("cron表达式不正确: %s\n", err)
			}

			if *tableURL == "" {
				log.Fatalf("error: need --tableURL arg\n")
			}

			if *name == "" {
				log.Fatalf("error: need --name arg\n")
			}

			urlParser, err := url.Parse(*tableURL)
			if err != nil {
				log.Fatalf("error: parse url(%s) %s\n", *tableURL, err)
			}

			err = Update(*tableURL, *siteName, urlParser.Host, *novelCronTab)
			if err != nil {
				log.Fatalf("error: %s\n", err)
			}

		case "delete":

			if *name == "" {
				log.Fatalf("error: need --name arg\n")
			}
			err := Delete(*name, *tableURL)
			if err != nil {
				log.Fatalf("error: %s\n", err)
			}

		case "debug":
			err := Debug()
			if err != nil {
				log.Fatalf("error: %s\n", err)
			}
		case "run":
			engine.Run()
			select {}
		}
	case "parser":
		nparser.ParserList()
	case "cron":
		next, err := util.CheckCron(*cronString)
		if err != nil {
			log.Fatalf("%s", err)
		}
		fmt.Println(next)
	}
}
