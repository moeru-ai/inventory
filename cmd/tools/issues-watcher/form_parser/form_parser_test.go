package form_parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	a := assert.New(t)

	formStr := `
### Field 1

Value

Another value

### Field 2

- [x] Option 1
- [ ] Option 2
- [ ] Option 3

- [x] Option 4
	`

	fields, err := Parse(formStr)
	if err != nil {
		t.Fatalf("failed to parse form: %v", err)
	}

	a.NoError(err)

	a.Equal([]Field{
		{
			Label: "Field 1",
			Value: "Value\nAnother value",
		},
		{
			Label: "Field 2",
			Values: []string{
				"Option 1",
				"Option 4",
			},
		},
	}, fields)
}
