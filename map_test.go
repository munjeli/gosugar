package gosugar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveFromMapByKey(t *testing.T) {
	type args struct {
		m   map[interface{}]interface{}
		key interface{}
	}
	tests := []struct {
		name string
		args args
		want map[interface{}]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, RemoveFromMapByKey(tt.args.m, tt.args.key))
		})
	}
}

func TestRemoveFromMapByValue(t *testing.T) {
	type args struct {
		m   map[interface{}]interface{}
		val interface{}
	}
	tests := []struct {
		name string
		args args
		want map[interface{}]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, RemoveFromMapByValue(tt.args.m, tt.args.val))
		})
	}
}

func TestFilterMapByKeyCondition(t *testing.T) {
	type args struct {
		m map[interface{}]interface{}
		f func(i interface{}) bool
	}
	tests := []struct {
		name  string
		args  args
		want  map[interface{}]interface{}
		want1 map[interface{}]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FilterMapByKeyCondition(tt.args.m, tt.args.f)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func TestFilterMapByValueCondition(t *testing.T) {
	type args struct {
		m map[interface{}]interface{}
		f func(i interface{}) bool
	}
	tests := []struct {
		name  string
		args  args
		want  map[interface{}]interface{}
		want1 map[interface{}]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FilterMapByValueCondition(tt.args.m, tt.args.f)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func TestSameMap(t *testing.T) {
	type args struct {
		m1 map[interface{}]interface{}
		m2 map[interface{}]interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, SameMap(tt.args.m1, tt.args.m2))
		})
	}
}

func TestSortMapByStringKey(t *testing.T) {
	type args struct {
		m map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, SortMapByStringKey(tt.args.m))
		})
	}
}

func TestSortMapByIntKey(t *testing.T) {
	type args struct {
		m map[int]interface{}
	}
	tests := []struct {
		name string
		args args
		want map[int]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, SortMapByIntKey(tt.args.m))
		})
	}
}

func TestSortMapByStringValue(t *testing.T) {
	type args struct {
		m map[interface{}]string
	}
	tests := []struct {
		name string
		args args
		want map[interface{}]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, SortMapByStringValue(tt.args.m))
		})
	}
}
