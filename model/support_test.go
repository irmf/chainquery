// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package model

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

func testSupports(t *testing.T) {
	t.Parallel()

	query := Supports(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testSupportsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	support := &Support{}
	if err = randomize.Struct(seed, support, supportDBTypes, true, supportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Support struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = support.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = support.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Supports(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSupportsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	support := &Support{}
	if err = randomize.Struct(seed, support, supportDBTypes, true, supportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Support struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = support.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Supports(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Supports(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSupportsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	support := &Support{}
	if err = randomize.Struct(seed, support, supportDBTypes, true, supportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Support struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = support.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := SupportSlice{support}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Supports(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testSupportsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	support := &Support{}
	if err = randomize.Struct(seed, support, supportDBTypes, true, supportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Support struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = support.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := SupportExists(tx, support.ID)
	if err != nil {
		t.Errorf("Unable to check if Support exists: %s", err)
	}
	if !e {
		t.Errorf("Expected SupportExistsG to return true, but got false.")
	}
}
func testSupportsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	support := &Support{}
	if err = randomize.Struct(seed, support, supportDBTypes, true, supportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Support struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = support.Insert(tx); err != nil {
		t.Error(err)
	}

	supportFound, err := FindSupport(tx, support.ID)
	if err != nil {
		t.Error(err)
	}

	if supportFound == nil {
		t.Error("want a record, got nil")
	}
}
func testSupportsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	support := &Support{}
	if err = randomize.Struct(seed, support, supportDBTypes, true, supportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Support struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = support.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Supports(tx).Bind(support); err != nil {
		t.Error(err)
	}
}

func testSupportsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	support := &Support{}
	if err = randomize.Struct(seed, support, supportDBTypes, true, supportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Support struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = support.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Supports(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSupportsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	supportOne := &Support{}
	supportTwo := &Support{}
	if err = randomize.Struct(seed, supportOne, supportDBTypes, false, supportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Support struct: %s", err)
	}
	if err = randomize.Struct(seed, supportTwo, supportDBTypes, false, supportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Support struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = supportOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = supportTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Supports(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSupportsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	supportOne := &Support{}
	supportTwo := &Support{}
	if err = randomize.Struct(seed, supportOne, supportDBTypes, false, supportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Support struct: %s", err)
	}
	if err = randomize.Struct(seed, supportTwo, supportDBTypes, false, supportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Support struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = supportOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = supportTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Supports(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testSupportsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	support := &Support{}
	if err = randomize.Struct(seed, support, supportDBTypes, true, supportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Support struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = support.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Supports(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSupportsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	support := &Support{}
	if err = randomize.Struct(seed, support, supportDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Support struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = support.Insert(tx, supportColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := Supports(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSupportToOneClaimUsingSupportedClaim(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Support
	var foreign Claim

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, supportDBTypes, false, supportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Support struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, claimDBTypes, false, claimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Claim struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.SupportedClaimID = foreign.ClaimID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.SupportedClaim(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ClaimID != foreign.ClaimID {
		t.Errorf("want: %v, got %v", foreign.ClaimID, check.ClaimID)
	}

	slice := SupportSlice{&local}
	if err = local.L.LoadSupportedClaim(tx, false, (*[]*Support)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.SupportedClaim == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.SupportedClaim = nil
	if err = local.L.LoadSupportedClaim(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.SupportedClaim == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testSupportToOneSetOpClaimUsingSupportedClaim(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Support
	var b, c Claim

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, supportDBTypes, false, strmangle.SetComplement(supportPrimaryKeyColumns, supportColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, claimDBTypes, false, strmangle.SetComplement(claimPrimaryKeyColumns, claimColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, claimDBTypes, false, strmangle.SetComplement(claimPrimaryKeyColumns, claimColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Claim{&b, &c} {
		err = a.SetSupportedClaim(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.SupportedClaim != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.SupportedClaimSupports[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.SupportedClaimID != x.ClaimID {
			t.Error("foreign key was wrong value", a.SupportedClaimID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.SupportedClaimID))
		reflect.Indirect(reflect.ValueOf(&a.SupportedClaimID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.SupportedClaimID != x.ClaimID {
			t.Error("foreign key was wrong value", a.SupportedClaimID, x.ClaimID)
		}
	}
}
func testSupportsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	support := &Support{}
	if err = randomize.Struct(seed, support, supportDBTypes, true, supportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Support struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = support.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = support.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testSupportsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	support := &Support{}
	if err = randomize.Struct(seed, support, supportDBTypes, true, supportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Support struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = support.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := SupportSlice{support}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testSupportsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	support := &Support{}
	if err = randomize.Struct(seed, support, supportDBTypes, true, supportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Support struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = support.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Supports(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	supportDBTypes = map[string]string{`BidState`: `varchar`, `ID`: `bigint`, `SupportAmount`: `double`, `SupportedClaimID`: `char`, `TransactionHash`: `varchar`, `Vout`: `int`}
	_              = bytes.MinRead
)

func testSupportsUpdate(t *testing.T) {
	t.Parallel()

	if len(supportColumns) == len(supportPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	support := &Support{}
	if err = randomize.Struct(seed, support, supportDBTypes, true, supportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Support struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = support.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Supports(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, support, supportDBTypes, true, supportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Support struct: %s", err)
	}

	if err = support.Update(tx); err != nil {
		t.Error(err)
	}
}

func testSupportsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(supportColumns) == len(supportPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	support := &Support{}
	if err = randomize.Struct(seed, support, supportDBTypes, true, supportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Support struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = support.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Supports(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, support, supportDBTypes, true, supportPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Support struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(supportColumns, supportPrimaryKeyColumns) {
		fields = supportColumns
	} else {
		fields = strmangle.SetComplement(
			supportColumns,
			supportPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(support))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := SupportSlice{support}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testSupportsUpsert(t *testing.T) {
	t.Parallel()

	if len(supportColumns) == len(supportPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	support := Support{}
	if err = randomize.Struct(seed, &support, supportDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Support struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = support.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert Support: %s", err)
	}

	count, err := Supports(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &support, supportDBTypes, false, supportPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Support struct: %s", err)
	}

	if err = support.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert Support: %s", err)
	}

	count, err = Supports(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}