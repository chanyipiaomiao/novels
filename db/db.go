package db

import (
	"fmt"
	"log"
	"os"
	"path"

	"novels/dtype"

	"github.com/chanyipiaomiao/hltool"
)

// NovelCURD NovelCURD
type NovelCURD struct {
	DB *hltool.BoltDB
}

// NewNovelCURD NewNovelCURD
func NewNovelCURD() *NovelCURD {
	dbPath := path.Join(path.Dir(os.Args[0]), "data", "novel.db")
	db, err := hltool.NewBoltDB(dbPath, "novel")
	if err != nil {
		log.Fatalf("db error: %s\n", err)
	}
	return &NovelCURD{DB: db}
}

// Get Get
func (n *NovelCURD) Get(name string) (*dtype.Novel, error) {
	novels, err := n.DB.Get([]string{name})
	if err != nil {
		return nil, fmt.Errorf("db get error: %s", err)
	}

	if _, ok := novels[name]; !ok {
		return nil, nil
	}

	novel := &dtype.Novel{}
	err = hltool.BytesToStruct(novels[name], novel)
	if err != nil {
		return nil, fmt.Errorf("BytesToStruct error: %s", err)
	}

	return novel, nil
}

// GetAll GetAll
func (n *NovelCURD) GetAll() (map[string][]byte, error) {
	m, err := n.DB.GetAll()
	if err != nil {
		return nil, err
	}
	return m, nil
}

// GetAllNovel 获取所有的小说
func GetAllNovel() ([]*dtype.Novel, error) {
	d := NewNovelCURD()
	m, err := d.GetAll()
	if err != nil {
		return nil, err
	}
	var novels []*dtype.Novel
	for _, v := range m {
		novel := &dtype.Novel{}
		err := hltool.BytesToStruct(v, novel)
		if err != nil {
			return nil, fmt.Errorf("BytesToStruct error: %s\n", err)
		}
		novels = append(novels, novel)
	}
	return novels, nil
}

// Update 更新
func (n *NovelCURD) Update(novel *dtype.Novel) error {
	b, err := hltool.StructToBytes(novel)
	if err != nil {
		return fmt.Errorf("StructToBytes error: %s", err)
	}

	n.DB.Set(map[string][]byte{
		novel.Name: b,
	})
	return nil
}

// Delete 删除
func (n *NovelCURD) Delete(name string) error {
	err := n.DB.Delete([]string{name})
	if err != nil {
		return fmt.Errorf("db error: %s", err)
	}
	return nil
}
