package model_v1_nature

import (
	// "fmt"

	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Nature struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
	DeletedAt   int64  `json:"deleted_at"`
}

func (t *Nature) TableName() string {
	return "natures"
}

func (t *Nature) TableUnique() [][]string {
	return [][]string{
		{"name"},
	}
}

func (t *Nature) CreateOrUpdate(stringFilter map[string]string, intFilter map[string]int64) (err error, num int64) {

	o := orm.NewOrm()
	newType := &Nature{
		UpdatedAt: time.Now().Unix(),
	}

	for k, v := range stringFilter {
		switch k {
		case "name":
			newType.Name = v
		case "description":
			newType.Description = v
		}
	}

	if _, ok := intFilter["deleted_at"]; ok {
		newType.DeletedAt = time.Now().Unix()
	}

	if v, ok := intFilter["id"]; ok {
		newType.Id = v
		newType.UpdatedAt = time.Now().Unix()
		_, err = o.Update(newType)
		num = v
	} else {
		newType.CreatedAt = time.Now().Unix()
		num, err = o.Insert(newType)
	}

	return
}

func (t *Nature) GetList(stringFilter map[string]string, intFilter map[string]int64, limit int64, offset int64) (err error, num int64, list []Nature) {

	o := orm.NewOrm()
	qs := o.QueryTable(t.TableName())

	for k, v := range stringFilter {
		qs = qs.Filter(k, v)
	}

	for k, v := range intFilter {
		qs = qs.Filter(k, v)
	}

	num, err = qs.Limit(limit, offset).All(&list)

	return
}

func (t *Nature) GetById(id int64) (err error, newNature Nature) {

	o := orm.NewOrm()
	newNature.Id = id
	err = o.Read(&newNature)

	return
}

func (t *Nature) GetTotal() (err error, total int64) {

	o := orm.NewOrm()
	total, err = o.QueryTable(t.TableName()).Count()
	return
}
