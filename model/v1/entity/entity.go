package model_v1_entity

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Entity struct {
	Id          int64  `json:"id"`
	Key         string `json:"key"`
	Description string `json:"description"`
	ParentId    int64  `json:"parent_id"`
	Type        int64  `json:"nature"`
	Order       int64  `json:"order"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
	DeletedAt   int64  `json:"deleted_at"`
}

func (e *Entity) TableName() string {
	return "entities"
}

func (e *Entity) Create(stringFilters map[string]string, intFilters map[string]int64) (num int64, err error) {

	o := orm.NewOrm()
	newEntity := &Entity{}

	for k, v := range stringFilters {
		switch k {
		case "key":
			newEntity.Key = v
		case "description":
			newEntity.Description = v
		}
	}

	for k, v := range intFilters {
		switch k {
		case "parent_id":
			newEntity.ParentId = v
		case "type":
			newEntity.Type = v
		case "order":
			newEntity.Order = v
		}
	}

	newEntity.CreatedAt = time.Now().Unix()
	newEntity.UpdatedAt = time.Now().Unix()

	num, err = o.Insert(newEntity)

	return
}

func (e *Entity) UpdateById(stringFilters map[string]string, intFilters map[string]int64) (err error, num int64) {

	o := orm.NewOrm()
	qs := o.QueryTable(e.TableName()).Filter("id", intFilters["id"])

	foundEntity := Entity{}
	err = qs.One(&foundEntity)

	if err != nil {
		return
	}

	for k, v := range stringFilters {

		switch k {
		case "key":
			foundEntity.Key = v
		case "description":
			foundEntity.Description = v
		}
	}

	for k, v := range intFilters {

		switch k {
		case "parent_id":
			foundEntity.ParentId = v
		case "order":
			foundEntity.Order = v
		case "deleted_at":
			foundEntity.DeletedAt = v
		case "type":
			foundEntity.Type = v
		}
	}

	num, err = o.Update(&foundEntity)

	return
}

func (e *Entity) ReadById(id int64) (err error, num int64, entity Entity) {

	o := orm.NewOrm()
	qs := o.QueryTable(e.TableName()).Filter("id")
	err = qs.One(&entity)

	return
}

func (e *Entity) ReadByIds(ids []int64) (err error, num int64, entities []Entity) {

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

func (e *Entity) GetAll(offset int, limit int, asc bool) (err error, num int64, list []Entity) {

	o := orm.NewOrm()
	qs := o.QueryTable(e.TableName())

	if asc {
		qs = qs.OrderBy("id")
	} else {
		qs = qs.OrderBy("-id")
	}

	num, err = qs.Limit(limit, offset).All(&list)

	return
}

func (e *Entity) Search(stringFilters map[string]string, intFiltfers map[string]int64) (err error, num int64, list []Entity) {

	o := orm.NewOrm()
	qs := o.QueryTable(e.TableName())

	for k, v := range stringFilters {
		qs = qs.Filter(k, v)
	}

	num, err = qs.All(&list)

	return
}
