package parser

import (
	"fmt"
	"novels/dtype"
)

type SiteParser struct {
	SiteName   string                                                     // 名称
	ParserFunc func(contents []byte, site *dtype.Site) *dtype.ParseResult // Parser
}

func (s *SiteParser) String() string {
	return fmt.Sprintf("网站名称: %s", s.SiteName)
}

var (
	supportParser = map[string]*SiteParser{
		"book.qidian.com":   {ParserFunc: ParseQidianTable, SiteName: "起点中文网"},
		"www.biquge.com.tw": {ParserFunc: ParseBiquGeTwTable, SiteName: "笔趣阁"},
		"www.80s.tw":        {ParserFunc: Parse80stwTable, SiteName: "80s手机电影"},
		"www.piaotian.com":  {ParserFunc: ParsePiaotianTable, SiteName: "飘天文学网"},
		"www.zwdu.com":      {ParserFunc: ParseZwduTable, SiteName: "八一中文网"},
	}
)

func ParserList() {
	for p, s := range supportParser {
		fmt.Printf("解析器: %s ", p)
		fmt.Println(s)
	}
}

func GetParser(parserName string) *SiteParser {
	return supportParser[parserName]
}

func ExistParser(parserName string) bool {
	if _, ok := supportParser[parserName]; ok {
		return true
	}
	return false
}
