// Code generated by go-queryset. DO NOT EDIT.
package models

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set MemosQuerySet

// MemosQuerySet is an queryset type for Memos
type MemosQuerySet struct {
	db *gorm.DB
}

// NewMemosQuerySet constructs new MemosQuerySet
func NewMemosQuerySet(db *gorm.DB) MemosQuerySet {
	return MemosQuerySet{
		db: db.Model(&Memos{}),
	}
}

func (qs MemosQuerySet) w(db *gorm.DB) MemosQuerySet {
	return NewMemosQuerySet(db)
}

func (qs MemosQuerySet) Select(fields ...MemosDBSchemaField) MemosQuerySet {
	names := []string{}
	for _, f := range fields {
		names = append(names, f.String())
	}

	return qs.w(qs.db.Select(strings.Join(names, ",")))
}

// Create is an autogenerated method
// nolint: dupl
func (o *Memos) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (o *Memos) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// All is an autogenerated method
// nolint: dupl
func (qs MemosQuerySet) All(ret *[]Memos) error {
	return qs.db.Find(ret).Error
}

// ContentEq is an autogenerated method
// nolint: dupl
func (qs MemosQuerySet) ContentEq(content string) MemosQuerySet {
	return qs.w(qs.db.Where("content = ?", content))
}

// ContentGt is an autogenerated method
// nolint: dupl
func (qs MemosQuerySet) ContentGt(content string) MemosQuerySet {
	return qs.w(qs.db.Where("content > ?", content))
}

// ContentGte is an autogenerated method
// nolint: dupl
func (qs MemosQuerySet) ContentGte(content string) MemosQuerySet {
	return qs.w(qs.db.Where("content >= ?", content))
}

// ContentIn is an autogenerated method
// nolint: dupl
func (qs MemosQuerySet) ContentIn(content ...string) MemosQuerySet {
	if len(content) == 0 {
		qs.db.AddError(errors.New("must at least pass one content in ContentIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("content IN (?)", content))
}

// ContentLike is an autogenerated method
// nolint: dupl
func (qs MemosQuerySet) ContentLike(content string) MemosQuerySet {
	return qs.w(qs.db.Where("content LIKE ?", content))
}

// ContentLt is an autogenerated method
// nolint: dupl
func (qs MemosQuerySet) ContentLt(content string) MemosQuerySet {
	return qs.w(qs.db.Where("content < ?", content))
}

// ContentLte is an autogenerated method
// nolint: dupl
func (qs MemosQuerySet) ContentLte(content string) MemosQuerySet {
	return qs.w(qs.db.Where("content <= ?", content))
}

// ContentNe is an autogenerated method
// nolint: dupl
func (qs MemosQuerySet) ContentNe(content string) MemosQuerySet {
	return qs.w(qs.db.Where("content != ?", content))
}

// ContentNotIn is an autogenerated method
// nolint: dupl
func (qs MemosQuerySet) ContentNotIn(content ...string) MemosQuerySet {
	if len(content) == 0 {
		qs.db.AddError(errors.New("must at least pass one content in ContentNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("content NOT IN (?)", content))
}

// ContentNotlike is an autogenerated method
// nolint: dupl
func (qs MemosQuerySet) ContentNotlike(content string) MemosQuerySet {
	return qs.w(qs.db.Where("content NOT LIKE ?", content))
}

// Count is an autogenerated method
// nolint: dupl
func (qs MemosQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// CreatedAtEq is an autogenerated method
// nolint: dupl
func (qs MemosQuerySet) CreatedAtEq(createdAt time.Time) MemosQuerySet {
	return qs.w(qs.db.Where("created_at = ?", createdAt))
}

// CreatedAtGt is an autogenerated method
// nolint: dupl
func (qs MemosQuerySet) CreatedAtGt(createdAt time.Time) MemosQuerySet {
	return qs.w(qs.db.Where("created_at > ?", createdAt))
}

// CreatedAtGte is an autogenerated method
// nolint: dupl
func (qs MemosQuerySet) CreatedAtGte(createdAt time.Time) MemosQuerySet {
	return qs.w(qs.db.Where("created_at >= ?", createdAt))
}

// CreatedAtLt is an autogenerated method
// nolint: dupl
func (qs MemosQuerySet) CreatedAtLt(createdAt time.Time) MemosQuerySet {
	return qs.w(qs.db.Where("created_at < ?", createdAt))
}

// CreatedAtLte is an autogenerated method
// nolint: dupl
func (qs MemosQuerySet) CreatedAtLte(createdAt time.Time) MemosQuerySet {
	return qs.w(qs.db.Where("created_at <= ?", createdAt))
}

// CreatedAtNe is an autogenerated method
// nolint: dupl
func (qs MemosQuerySet) CreatedAtNe(createdAt time.Time) MemosQuerySet {
	return qs.w(qs.db.Where("created_at != ?", createdAt))
}

// Delete is an autogenerated method
// nolint: dupl
func (qs MemosQuerySet) Delete() error {
	return qs.db.Delete(Memos{}).Error
}

// DeleteNum is an autogenerated method
// nolint: dupl
func (qs MemosQuerySet) DeleteNum() (int64, error) {
	db := qs.db.Delete(Memos{})
	return db.RowsAffected, db.Error
}

// DeleteNumUnscoped is an autogenerated method
// nolint: dupl
func (qs MemosQuerySet) DeleteNumUnscoped() (int64, error) {
	db := qs.db.Unscoped().Delete(Memos{})
	return db.RowsAffected, db.Error
}

// GetDB is an autogenerated method
// nolint: dupl
func (qs MemosQuerySet) GetDB() *gorm.DB {
	return qs.db
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs MemosQuerySet) GetUpdater() MemosUpdater {
	return NewMemosUpdater(qs.db)
}

// IDEq is an autogenerated method
// nolint: dupl
func (qs MemosQuerySet) IDEq(ID uint) MemosQuerySet {
	return qs.w(qs.db.Where("id = ?", ID))
}

// IDGt is an autogenerated method
// nolint: dupl
func (qs MemosQuerySet) IDGt(ID uint) MemosQuerySet {
	return qs.w(qs.db.Where("id > ?", ID))
}

// IDGte is an autogenerated method
// nolint: dupl
func (qs MemosQuerySet) IDGte(ID uint) MemosQuerySet {
	return qs.w(qs.db.Where("id >= ?", ID))
}

// IDIn is an autogenerated method
// nolint: dupl
func (qs MemosQuerySet) IDIn(ID ...uint) MemosQuerySet {
	if len(ID) == 0 {
		qs.db.AddError(errors.New("must at least pass one ID in IDIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id IN (?)", ID))
}

// IDLt is an autogenerated method
// nolint: dupl
func (qs MemosQuerySet) IDLt(ID uint) MemosQuerySet {
	return qs.w(qs.db.Where("id < ?", ID))
}

// IDLte is an autogenerated method
// nolint: dupl
func (qs MemosQuerySet) IDLte(ID uint) MemosQuerySet {
	return qs.w(qs.db.Where("id <= ?", ID))
}

// IDNe is an autogenerated method
// nolint: dupl
func (qs MemosQuerySet) IDNe(ID uint) MemosQuerySet {
	return qs.w(qs.db.Where("id != ?", ID))
}

// IDNotIn is an autogenerated method
// nolint: dupl
func (qs MemosQuerySet) IDNotIn(ID ...uint) MemosQuerySet {
	if len(ID) == 0 {
		qs.db.AddError(errors.New("must at least pass one ID in IDNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id NOT IN (?)", ID))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs MemosQuerySet) Limit(limit int) MemosQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// Offset is an autogenerated method
// nolint: dupl
func (qs MemosQuerySet) Offset(offset int) MemosQuerySet {
	return qs.w(qs.db.Offset(offset))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs MemosQuerySet) One(ret *Memos) error {
	return qs.db.First(ret).Error
}

// OrderAscByContent is an autogenerated method
// nolint: dupl
func (qs MemosQuerySet) OrderAscByContent() MemosQuerySet {
	return qs.w(qs.db.Order("content ASC"))
}

// OrderAscByCreatedAt is an autogenerated method
// nolint: dupl
func (qs MemosQuerySet) OrderAscByCreatedAt() MemosQuerySet {
	return qs.w(qs.db.Order("created_at ASC"))
}

// OrderAscByID is an autogenerated method
// nolint: dupl
func (qs MemosQuerySet) OrderAscByID() MemosQuerySet {
	return qs.w(qs.db.Order("id ASC"))
}

// OrderDescByContent is an autogenerated method
// nolint: dupl
func (qs MemosQuerySet) OrderDescByContent() MemosQuerySet {
	return qs.w(qs.db.Order("content DESC"))
}

// OrderDescByCreatedAt is an autogenerated method
// nolint: dupl
func (qs MemosQuerySet) OrderDescByCreatedAt() MemosQuerySet {
	return qs.w(qs.db.Order("created_at DESC"))
}

// OrderDescByID is an autogenerated method
// nolint: dupl
func (qs MemosQuerySet) OrderDescByID() MemosQuerySet {
	return qs.w(qs.db.Order("id DESC"))
}

// SetContent is an autogenerated method
// nolint: dupl
func (u MemosUpdater) SetContent(content string) MemosUpdater {
	u.fields[string(MemosDBSchema.Content)] = content
	return u
}

// SetCreatedAt is an autogenerated method
// nolint: dupl
func (u MemosUpdater) SetCreatedAt(createdAt time.Time) MemosUpdater {
	u.fields[string(MemosDBSchema.CreatedAt)] = createdAt
	return u
}

// SetID is an autogenerated method
// nolint: dupl
func (u MemosUpdater) SetID(ID uint) MemosUpdater {
	u.fields[string(MemosDBSchema.ID)] = ID
	return u
}

// Update is an autogenerated method
// nolint: dupl
func (u MemosUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u MemosUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// ===== END of query set MemosQuerySet

// ===== BEGIN of Memos modifiers

// MemosDBSchemaField describes database schema field. It requires for method 'Update'
type MemosDBSchemaField string

// String method returns string representation of field.
// nolint: dupl
func (f MemosDBSchemaField) String() string {
	return string(f)
}

// MemosDBSchema stores db field names of Memos
var MemosDBSchema = struct {
	ID        MemosDBSchemaField
	Content   MemosDBSchemaField
	CreatedAt MemosDBSchemaField
}{

	ID:        MemosDBSchemaField("id"),
	Content:   MemosDBSchemaField("content"),
	CreatedAt: MemosDBSchemaField("created_at"),
}

// Update updates Memos fields by primary key
// nolint: dupl
func (o *Memos) Update(db *gorm.DB, fields ...MemosDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"id":         o.ID,
		"content":    o.Content,
		"created_at": o.CreatedAt,
	}
	u := map[string]interface{}{}
	for _, f := range fields {
		fs := f.String()
		u[fs] = dbNameToFieldName[fs]
	}
	if err := db.Model(o).Updates(u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}

		return fmt.Errorf("can't update Memos %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// MemosUpdater is an Memos updates manager
type MemosUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewMemosUpdater creates new Memos updater
// nolint: dupl
func NewMemosUpdater(db *gorm.DB) MemosUpdater {
	return MemosUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&Memos{}),
	}
}

// ===== END of Memos modifiers

// ===== END of all query sets