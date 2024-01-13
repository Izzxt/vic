package list_test

import (
	"testing"

	"github.com/Izzxt/vic/list"
)

func TestAdd(t *testing.T) {
	type args[T any] struct {
		list list.List[T]
	}
	tests := []struct {
		name string
		args list.List[string]
		want []string
	}{
		{
			name: "Add",
			args: list.New[string](0),
			want: []string{"john", "alex", "alicia"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.Add("john")
			tt.args.Add("alex")
			tt.args.Add("alicia")

			if got := tt.args.Values(); len(got) != len(tt.want) {
				t.Errorf("Add = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContains(t *testing.T) {
	type args[T any] struct {
		list list.List[T]
	}
	tests := []struct {
		name string
		args list.List[string]
		want bool
	}{
		{
			name: "Contains",
			args: list.New[string](0),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.Add("john")
			tt.args.Add("alex")
			tt.args.Add("alicia")

			if got := tt.args.Contains("alicia"); got != tt.want {
				t.Errorf("Contains = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	type args[T any] struct {
		list list.List[T]
	}
	tests := []struct {
		name string
		args list.List[string]
		want []string
	}{
		{
			name: "Add",
			args: list.New[string](0),
			want: []string{"john", "alicia"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.Add("john")
			tt.args.Add("alex")
			tt.args.Add("alicia")

			tt.args.Remove("alex")

			if got := tt.args.Values(); len(got) != len(tt.want) {
				t.Errorf("Add = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLen(t *testing.T) {
	type args[T any] struct {
		list list.List[T]
	}
	tests := []struct {
		name string
		args list.List[string]
		want int
	}{
		{
			name: "Len",
			args: list.New[string](0),
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.Add("john")
			tt.args.Add("alex")
			tt.args.Add("alicia")

			if got := tt.args.Len(); got != tt.want {
				t.Errorf("Len = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValues(t *testing.T) {
	type args[T any] struct {
		list list.List[T]
	}
	tests := []struct {
		name string
		args list.List[string]
		want []string
	}{
		{
			name: "Values",
			args: list.New[string](0),
			want: []string{"john", "alex", "alicia"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.Add("john")
			tt.args.Add("alex")
			tt.args.Add("alicia")

			if got := tt.args.Values(); len(got) != len(tt.want) {
				t.Errorf("Values = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPop(t *testing.T) {
	type args[T any] struct {
		list list.List[T]
	}
	tests := []struct {
		name string
		args list.List[string]
		want []string
	}{
		{
			name: "Pop",
			args: list.New[string](0),
			want: []string{"john", "alex"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.Add("john")
			tt.args.Add("alex")
			tt.args.Add("alicia")

			tt.args.Pop()

			if got := tt.args.Values(); len(got) != len(tt.want) {
				t.Errorf("Pop = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPopFront(t *testing.T) {
	type args[T any] struct {
		list list.List[T]
	}
	tests := []struct {
		name string
		args list.List[string]
		want []string
	}{
		{
			name: "PopFront",
			args: list.New[string](0),
			want: []string{"alex", "alicia"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.Add("john")
			tt.args.Add("alex")
			tt.args.Add("alicia")

			tt.args.PopFront()

			if got := tt.args.Values(); len(got) != len(tt.want) {
				t.Errorf("Pop = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	type args[T any] struct {
		list list.List[T]
	}
	tests := []struct {
		name string
		args list.List[string]
		want []string
	}{
		{
			name: "Reverse",
			args: list.New[string](0),
			want: []string{"alicia", "alex", "john"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.Add("john")
			tt.args.Add("alex")
			tt.args.Add("alicia")

			tt.args.Reverse()

			if got := tt.args.Values(); len(got) != len(tt.want) {
				t.Errorf("Pop = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLast(t *testing.T) {
	type args[T any] struct {
		list list.List[T]
	}
	tests := []struct {
		name string
		args list.List[string]
		want string
	}{
		{
			name: "Last",
			args: list.New[string](0),
			want: "alicia",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.Add("john")
			tt.args.Add("alex")
			tt.args.Add("alicia")

			if got := tt.args.Last(); got != tt.want {
				t.Errorf("Pop = %v, want %v", got, tt.want)
			}
		})
	}
}
