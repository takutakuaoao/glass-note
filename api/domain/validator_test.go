package domain_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/takutakuaoao/glass-note/api/domain"
	"testing"
)

func Test_ノートのバリデーション(t *testing.T) {
	cases := []struct {
		name     string
		title    string
		expected bool
	}{
		{
			name:     "ノートのタイトルが０文字のときは登録できない",
			title:    "",
			expected: false,
		},
		{
			name:     "ノートのタイトルが１文字以上ある場合は登録できる",
			title:    "a",
			expected: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := domain.NewNoteValidator(c.title).Validate()

			assert.Equal(t, c.expected, result)
		})
	}
}
