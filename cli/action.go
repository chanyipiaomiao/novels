package cli

import (
	"fmt"
	"github.com/chanyipiaomiao/hltool"
	"novel-update-notice/db"
	"novel-update-notice/dtype"
	"novel-update-notice/parser"
)

func Update(tableURL, siteName, parserStr, crontab string) error {

	if !parser.ExistParser(parserStr) {
		return fmt.Errorf("你增加的网站并不支持.")
	}

	novel := &dtype.Novel{
		Name: *name,
		Site: []dtype.Site{
			{URL: tableURL, SiteName: siteName, Parser: parserStr},
		},
		Crontab: crontab,
	}

	d := db.NewNovelCURD()
	dbNovel, err := d.Get(*name)
	if err != nil {
		return err
	}

	if dbNovel == nil {
		err := d.Update(novel)
		if err != nil {
			return err
		}
	} else {
		var dbSiteURL []string
		for _, d := range dbNovel.Site {
			dbSiteURL = append(dbSiteURL, d.URL)
		}

		if crontab != dbNovel.Crontab {
			dbNovel.Crontab = crontab
		}

		for _, site := range novel.Site {
			if ok, _ := hltool.InStringSlice(dbSiteURL, site.URL); ok {
				continue
			}
			dbNovel.Site = append(dbNovel.Site, site)
		}
		err := d.Update(dbNovel)
		if err != nil {
			return err
		}

	}
	return nil
}

func Delete(name, tableURL string) error {
	d := db.NewNovelCURD()
	if tableURL != "" {
		n, err := d.Get(name)
		if err != nil {
			return err
		}

		indexDeleteE := 0
		for i, v := range n.Site {
			if v.URL == tableURL {
				indexDeleteE = i
				break
			}
		}
		n.Site = append(n.Site[:indexDeleteE], n.Site[indexDeleteE+1:]...)
		err = d.Update(n)
		if err != nil {
			return err
		}
	} else {
		err := d.Delete(name)
		if err != nil {
			return err
		}
	}
	return nil
}

func Debug() error {
	novels, err := db.GetAllNovel()
	if err != nil {
		return err
	}

	for _, n := range novels {
		fmt.Println(n)
	}

	return nil
}
