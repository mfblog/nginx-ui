// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/0xJacky/Nginx-UI/server/model"
)

func newDnsCredential(db *gorm.DB, opts ...gen.DOOption) dnsCredential {
	_dnsCredential := dnsCredential{}

	_dnsCredential.dnsCredentialDo.UseDB(db, opts...)
	_dnsCredential.dnsCredentialDo.UseModel(&model.DnsCredential{})

	tableName := _dnsCredential.dnsCredentialDo.TableName()
	_dnsCredential.ALL = field.NewAsterisk(tableName)
	_dnsCredential.ID = field.NewInt(tableName, "id")
	_dnsCredential.CreatedAt = field.NewTime(tableName, "created_at")
	_dnsCredential.UpdatedAt = field.NewTime(tableName, "updated_at")
	_dnsCredential.DeletedAt = field.NewField(tableName, "deleted_at")
	_dnsCredential.Name = field.NewString(tableName, "name")
	_dnsCredential.Config = field.NewField(tableName, "config")
	_dnsCredential.Provider = field.NewString(tableName, "provider")

	_dnsCredential.fillFieldMap()

	return _dnsCredential
}

type dnsCredential struct {
	dnsCredentialDo

	ALL       field.Asterisk
	ID        field.Int
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	Name      field.String
	Config    field.Field
	Provider  field.String

	fieldMap map[string]field.Expr
}

func (d dnsCredential) Table(newTableName string) *dnsCredential {
	d.dnsCredentialDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d dnsCredential) As(alias string) *dnsCredential {
	d.dnsCredentialDo.DO = *(d.dnsCredentialDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *dnsCredential) updateTableName(table string) *dnsCredential {
	d.ALL = field.NewAsterisk(table)
	d.ID = field.NewInt(table, "id")
	d.CreatedAt = field.NewTime(table, "created_at")
	d.UpdatedAt = field.NewTime(table, "updated_at")
	d.DeletedAt = field.NewField(table, "deleted_at")
	d.Name = field.NewString(table, "name")
	d.Config = field.NewField(table, "config")
	d.Provider = field.NewString(table, "provider")

	d.fillFieldMap()

	return d
}

func (d *dnsCredential) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *dnsCredential) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 7)
	d.fieldMap["id"] = d.ID
	d.fieldMap["created_at"] = d.CreatedAt
	d.fieldMap["updated_at"] = d.UpdatedAt
	d.fieldMap["deleted_at"] = d.DeletedAt
	d.fieldMap["name"] = d.Name
	d.fieldMap["config"] = d.Config
	d.fieldMap["provider"] = d.Provider
}

func (d dnsCredential) clone(db *gorm.DB) dnsCredential {
	d.dnsCredentialDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d dnsCredential) replaceDB(db *gorm.DB) dnsCredential {
	d.dnsCredentialDo.ReplaceDB(db)
	return d
}

type dnsCredentialDo struct{ gen.DO }

// FirstByID Where("id=@id")
func (d dnsCredentialDo) FirstByID(id int) (result *model.DnsCredential, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("id=? ")

	var executeSQL *gorm.DB
	executeSQL = d.UnderlyingDB().Where(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// DeleteByID update @@table set deleted_at=strftime('%Y-%m-%d %H:%M:%S','now') where id=@id
func (d dnsCredentialDo) DeleteByID(id int) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("update dns_credentials set deleted_at=strftime('%Y-%m-%d %H:%M:%S','now') where id=? ")

	var executeSQL *gorm.DB
	executeSQL = d.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (d dnsCredentialDo) Debug() *dnsCredentialDo {
	return d.withDO(d.DO.Debug())
}

func (d dnsCredentialDo) WithContext(ctx context.Context) *dnsCredentialDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d dnsCredentialDo) ReadDB() *dnsCredentialDo {
	return d.Clauses(dbresolver.Read)
}

func (d dnsCredentialDo) WriteDB() *dnsCredentialDo {
	return d.Clauses(dbresolver.Write)
}

func (d dnsCredentialDo) Session(config *gorm.Session) *dnsCredentialDo {
	return d.withDO(d.DO.Session(config))
}

func (d dnsCredentialDo) Clauses(conds ...clause.Expression) *dnsCredentialDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d dnsCredentialDo) Returning(value interface{}, columns ...string) *dnsCredentialDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d dnsCredentialDo) Not(conds ...gen.Condition) *dnsCredentialDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d dnsCredentialDo) Or(conds ...gen.Condition) *dnsCredentialDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d dnsCredentialDo) Select(conds ...field.Expr) *dnsCredentialDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d dnsCredentialDo) Where(conds ...gen.Condition) *dnsCredentialDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d dnsCredentialDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *dnsCredentialDo {
	return d.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (d dnsCredentialDo) Order(conds ...field.Expr) *dnsCredentialDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d dnsCredentialDo) Distinct(cols ...field.Expr) *dnsCredentialDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d dnsCredentialDo) Omit(cols ...field.Expr) *dnsCredentialDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d dnsCredentialDo) Join(table schema.Tabler, on ...field.Expr) *dnsCredentialDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d dnsCredentialDo) LeftJoin(table schema.Tabler, on ...field.Expr) *dnsCredentialDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d dnsCredentialDo) RightJoin(table schema.Tabler, on ...field.Expr) *dnsCredentialDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d dnsCredentialDo) Group(cols ...field.Expr) *dnsCredentialDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d dnsCredentialDo) Having(conds ...gen.Condition) *dnsCredentialDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d dnsCredentialDo) Limit(limit int) *dnsCredentialDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d dnsCredentialDo) Offset(offset int) *dnsCredentialDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d dnsCredentialDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *dnsCredentialDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d dnsCredentialDo) Unscoped() *dnsCredentialDo {
	return d.withDO(d.DO.Unscoped())
}

func (d dnsCredentialDo) Create(values ...*model.DnsCredential) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d dnsCredentialDo) CreateInBatches(values []*model.DnsCredential, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d dnsCredentialDo) Save(values ...*model.DnsCredential) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d dnsCredentialDo) First() (*model.DnsCredential, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.DnsCredential), nil
	}
}

func (d dnsCredentialDo) Take() (*model.DnsCredential, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.DnsCredential), nil
	}
}

func (d dnsCredentialDo) Last() (*model.DnsCredential, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.DnsCredential), nil
	}
}

func (d dnsCredentialDo) Find() ([]*model.DnsCredential, error) {
	result, err := d.DO.Find()
	return result.([]*model.DnsCredential), err
}

func (d dnsCredentialDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DnsCredential, err error) {
	buf := make([]*model.DnsCredential, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d dnsCredentialDo) FindInBatches(result *[]*model.DnsCredential, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d dnsCredentialDo) Attrs(attrs ...field.AssignExpr) *dnsCredentialDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d dnsCredentialDo) Assign(attrs ...field.AssignExpr) *dnsCredentialDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d dnsCredentialDo) Joins(fields ...field.RelationField) *dnsCredentialDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d dnsCredentialDo) Preload(fields ...field.RelationField) *dnsCredentialDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d dnsCredentialDo) FirstOrInit() (*model.DnsCredential, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.DnsCredential), nil
	}
}

func (d dnsCredentialDo) FirstOrCreate() (*model.DnsCredential, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.DnsCredential), nil
	}
}

func (d dnsCredentialDo) FindByPage(offset int, limit int) (result []*model.DnsCredential, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d dnsCredentialDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d dnsCredentialDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d dnsCredentialDo) Delete(models ...*model.DnsCredential) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *dnsCredentialDo) withDO(do gen.Dao) *dnsCredentialDo {
	d.DO = *do.(*gen.DO)
	return d
}