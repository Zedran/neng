package neng

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

/* Handles case transformations. */
type caser struct {
	lower cases.Caser
	title cases.Caser
	upper cases.Caser
}

/* Transforms word to lower case. */
func (c *caser) toLower(word string) string {
	return c.lower.String(word)
}

/* Transforms word to title case. */
func (c *caser) toTitle(word string) string {
	return c.title.String(word)
}

/* Transforms word to upper case. */
func (c *caser) toUpper(word string) string {
	return c.upper.String(word)
}

/* Returns a pointer to new caser struct. */
func newCaser() *caser {
	return &caser{
		lower: cases.Lower(language.English),
		title: cases.Title(language.English),
		upper: cases.Upper(language.English),
	}
}
