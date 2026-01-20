// pinyin_test.go
package chinaid

import "testing"

func TestConvertPinyin(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"王", "wang"},
		{"李明", "liming"},
		{"张伟", "zhangwei"},
		{"欧阳", "ouyang"},
		{"司马", "sima"},
		{"", ""},
	}

	for _, tt := range tests {
		got := ConvertPinyin(tt.input)
		if got != tt.want {
			t.Errorf("ConvertPinyin(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestConvertPinyinFirst(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"小明", "xiao"},
		{"建国", "jian"},
		{"", ""},
	}

	for _, tt := range tests {
		got := ConvertPinyinFirst(tt.input)
		if got != tt.want {
			t.Errorf("ConvertPinyinFirst(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}
