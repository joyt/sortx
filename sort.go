package sort

import (
	"bytes"
	"reflect"
	"sort"
)

const (
	Ascending  = true
	Descending = false
)

type sortedDocs struct {
	// These are parallel arrays for the values actually being compared. It's better to store them here in a first pass so because that's O(N) Index,FieldByName ops while waiting to inspect the values in Less() will be O(Nlog(N)).
	// Should run some benchmarks to make sure though.
	ints      []int64
	uints     []uint64
	floats    []float64
	strings   []string
	size      int
	docs      interface{}
	fieldType reflect.Kind
	field     string // Name of the field to sort on. Only number and string types currently supported. Might implement sorting on inner fields later.
	order     bool   // true = ascending, false = descending
}

func SortByField(docs interface{}, field string, order bool) {
	v := reflect.ValueOf(docs)
	if k := v.Kind(); k != reflect.Array && k != reflect.Slice {
		panic("Items to sort must be array or slice of structs")
	}
	sdocs := sortedDocs{field: field, order: order, docs: docs}
	sdocs.size = v.Len()
	if sdocs.size == 0 {
		return
	}
	e := v.Index(0)
	if e.Kind() != reflect.Struct {
		panic("Items to sort must be array or slice of structs")
	}
	sdocs.fieldType = e.FieldByName(field).Kind()
	switch sdocs.fieldType {
	case reflect.Invalid:
		panic("Field provided to sort is invalid")
	case reflect.String:
		sdocs.strings = make([]string, sdocs.size)
		for i := 0; i < sdocs.size; i++ {
			val := v.Index(i).FieldByName(field)
			sdocs.strings[i] = val.String()
		}
	case reflect.Bool:
		sdocs.ints = make([]int64, sdocs.size)
		for i := 0; i < sdocs.size; i++ {
			val := v.Index(i).FieldByName(field)
			if val.Bool() {
				sdocs.ints[i] = 1
			} else {
				sdocs.ints[i] = 0
			}
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		sdocs.ints = make([]int64, sdocs.size)
		for i := 0; i < sdocs.size; i++ {
			val := v.Index(i).FieldByName(field)
			sdocs.ints[i] = val.Int()
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		sdocs.uints = make([]uint64, sdocs.size)
		for i := 0; i < sdocs.size; i++ {
			val := v.Index(i).FieldByName(field)
			sdocs.uints[i] = val.Uint()
		}
	case reflect.Float32, reflect.Float64:
		sdocs.floats = make([]float64, sdocs.size)
		for i := 0; i < sdocs.size; i++ {
			val := v.Index(i).FieldByName(field)
			sdocs.floats[i] = val.Float()
		}
	default:
		panic("Field type is not supported comparable type (string, bool, int or float type)")
	}
	sort.Sort(sdocs)
}

func (s sortedDocs) Len() int {
	return s.size
}

func (s sortedDocs) Less(i, j int) bool {
	var b bool
	switch s.fieldType {
	case reflect.String:
		b = bytes.Compare([]byte(s.strings[i]), []byte(s.strings[j])) < 0
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		b = s.ints[i] < s.ints[j]
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		b = s.uints[i] < s.uints[j]
	case reflect.Float32, reflect.Float64:
		b = s.floats[i] < s.floats[j]
	default:
		return false
	}
	return b == s.order
}

func (s sortedDocs) Swap(i, j int) {
	// This function has to be a little roundabout to actually manage to swap the values generically in the original array.
	// Essentially since calling Set will erase the original value we need to save the first value by using Interface{} to basically create a copy.
	v := reflect.ValueOf(s.docs)
	v1 := v.Index(i)
	v2 := v.Index(j)
	temp := reflect.ValueOf(v1.Interface())
	v1.Set(v2)
	v2.Set(temp)
	switch s.fieldType {
	case reflect.String:
		s.strings[i], s.strings[j] = s.strings[j], s.strings[i]
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		s.ints[i], s.ints[j] = s.ints[j], s.ints[i]
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		s.uints[i], s.uints[j] = s.uints[j], s.uints[i]
	case reflect.Float32, reflect.Float64:
		s.floats[i], s.floats[j] = s.floats[j], s.floats[i]
	}
}
