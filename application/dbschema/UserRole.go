// @generated Do not edit a file, which is automatically generated by the generator.

package dbschema

import (
	"fmt"

	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/param"

	"time"
)

type Slice_UserRole []*UserRole

func (s Slice_UserRole) Range(fn func(m factory.Model) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_UserRole) RangeRaw(fn func(m *UserRole) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

// UserRole 用户角色
type UserRole struct {
	base    factory.Base
	objects []*UserRole

	Id          uint   `db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	Name        string `db:"name" bson:"name" comment:"名称" json:"name" xml:"name"`
	Description string `db:"description" bson:"description" comment:"说明" json:"description" xml:"description"`
	Created     uint   `db:"created" bson:"created" comment:"添加时间" json:"created" xml:"created"`
	Updated     uint   `db:"updated" bson:"updated" comment:"修改时间" json:"updated" xml:"updated"`
	Disabled    string `db:"disabled" bson:"disabled" comment:"是否禁用" json:"disabled" xml:"disabled"`
	ParentId    uint   `db:"parent_id" bson:"parent_id" comment:"父级ID" json:"parent_id" xml:"parent_id"`
	PermCmd     string `db:"perm_cmd" bson:"perm_cmd" comment:"指令集权限(多个用“,”分隔开)" json:"perm_cmd" xml:"perm_cmd"`
	PermAction  string `db:"perm_action" bson:"perm_action" comment:"行为权限(多个用“,”分隔开)" json:"perm_action" xml:"perm_action"`
}

// - base function

func (a *UserRole) Trans() *factory.Transaction {
	return a.base.Trans()
}

func (a *UserRole) Use(trans *factory.Transaction) factory.Model {
	a.base.Use(trans)
	return a
}

func (a *UserRole) SetContext(ctx echo.Context) factory.Model {
	a.base.SetContext(ctx)
	return a
}

func (a *UserRole) Context() echo.Context {
	return a.base.Context()
}

func (a *UserRole) SetConnID(connID int) factory.Model {
	a.base.SetConnID(connID)
	return a
}

func (a *UserRole) SetNamer(namer func(string) string) factory.Model {
	a.base.SetNamer(namer)
	return a
}

func (a *UserRole) Namer() func(string) string {
	return a.base.Namer()
}

func (a *UserRole) SetParam(param *factory.Param) factory.Model {
	a.base.SetParam(param)
	return a
}

func (a *UserRole) Param() *factory.Param {
	if a.base.Param() == nil {
		return a.NewParam()
	}
	return a.base.Param()
}

// - current function

func (a *UserRole) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName, connID[0]).Use(a.base.Trans())
	}
	return factory.NewModel(structName, a.base.ConnID()).Use(a.base.Trans())
}

func (a *UserRole) Objects() []*UserRole {
	if a.objects == nil {
		return nil
	}
	return a.objects[:]
}

func (a *UserRole) NewObjects() factory.Ranger {
	return &Slice_UserRole{}
}

func (a *UserRole) InitObjects() *[]*UserRole {
	a.objects = []*UserRole{}
	return &a.objects
}

func (a *UserRole) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(a.base.ConnID()).SetTrans(a.base.Trans()).SetCollection(a.Name_()).SetModel(a)
}

func (a *UserRole) Short_() string {
	return "user_role"
}

func (a *UserRole) Struct_() string {
	return "UserRole"
}

func (a *UserRole) Name_() string {
	if a.base.Namer() != nil {
		return WithPrefix(a.base.Namer()(a.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(a.Short_())(a))
}

func (a *UserRole) CPAFrom(source factory.Model) factory.Model {
	a.SetContext(source.Context())
	a.Use(source.Trans())
	a.SetNamer(source.Namer())
	return a
}

func (a *UserRole) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	base := a.base
	err := a.Param().SetArgs(args...).SetRecv(a).SetMiddleware(mw).One()
	a.base = base
	return err
}

func (a *UserRole) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	return a.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (a *UserRole) GroupBy(keyField string, inputRows ...[]*UserRole) map[string][]*UserRole {
	var rows []*UserRole
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = a.Objects()
	}
	r := map[string][]*UserRole{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*UserRole{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (a *UserRole) KeyBy(keyField string, inputRows ...[]*UserRole) map[string]*UserRole {
	var rows []*UserRole
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = a.Objects()
	}
	r := map[string]*UserRole{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (a *UserRole) AsKV(keyField string, valueField string, inputRows ...[]*UserRole) map[string]interface{} {
	var rows []*UserRole
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = a.Objects()
	}
	r := map[string]interface{}{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (a *UserRole) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	return a.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (a *UserRole) Add() (pk interface{}, err error) {
	a.Created = uint(time.Now().Unix())
	a.Id = 0
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
	}
	err = DBI.Fire("creating", a, nil)
	if err != nil {
		return
	}
	pk, err = a.Param().SetSend(a).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint(v)
		}
	}
	if err == nil {
		err = DBI.Fire("created", a, nil)
	}
	return
}

func (a *UserRole) Edit(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	a.Updated = uint(time.Now().Unix())
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
	}
	if err = DBI.Fire("updating", a, mw, args...); err != nil {
		return
	}
	if err = a.Setter(mw, args...).SetSend(a).Update(); err != nil {
		return
	}
	return DBI.Fire("updated", a, mw, args...)
}

func (a *UserRole) Setter(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	return a.Param().SetArgs(args...).SetMiddleware(mw)
}

func (a *UserRole) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return a.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *UserRole) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {

	if val, ok := kvset["disabled"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["disabled"] = "N"
		}
	}
	m := *a
	m.FromRow(kvset)
	var editColumns []string
	for column := range kvset {
		editColumns = append(editColumns, column)
	}
	if err = DBI.FireUpdate("updating", &m, editColumns, mw, args...); err != nil {
		return
	}
	if err = a.Setter(mw, args...).SetSend(kvset).Update(); err != nil {
		return
	}
	return DBI.FireUpdate("updated", &m, editColumns, mw, args...)
}

func (a *UserRole) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = a.Param().SetArgs(args...).SetSend(a).SetMiddleware(mw).Upsert(func() error {
		a.Updated = uint(time.Now().Unix())
		if len(a.Disabled) == 0 {
			a.Disabled = "N"
		}
		return DBI.Fire("updating", a, mw, args...)
	}, func() error {
		a.Created = uint(time.Now().Unix())
		a.Id = 0
		if len(a.Disabled) == 0 {
			a.Disabled = "N"
		}
		return DBI.Fire("creating", a, nil)
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint(v)
		}
	}
	if err == nil {
		if pk == nil {
			err = DBI.Fire("updated", a, mw, args...)
		} else {
			err = DBI.Fire("created", a, nil)
		}
	}
	return
}

func (a *UserRole) Delete(mw func(db.Result) db.Result, args ...interface{}) (err error) {

	if err = DBI.Fire("deleting", a, mw, args...); err != nil {
		return
	}
	if err = a.Param().SetArgs(args...).SetMiddleware(mw).Delete(); err != nil {
		return
	}
	return DBI.Fire("deleted", a, mw, args...)
}

func (a *UserRole) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return a.Param().SetArgs(args...).SetMiddleware(mw).Count()
}

func (a *UserRole) Reset() *UserRole {
	a.Id = 0
	a.Name = ``
	a.Description = ``
	a.Created = 0
	a.Updated = 0
	a.Disabled = ``
	a.ParentId = 0
	a.PermCmd = ``
	a.PermAction = ``
	return a
}

func (a *UserRole) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = a.Id
	r["Name"] = a.Name
	r["Description"] = a.Description
	r["Created"] = a.Created
	r["Updated"] = a.Updated
	r["Disabled"] = a.Disabled
	r["ParentId"] = a.ParentId
	r["PermCmd"] = a.PermCmd
	r["PermAction"] = a.PermAction
	return r
}

func (a *UserRole) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
		case "id":
			a.Id = param.AsUint(value)
		case "name":
			a.Name = param.AsString(value)
		case "description":
			a.Description = param.AsString(value)
		case "created":
			a.Created = param.AsUint(value)
		case "updated":
			a.Updated = param.AsUint(value)
		case "disabled":
			a.Disabled = param.AsString(value)
		case "parent_id":
			a.ParentId = param.AsUint(value)
		case "perm_cmd":
			a.PermCmd = param.AsString(value)
		case "perm_action":
			a.PermAction = param.AsString(value)
		}
	}
}

func (a *UserRole) Set(key interface{}, value ...interface{}) {
	switch k := key.(type) {
	case map[string]interface{}:
		for kk, vv := range k {
			a.Set(kk, vv)
		}
	default:
		var (
			kk string
			vv interface{}
		)
		if k, y := key.(string); y {
			kk = k
		} else {
			kk = fmt.Sprint(key)
		}
		if len(value) > 0 {
			vv = value[0]
		}
		switch kk {
		case "Id":
			a.Id = param.AsUint(vv)
		case "Name":
			a.Name = param.AsString(vv)
		case "Description":
			a.Description = param.AsString(vv)
		case "Created":
			a.Created = param.AsUint(vv)
		case "Updated":
			a.Updated = param.AsUint(vv)
		case "Disabled":
			a.Disabled = param.AsString(vv)
		case "ParentId":
			a.ParentId = param.AsUint(vv)
		case "PermCmd":
			a.PermCmd = param.AsString(vv)
		case "PermAction":
			a.PermAction = param.AsString(vv)
		}
	}
}

func (a *UserRole) AsRow() map[string]interface{} {
	r := map[string]interface{}{}
	r["id"] = a.Id
	r["name"] = a.Name
	r["description"] = a.Description
	r["created"] = a.Created
	r["updated"] = a.Updated
	r["disabled"] = a.Disabled
	r["parent_id"] = a.ParentId
	r["perm_cmd"] = a.PermCmd
	r["perm_action"] = a.PermAction
	return r
}

func (a *UserRole) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = a.AsRow()
	}
	return factory.BatchValidate(a.Short_(), kvset)
}

func (a *UserRole) Validate(field string, value interface{}) error {
	return factory.Validate(a.Short_(), field, value)
}
