package model_v1_entity

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Entity struct {
	Id          int64
	Key         string
	Description string
	ParentId    int64
	Type        int64
	Order       int64
	CreatedAt   int64
	UpdatedAt   int64
	DeletedAt   int64
}

func (e *Entity) TableName() string {
	return "entities"
}

func (e *Entity) Create(key string, description string, parentId int64, _type int64, order int64) (err error, id int64) {

	if parentId < 0 {
		return errors.New("ParentId needs to be greater than 0.")
	}

	if order < 0 {
		return errors.New("Order needs to be greated than 0.")
	}

	o := orm.NewOrm()
	newEntity := &Entity{
		Key:         key,
		Description: description,
		ParentId:    parentId,
		Type:        _type,
		Order:       order,
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}

	id, err = o.Insert(newEntity)

	return
}

func (e *Entity) ReadByIds(ids []int) (err error, num int64, entities []Entity) {

	o := orm.NewOrm()
	qs := o.QueryTable(e.TableName())

	for _, e := range ids {
		qs = qs.Filter("id__in", e)
	}

	num, err = qs.All(&entities)

	return
}

func (e *Entity) DeleteByIds(ids []int) (err error, num int64) {

	o := orm.NewOrm()
	qs := o.QueryTable(e.TableName())

	for _, e := range ids {
		qs = qs.Filter("id__in", e)
	}

	num, err = qs.Delete()
	return
}
