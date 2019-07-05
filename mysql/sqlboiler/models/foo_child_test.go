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

func testFooChildren(t *testing.T) {
	t.Parallel()

	query := FooChildren()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testFooChildrenDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FooChild{}
	if err = randomize.Struct(seed, o, fooChildDBTypes, true, fooChildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FooChild struct: %s", err)
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

	count, err := FooChildren().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFooChildrenQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FooChild{}
	if err = randomize.Struct(seed, o, fooChildDBTypes, true, fooChildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FooChild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := FooChildren().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := FooChildren().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFooChildrenSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FooChild{}
	if err = randomize.Struct(seed, o, fooChildDBTypes, true, fooChildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FooChild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := FooChildSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := FooChildren().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFooChildrenExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FooChild{}
	if err = randomize.Struct(seed, o, fooChildDBTypes, true, fooChildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FooChild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := FooChildExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if FooChild exists: %s", err)
	}
	if !e {
		t.Errorf("Expected FooChildExists to return true, but got false.")
	}
}

func testFooChildrenFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FooChild{}
	if err = randomize.Struct(seed, o, fooChildDBTypes, true, fooChildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FooChild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	fooChildFound, err := FindFooChild(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if fooChildFound == nil {
		t.Error("want a record, got nil")
	}
}

func testFooChildrenBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FooChild{}
	if err = randomize.Struct(seed, o, fooChildDBTypes, true, fooChildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FooChild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = FooChildren().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testFooChildrenOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FooChild{}
	if err = randomize.Struct(seed, o, fooChildDBTypes, true, fooChildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FooChild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := FooChildren().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testFooChildrenAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	fooChildOne := &FooChild{}
	fooChildTwo := &FooChild{}
	if err = randomize.Struct(seed, fooChildOne, fooChildDBTypes, false, fooChildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FooChild struct: %s", err)
	}
	if err = randomize.Struct(seed, fooChildTwo, fooChildDBTypes, false, fooChildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FooChild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = fooChildOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = fooChildTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := FooChildren().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testFooChildrenCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	fooChildOne := &FooChild{}
	fooChildTwo := &FooChild{}
	if err = randomize.Struct(seed, fooChildOne, fooChildDBTypes, false, fooChildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FooChild struct: %s", err)
	}
	if err = randomize.Struct(seed, fooChildTwo, fooChildDBTypes, false, fooChildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FooChild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = fooChildOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = fooChildTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := FooChildren().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func fooChildBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *FooChild) error {
	*o = FooChild{}
	return nil
}

func fooChildAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *FooChild) error {
	*o = FooChild{}
	return nil
}

func fooChildAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *FooChild) error {
	*o = FooChild{}
	return nil
}

func fooChildBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *FooChild) error {
	*o = FooChild{}
	return nil
}

func fooChildAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *FooChild) error {
	*o = FooChild{}
	return nil
}

func fooChildBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *FooChild) error {
	*o = FooChild{}
	return nil
}

func fooChildAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *FooChild) error {
	*o = FooChild{}
	return nil
}

func fooChildBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *FooChild) error {
	*o = FooChild{}
	return nil
}

func fooChildAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *FooChild) error {
	*o = FooChild{}
	return nil
}

func testFooChildrenHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &FooChild{}
	o := &FooChild{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, fooChildDBTypes, false); err != nil {
		t.Errorf("Unable to randomize FooChild object: %s", err)
	}

	AddFooChildHook(boil.BeforeInsertHook, fooChildBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	fooChildBeforeInsertHooks = []FooChildHook{}

	AddFooChildHook(boil.AfterInsertHook, fooChildAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	fooChildAfterInsertHooks = []FooChildHook{}

	AddFooChildHook(boil.AfterSelectHook, fooChildAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	fooChildAfterSelectHooks = []FooChildHook{}

	AddFooChildHook(boil.BeforeUpdateHook, fooChildBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	fooChildBeforeUpdateHooks = []FooChildHook{}

	AddFooChildHook(boil.AfterUpdateHook, fooChildAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	fooChildAfterUpdateHooks = []FooChildHook{}

	AddFooChildHook(boil.BeforeDeleteHook, fooChildBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	fooChildBeforeDeleteHooks = []FooChildHook{}

	AddFooChildHook(boil.AfterDeleteHook, fooChildAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	fooChildAfterDeleteHooks = []FooChildHook{}

	AddFooChildHook(boil.BeforeUpsertHook, fooChildBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	fooChildBeforeUpsertHooks = []FooChildHook{}

	AddFooChildHook(boil.AfterUpsertHook, fooChildAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	fooChildAfterUpsertHooks = []FooChildHook{}
}

func testFooChildrenInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FooChild{}
	if err = randomize.Struct(seed, o, fooChildDBTypes, true, fooChildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FooChild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := FooChildren().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFooChildrenInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FooChild{}
	if err = randomize.Struct(seed, o, fooChildDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FooChild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(fooChildColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := FooChildren().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFooChildToOneFooUsingFoo(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local FooChild
	var foreign Foo

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, fooChildDBTypes, false, fooChildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FooChild struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, fooDBTypes, false, fooColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Foo struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.FooID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Foo().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := FooChildSlice{&local}
	if err = local.L.LoadFoo(ctx, tx, false, (*[]*FooChild)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Foo == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Foo = nil
	if err = local.L.LoadFoo(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Foo == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFooChildToOneSetOpFooUsingFoo(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a FooChild
	var b, c Foo

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, fooChildDBTypes, false, strmangle.SetComplement(fooChildPrimaryKeyColumns, fooChildColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, fooDBTypes, false, strmangle.SetComplement(fooPrimaryKeyColumns, fooColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, fooDBTypes, false, strmangle.SetComplement(fooPrimaryKeyColumns, fooColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Foo{&b, &c} {
		err = a.SetFoo(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Foo != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.FooChildren[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.FooID != x.ID {
			t.Error("foreign key was wrong value", a.FooID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.FooID))
		reflect.Indirect(reflect.ValueOf(&a.FooID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.FooID != x.ID {
			t.Error("foreign key was wrong value", a.FooID, x.ID)
		}
	}
}

func testFooChildrenReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FooChild{}
	if err = randomize.Struct(seed, o, fooChildDBTypes, true, fooChildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FooChild struct: %s", err)
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

func testFooChildrenReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FooChild{}
	if err = randomize.Struct(seed, o, fooChildDBTypes, true, fooChildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FooChild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := FooChildSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testFooChildrenSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FooChild{}
	if err = randomize.Struct(seed, o, fooChildDBTypes, true, fooChildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FooChild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := FooChildren().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	fooChildDBTypes = map[string]string{`FooID`: `int`, `ID`: `int`, `Name`: `varchar`}
	_               = bytes.MinRead
)

func testFooChildrenUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(fooChildPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(fooChildColumns) == len(fooChildPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &FooChild{}
	if err = randomize.Struct(seed, o, fooChildDBTypes, true, fooChildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FooChild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := FooChildren().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, fooChildDBTypes, true, fooChildPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FooChild struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testFooChildrenSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(fooChildColumns) == len(fooChildPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &FooChild{}
	if err = randomize.Struct(seed, o, fooChildDBTypes, true, fooChildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FooChild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := FooChildren().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, fooChildDBTypes, true, fooChildPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FooChild struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(fooChildColumns, fooChildPrimaryKeyColumns) {
		fields = fooChildColumns
	} else {
		fields = strmangle.SetComplement(
			fooChildColumns,
			fooChildPrimaryKeyColumns,
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

	slice := FooChildSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testFooChildrenUpsert(t *testing.T) {
	t.Parallel()

	if len(fooChildColumns) == len(fooChildPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLFooChildUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := FooChild{}
	if err = randomize.Struct(seed, &o, fooChildDBTypes, false); err != nil {
		t.Errorf("Unable to randomize FooChild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert FooChild: %s", err)
	}

	count, err := FooChildren().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, fooChildDBTypes, false, fooChildPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FooChild struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert FooChild: %s", err)
	}

	count, err = FooChildren().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}