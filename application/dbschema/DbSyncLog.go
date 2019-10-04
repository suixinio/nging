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

type Slice_DbSyncLog []*DbSyncLog

func (s Slice_DbSyncLog) Range(fn func(m factory.Model) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_DbSyncLog) RangeRaw(fn func(m *DbSyncLog) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

// DbSyncLog 数据表同步日志
type DbSyncLog struct {
	base    factory.Base
	objects []*DbSyncLog

	Id             uint64 `db:"id,omitempty,pk" bson:"id,omitempty" comment:"" json:"id" xml:"id"`
	SyncId         uint   `db:"sync_id" bson:"sync_id" comment:"同步方案ID" json:"sync_id" xml:"sync_id"`
	Created        uint   `db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	Result         string `db:"result" bson:"result" comment:"结果报表" json:"result" xml:"result"`
	ChangeTables   string `db:"change_tables" bson:"change_tables" comment:"被更改的表" json:"change_tables" xml:"change_tables"`
	ChangeTableNum uint   `db:"change_table_num" bson:"change_table_num" comment:"被更改的表的数量" json:"change_table_num" xml:"change_table_num"`
	Elapsed        uint64 `db:"elapsed" bson:"elapsed" comment:"总共耗时" json:"elapsed" xml:"elapsed"`
	Failed         uint   `db:"failed" bson:"failed" comment:"失败次数" json:"failed" xml:"failed"`
}

// - base function

func (a *DbSyncLog) Trans() *factory.Transaction {
	return a.base.Trans()
}

func (a *DbSyncLog) Use(trans *factory.Transaction) factory.Model {
	a.base.Use(trans)
	return a
}

func (a *DbSyncLog) SetContext(ctx echo.Context) factory.Model {
	a.base.SetContext(ctx)
	return a
}

func (a *DbSyncLog) Context() echo.Context {
	return a.base.Context()
}

func (a *DbSyncLog) SetConnID(connID int) factory.Model {
	a.base.SetConnID(connID)
	return a
}

func (a *DbSyncLog) SetNamer(namer func(string) string) factory.Model {
	a.base.SetNamer(namer)
	return a
}

func (a *DbSyncLog) Namer() func(string) string {
	return a.base.Namer()
}

func (a *DbSyncLog) SetParam(param *factory.Param) factory.Model {
	a.base.SetParam(param)
	return a
}

func (a *DbSyncLog) Param() *factory.Param {
	if a.base.Param() == nil {
		return a.NewParam()
	}
	return a.base.Param()
}

// - current function

func (a *DbSyncLog) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName, connID[0]).Use(a.base.Trans())
	}
	return factory.NewModel(structName, a.base.ConnID()).Use(a.base.Trans())
}

func (a *DbSyncLog) Objects() []*DbSyncLog {
	if a.objects == nil {
		return nil
	}
	return a.objects[:]
}

func (a *DbSyncLog) NewObjects() factory.Ranger {
	return &Slice_DbSyncLog{}
}

func (a *DbSyncLog) InitObjects() *[]*DbSyncLog {
	a.objects = []*DbSyncLog{}
	return &a.objects
}

func (a *DbSyncLog) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(a.base.ConnID()).SetTrans(a.base.Trans()).SetCollection(a.Name_()).SetModel(a)
}

func (a *DbSyncLog) Short_() string {
	return "db_sync_log"
}

func (a *DbSyncLog) Struct_() string {
	return "DbSyncLog"
}

func (a *DbSyncLog) Name_() string {
	if a.base.Namer() != nil {
		return WithPrefix(a.base.Namer()(a.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(a.Short_())(a))
}

func (a *DbSyncLog) CPAFrom(source factory.Model) factory.Model {
	a.SetContext(source.Context())
	a.Use(source.Trans())
	a.SetNamer(source.Namer())
	return a
}

func (a *DbSyncLog) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	base := a.base
	err := a.Param().SetArgs(args...).SetRecv(a).SetMiddleware(mw).One()
	a.base = base
	return err
}

func (a *DbSyncLog) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	return a.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (a *DbSyncLog) GroupBy(keyField string, inputRows ...[]*DbSyncLog) map[string][]*DbSyncLog {
	var rows []*DbSyncLog
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = a.Objects()
	}
	r := map[string][]*DbSyncLog{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*DbSyncLog{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (a *DbSyncLog) KeyBy(keyField string, inputRows ...[]*DbSyncLog) map[string]*DbSyncLog {
	var rows []*DbSyncLog
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = a.Objects()
	}
	r := map[string]*DbSyncLog{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (a *DbSyncLog) AsKV(keyField string, valueField string, inputRows ...[]*DbSyncLog) map[string]interface{} {
	var rows []*DbSyncLog
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

func (a *DbSyncLog) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	return a.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (a *DbSyncLog) Add() (pk interface{}, err error) {
	a.Created = uint(time.Now().Unix())
	a.Id = 0
	err = DBI.Fire("creating", a, nil)
	if err != nil {
		return
	}
	pk, err = a.Param().SetSend(a).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint64); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint64(v)
		}
	}
	if err == nil {
		err = DBI.Fire("created", a, nil)
	}
	return
}

func (a *DbSyncLog) Edit(mw func(db.Result) db.Result, args ...interface{}) (err error) {

	if err = DBI.Fire("updating", a, mw, args...); err != nil {
		return
	}
	if err = a.Setter(mw, args...).SetSend(a).Update(); err != nil {
		return
	}
	return DBI.Fire("updated", a, mw, args...)
}

func (a *DbSyncLog) Setter(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	return a.Param().SetArgs(args...).SetMiddleware(mw)
}

func (a *DbSyncLog) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return a.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *DbSyncLog) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {

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

func (a *DbSyncLog) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = a.Param().SetArgs(args...).SetSend(a).SetMiddleware(mw).Upsert(func() error {
		return DBI.Fire("updating", a, mw, args...)
	}, func() error {
		a.Created = uint(time.Now().Unix())
		a.Id = 0
		return DBI.Fire("creating", a, nil)
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint64); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint64(v)
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

func (a *DbSyncLog) Delete(mw func(db.Result) db.Result, args ...interface{}) (err error) {

	if err = DBI.Fire("deleting", a, mw, args...); err != nil {
		return
	}
	if err = a.Param().SetArgs(args...).SetMiddleware(mw).Delete(); err != nil {
		return
	}
	return DBI.Fire("deleted", a, mw, args...)
}

func (a *DbSyncLog) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return a.Param().SetArgs(args...).SetMiddleware(mw).Count()
}

func (a *DbSyncLog) Reset() *DbSyncLog {
	a.Id = 0
	a.SyncId = 0
	a.Created = 0
	a.Result = ``
	a.ChangeTables = ``
	a.ChangeTableNum = 0
	a.Elapsed = 0
	a.Failed = 0
	return a
}

func (a *DbSyncLog) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = a.Id
	r["SyncId"] = a.SyncId
	r["Created"] = a.Created
	r["Result"] = a.Result
	r["ChangeTables"] = a.ChangeTables
	r["ChangeTableNum"] = a.ChangeTableNum
	r["Elapsed"] = a.Elapsed
	r["Failed"] = a.Failed
	return r
}

func (a *DbSyncLog) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
		case "id":
			a.Id = param.AsUint64(value)
		case "sync_id":
			a.SyncId = param.AsUint(value)
		case "created":
			a.Created = param.AsUint(value)
		case "result":
			a.Result = param.AsString(value)
		case "change_tables":
			a.ChangeTables = param.AsString(value)
		case "change_table_num":
			a.ChangeTableNum = param.AsUint(value)
		case "elapsed":
			a.Elapsed = param.AsUint64(value)
		case "failed":
			a.Failed = param.AsUint(value)
		}
	}
}

func (a *DbSyncLog) Set(key interface{}, value ...interface{}) {
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
			a.Id = param.AsUint64(vv)
		case "SyncId":
			a.SyncId = param.AsUint(vv)
		case "Created":
			a.Created = param.AsUint(vv)
		case "Result":
			a.Result = param.AsString(vv)
		case "ChangeTables":
			a.ChangeTables = param.AsString(vv)
		case "ChangeTableNum":
			a.ChangeTableNum = param.AsUint(vv)
		case "Elapsed":
			a.Elapsed = param.AsUint64(vv)
		case "Failed":
			a.Failed = param.AsUint(vv)
		}
	}
}

func (a *DbSyncLog) AsRow() map[string]interface{} {
	r := map[string]interface{}{}
	r["id"] = a.Id
	r["sync_id"] = a.SyncId
	r["created"] = a.Created
	r["result"] = a.Result
	r["change_tables"] = a.ChangeTables
	r["change_table_num"] = a.ChangeTableNum
	r["elapsed"] = a.Elapsed
	r["failed"] = a.Failed
	return r
}

func (a *DbSyncLog) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = a.AsRow()
	}
	return factory.BatchValidate(a.Short_(), kvset)
}

func (a *DbSyncLog) Validate(field string, value interface{}) error {
	return factory.Validate(a.Short_(), field, value)
}
