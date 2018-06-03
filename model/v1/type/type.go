package model_v1_type

import (
	"fmt"

	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Type struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
	DeletedAt   int64  `json:"deleted_at"`
}

func (t *Type) TableName() string {
	return "types"
}

func (t *Type) TableUnique() [][]string {
	return [][]string{
		{"name"},
	}
}

func (t *Type) CreateOrUpdate(stringFilter map[string]string, intFilter map[string]int64) (err error, num int64) {

	o := orm.NewOrm()
	newType := &Type{
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

func (t *Type) GetList(stringFilter map[string]string, intFilter map[string]int64, limit int64, offset int64) (err error, num int64, list []Type) {

	fmt.Println(limit)
	fmt.Println(offset)

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

func (t *Type) GetById(id int64) (err error, newType Type) {

	o := orm.NewOrm()
	newType.Id = id
	err = o.Read(&newType)

	return
}
