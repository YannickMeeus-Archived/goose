package mock_text_generator

import (
	"strings"
	"testing"
)

func Test_ToMockText(t *testing.T) {
	t.Run("should return a string that when lower-cased is the same string as the original", func(t *testing.T) {
		if got := ToMockText("Whatever With Whatever Case"); strings.ToLower(got) != "whatever with whatever case" {
			t.Errorf("ToMockText() = %v, want %v", got, "hello world")
		}
	})

	t.Run("should return a string that contains some uppercase and some lowercase characters", func(t *testing.T) {
		if got := ToMockText("all lower case to start with"); !strings.ContainsAny(got, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
			t.Errorf("ToMockText() = %v, needs some uppercase characters", got)
		}
	})
	t.Run("should return a string that contains some lowercase characters", func(t *testing.T) {
		if got := ToMockText("ALL UPPER CASE TO START WITH"); !strings.ContainsAny(got, "abcdefghijklmnopqrstuvwxyz") {
			t.Errorf("ToMockText() = %v, needs some lowercase characters", got)
		}
	})
}
