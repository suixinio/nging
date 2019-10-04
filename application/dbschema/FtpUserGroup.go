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

type Slice_FtpUserGroup []*FtpUserGroup

func (s Slice_FtpUserGroup) Range(fn func(m factory.Model) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_FtpUserGroup) RangeRaw(fn func(m *FtpUserGroup) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

// FtpUserGroup FTP用户组
type FtpUserGroup struct {
	base    factory.Base
	objects []*FtpUserGroup

	Id          uint   `db:"id,omitempty,pk" bson:"id,omitempty" comment:"" json:"id" xml:"id"`
	Name        string `db:"name" bson:"name" comment:"组名称" json:"name" xml:"name"`
	Created     uint   `db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	Updated     uint   `db:"updated" bson:"updated" comment:"修改时间" json:"updated" xml:"updated"`
	Disabled    string `db:"disabled" bson:"disabled" comment:"是否禁用" json:"disabled" xml:"disabled"`
	Banned      string `db:"banned" bson:"banned" comment:"是否禁止组内用户连接" json:"banned" xml:"banned"`
	Directory   string `db:"directory" bson:"directory" comment:"授权目录" json:"directory" xml:"directory"`
	IpWhitelist string `db:"ip_whitelist" bson:"ip_whitelist" comment:"IP白名单(一行一个)" json:"ip_whitelist" xml:"ip_whitelist"`
	IpBlacklist string `db:"ip_blacklist" bson:"ip_blacklist" comment:"IP黑名单(一行一个)" json:"ip_blacklist" xml:"ip_blacklist"`
}

// - base function

func (a *FtpUserGroup) Trans() *factory.Transaction {
	return a.base.Trans()
}

func (a *FtpUserGroup) Use(trans *factory.Transaction) factory.Model {
	a.base.Use(trans)
	return a
}

func (a *FtpUserGroup) SetContext(ctx echo.Context) factory.Model {
	a.base.SetContext(ctx)
	return a
}

func (a *FtpUserGroup) Context() echo.Context {
	return a.base.Context()
}

func (a *FtpUserGroup) SetConnID(connID int) factory.Model {
	a.base.SetConnID(connID)
	return a
}

func (a *FtpUserGroup) SetNamer(namer func(string) string) factory.Model {
	a.base.SetNamer(namer)
	return a
}

func (a *FtpUserGroup) Namer() func(string) string {
	return a.base.Namer()
}

func (a *FtpUserGroup) SetParam(param *factory.Param) factory.Model {
	a.base.SetParam(param)
	return a
}

func (a *FtpUserGroup) Param() *factory.Param {
	if a.base.Param() == nil {
		return a.NewParam()
	}
	return a.base.Param()
}

// - current function

func (a *FtpUserGroup) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName, connID[0]).Use(a.base.Trans())
	}
	return factory.NewModel(structName, a.base.ConnID()).Use(a.base.Trans())
}

func (a *FtpUserGroup) Objects() []*FtpUserGroup {
	if a.objects == nil {
		return nil
	}
	return a.objects[:]
}

func (a *FtpUserGroup) NewObjects() factory.Ranger {
	return &Slice_FtpUserGroup{}
}

func (a *FtpUserGroup) InitObjects() *[]*FtpUserGroup {
	a.objects = []*FtpUserGroup{}
	return &a.objects
}

func (a *FtpUserGroup) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(a.base.ConnID()).SetTrans(a.base.Trans()).SetCollection(a.Name_()).SetModel(a)
}

func (a *FtpUserGroup) Short_() string {
	return "ftp_user_group"
}

func (a *FtpUserGroup) Struct_() string {
	return "FtpUserGroup"
}

func (a *FtpUserGroup) Name_() string {
	if a.base.Namer() != nil {
		return WithPrefix(a.base.Namer()(a.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(a.Short_())(a))
}

func (a *FtpUserGroup) CPAFrom(source factory.Model) factory.Model {
	a.SetContext(source.Context())
	a.Use(source.Trans())
	a.SetNamer(source.Namer())
	return a
}

func (a *FtpUserGroup) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	base := a.base
	err := a.Param().SetArgs(args...).SetRecv(a).SetMiddleware(mw).One()
	a.base = base
	return err
}

func (a *FtpUserGroup) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	return a.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (a *FtpUserGroup) GroupBy(keyField string, inputRows ...[]*FtpUserGroup) map[string][]*FtpUserGroup {
	var rows []*FtpUserGroup
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = a.Objects()
	}
	r := map[string][]*FtpUserGroup{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*FtpUserGroup{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (a *FtpUserGroup) KeyBy(keyField string, inputRows ...[]*FtpUserGroup) map[string]*FtpUserGroup {
	var rows []*FtpUserGroup
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = a.Objects()
	}
	r := map[string]*FtpUserGroup{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (a *FtpUserGroup) AsKV(keyField string, valueField string, inputRows ...[]*FtpUserGroup) map[string]interface{} {
	var rows []*FtpUserGroup
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

func (a *FtpUserGroup) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	return a.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (a *FtpUserGroup) Add() (pk interface{}, err error) {
	a.Created = uint(time.Now().Unix())
	a.Id = 0
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
	}
	if len(a.Banned) == 0 {
		a.Banned = "N"
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

func (a *FtpUserGroup) Edit(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	a.Updated = uint(time.Now().Unix())
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
	}
	if len(a.Banned) == 0 {
		a.Banned = "N"
	}
	if err = DBI.Fire("updating", a, mw, args...); err != nil {
		return
	}
	if err = a.Setter(mw, args...).SetSend(a).Update(); err != nil {
		return
	}
	return DBI.Fire("updated", a, mw, args...)
}

func (a *FtpUserGroup) Setter(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	return a.Param().SetArgs(args...).SetMiddleware(mw)
}

func (a *FtpUserGroup) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return a.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *FtpUserGroup) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {

	if val, ok := kvset["disabled"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["disabled"] = "N"
		}
	}
	if val, ok := kvset["banned"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["banned"] = "N"
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

func (a *FtpUserGroup) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = a.Param().SetArgs(args...).SetSend(a).SetMiddleware(mw).Upsert(func() error {
		a.Updated = uint(time.Now().Unix())
		if len(a.Disabled) == 0 {
			a.Disabled = "N"
		}
		if len(a.Banned) == 0 {
			a.Banned = "N"
		}
		return DBI.Fire("updating", a, mw, args...)
	}, func() error {
		a.Created = uint(time.Now().Unix())
		a.Id = 0
		if len(a.Disabled) == 0 {
			a.Disabled = "N"
		}
		if len(a.Banned) == 0 {
			a.Banned = "N"
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

func (a *FtpUserGroup) Delete(mw func(db.Result) db.Result, args ...interface{}) (err error) {

	if err = DBI.Fire("deleting", a, mw, args...); err != nil {
		return
	}
	if err = a.Param().SetArgs(args...).SetMiddleware(mw).Delete(); err != nil {
		return
	}
	return DBI.Fire("deleted", a, mw, args...)
}

func (a *FtpUserGroup) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return a.Param().SetArgs(args...).SetMiddleware(mw).Count()
}

func (a *FtpUserGroup) Reset() *FtpUserGroup {
	a.Id = 0
	a.Name = ``
	a.Created = 0
	a.Updated = 0
	a.Disabled = ``
	a.Banned = ``
	a.Directory = ``
	a.IpWhitelist = ``
	a.IpBlacklist = ``
	return a
}

func (a *FtpUserGroup) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = a.Id
	r["Name"] = a.Name
	r["Created"] = a.Created
	r["Updated"] = a.Updated
	r["Disabled"] = a.Disabled
	r["Banned"] = a.Banned
	r["Directory"] = a.Directory
	r["IpWhitelist"] = a.IpWhitelist
	r["IpBlacklist"] = a.IpBlacklist
	return r
}

func (a *FtpUserGroup) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
		case "id":
			a.Id = param.AsUint(value)
		case "name":
			a.Name = param.AsString(value)
		case "created":
			a.Created = param.AsUint(value)
		case "updated":
			a.Updated = param.AsUint(value)
		case "disabled":
			a.Disabled = param.AsString(value)
		case "banned":
			a.Banned = param.AsString(value)
		case "directory":
			a.Directory = param.AsString(value)
		case "ip_whitelist":
			a.IpWhitelist = param.AsString(value)
		case "ip_blacklist":
			a.IpBlacklist = param.AsString(value)
		}
	}
}

func (a *FtpUserGroup) Set(key interface{}, value ...interface{}) {
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
		case "Created":
			a.Created = param.AsUint(vv)
		case "Updated":
			a.Updated = param.AsUint(vv)
		case "Disabled":
			a.Disabled = param.AsString(vv)
		case "Banned":
			a.Banned = param.AsString(vv)
		case "Directory":
			a.Directory = param.AsString(vv)
		case "IpWhitelist":
			a.IpWhitelist = param.AsString(vv)
		case "IpBlacklist":
			a.IpBlacklist = param.AsString(vv)
		}
	}
}

func (a *FtpUserGroup) AsRow() map[string]interface{} {
	r := map[string]interface{}{}
	r["id"] = a.Id
	r["name"] = a.Name
	r["created"] = a.Created
	r["updated"] = a.Updated
	r["disabled"] = a.Disabled
	r["banned"] = a.Banned
	r["directory"] = a.Directory
	r["ip_whitelist"] = a.IpWhitelist
	r["ip_blacklist"] = a.IpBlacklist
	return r
}

func (a *FtpUserGroup) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = a.AsRow()
	}
	return factory.BatchValidate(a.Short_(), kvset)
}

func (a *FtpUserGroup) Validate(field string, value interface{}) error {
	return factory.Validate(a.Short_(), field, value)
}
