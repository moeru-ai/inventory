package form_parser

import (
	"bytes"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

type Field struct {
	Label  string   `json:"label"`
	Value  string   `json:"value"`
	Values []string `json:"values"` // for checkboxes and dropdown
}

func Parse(form string) ([]Field, error) {
	formBytes := []byte(form)

	reader := text.NewReader(formBytes)
	rootNode := goldmark.DefaultParser().Parse(reader)
	fields := []Field{}
	var field *Field

	ast.Walk(rootNode, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		if node.Kind() == ast.KindHeading && node.(*ast.Heading).Level == 3 {
			if field != nil {
				fields = append(fields, *field)
				field = nil
			}

			field = &Field{
				Label: string(node.Text((formBytes))),
			}

			return ast.WalkSkipChildren, nil
		}

		if node.Kind() == ast.KindParagraph {
			if field == nil { // ignore paragraphs before field
				return ast.WalkSkipChildren, nil
			}

			paragraph := node.(*ast.Paragraph)
			paragraphText := string(paragraph.Lines().Value(formBytes))

			if field.Value != "" { // append to field value
				field.Value += "\n" + paragraphText
				return ast.WalkSkipChildren, nil
			}

			field.Value = paragraphText

			return ast.WalkSkipChildren, nil
		}

		if node.Kind() == ast.KindListItem {
			if field == nil { // ignore list items before field
				return ast.WalkSkipChildren, nil
			}

			listItem := node.(*ast.ListItem)
			text := listItem.Text(formBytes)

			if bytes.HasPrefix(text, []byte("[x]")) {
				text = bytes.TrimPrefix(text, []byte("[x]"))
				text = bytes.TrimSpace(text)
				field.Values = append(field.Values, string(text))
			}

			return ast.WalkSkipChildren, nil
		}

		return ast.WalkContinue, nil
	})

	if field != nil { // add last field
		fields = append(fields, *field)
	}

	return fields, nil
}
