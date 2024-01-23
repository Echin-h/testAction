// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"gin-rush-template/internal/model"
)

func newJob(db *gorm.DB, opts ...gen.DOOption) job {
	_job := job{}

	_job.jobDo.UseDB(db, opts...)
	_job.jobDo.UseModel(&model.Job{})

	tableName := _job.jobDo.TableName()
	_job.ALL = field.NewAsterisk(tableName)
	_job.ID = field.NewUint(tableName, "id")
	_job.CreatedAt = field.NewTime(tableName, "created_at")
	_job.UpdatedAt = field.NewTime(tableName, "updated_at")
	_job.DeletedAt = field.NewField(tableName, "deleted_at")
	_job.Name = field.NewString(tableName, "name")
	_job.Description = field.NewString(tableName, "description")
	_job.Applications = jobHasManyApplications{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Applications", "model.Application"),
		Interview: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Applications.Interview", "model.Interview"),
		},
	}

	_job.Slots = jobManyToManySlots{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Slots", "model.Slot"),
		Interviews: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Slots.Interviews", "model.Interview"),
		},
		Jobs: struct {
			field.RelationField
			Applications struct {
				field.RelationField
			}
			Slots struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("Slots.Jobs", "model.Job"),
			Applications: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Slots.Jobs.Applications", "model.Application"),
			},
			Slots: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Slots.Jobs.Slots", "model.Slot"),
			},
		},
	}

	_job.fillFieldMap()

	return _job
}

type job struct {
	jobDo

	ALL          field.Asterisk
	ID           field.Uint
	CreatedAt    field.Time
	UpdatedAt    field.Time
	DeletedAt    field.Field
	Name         field.String
	Description  field.String
	Applications jobHasManyApplications

	Slots jobManyToManySlots

	fieldMap map[string]field.Expr
}

func (j job) Table(newTableName string) *job {
	j.jobDo.UseTable(newTableName)
	return j.updateTableName(newTableName)
}

func (j job) As(alias string) *job {
	j.jobDo.DO = *(j.jobDo.As(alias).(*gen.DO))
	return j.updateTableName(alias)
}

func (j *job) updateTableName(table string) *job {
	j.ALL = field.NewAsterisk(table)
	j.ID = field.NewUint(table, "id")
	j.CreatedAt = field.NewTime(table, "created_at")
	j.UpdatedAt = field.NewTime(table, "updated_at")
	j.DeletedAt = field.NewField(table, "deleted_at")
	j.Name = field.NewString(table, "name")
	j.Description = field.NewString(table, "description")

	j.fillFieldMap()

	return j
}

func (j *job) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := j.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (j *job) fillFieldMap() {
	j.fieldMap = make(map[string]field.Expr, 8)
	j.fieldMap["id"] = j.ID
	j.fieldMap["created_at"] = j.CreatedAt
	j.fieldMap["updated_at"] = j.UpdatedAt
	j.fieldMap["deleted_at"] = j.DeletedAt
	j.fieldMap["name"] = j.Name
	j.fieldMap["description"] = j.Description

}

func (j job) clone(db *gorm.DB) job {
	j.jobDo.ReplaceConnPool(db.Statement.ConnPool)
	return j
}

func (j job) replaceDB(db *gorm.DB) job {
	j.jobDo.ReplaceDB(db)
	return j
}

type jobHasManyApplications struct {
	db *gorm.DB

	field.RelationField

	Interview struct {
		field.RelationField
	}
}

func (a jobHasManyApplications) Where(conds ...field.Expr) *jobHasManyApplications {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a jobHasManyApplications) WithContext(ctx context.Context) *jobHasManyApplications {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a jobHasManyApplications) Session(session *gorm.Session) *jobHasManyApplications {
	a.db = a.db.Session(session)
	return &a
}

func (a jobHasManyApplications) Model(m *model.Job) *jobHasManyApplicationsTx {
	return &jobHasManyApplicationsTx{a.db.Model(m).Association(a.Name())}
}

type jobHasManyApplicationsTx struct{ tx *gorm.Association }

func (a jobHasManyApplicationsTx) Find() (result []*model.Application, err error) {
	return result, a.tx.Find(&result)
}

func (a jobHasManyApplicationsTx) Append(values ...*model.Application) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a jobHasManyApplicationsTx) Replace(values ...*model.Application) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a jobHasManyApplicationsTx) Delete(values ...*model.Application) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a jobHasManyApplicationsTx) Clear() error {
	return a.tx.Clear()
}

func (a jobHasManyApplicationsTx) Count() int64 {
	return a.tx.Count()
}

type jobManyToManySlots struct {
	db *gorm.DB

	field.RelationField

	Interviews struct {
		field.RelationField
	}
	Jobs struct {
		field.RelationField
		Applications struct {
			field.RelationField
		}
		Slots struct {
			field.RelationField
		}
	}
}

func (a jobManyToManySlots) Where(conds ...field.Expr) *jobManyToManySlots {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a jobManyToManySlots) WithContext(ctx context.Context) *jobManyToManySlots {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a jobManyToManySlots) Session(session *gorm.Session) *jobManyToManySlots {
	a.db = a.db.Session(session)
	return &a
}

func (a jobManyToManySlots) Model(m *model.Job) *jobManyToManySlotsTx {
	return &jobManyToManySlotsTx{a.db.Model(m).Association(a.Name())}
}

type jobManyToManySlotsTx struct{ tx *gorm.Association }

func (a jobManyToManySlotsTx) Find() (result []*model.Slot, err error) {
	return result, a.tx.Find(&result)
}

func (a jobManyToManySlotsTx) Append(values ...*model.Slot) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a jobManyToManySlotsTx) Replace(values ...*model.Slot) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a jobManyToManySlotsTx) Delete(values ...*model.Slot) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a jobManyToManySlotsTx) Clear() error {
	return a.tx.Clear()
}

func (a jobManyToManySlotsTx) Count() int64 {
	return a.tx.Count()
}

type jobDo struct{ gen.DO }

func (j jobDo) Debug() *jobDo {
	return j.withDO(j.DO.Debug())
}

func (j jobDo) WithContext(ctx context.Context) *jobDo {
	return j.withDO(j.DO.WithContext(ctx))
}

func (j jobDo) ReadDB() *jobDo {
	return j.Clauses(dbresolver.Read)
}

func (j jobDo) WriteDB() *jobDo {
	return j.Clauses(dbresolver.Write)
}

func (j jobDo) Session(config *gorm.Session) *jobDo {
	return j.withDO(j.DO.Session(config))
}

func (j jobDo) Clauses(conds ...clause.Expression) *jobDo {
	return j.withDO(j.DO.Clauses(conds...))
}

func (j jobDo) Returning(value interface{}, columns ...string) *jobDo {
	return j.withDO(j.DO.Returning(value, columns...))
}

func (j jobDo) Not(conds ...gen.Condition) *jobDo {
	return j.withDO(j.DO.Not(conds...))
}

func (j jobDo) Or(conds ...gen.Condition) *jobDo {
	return j.withDO(j.DO.Or(conds...))
}

func (j jobDo) Select(conds ...field.Expr) *jobDo {
	return j.withDO(j.DO.Select(conds...))
}

func (j jobDo) Where(conds ...gen.Condition) *jobDo {
	return j.withDO(j.DO.Where(conds...))
}

func (j jobDo) Order(conds ...field.Expr) *jobDo {
	return j.withDO(j.DO.Order(conds...))
}

func (j jobDo) Distinct(cols ...field.Expr) *jobDo {
	return j.withDO(j.DO.Distinct(cols...))
}

func (j jobDo) Omit(cols ...field.Expr) *jobDo {
	return j.withDO(j.DO.Omit(cols...))
}

func (j jobDo) Join(table schema.Tabler, on ...field.Expr) *jobDo {
	return j.withDO(j.DO.Join(table, on...))
}

func (j jobDo) LeftJoin(table schema.Tabler, on ...field.Expr) *jobDo {
	return j.withDO(j.DO.LeftJoin(table, on...))
}

func (j jobDo) RightJoin(table schema.Tabler, on ...field.Expr) *jobDo {
	return j.withDO(j.DO.RightJoin(table, on...))
}

func (j jobDo) Group(cols ...field.Expr) *jobDo {
	return j.withDO(j.DO.Group(cols...))
}

func (j jobDo) Having(conds ...gen.Condition) *jobDo {
	return j.withDO(j.DO.Having(conds...))
}

func (j jobDo) Limit(limit int) *jobDo {
	return j.withDO(j.DO.Limit(limit))
}

func (j jobDo) Offset(offset int) *jobDo {
	return j.withDO(j.DO.Offset(offset))
}

func (j jobDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *jobDo {
	return j.withDO(j.DO.Scopes(funcs...))
}

func (j jobDo) Unscoped() *jobDo {
	return j.withDO(j.DO.Unscoped())
}

func (j jobDo) Create(values ...*model.Job) error {
	if len(values) == 0 {
		return nil
	}
	return j.DO.Create(values)
}

func (j jobDo) CreateInBatches(values []*model.Job, batchSize int) error {
	return j.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (j jobDo) Save(values ...*model.Job) error {
	if len(values) == 0 {
		return nil
	}
	return j.DO.Save(values)
}

func (j jobDo) First() (*model.Job, error) {
	if result, err := j.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Job), nil
	}
}

func (j jobDo) Take() (*model.Job, error) {
	if result, err := j.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Job), nil
	}
}

func (j jobDo) Last() (*model.Job, error) {
	if result, err := j.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Job), nil
	}
}

func (j jobDo) Find() ([]*model.Job, error) {
	result, err := j.DO.Find()
	return result.([]*model.Job), err
}

func (j jobDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Job, err error) {
	buf := make([]*model.Job, 0, batchSize)
	err = j.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (j jobDo) FindInBatches(result *[]*model.Job, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return j.DO.FindInBatches(result, batchSize, fc)
}

func (j jobDo) Attrs(attrs ...field.AssignExpr) *jobDo {
	return j.withDO(j.DO.Attrs(attrs...))
}

func (j jobDo) Assign(attrs ...field.AssignExpr) *jobDo {
	return j.withDO(j.DO.Assign(attrs...))
}

func (j jobDo) Joins(fields ...field.RelationField) *jobDo {
	for _, _f := range fields {
		j = *j.withDO(j.DO.Joins(_f))
	}
	return &j
}

func (j jobDo) Preload(fields ...field.RelationField) *jobDo {
	for _, _f := range fields {
		j = *j.withDO(j.DO.Preload(_f))
	}
	return &j
}

func (j jobDo) FirstOrInit() (*model.Job, error) {
	if result, err := j.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Job), nil
	}
}

func (j jobDo) FirstOrCreate() (*model.Job, error) {
	if result, err := j.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Job), nil
	}
}

func (j jobDo) FindByPage(offset int, limit int) (result []*model.Job, count int64, err error) {
	result, err = j.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = j.Offset(-1).Limit(-1).Count()
	return
}

func (j jobDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = j.Count()
	if err != nil {
		return
	}

	err = j.Offset(offset).Limit(limit).Scan(result)
	return
}

func (j jobDo) Scan(result interface{}) (err error) {
	return j.DO.Scan(result)
}

func (j jobDo) Delete(models ...*model.Job) (result gen.ResultInfo, err error) {
	return j.DO.Delete(models)
}

func (j *jobDo) withDO(do gen.Dao) *jobDo {
	j.DO = *do.(*gen.DO)
	return j
}