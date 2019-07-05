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

// BarChild is an object representing the database table.
type BarChild struct {
	ID    int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name  null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	BarID int         `boil:"bar_id" json:"bar_id" toml:"bar_id" yaml:"bar_id"`

	R *barChildR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L barChildL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BarChildColumns = struct {
	ID    string
	Name  string
	BarID string
}{
	ID:    "id",
	Name:  "name",
	BarID: "bar_id",
}

// BarChildRels is where relationship names are stored.
var BarChildRels = struct {
	Bar string
}{
	Bar: "Bar",
}

// barChildR is where relationships are stored.
type barChildR struct {
	Bar *Bar
}

// NewStruct creates a new relationship struct
func (*barChildR) NewStruct() *barChildR {
	return &barChildR{}
}

// barChildL is where Load methods for each relationship are stored.
type barChildL struct{}

var (
	barChildColumns               = []string{"id", "name", "bar_id"}
	barChildColumnsWithoutDefault = []string{"name", "bar_id"}
	barChildColumnsWithDefault    = []string{"id"}
	barChildPrimaryKeyColumns     = []string{"id"}
)

type (
	// BarChildSlice is an alias for a slice of pointers to BarChild.
	// This should generally be used opposed to []BarChild.
	BarChildSlice []*BarChild
	// BarChildHook is the signature for custom BarChild hook methods
	BarChildHook func(context.Context, boil.ContextExecutor, *BarChild) error

	barChildQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	barChildType                 = reflect.TypeOf(&BarChild{})
	barChildMapping              = queries.MakeStructMapping(barChildType)
	barChildPrimaryKeyMapping, _ = queries.BindMapping(barChildType, barChildMapping, barChildPrimaryKeyColumns)
	barChildInsertCacheMut       sync.RWMutex
	barChildInsertCache          = make(map[string]insertCache)
	barChildUpdateCacheMut       sync.RWMutex
	barChildUpdateCache          = make(map[string]updateCache)
	barChildUpsertCacheMut       sync.RWMutex
	barChildUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var barChildBeforeInsertHooks []BarChildHook
var barChildBeforeUpdateHooks []BarChildHook
var barChildBeforeDeleteHooks []BarChildHook
var barChildBeforeUpsertHooks []BarChildHook

var barChildAfterInsertHooks []BarChildHook
var barChildAfterSelectHooks []BarChildHook
var barChildAfterUpdateHooks []BarChildHook
var barChildAfterDeleteHooks []BarChildHook
var barChildAfterUpsertHooks []BarChildHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *BarChild) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range barChildBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *BarChild) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range barChildBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *BarChild) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range barChildBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *BarChild) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range barChildBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *BarChild) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range barChildAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *BarChild) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range barChildAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *BarChild) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range barChildAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *BarChild) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range barChildAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *BarChild) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range barChildAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddBarChildHook registers your hook function for all future operations.
func AddBarChildHook(hookPoint boil.HookPoint, barChildHook BarChildHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		barChildBeforeInsertHooks = append(barChildBeforeInsertHooks, barChildHook)
	case boil.BeforeUpdateHook:
		barChildBeforeUpdateHooks = append(barChildBeforeUpdateHooks, barChildHook)
	case boil.BeforeDeleteHook:
		barChildBeforeDeleteHooks = append(barChildBeforeDeleteHooks, barChildHook)
	case boil.BeforeUpsertHook:
		barChildBeforeUpsertHooks = append(barChildBeforeUpsertHooks, barChildHook)
	case boil.AfterInsertHook:
		barChildAfterInsertHooks = append(barChildAfterInsertHooks, barChildHook)
	case boil.AfterSelectHook:
		barChildAfterSelectHooks = append(barChildAfterSelectHooks, barChildHook)
	case boil.AfterUpdateHook:
		barChildAfterUpdateHooks = append(barChildAfterUpdateHooks, barChildHook)
	case boil.AfterDeleteHook:
		barChildAfterDeleteHooks = append(barChildAfterDeleteHooks, barChildHook)
	case boil.AfterUpsertHook:
		barChildAfterUpsertHooks = append(barChildAfterUpsertHooks, barChildHook)
	}
}

// One returns a single barChild record from the query.
func (q barChildQuery) One(ctx context.Context, exec boil.ContextExecutor) (*BarChild, error) {
	o := &BarChild{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for bar_child")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all BarChild records from the query.
func (q barChildQuery) All(ctx context.Context, exec boil.ContextExecutor) (BarChildSlice, error) {
	var o []*BarChild

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to BarChild slice")
	}

	if len(barChildAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all BarChild records in the query.
func (q barChildQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count bar_child rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q barChildQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if bar_child exists")
	}

	return count > 0, nil
}

// Bar pointed to by the foreign key.
func (o *BarChild) Bar(mods ...qm.QueryMod) barQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.BarID),
	}

	queryMods = append(queryMods, mods...)

	query := Bars(queryMods...)
	queries.SetFrom(query.Query, "`bar`")

	return query
}

// LoadBar allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (barChildL) LoadBar(ctx context.Context, e boil.ContextExecutor, singular bool, maybeBarChild interface{}, mods queries.Applicator) error {
	var slice []*BarChild
	var object *BarChild

	if singular {
		object = maybeBarChild.(*BarChild)
	} else {
		slice = *maybeBarChild.(*[]*BarChild)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &barChildR{}
		}
		args = append(args, object.BarID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &barChildR{}
			}

			for _, a := range args {
				if a == obj.BarID {
					continue Outer
				}
			}

			args = append(args, obj.BarID)

		}
	}

	query := NewQuery(qm.From(`bar`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Bar")
	}

	var resultSlice []*Bar
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Bar")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for bar")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for bar")
	}

	if len(barChildAfterSelectHooks) != 0 {
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
		object.R.Bar = foreign
		if foreign.R == nil {
			foreign.R = &barR{}
		}
		foreign.R.BarChildren = append(foreign.R.BarChildren, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.BarID == foreign.ID {
				local.R.Bar = foreign
				if foreign.R == nil {
					foreign.R = &barR{}
				}
				foreign.R.BarChildren = append(foreign.R.BarChildren, local)
				break
			}
		}
	}

	return nil
}

// SetBar of the barChild to the related item.
// Sets o.R.Bar to related.
// Adds o to related.R.BarChildren.
func (o *BarChild) SetBar(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Bar) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `bar_child` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"bar_id"}),
		strmangle.WhereClause("`", "`", 0, barChildPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.BarID = related.ID
	if o.R == nil {
		o.R = &barChildR{
			Bar: related,
		}
	} else {
		o.R.Bar = related
	}

	if related.R == nil {
		related.R = &barR{
			BarChildren: BarChildSlice{o},
		}
	} else {
		related.R.BarChildren = append(related.R.BarChildren, o)
	}

	return nil
}

// BarChildren retrieves all the records using an executor.
func BarChildren(mods ...qm.QueryMod) barChildQuery {
	mods = append(mods, qm.From("`bar_child`"))
	return barChildQuery{NewQuery(mods...)}
}

// FindBarChild retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBarChild(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*BarChild, error) {
	barChildObj := &BarChild{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `bar_child` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, barChildObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from bar_child")
	}

	return barChildObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *BarChild) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no bar_child provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(barChildColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	barChildInsertCacheMut.RLock()
	cache, cached := barChildInsertCache[key]
	barChildInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			barChildColumns,
			barChildColumnsWithDefault,
			barChildColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(barChildType, barChildMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(barChildType, barChildMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `bar_child` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `bar_child` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `bar_child` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, barChildPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into bar_child")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == barChildMapping["ID"] {
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
		return errors.Wrap(err, "models: unable to populate default values for bar_child")
	}

CacheNoHooks:
	if !cached {
		barChildInsertCacheMut.Lock()
		barChildInsertCache[key] = cache
		barChildInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the BarChild.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *BarChild) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	barChildUpdateCacheMut.RLock()
	cache, cached := barChildUpdateCache[key]
	barChildUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			barChildColumns,
			barChildPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update bar_child, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `bar_child` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, barChildPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(barChildType, barChildMapping, append(wl, barChildPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update bar_child row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for bar_child")
	}

	if !cached {
		barChildUpdateCacheMut.Lock()
		barChildUpdateCache[key] = cache
		barChildUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q barChildQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for bar_child")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for bar_child")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BarChildSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), barChildPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `bar_child` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, barChildPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in barChild slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all barChild")
	}
	return rowsAff, nil
}

var mySQLBarChildUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *BarChild) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no bar_child provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(barChildColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLBarChildUniqueColumns, o)

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

	barChildUpsertCacheMut.RLock()
	cache, cached := barChildUpsertCache[key]
	barChildUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			barChildColumns,
			barChildColumnsWithDefault,
			barChildColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			barChildColumns,
			barChildPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert bar_child, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "bar_child", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `bar_child` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(barChildType, barChildMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(barChildType, barChildMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for bar_child")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == barChildMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(barChildType, barChildMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for bar_child")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for bar_child")
	}

CacheNoHooks:
	if !cached {
		barChildUpsertCacheMut.Lock()
		barChildUpsertCache[key] = cache
		barChildUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single BarChild record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *BarChild) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no BarChild provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), barChildPrimaryKeyMapping)
	sql := "DELETE FROM `bar_child` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from bar_child")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for bar_child")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q barChildQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no barChildQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from bar_child")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for bar_child")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BarChildSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no BarChild slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(barChildBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), barChildPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `bar_child` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, barChildPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from barChild slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for bar_child")
	}

	if len(barChildAfterDeleteHooks) != 0 {
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
func (o *BarChild) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBarChild(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BarChildSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BarChildSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), barChildPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `bar_child`.* FROM `bar_child` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, barChildPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BarChildSlice")
	}

	*o = slice

	return nil
}

// BarChildExists checks if the BarChild row exists.
func BarChildExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `bar_child` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if bar_child exists")
	}

	return exists, nil
}