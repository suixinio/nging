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

type Slice_SendingLog []*SendingLog

func (s Slice_SendingLog) Range(fn func(m factory.Model) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_SendingLog) RangeRaw(fn func(m *SendingLog) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

// SendingLog 邮件短信等发送日志
type SendingLog struct {
	base    factory.Base
	objects []*SendingLog

	Id              uint64 `db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	Created         uint   `db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	SentAt          uint   `db:"sent_at" bson:"sent_at" comment:"发送时间" json:"sent_at" xml:"sent_at"`
	SourceId        uint64 `db:"source_id" bson:"source_id" comment:"来源ID" json:"source_id" xml:"source_id"`
	SourceType      string `db:"source_type" bson:"source_type" comment:"来源类型" json:"source_type" xml:"source_type"`
	Disabled        string `db:"disabled" bson:"disabled" comment:"是否禁用" json:"disabled" xml:"disabled"`
	Method          string `db:"method" bson:"method" comment:"发送方式(mobile-手机;email-邮箱)" json:"method" xml:"method"`
	To              string `db:"to" bson:"to" comment:"发送目标" json:"to" xml:"to"`
	Provider        string `db:"provider" bson:"provider" comment:"发送平台" json:"provider" xml:"provider"`
	Result          string `db:"result" bson:"result" comment:"发送结果描述" json:"result" xml:"result"`
	Status          string `db:"status" bson:"status" comment:"发送状态(none-无需发送)" json:"status" xml:"status"`
	Retries         uint   `db:"retries" bson:"retries" comment:"重试次数" json:"retries" xml:"retries"`
	Content         string `db:"content" bson:"content" comment:"发送消息内容" json:"content" xml:"content"`
	Params          string `db:"params" bson:"params" comment:"发送消息参数(JSON)" json:"params" xml:"params"`
	AppointmentTime uint   `db:"appointment_time" bson:"appointment_time" comment:"预约发送时间" json:"appointment_time" xml:"appointment_time"`
}

// - base function

func (a *SendingLog) Trans() *factory.Transaction {
	return a.base.Trans()
}

func (a *SendingLog) Use(trans *factory.Transaction) factory.Model {
	a.base.Use(trans)
	return a
}

func (a *SendingLog) SetContext(ctx echo.Context) factory.Model {
	a.base.SetContext(ctx)
	return a
}

func (a *SendingLog) Context() echo.Context {
	return a.base.Context()
}

func (a *SendingLog) SetConnID(connID int) factory.Model {
	a.base.SetConnID(connID)
	return a
}

func (a *SendingLog) SetNamer(namer func(string) string) factory.Model {
	a.base.SetNamer(namer)
	return a
}

func (a *SendingLog) Namer() func(string) string {
	return a.base.Namer()
}

func (a *SendingLog) SetParam(param *factory.Param) factory.Model {
	a.base.SetParam(param)
	return a
}

func (a *SendingLog) Param() *factory.Param {
	if a.base.Param() == nil {
		return a.NewParam()
	}
	return a.base.Param()
}

// - current function

func (a *SendingLog) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName, connID[0]).Use(a.base.Trans())
	}
	return factory.NewModel(structName, a.base.ConnID()).Use(a.base.Trans())
}

func (a *SendingLog) Objects() []*SendingLog {
	if a.objects == nil {
		return nil
	}
	return a.objects[:]
}

func (a *SendingLog) NewObjects() factory.Ranger {
	return &Slice_SendingLog{}
}

func (a *SendingLog) InitObjects() *[]*SendingLog {
	a.objects = []*SendingLog{}
	return &a.objects
}

func (a *SendingLog) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(a.base.ConnID()).SetTrans(a.base.Trans()).SetCollection(a.Name_()).SetModel(a)
}

func (a *SendingLog) Short_() string {
	return "sending_log"
}

func (a *SendingLog) Struct_() string {
	return "SendingLog"
}

func (a *SendingLog) Name_() string {
	if a.base.Namer() != nil {
		return WithPrefix(a.base.Namer()(a.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(a.Short_())(a))
}

func (a *SendingLog) CPAFrom(source factory.Model) factory.Model {
	a.SetContext(source.Context())
	a.Use(source.Trans())
	a.SetNamer(source.Namer())
	return a
}

func (a *SendingLog) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	base := a.base
	err := a.Param().SetArgs(args...).SetRecv(a).SetMiddleware(mw).One()
	a.base = base
	return err
}

func (a *SendingLog) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	return a.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (a *SendingLog) GroupBy(keyField string, inputRows ...[]*SendingLog) map[string][]*SendingLog {
	var rows []*SendingLog
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = a.Objects()
	}
	r := map[string][]*SendingLog{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*SendingLog{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (a *SendingLog) KeyBy(keyField string, inputRows ...[]*SendingLog) map[string]*SendingLog {
	var rows []*SendingLog
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = a.Objects()
	}
	r := map[string]*SendingLog{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (a *SendingLog) AsKV(keyField string, valueField string, inputRows ...[]*SendingLog) map[string]interface{} {
	var rows []*SendingLog
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

func (a *SendingLog) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	return a.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (a *SendingLog) Add() (pk interface{}, err error) {
	a.Created = uint(time.Now().Unix())
	a.Id = 0
	if len(a.SourceType) == 0 {
		a.SourceType = "user"
	}
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
	}
	if len(a.Method) == 0 {
		a.Method = "mobile"
	}
	if len(a.Status) == 0 {
		a.Status = "waiting"
	}
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

func (a *SendingLog) Edit(mw func(db.Result) db.Result, args ...interface{}) (err error) {

	if len(a.SourceType) == 0 {
		a.SourceType = "user"
	}
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
	}
	if len(a.Method) == 0 {
		a.Method = "mobile"
	}
	if len(a.Status) == 0 {
		a.Status = "waiting"
	}
	if err = DBI.Fire("updating", a, mw, args...); err != nil {
		return
	}
	if err = a.Setter(mw, args...).SetSend(a).Update(); err != nil {
		return
	}
	return DBI.Fire("updated", a, mw, args...)
}

func (a *SendingLog) Setter(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	return a.Param().SetArgs(args...).SetMiddleware(mw)
}

func (a *SendingLog) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return a.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *SendingLog) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {

	if val, ok := kvset["source_type"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["source_type"] = "user"
		}
	}
	if val, ok := kvset["disabled"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["disabled"] = "N"
		}
	}
	if val, ok := kvset["method"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["method"] = "mobile"
		}
	}
	if val, ok := kvset["status"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["status"] = "waiting"
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

func (a *SendingLog) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = a.Param().SetArgs(args...).SetSend(a).SetMiddleware(mw).Upsert(func() error {
		if len(a.SourceType) == 0 {
			a.SourceType = "user"
		}
		if len(a.Disabled) == 0 {
			a.Disabled = "N"
		}
		if len(a.Method) == 0 {
			a.Method = "mobile"
		}
		if len(a.Status) == 0 {
			a.Status = "waiting"
		}
		return DBI.Fire("updating", a, mw, args...)
	}, func() error {
		a.Created = uint(time.Now().Unix())
		a.Id = 0
		if len(a.SourceType) == 0 {
			a.SourceType = "user"
		}
		if len(a.Disabled) == 0 {
			a.Disabled = "N"
		}
		if len(a.Method) == 0 {
			a.Method = "mobile"
		}
		if len(a.Status) == 0 {
			a.Status = "waiting"
		}
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

func (a *SendingLog) Delete(mw func(db.Result) db.Result, args ...interface{}) (err error) {

	if err = DBI.Fire("deleting", a, mw, args...); err != nil {
		return
	}
	if err = a.Param().SetArgs(args...).SetMiddleware(mw).Delete(); err != nil {
		return
	}
	return DBI.Fire("deleted", a, mw, args...)
}

func (a *SendingLog) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return a.Param().SetArgs(args...).SetMiddleware(mw).Count()
}

func (a *SendingLog) Reset() *SendingLog {
	a.Id = 0
	a.Created = 0
	a.SentAt = 0
	a.SourceId = 0
	a.SourceType = ``
	a.Disabled = ``
	a.Method = ``
	a.To = ``
	a.Provider = ``
	a.Result = ``
	a.Status = ``
	a.Retries = 0
	a.Content = ``
	a.Params = ``
	a.AppointmentTime = 0
	return a
}

func (a *SendingLog) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = a.Id
	r["Created"] = a.Created
	r["SentAt"] = a.SentAt
	r["SourceId"] = a.SourceId
	r["SourceType"] = a.SourceType
	r["Disabled"] = a.Disabled
	r["Method"] = a.Method
	r["To"] = a.To
	r["Provider"] = a.Provider
	r["Result"] = a.Result
	r["Status"] = a.Status
	r["Retries"] = a.Retries
	r["Content"] = a.Content
	r["Params"] = a.Params
	r["AppointmentTime"] = a.AppointmentTime
	return r
}

func (a *SendingLog) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
		case "id":
			a.Id = param.AsUint64(value)
		case "created":
			a.Created = param.AsUint(value)
		case "sent_at":
			a.SentAt = param.AsUint(value)
		case "source_id":
			a.SourceId = param.AsUint64(value)
		case "source_type":
			a.SourceType = param.AsString(value)
		case "disabled":
			a.Disabled = param.AsString(value)
		case "method":
			a.Method = param.AsString(value)
		case "to":
			a.To = param.AsString(value)
		case "provider":
			a.Provider = param.AsString(value)
		case "result":
			a.Result = param.AsString(value)
		case "status":
			a.Status = param.AsString(value)
		case "retries":
			a.Retries = param.AsUint(value)
		case "content":
			a.Content = param.AsString(value)
		case "params":
			a.Params = param.AsString(value)
		case "appointment_time":
			a.AppointmentTime = param.AsUint(value)
		}
	}
}

func (a *SendingLog) Set(key interface{}, value ...interface{}) {
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
		case "Created":
			a.Created = param.AsUint(vv)
		case "SentAt":
			a.SentAt = param.AsUint(vv)
		case "SourceId":
			a.SourceId = param.AsUint64(vv)
		case "SourceType":
			a.SourceType = param.AsString(vv)
		case "Disabled":
			a.Disabled = param.AsString(vv)
		case "Method":
			a.Method = param.AsString(vv)
		case "To":
			a.To = param.AsString(vv)
		case "Provider":
			a.Provider = param.AsString(vv)
		case "Result":
			a.Result = param.AsString(vv)
		case "Status":
			a.Status = param.AsString(vv)
		case "Retries":
			a.Retries = param.AsUint(vv)
		case "Content":
			a.Content = param.AsString(vv)
		case "Params":
			a.Params = param.AsString(vv)
		case "AppointmentTime":
			a.AppointmentTime = param.AsUint(vv)
		}
	}
}

func (a *SendingLog) AsRow() map[string]interface{} {
	r := map[string]interface{}{}
	r["id"] = a.Id
	r["created"] = a.Created
	r["sent_at"] = a.SentAt
	r["source_id"] = a.SourceId
	r["source_type"] = a.SourceType
	r["disabled"] = a.Disabled
	r["method"] = a.Method
	r["to"] = a.To
	r["provider"] = a.Provider
	r["result"] = a.Result
	r["status"] = a.Status
	r["retries"] = a.Retries
	r["content"] = a.Content
	r["params"] = a.Params
	r["appointment_time"] = a.AppointmentTime
	return r
}

func (a *SendingLog) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = a.AsRow()
	}
	return factory.BatchValidate(a.Short_(), kvset)
}

func (a *SendingLog) Validate(field string, value interface{}) error {
	return factory.Validate(a.Short_(), field, value)
}
