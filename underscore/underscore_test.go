package underscore

// Testing multiple cases
// An example of using table driven tests
// As well as using testing.Run to spawn subtests

import "testing"

// func TestCamel(t *testing.T) {
// 	testCases := []struct {
// 		arg  string
// 		want string
// 	}{
// 		{"thisIsACamelCaseString", "this_is_a_camel_case_string"},
// 		{"thisIsACamelCase string", "this_is_a_camel_case string"},
// 	}
// 	for _, tc := range testCases {
// 		got := Camel(tc.arg)
// 		if got != tc.want {
// 			t.Errorf("Camel(%q) = %q; want %q", tc.arg, got, tc.want)
// 		}

// 	}
// }

func TestCamel(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"simple", args{"thisIsACamelCaseString"}, "this_is_a_camel_case_string"},
		{"spaces", args{"with a space"}, "with a space"},
		{"ends with capital", args{"endsWithA"}, "ends_with_a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Camel(tt.args.str); got != tt.want {
				t.Errorf("Camel() = %v, want %v", got, tt.want)
			}
		})
	}
}
