package main

import (
    list "./testlist"
    "testing"
    "reflect"
    "strings"
)

func notEquals(a interface{}, b interface{}) bool {
    return !reflect.DeepEqual(a, b)
}

func isFourLtr(s string) bool {
    return len(s) == 4
}

func TestEmptySizeIs0(t *testing.T) {
    var lst = list.MEmpty()
    if lst.Size() != 0 {
        t.Error("MEmpty should create empty list")
    }
}

func TestAddToEmpty(t *testing.T) {
    var lst = list.MEmpty().Add("blah")
    if lst.Size() != 1 {
        t.Error("List with 1 element should have size 1")
    }
}

func TestRemove(t *testing.T) {
    cases := [][]string{
        []string{}, []string{},
        []string{"one"}, []string{"one"},
        []string{"one", "two"}, []string{"one", "two"},
        []string{"zero", "one", "two"}, []string{"one", "two"},
        []string{"one", "zero", "two"}, []string{"one", "two"},
        []string{"one", "two", "zero"}, []string{"one", "two"},
        []string{"one", "zero", "two", "zero"}, []string{"one", "two", "zero"},
    }
    for i := 0; i < len(cases); i += 2 {
        if notEquals(list.FromSlice(cases[i]).Remove("zero").ToSlice(), cases[i + 1]) {
            t.Error("Failed removing 'zero' from " + list.FromSlice(cases[i]).String())
        }
    }
}

func TestToSlice(t *testing.T) {
    if notEquals(list.MEmpty().ToSlice(), []string{}) {
        t.Error("Empty list not matching empty slice")
    }
    if notEquals(
            list.MEmpty().Add("some").Add("other").ToSlice(),
            []string{"other", "some"}) {
        t.Error("Non-empty list not matching non-empty slice")
    }
}

func TestFromAndToSlice(t *testing.T) {
    if notEquals(list.FromSlice([]string{}).ToSlice(), []string{}) {
        t.Error("Empty list to-fro conversion failed")
    }
    sample := []string{"humans", "aliens"};
    if notEquals(list.FromSlice(sample).ToSlice(), sample) {
        t.Error("Non-Empty list to-fro conversion failed")
    }
}

func TestFind(t *testing.T) {
    if r, ok := list.FromSlice([]string{"bla", "zvon", "clown"}).Find(isFourLtr); r != "zvon" || !ok {
        t.Error("Find existing element fails")
    }
    if _, ok := list.FromSlice([]string{"bla", "clown"}).Find(isFourLtr); ok {
        t.Error("Finding non-existing element shouldn't be ok")
    }
}

func TestFilter(t *testing.T) {
    src := list.FromSlice([]string{"who", "when", "where", "what"})
    if notEquals(src.Filter(isFourLtr).ToSlice(), []string{"when", "what"}) {
        t.Error("Filtering fails")
    }
}

func TestMap(t *testing.T) {
    src := list.FromSlice([]string{"Alpha", "bEta"})
    if notEquals(src.Map(strings.ToUpper).ToSlice(), []string{"ALPHA", "BETA"}) {
        t.Error("Map function fails")
    }
}

func TestReduce(t *testing.T) {
    f := func(a, b string) (string) {
        return b + a
    }
    if list.FromSlice([]string{"calbus", "salo", "ve"}).Reduce(f) != "vesalocalbus" {
        t.Error("Reduce fails")
    }
}

func TestMConcant(t *testing.T) {
    x := list.FromSlice([]string{"ax", "bx"})
    y := list.FromSlice([]string{"cs", "ds"})
    if notEquals(x.MConcat(y).ToSlice(), []string{"ax", "bx", "cs", "ds"}) {
        t.Error("MConcat fails")
    }
}

func TestToString(t *testing.T) {
    if list.MEmpty().String() != "[]" {
        t.Error("Empty list to string conversion failed")
    }
    if list.FromSlice([]string{"aha", "bom"}).String() != "[aha, bom]" {
        t.Error("Non-Empty list to string conversion failed")
    }
}
