package model_v1_type

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Type struct {
	Id          int64
	Name        string
	Description string
	CreatedAt   int64
	UpdatedAt   int64
	DeletedAt   int64
}

func (t *Type) TableName() string {
	return "types"
}

func (t *Type) Create(name string, description string) (err error) {

	o := orm.NewOrm()

	newType := &Type{
		Name:        name,
		Description: description,
		UpdatedAt:   time.Now().Unix(),
		CreatedAt:   time.Now().Unix(),
	}

	o.Insert(&newType)

	return
}

func (t *Type) ReadList(start int, count int, asc bool) (err error, num int64, list []Type) {

	o := orm.NewOrm()
	qs := o.QueryTable(t.TableName()).Limit(count, start)

	if asc {
		qs = qs.OrderBy("id")
	} else {
		qs = qs.OrderBy("-id")
	}

	num, err = qs.All(list)

	return
}

func (t *Type) Update(id int64) (err error) {

	return
}

func (t *Type) DeleteById(id int64) (err error, num int64) {

	o := orm.NewOrm()
	qs := o.QueryTable(t.TableName()).Filter("id", id)

	num, err = qs.Delete()

	return
}
