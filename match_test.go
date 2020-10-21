package jcheck

import "testing"

func TestMatch(t *testing.T) {
	testcases := []struct {
		path    string
		pattern string
		expect  bool
	}{
		{"", "#", true},
		{"", "*", true},
		{"hello.world", "hello.world", true},
		{"hello.world", "#", true},
		{"hello.world", "h#", true},
		{"hello.world", "#.world", true},
		{"hello.world", "#world", true},
		{"hello.world", "#world*", true},
		{"hello.world", "#ello.world", true},
		{"hello.world", "h#lo.world", true},
		{"hello.world", "h#rld", true},
		{"hello.world", "he*.#", true},
		{"hello.world", "*.world", true},
		{"hello.world", "hello.*", true},
		{"hello.world", "h*.#d", true},
		{"hello.world", "h#.*d", true},
		{"hello.world", "he??o.world", true},
		{"hello.world", "he??o.?????", true},
		{"hello.world", "?ello.?orl?", true},
		{"hello.world", "?ello.?#", true},
		{"hello.world", "?ello.?*", true},
		{"hello.10.world", "h#rld", true},
		{"hello.10.world", "h#rl*", true},
		{"hello.10.world", "hello.*.world", true},
		{"hello.10.world", "hello.1*.world", true},
		{"hello.10.world", "hello.*0.world", true},

		{"", "?", false},
		{"hello.world", "hello?world", false},
		{"hello.world", "hel?.world", false},
		{"hello.world", "hello.wo?", false},
		{"hello.world", "?llo.world", false},
		{"hello.world2", "hello.world", false},
		{"hello.world", "h*d", false},
		{"hello.world", "hell*d", false},
		{"hello.world", "h*world", false},
		{"hello.world", "h*d", false},
		{"hello.10.world", "hello.2*.world", false},
	}

	for _, tc := range testcases {
		if match(tc.pattern, tc.path) != tc.expect {
			t.Errorf("path:%q pattern:%q expected:%v", tc.path, tc.pattern, tc.expect)
		}
	}
}
