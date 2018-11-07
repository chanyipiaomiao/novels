package parser

import (
	"bytes"
	"fmt"
	"log"
	"net/url"
	"novels/dtype"
	"path"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var (
	NotFountHref = "没找到链接"

	doc     *goquery.Document
	err     error
	chapter *goquery.Selection
	result  *dtype.ParseResult
	link    string
	ok      bool
)

// ParseQidianTable 解析起点中文网小说目录 book.qidian.com
func ParseQidianTable(contents []byte, site *dtype.Site) *dtype.ParseResult {

	if doc, err = goquery.NewDocumentFromReader(bytes.NewReader(contents)); err != nil {
		log.Printf("Parse book.qidian.com error: %s\n", err)
		return nil
	}

	result = &dtype.ParseResult{}
	chapter = doc.Find("#j-catalogWrap .volume-wrap .volume").Last().Find("ul li").Last().Find("a")

	result.Items = append(result.Items, strings.TrimSpace(chapter.Text()))

	if link, ok = chapter.Attr("href"); ok {
		result.Items = append(result.Items, fmt.Sprintf("https:%s", link))
	} else {
		result.Items = append(result.Items, NotFountHref)
	}

	return result
}

// ParseBiquGeTwTable 解析www.biquge.com.tw小说网站目录
func ParseBiquGeTwTable(contents []byte, site *dtype.Site) *dtype.ParseResult {

	if doc, err = goquery.NewDocumentFromReader(bytes.NewReader(contents)); err != nil {
		log.Printf("Parse www.biquge.com.tw error: %s\n", err)
		return nil
	}

	result = &dtype.ParseResult{}
	chapter = doc.Find("#list").Find("dl dd").Last().Find("a")

	result.Items = append(result.Items, strings.TrimSpace(chapter.Text()))

	if link, ok = chapter.Attr("href"); ok {
		result.Items = append(result.Items, fmt.Sprintf("http://www.biquge.com.tw%s", link))
	} else {
		result.Items = append(result.Items, NotFountHref)
	}

	return result
}

// Parse80stwTable 解析www.80s.tw手机电影网站影视目录
func Parse80stwTable(contents []byte, site *dtype.Site) *dtype.ParseResult {

	if doc, err = goquery.NewDocumentFromReader(bytes.NewReader(contents)); err != nil {
		log.Printf("Parse www.80s.tw error: %s\n", err)
		return nil
	}

	result = &dtype.ParseResult{}
	chapter = doc.Find("#myform ul li").First().Next().Find("a").First()

	result.Items = append(result.Items, strings.TrimSpace(chapter.Text()))

	if link, ok = chapter.Attr("href"); ok {
		result.Items = append(result.Items, fmt.Sprintf("%s", link))
	} else {
		result.Items = append(result.Items, NotFountHref)
	}

	return result
}

// ParsePiaotianTable  解析www.piaotian.com小说网站目录
func ParsePiaotianTable(contents []byte, site *dtype.Site) *dtype.ParseResult {

	var (
		tableUrlParse *url.URL
		linkPrefix    string
	)

	if doc, err = goquery.NewDocumentFromReader(bytes.NewReader(contents)); err != nil {
		log.Printf("Parse www.piaotian.com error: %s\n", err)
		return nil
	}

	result = &dtype.ParseResult{}

	chapter = doc.Find(".mainbody ul").Last().Find("a").Last()

	result.Items = append(result.Items, strings.TrimSpace(chapter.Text()))

	if tableUrlParse, err = url.Parse(site.URL); err != nil {
		log.Printf("Parse www.piaotian.com Table error: %s\n", err)
		return nil
	}

	linkPrefix = path.Dir(tableUrlParse.Path)

	if link, ok = chapter.Attr("href"); ok {
		result.Items = append(result.Items, fmt.Sprintf("%s://%s%s/%s", tableUrlParse.Scheme, tableUrlParse.Host, linkPrefix, link))
	} else {
		result.Items = append(result.Items, NotFountHref)
	}

	return result
}

// ParseZwduTable 解析 www.zwdu.com 小说目录
func ParseZwduTable(contents []byte, site *dtype.Site) *dtype.ParseResult {

	if doc, err = goquery.NewDocumentFromReader(bytes.NewReader(contents)); err != nil {
		log.Printf("Parse www.zwdu.com error: %s\n", err)
		return nil
	}

	result = &dtype.ParseResult{}

	chapter := doc.Find("#list dl").Last().Find("dd").Last().Find("a")

	result.Items = append(result.Items, strings.TrimSpace(chapter.Text()))

	if link, ok = chapter.Attr("href"); ok {
		result.Items = append(result.Items, fmt.Sprintf("https://www.zwdu.com%s", link))
	} else {
		result.Items = append(result.Items, NotFountHref)
	}

	return result
}
