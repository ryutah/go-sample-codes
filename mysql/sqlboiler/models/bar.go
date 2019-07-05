// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Bar is an object representing the database table.
type Bar struct {
	ID    int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name  null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	FooID int         `boil:"foo_id" json:"foo_id" toml:"foo_id" yaml:"foo_id"`

	R *barR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L barL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BarColumns = struct {
	ID    string
	Name  string
	FooID string
}{
	ID:    "id",
	Name:  "name",
	FooID: "foo_id",
}

// BarRels is where relationship names are stored.
var BarRels = struct {
	Foo            string
	BarChildren    string
	BarOneChildren string
}{
	Foo:            "Foo",
	BarChildren:    "BarChildren",
	BarOneChildren: "BarOneChildren",
}

// barR is where relationships are stored.
type barR struct {
	Foo            *Foo
	BarChildren    BarChildSlice
	BarOneChildren BarOneChildSlice
}

// NewStruct creates a new relationship struct
func (*barR) NewStruct() *barR {
	return &barR{}
}

// barL is where Load methods for each relationship are stored.
type barL struct{}

var (
	barColumns               = []string{"id", "name", "foo_id"}
	barColumnsWithoutDefault = []string{"name", "foo_id"}
	barColumnsWithDefault    = []string{"id"}
	barPrimaryKeyColumns     = []string{"id"}
)

type (
	// BarSlice is an alias for a slice of pointers to Bar.
	// This should generally be used opposed to []Bar.
	BarSlice []*Bar
	// BarHook is the signature for custom Bar hook methods
	BarHook func(context.Context, boil.ContextExecutor, *Bar) error

	barQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	barType                 = reflect.TypeOf(&Bar{})
	barMapping              = queries.MakeStructMapping(barType)
	barPrimaryKeyMapping, _ = queries.BindMapping(barType, barMapping, barPrimaryKeyColumns)
	barInsertCacheMut       sync.RWMutex
	barInsertCache          = make(map[string]insertCache)
	barUpdateCacheMut       sync.RWMutex
	barUpdateCache          = make(map[string]updateCache)
	barUpsertCacheMut       sync.RWMutex
	barUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var barBeforeInsertHooks []BarHook
var barBeforeUpdateHooks []BarHook
var barBeforeDeleteHooks []BarHook
var barBeforeUpsertHooks []BarHook

var barAfterInsertHooks []BarHook
var barAfterSelectHooks []BarHook
var barAfterUpdateHooks []BarHook
var barAfterDeleteHooks []BarHook
var barAfterUpsertHooks []BarHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Bar) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range barBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Bar) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range barBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Bar) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range barBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Bar) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range barBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Bar) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range barAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Bar) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range barAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Bar) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range barAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Bar) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range barAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Bar) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range barAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddBarHook registers your hook function for all future operations.
func AddBarHook(hookPoint boil.HookPoint, barHook BarHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		barBeforeInsertHooks = append(barBeforeInsertHooks, barHook)
	case boil.BeforeUpdateHook:
		barBeforeUpdateHooks = append(barBeforeUpdateHooks, barHook)
	case boil.BeforeDeleteHook:
		barBeforeDeleteHooks = append(barBeforeDeleteHooks, barHook)
	case boil.BeforeUpsertHook:
		barBeforeUpsertHooks = append(barBeforeUpsertHooks, barHook)
	case boil.AfterInsertHook:
		barAfterInsertHooks = append(barAfterInsertHooks, barHook)
	case boil.AfterSelectHook:
		barAfterSelectHooks = append(barAfterSelectHooks, barHook)
	case boil.AfterUpdateHook:
		barAfterUpdateHooks = append(barAfterUpdateHooks, barHook)
	case boil.AfterDeleteHook:
		barAfterDeleteHooks = append(barAfterDeleteHooks, barHook)
	case boil.AfterUpsertHook:
		barAfterUpsertHooks = append(barAfterUpsertHooks, barHook)
	}
}

// One returns a single bar record from the query.
func (q barQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Bar, error) {
	o := &Bar{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for bar")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Bar records from the query.
func (q barQuery) All(ctx context.Context, exec boil.ContextExecutor) (BarSlice, error) {
	var o []*Bar

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Bar slice")
	}

	if len(barAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Bar records in the query.
func (q barQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count bar rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q barQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if bar exists")
	}

	return count > 0, nil
}

// Foo pointed to by the foreign key.
func (o *Bar) Foo(mods ...qm.QueryMod) fooQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.FooID),
	}

	queryMods = append(queryMods, mods...)

	query := Foos(queryMods...)
	queries.SetFrom(query.Query, "`foo`")

	return query
}

// BarChildren retrieves all the bar_child's BarChildren with an executor.
func (o *Bar) BarChildren(mods ...qm.QueryMod) barChildQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`bar_child`.`bar_id`=?", o.ID),
	)

	query := BarChildren(queryMods...)
	queries.SetFrom(query.Query, "`bar_child`")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"`bar_child`.*"})
	}

	return query
}

// BarOneChildren retrieves all the bar_one_child's BarOneChildren with an executor.
func (o *Bar) BarOneChildren(mods ...qm.QueryMod) barOneChildQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`bar_one_child`.`bar_id`=?", o.ID),
	)

	query := BarOneChildren(queryMods...)
	queries.SetFrom(query.Query, "`bar_one_child`")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"`bar_one_child`.*"})
	}

	return query
}

// LoadFoo allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (barL) LoadFoo(ctx context.Context, e boil.ContextExecutor, singular bool, maybeBar interface{}, mods queries.Applicator) error {
	var slice []*Bar
	var object *Bar

	if singular {
		object = maybeBar.(*Bar)
	} else {
		slice = *maybeBar.(*[]*Bar)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &barR{}
		}
		args = append(args, object.FooID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &barR{}
			}

			for _, a := range args {
				if a == obj.FooID {
					continue Outer
				}
			}

			args = append(args, obj.FooID)

		}
	}

	query := NewQuery(qm.From(`foo`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Foo")
	}

	var resultSlice []*Foo
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Foo")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for foo")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for foo")
	}

	if len(barAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Foo = foreign
		if foreign.R == nil {
			foreign.R = &fooR{}
		}
		foreign.R.Bars = append(foreign.R.Bars, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.FooID == foreign.ID {
				local.R.Foo = foreign
				if foreign.R == nil {
					foreign.R = &fooR{}
				}
				foreign.R.Bars = append(foreign.R.Bars, local)
				break
			}
		}
	}

	return nil
}

// LoadBarChildren allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (barL) LoadBarChildren(ctx context.Context, e boil.ContextExecutor, singular bool, maybeBar interface{}, mods queries.Applicator) error {
	var slice []*Bar
	var object *Bar

	if singular {
		object = maybeBar.(*Bar)
	} else {
		slice = *maybeBar.(*[]*Bar)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &barR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &barR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	query := NewQuery(qm.From(`bar_child`), qm.WhereIn(`bar_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load bar_child")
	}

	var resultSlice []*BarChild
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice bar_child")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on bar_child")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for bar_child")
	}

	if len(barChildAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.BarChildren = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &barChildR{}
			}
			foreign.R.Bar = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.BarID {
				local.R.BarChildren = append(local.R.BarChildren, foreign)
				if foreign.R == nil {
					foreign.R = &barChildR{}
				}
				foreign.R.Bar = local
				break
			}
		}
	}

	return nil
}

// LoadBarOneChildren allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (barL) LoadBarOneChildren(ctx context.Context, e boil.ContextExecutor, singular bool, maybeBar interface{}, mods queries.Applicator) error {
	var slice []*Bar
	var object *Bar

	if singular {
		object = maybeBar.(*Bar)
	} else {
		slice = *maybeBar.(*[]*Bar)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &barR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &barR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	query := NewQuery(qm.From(`bar_one_child`), qm.WhereIn(`bar_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load bar_one_child")
	}

	var resultSlice []*BarOneChild
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice bar_one_child")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on bar_one_child")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for bar_one_child")
	}

	if len(barOneChildAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.BarOneChildren = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &barOneChildR{}
			}
			foreign.R.Bar = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.BarID {
				local.R.BarOneChildren = append(local.R.BarOneChildren, foreign)
				if foreign.R == nil {
					foreign.R = &barOneChildR{}
				}
				foreign.R.Bar = local
				break
			}
		}
	}

	return nil
}

// SetFoo of the bar to the related item.
// Sets o.R.Foo to related.
// Adds o to related.R.Bars.
func (o *Bar) SetFoo(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Foo) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `bar` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"foo_id"}),
		strmangle.WhereClause("`", "`", 0, barPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.FooID = related.ID
	if o.R == nil {
		o.R = &barR{
			Foo: related,
		}
	} else {
		o.R.Foo = related
	}

	if related.R == nil {
		related.R = &fooR{
			Bars: BarSlice{o},
		}
	} else {
		related.R.Bars = append(related.R.Bars, o)
	}

	return nil
}

// AddBarChildren adds the given related objects to the existing relationships
// of the bar, optionally inserting them as new records.
// Appends related to o.R.BarChildren.
// Sets related.R.Bar appropriately.
func (o *Bar) AddBarChildren(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*BarChild) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.BarID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `bar_child` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"bar_id"}),
				strmangle.WhereClause("`", "`", 0, barChildPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.BarID = o.ID
		}
	}

	if o.R == nil {
		o.R = &barR{
			BarChildren: related,
		}
	} else {
		o.R.BarChildren = append(o.R.BarChildren, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &barChildR{
				Bar: o,
			}
		} else {
			rel.R.Bar = o
		}
	}
	return nil
}

// AddBarOneChildren adds the given related objects to the existing relationships
// of the bar, optionally inserting them as new records.
// Appends related to o.R.BarOneChildren.
// Sets related.R.Bar appropriately.
func (o *Bar) AddBarOneChildren(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*BarOneChild) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.BarID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `bar_one_child` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"bar_id"}),
				strmangle.WhereClause("`", "`", 0, barOneChildPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.BarID = o.ID
		}
	}

	if o.R == nil {
		o.R = &barR{
			BarOneChildren: related,
		}
	} else {
		o.R.BarOneChildren = append(o.R.BarOneChildren, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &barOneChildR{
				Bar: o,
			}
		} else {
			rel.R.Bar = o
		}
	}
	return nil
}

// Bars retrieves all the records using an executor.
func Bars(mods ...qm.QueryMod) barQuery {
	mods = append(mods, qm.From("`bar`"))
	return barQuery{NewQuery(mods...)}
}

// FindBar retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBar(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Bar, error) {
	barObj := &Bar{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `bar` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, barObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from bar")
	}

	return barObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Bar) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no bar provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(barColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	barInsertCacheMut.RLock()
	cache, cached := barInsertCache[key]
	barInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			barColumns,
			barColumnsWithDefault,
			barColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(barType, barMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(barType, barMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `bar` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `bar` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `bar` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, barPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into bar")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == barMapping["ID"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for bar")
	}

CacheNoHooks:
	if !cached {
		barInsertCacheMut.Lock()
		barInsertCache[key] = cache
		barInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Bar.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Bar) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	barUpdateCacheMut.RLock()
	cache, cached := barUpdateCache[key]
	barUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			barColumns,
			barPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update bar, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `bar` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, barPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(barType, barMapping, append(wl, barPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update bar row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for bar")
	}

	if !cached {
		barUpdateCacheMut.Lock()
		barUpdateCache[key] = cache
		barUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q barQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for bar")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for bar")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BarSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), barPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `bar` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, barPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in bar slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all bar")
	}
	return rowsAff, nil
}

var mySQLBarUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Bar) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no bar provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(barColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLBarUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	barUpsertCacheMut.RLock()
	cache, cached := barUpsertCache[key]
	barUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			barColumns,
			barColumnsWithDefault,
			barColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			barColumns,
			barPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert bar, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "bar", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `bar` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(barType, barMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(barType, barMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for bar")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == barMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(barType, barMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for bar")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for bar")
	}

CacheNoHooks:
	if !cached {
		barUpsertCacheMut.Lock()
		barUpsertCache[key] = cache
		barUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Bar record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Bar) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Bar provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), barPrimaryKeyMapping)
	sql := "DELETE FROM `bar` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from bar")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for bar")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q barQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no barQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from bar")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for bar")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BarSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Bar slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(barBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), barPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `bar` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, barPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from bar slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for bar")
	}

	if len(barAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Bar) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBar(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BarSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BarSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), barPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `bar`.* FROM `bar` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, barPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BarSlice")
	}

	*o = slice

	return nil
}

// BarExists checks if the Bar row exists.
func BarExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `bar` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if bar exists")
	}

	return exists, nil
}