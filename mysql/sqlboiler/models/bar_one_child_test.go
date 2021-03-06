// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testBarOneChildren(t *testing.T) {
	t.Parallel()

	query := BarOneChildren()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testBarOneChildrenDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BarOneChild{}
	if err = randomize.Struct(seed, o, barOneChildDBTypes, true, barOneChildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BarOneChild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BarOneChildren().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBarOneChildrenQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BarOneChild{}
	if err = randomize.Struct(seed, o, barOneChildDBTypes, true, barOneChildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BarOneChild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := BarOneChildren().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BarOneChildren().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBarOneChildrenSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BarOneChild{}
	if err = randomize.Struct(seed, o, barOneChildDBTypes, true, barOneChildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BarOneChild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BarOneChildSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BarOneChildren().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBarOneChildrenExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BarOneChild{}
	if err = randomize.Struct(seed, o, barOneChildDBTypes, true, barOneChildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BarOneChild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := BarOneChildExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if BarOneChild exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BarOneChildExists to return true, but got false.")
	}
}

func testBarOneChildrenFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BarOneChild{}
	if err = randomize.Struct(seed, o, barOneChildDBTypes, true, barOneChildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BarOneChild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	barOneChildFound, err := FindBarOneChild(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if barOneChildFound == nil {
		t.Error("want a record, got nil")
	}
}

func testBarOneChildrenBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BarOneChild{}
	if err = randomize.Struct(seed, o, barOneChildDBTypes, true, barOneChildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BarOneChild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = BarOneChildren().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testBarOneChildrenOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BarOneChild{}
	if err = randomize.Struct(seed, o, barOneChildDBTypes, true, barOneChildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BarOneChild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := BarOneChildren().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBarOneChildrenAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	barOneChildOne := &BarOneChild{}
	barOneChildTwo := &BarOneChild{}
	if err = randomize.Struct(seed, barOneChildOne, barOneChildDBTypes, false, barOneChildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BarOneChild struct: %s", err)
	}
	if err = randomize.Struct(seed, barOneChildTwo, barOneChildDBTypes, false, barOneChildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BarOneChild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = barOneChildOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = barOneChildTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BarOneChildren().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBarOneChildrenCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	barOneChildOne := &BarOneChild{}
	barOneChildTwo := &BarOneChild{}
	if err = randomize.Struct(seed, barOneChildOne, barOneChildDBTypes, false, barOneChildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BarOneChild struct: %s", err)
	}
	if err = randomize.Struct(seed, barOneChildTwo, barOneChildDBTypes, false, barOneChildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BarOneChild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = barOneChildOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = barOneChildTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BarOneChildren().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func barOneChildBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *BarOneChild) error {
	*o = BarOneChild{}
	return nil
}

func barOneChildAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *BarOneChild) error {
	*o = BarOneChild{}
	return nil
}

func barOneChildAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *BarOneChild) error {
	*o = BarOneChild{}
	return nil
}

func barOneChildBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *BarOneChild) error {
	*o = BarOneChild{}
	return nil
}

func barOneChildAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *BarOneChild) error {
	*o = BarOneChild{}
	return nil
}

func barOneChildBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *BarOneChild) error {
	*o = BarOneChild{}
	return nil
}

func barOneChildAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *BarOneChild) error {
	*o = BarOneChild{}
	return nil
}

func barOneChildBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *BarOneChild) error {
	*o = BarOneChild{}
	return nil
}

func barOneChildAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *BarOneChild) error {
	*o = BarOneChild{}
	return nil
}

func testBarOneChildrenHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &BarOneChild{}
	o := &BarOneChild{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, barOneChildDBTypes, false); err != nil {
		t.Errorf("Unable to randomize BarOneChild object: %s", err)
	}

	AddBarOneChildHook(boil.BeforeInsertHook, barOneChildBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	barOneChildBeforeInsertHooks = []BarOneChildHook{}

	AddBarOneChildHook(boil.AfterInsertHook, barOneChildAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	barOneChildAfterInsertHooks = []BarOneChildHook{}

	AddBarOneChildHook(boil.AfterSelectHook, barOneChildAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	barOneChildAfterSelectHooks = []BarOneChildHook{}

	AddBarOneChildHook(boil.BeforeUpdateHook, barOneChildBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	barOneChildBeforeUpdateHooks = []BarOneChildHook{}

	AddBarOneChildHook(boil.AfterUpdateHook, barOneChildAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	barOneChildAfterUpdateHooks = []BarOneChildHook{}

	AddBarOneChildHook(boil.BeforeDeleteHook, barOneChildBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	barOneChildBeforeDeleteHooks = []BarOneChildHook{}

	AddBarOneChildHook(boil.AfterDeleteHook, barOneChildAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	barOneChildAfterDeleteHooks = []BarOneChildHook{}

	AddBarOneChildHook(boil.BeforeUpsertHook, barOneChildBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	barOneChildBeforeUpsertHooks = []BarOneChildHook{}

	AddBarOneChildHook(boil.AfterUpsertHook, barOneChildAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	barOneChildAfterUpsertHooks = []BarOneChildHook{}
}

func testBarOneChildrenInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BarOneChild{}
	if err = randomize.Struct(seed, o, barOneChildDBTypes, true, barOneChildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BarOneChild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BarOneChildren().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBarOneChildrenInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BarOneChild{}
	if err = randomize.Struct(seed, o, barOneChildDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BarOneChild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(barOneChildColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := BarOneChildren().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBarOneChildToOneBarUsingBar(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local BarOneChild
	var foreign Bar

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, barOneChildDBTypes, false, barOneChildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BarOneChild struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, barDBTypes, false, barColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bar struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.BarID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Bar().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := BarOneChildSlice{&local}
	if err = local.L.LoadBar(ctx, tx, false, (*[]*BarOneChild)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Bar == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Bar = nil
	if err = local.L.LoadBar(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Bar == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testBarOneChildToOneSetOpBarUsingBar(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BarOneChild
	var b, c Bar

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, barOneChildDBTypes, false, strmangle.SetComplement(barOneChildPrimaryKeyColumns, barOneChildColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, barDBTypes, false, strmangle.SetComplement(barPrimaryKeyColumns, barColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, barDBTypes, false, strmangle.SetComplement(barPrimaryKeyColumns, barColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Bar{&b, &c} {
		err = a.SetBar(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Bar != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.BarOneChildren[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.BarID != x.ID {
			t.Error("foreign key was wrong value", a.BarID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.BarID))
		reflect.Indirect(reflect.ValueOf(&a.BarID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.BarID != x.ID {
			t.Error("foreign key was wrong value", a.BarID, x.ID)
		}
	}
}

func testBarOneChildrenReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BarOneChild{}
	if err = randomize.Struct(seed, o, barOneChildDBTypes, true, barOneChildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BarOneChild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBarOneChildrenReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BarOneChild{}
	if err = randomize.Struct(seed, o, barOneChildDBTypes, true, barOneChildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BarOneChild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BarOneChildSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBarOneChildrenSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BarOneChild{}
	if err = randomize.Struct(seed, o, barOneChildDBTypes, true, barOneChildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BarOneChild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BarOneChildren().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	barOneChildDBTypes = map[string]string{`BarID`: `int`, `ID`: `int`, `Name`: `varchar`}
	_                  = bytes.MinRead
)

func testBarOneChildrenUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(barOneChildPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(barOneChildColumns) == len(barOneChildPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BarOneChild{}
	if err = randomize.Struct(seed, o, barOneChildDBTypes, true, barOneChildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BarOneChild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BarOneChildren().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, barOneChildDBTypes, true, barOneChildPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BarOneChild struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testBarOneChildrenSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(barOneChildColumns) == len(barOneChildPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BarOneChild{}
	if err = randomize.Struct(seed, o, barOneChildDBTypes, true, barOneChildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BarOneChild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BarOneChildren().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, barOneChildDBTypes, true, barOneChildPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BarOneChild struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(barOneChildColumns, barOneChildPrimaryKeyColumns) {
		fields = barOneChildColumns
	} else {
		fields = strmangle.SetComplement(
			barOneChildColumns,
			barOneChildPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := BarOneChildSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testBarOneChildrenUpsert(t *testing.T) {
	t.Parallel()

	if len(barOneChildColumns) == len(barOneChildPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLBarOneChildUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := BarOneChild{}
	if err = randomize.Struct(seed, &o, barOneChildDBTypes, false); err != nil {
		t.Errorf("Unable to randomize BarOneChild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BarOneChild: %s", err)
	}

	count, err := BarOneChildren().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, barOneChildDBTypes, false, barOneChildPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BarOneChild struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BarOneChild: %s", err)
	}

	count, err = BarOneChildren().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
