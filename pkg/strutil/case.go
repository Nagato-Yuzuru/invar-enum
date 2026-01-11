package strutil

import "github.com/iancoleman/strcase"

// CaseSet All case set.
// Use the corresponding fields directly in the template.
type CaseSet struct {
	Kebab          string
	Snake          string
	Pascal         string
	Camel          string
	ScreamingSnake string
}

// CaseSetMaker factory about CaseSet
func CaseSetMaker(s string) CaseSet {
	return CaseSet{
		Kebab:          strcase.ToKebab(s),
		Snake:          strcase.ToSnake(s),
		Pascal:         strcase.ToCamel(s),
		Camel:          strcase.ToLowerCamel(s),
		ScreamingSnake: strcase.ToScreamingSnake(s),
	}
}
