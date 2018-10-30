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

func testMydiscordGuilds(t *testing.T) {
	t.Parallel()

	query := MydiscordGuilds()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testMydiscordGuildsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MydiscordGuild{}
	if err = randomize.Struct(seed, o, mydiscordGuildDBTypes, true, mydiscordGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuild struct: %s", err)
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

	count, err := MydiscordGuilds().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMydiscordGuildsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MydiscordGuild{}
	if err = randomize.Struct(seed, o, mydiscordGuildDBTypes, true, mydiscordGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := MydiscordGuilds().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MydiscordGuilds().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMydiscordGuildsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MydiscordGuild{}
	if err = randomize.Struct(seed, o, mydiscordGuildDBTypes, true, mydiscordGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MydiscordGuildSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MydiscordGuilds().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMydiscordGuildsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MydiscordGuild{}
	if err = randomize.Struct(seed, o, mydiscordGuildDBTypes, true, mydiscordGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := MydiscordGuildExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if MydiscordGuild exists: %s", err)
	}
	if !e {
		t.Errorf("Expected MydiscordGuildExists to return true, but got false.")
	}
}

func testMydiscordGuildsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MydiscordGuild{}
	if err = randomize.Struct(seed, o, mydiscordGuildDBTypes, true, mydiscordGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	mydiscordGuildFound, err := FindMydiscordGuild(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if mydiscordGuildFound == nil {
		t.Error("want a record, got nil")
	}
}

func testMydiscordGuildsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MydiscordGuild{}
	if err = randomize.Struct(seed, o, mydiscordGuildDBTypes, true, mydiscordGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = MydiscordGuilds().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testMydiscordGuildsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MydiscordGuild{}
	if err = randomize.Struct(seed, o, mydiscordGuildDBTypes, true, mydiscordGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := MydiscordGuilds().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testMydiscordGuildsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	mydiscordGuildOne := &MydiscordGuild{}
	mydiscordGuildTwo := &MydiscordGuild{}
	if err = randomize.Struct(seed, mydiscordGuildOne, mydiscordGuildDBTypes, false, mydiscordGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuild struct: %s", err)
	}
	if err = randomize.Struct(seed, mydiscordGuildTwo, mydiscordGuildDBTypes, false, mydiscordGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = mydiscordGuildOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = mydiscordGuildTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := MydiscordGuilds().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testMydiscordGuildsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	mydiscordGuildOne := &MydiscordGuild{}
	mydiscordGuildTwo := &MydiscordGuild{}
	if err = randomize.Struct(seed, mydiscordGuildOne, mydiscordGuildDBTypes, false, mydiscordGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuild struct: %s", err)
	}
	if err = randomize.Struct(seed, mydiscordGuildTwo, mydiscordGuildDBTypes, false, mydiscordGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = mydiscordGuildOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = mydiscordGuildTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MydiscordGuilds().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func mydiscordGuildBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *MydiscordGuild) error {
	*o = MydiscordGuild{}
	return nil
}

func mydiscordGuildAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *MydiscordGuild) error {
	*o = MydiscordGuild{}
	return nil
}

func mydiscordGuildAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *MydiscordGuild) error {
	*o = MydiscordGuild{}
	return nil
}

func mydiscordGuildBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *MydiscordGuild) error {
	*o = MydiscordGuild{}
	return nil
}

func mydiscordGuildAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *MydiscordGuild) error {
	*o = MydiscordGuild{}
	return nil
}

func mydiscordGuildBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *MydiscordGuild) error {
	*o = MydiscordGuild{}
	return nil
}

func mydiscordGuildAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *MydiscordGuild) error {
	*o = MydiscordGuild{}
	return nil
}

func mydiscordGuildBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *MydiscordGuild) error {
	*o = MydiscordGuild{}
	return nil
}

func mydiscordGuildAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *MydiscordGuild) error {
	*o = MydiscordGuild{}
	return nil
}

func testMydiscordGuildsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &MydiscordGuild{}
	o := &MydiscordGuild{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, mydiscordGuildDBTypes, false); err != nil {
		t.Errorf("Unable to randomize MydiscordGuild object: %s", err)
	}

	AddMydiscordGuildHook(boil.BeforeInsertHook, mydiscordGuildBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	mydiscordGuildBeforeInsertHooks = []MydiscordGuildHook{}

	AddMydiscordGuildHook(boil.AfterInsertHook, mydiscordGuildAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	mydiscordGuildAfterInsertHooks = []MydiscordGuildHook{}

	AddMydiscordGuildHook(boil.AfterSelectHook, mydiscordGuildAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	mydiscordGuildAfterSelectHooks = []MydiscordGuildHook{}

	AddMydiscordGuildHook(boil.BeforeUpdateHook, mydiscordGuildBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	mydiscordGuildBeforeUpdateHooks = []MydiscordGuildHook{}

	AddMydiscordGuildHook(boil.AfterUpdateHook, mydiscordGuildAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	mydiscordGuildAfterUpdateHooks = []MydiscordGuildHook{}

	AddMydiscordGuildHook(boil.BeforeDeleteHook, mydiscordGuildBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	mydiscordGuildBeforeDeleteHooks = []MydiscordGuildHook{}

	AddMydiscordGuildHook(boil.AfterDeleteHook, mydiscordGuildAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	mydiscordGuildAfterDeleteHooks = []MydiscordGuildHook{}

	AddMydiscordGuildHook(boil.BeforeUpsertHook, mydiscordGuildBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	mydiscordGuildBeforeUpsertHooks = []MydiscordGuildHook{}

	AddMydiscordGuildHook(boil.AfterUpsertHook, mydiscordGuildAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	mydiscordGuildAfterUpsertHooks = []MydiscordGuildHook{}
}

func testMydiscordGuildsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MydiscordGuild{}
	if err = randomize.Struct(seed, o, mydiscordGuildDBTypes, true, mydiscordGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MydiscordGuilds().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMydiscordGuildsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MydiscordGuild{}
	if err = randomize.Struct(seed, o, mydiscordGuildDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MydiscordGuild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(mydiscordGuildColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := MydiscordGuilds().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMydiscordGuildToManyGuildMydiscordAliases(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MydiscordGuild
	var b, c MydiscordAlias

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, mydiscordGuildDBTypes, true, mydiscordGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuild struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, mydiscordAliasDBTypes, false, mydiscordAliasColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, mydiscordAliasDBTypes, false, mydiscordAliasColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.GuildID = a.ID
	c.GuildID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	mydiscordAlias, err := a.GuildMydiscordAliases().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range mydiscordAlias {
		if v.GuildID == b.GuildID {
			bFound = true
		}
		if v.GuildID == c.GuildID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := MydiscordGuildSlice{&a}
	if err = a.L.LoadGuildMydiscordAliases(ctx, tx, false, (*[]*MydiscordGuild)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.GuildMydiscordAliases); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.GuildMydiscordAliases = nil
	if err = a.L.LoadGuildMydiscordAliases(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.GuildMydiscordAliases); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", mydiscordAlias)
	}
}

func testMydiscordGuildToManyGuildMydiscordGuildModules(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MydiscordGuild
	var b, c MydiscordGuildModule

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, mydiscordGuildDBTypes, true, mydiscordGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuild struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, mydiscordGuildModuleDBTypes, false, mydiscordGuildModuleColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, mydiscordGuildModuleDBTypes, false, mydiscordGuildModuleColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.GuildID = a.ID
	c.GuildID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	mydiscordGuildModule, err := a.GuildMydiscordGuildModules().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range mydiscordGuildModule {
		if v.GuildID == b.GuildID {
			bFound = true
		}
		if v.GuildID == c.GuildID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := MydiscordGuildSlice{&a}
	if err = a.L.LoadGuildMydiscordGuildModules(ctx, tx, false, (*[]*MydiscordGuild)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.GuildMydiscordGuildModules); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.GuildMydiscordGuildModules = nil
	if err = a.L.LoadGuildMydiscordGuildModules(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.GuildMydiscordGuildModules); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", mydiscordGuildModule)
	}
}

func testMydiscordGuildToManyGuildPublicrolesPublicroles(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MydiscordGuild
	var b, c PublicrolesPublicrole

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, mydiscordGuildDBTypes, true, mydiscordGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuild struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, publicrolesPublicroleDBTypes, false, publicrolesPublicroleColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, publicrolesPublicroleDBTypes, false, publicrolesPublicroleColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.GuildID = a.ID
	c.GuildID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	publicrolesPublicrole, err := a.GuildPublicrolesPublicroles().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range publicrolesPublicrole {
		if v.GuildID == b.GuildID {
			bFound = true
		}
		if v.GuildID == c.GuildID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := MydiscordGuildSlice{&a}
	if err = a.L.LoadGuildPublicrolesPublicroles(ctx, tx, false, (*[]*MydiscordGuild)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.GuildPublicrolesPublicroles); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.GuildPublicrolesPublicroles = nil
	if err = a.L.LoadGuildPublicrolesPublicroles(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.GuildPublicrolesPublicroles); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", publicrolesPublicrole)
	}
}

func testMydiscordGuildToManyAddOpGuildMydiscordAliases(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MydiscordGuild
	var b, c, d, e MydiscordAlias

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, mydiscordGuildDBTypes, false, strmangle.SetComplement(mydiscordGuildPrimaryKeyColumns, mydiscordGuildColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*MydiscordAlias{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, mydiscordAliasDBTypes, false, strmangle.SetComplement(mydiscordAliasPrimaryKeyColumns, mydiscordAliasColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*MydiscordAlias{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddGuildMydiscordAliases(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.GuildID {
			t.Error("foreign key was wrong value", a.ID, first.GuildID)
		}
		if a.ID != second.GuildID {
			t.Error("foreign key was wrong value", a.ID, second.GuildID)
		}

		if first.R.Guild != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Guild != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.GuildMydiscordAliases[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.GuildMydiscordAliases[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.GuildMydiscordAliases().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testMydiscordGuildToManyAddOpGuildMydiscordGuildModules(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MydiscordGuild
	var b, c, d, e MydiscordGuildModule

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, mydiscordGuildDBTypes, false, strmangle.SetComplement(mydiscordGuildPrimaryKeyColumns, mydiscordGuildColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*MydiscordGuildModule{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, mydiscordGuildModuleDBTypes, false, strmangle.SetComplement(mydiscordGuildModulePrimaryKeyColumns, mydiscordGuildModuleColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*MydiscordGuildModule{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddGuildMydiscordGuildModules(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.GuildID {
			t.Error("foreign key was wrong value", a.ID, first.GuildID)
		}
		if a.ID != second.GuildID {
			t.Error("foreign key was wrong value", a.ID, second.GuildID)
		}

		if first.R.Guild != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Guild != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.GuildMydiscordGuildModules[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.GuildMydiscordGuildModules[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.GuildMydiscordGuildModules().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testMydiscordGuildToManyAddOpGuildPublicrolesPublicroles(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MydiscordGuild
	var b, c, d, e PublicrolesPublicrole

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, mydiscordGuildDBTypes, false, strmangle.SetComplement(mydiscordGuildPrimaryKeyColumns, mydiscordGuildColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*PublicrolesPublicrole{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, publicrolesPublicroleDBTypes, false, strmangle.SetComplement(publicrolesPublicrolePrimaryKeyColumns, publicrolesPublicroleColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*PublicrolesPublicrole{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddGuildPublicrolesPublicroles(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.GuildID {
			t.Error("foreign key was wrong value", a.ID, first.GuildID)
		}
		if a.ID != second.GuildID {
			t.Error("foreign key was wrong value", a.ID, second.GuildID)
		}

		if first.R.Guild != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Guild != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.GuildPublicrolesPublicroles[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.GuildPublicrolesPublicroles[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.GuildPublicrolesPublicroles().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testMydiscordGuildsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MydiscordGuild{}
	if err = randomize.Struct(seed, o, mydiscordGuildDBTypes, true, mydiscordGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuild struct: %s", err)
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

func testMydiscordGuildsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MydiscordGuild{}
	if err = randomize.Struct(seed, o, mydiscordGuildDBTypes, true, mydiscordGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MydiscordGuildSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMydiscordGuildsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MydiscordGuild{}
	if err = randomize.Struct(seed, o, mydiscordGuildDBTypes, true, mydiscordGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := MydiscordGuilds().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	mydiscordGuildDBTypes = map[string]string{`DiscordID`: `text`, `ID`: `integer`, `Name`: `text`, `Trigger`: `character varying`}
	_                     = bytes.MinRead
)

func testMydiscordGuildsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(mydiscordGuildPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(mydiscordGuildColumns) == len(mydiscordGuildPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &MydiscordGuild{}
	if err = randomize.Struct(seed, o, mydiscordGuildDBTypes, true, mydiscordGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MydiscordGuilds().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, mydiscordGuildDBTypes, true, mydiscordGuildPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuild struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testMydiscordGuildsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(mydiscordGuildColumns) == len(mydiscordGuildPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &MydiscordGuild{}
	if err = randomize.Struct(seed, o, mydiscordGuildDBTypes, true, mydiscordGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MydiscordGuilds().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, mydiscordGuildDBTypes, true, mydiscordGuildPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuild struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(mydiscordGuildColumns, mydiscordGuildPrimaryKeyColumns) {
		fields = mydiscordGuildColumns
	} else {
		fields = strmangle.SetComplement(
			mydiscordGuildColumns,
			mydiscordGuildPrimaryKeyColumns,
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

	slice := MydiscordGuildSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testMydiscordGuildsUpsert(t *testing.T) {
	t.Parallel()

	if len(mydiscordGuildColumns) == len(mydiscordGuildPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := MydiscordGuild{}
	if err = randomize.Struct(seed, &o, mydiscordGuildDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MydiscordGuild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert MydiscordGuild: %s", err)
	}

	count, err := MydiscordGuilds().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, mydiscordGuildDBTypes, false, mydiscordGuildPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuild struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert MydiscordGuild: %s", err)
	}

	count, err = MydiscordGuilds().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}