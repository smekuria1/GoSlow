package unionfind

import (
	"reflect"
	"testing"

	darray "github.com/smekuria1/GoSlow/dArray"
)

func TestNewUnionFind(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want *UnionFind[int]
	}{
		// TODO: Add test cases.
		{
			"TestNewUnionFind",
			args{10},
			helperUnionFind(10),
		},
		// negative size
		{
			"TestNewUnionFind",
			args{-10},
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//test for nil
			if tt.want == nil {
				if got := NewUnionFind[int](tt.args.size); got != nil {
					t.Errorf("NewUnionFind() = %v, want %v", got, tt.want)
				}
				return
			}
			if got := NewUnionFind[int](tt.args.size); !reflect.DeepEqual(got.id.ToString(), tt.want.id.ToString()) {
				t.Errorf("NewUnionFind() = %v, want %v", got.id.ToString(), tt.want.id.ToString())
			}
		})
	}
}

func helperUnionFind(size int) *UnionFind[int] {
	return NewUnionFind[int](size)
}

func TestUnionFind_Find(t *testing.T) {
	type fields struct {
		size          int
		sz            *darray.DynamicArray[int]
		id            *darray.DynamicArray[int]
		numComponents int
	}
	type args struct {
		p int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			"TestUnionFind_Find",
			fields{
				10,
				helperDynamicArray(10),
				helperDynamicArray(10),
				10,
			},
			args{1},
			1,
		},
		{
			"TestUnionFind_Find",
			fields{
				10,
				helperDynamicArray(10),
				helperDynamicArray(10),
				10,
			},
			args{2},
			2,
		},
		{
			"TestUnionFind_Find",
			fields{
				10,
				helperDynamicArray(10),
				helperDynamicArray(10),
				10,
			},
			args{3},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uf := &UnionFind[int]{
				size:          tt.fields.size,
				sz:            tt.fields.sz,
				id:            tt.fields.id,
				numComponents: tt.fields.numComponents,
			}
			if got := uf.Find(tt.args.p); got != tt.want {
				t.Errorf("UnionFind.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
func helperDynamicArray(size int) *darray.DynamicArray[int] {
	arr := darray.NewDynamicArray[int](size)
	for i := 0; i < size; i++ {
		arr.Add(i)
	}
	return arr
}

func TestUnionFind_Connected(t *testing.T) {

}

func TestUnionFind_ComponentSize(t *testing.T) {
	type args struct {
		p int
	}
	tests := []struct {
		name   string
		fields *UnionFind[int]
		args   args
		want   int
	}{
		{
			"TestUnionFind_ComponentSize",
			helperUnionFind(10),
			args{1},
			1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uf := tt.fields
			if got := uf.ComponentSize(tt.args.p); got != tt.want {
				t.Errorf("UnionFind.ComponentSize() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestUnionFind_Unify(t *testing.T) {
	type fields struct {
		size          int
		sz            *darray.DynamicArray[int]
		id            *darray.DynamicArray[int]
		numComponents int
	}
	type args struct {
		p int
		q int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"TestUnionFind_Unify",
			fields{
				10,
				helperDynamicArray(10),
				helperDynamicArray(10),
				10,
			},
			args{1, 2},
			true,
		},
		{
			"TestUnionFind_Unify",
			fields{
				10,
				helperDynamicArray(10),
				helperDynamicArray(10),
				10,
			},
			args{3, 4},
			true,
		},
		{
			"TestUnionFind_Unify",
			fields{
				10,
				helperDynamicArray(10),
				helperDynamicArray(10),
				10,
			},
			args{5, 6},
			true,
		},
		{
			"TestUnionFind_Unify",
			fields{
				10,
				helperDynamicArray(10),
				helperDynamicArray(10),
				10,
			},
			args{5, 6},
			true,
		},
		// // false cases
		// {
		// 	"TestUnionFind_UnifyFalse",
		// 	fields{
		// 		10,
		// 		helperDynamicArray(10),
		// 		helperDynamicArray(10),
		// 		10,
		// 	},
		// 	args{1, 3},
		// 	false,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uf := &UnionFind[int]{
				size:          tt.fields.size,
				sz:            tt.fields.sz,
				id:            tt.fields.id,
				numComponents: tt.fields.numComponents,
			}
			// unify does not return anything
			uf.Unify(tt.args.p, tt.args.q)
			if got := uf.Connected(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("UnionFind.Unify() = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestUnionFind_Size(t *testing.T) {
	tests := []struct {
		name   string
		fields *UnionFind[int]
		want   int
	}{
		{
			"TestUnionFind_Size",
			helperUnionFind(10),
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uf := tt.fields
			if got := uf.Size(); got != tt.want {
				t.Errorf("UnionFind.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnionFind_Components(t *testing.T) {
	tests := []struct {
		name   string
		fields *UnionFind[int]
		want   int
	}{
		{
			"TestUnionFind_Components",
			helperUnionFind(10),
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uf := tt.fields
			if got := uf.Components(); got != tt.want {
				t.Errorf("UnionFind.Components() = %v, want %v", got, tt.want)
			}
		})
	}
}
