package search

import (
	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
)

// CopyRules copies the rules from the given source index into the destination one.
//
// This method can only be used with indices which belong to the same Algolia application.
func (c *Client) CopyRules(source, destination string, opts ...interface{}) (UpdateTaskRes, error) {
	opts = opt.InsertOrReplaceOption(opts, opt.Scopes("rules"))
	return c.CopyIndex(source, destination, opts...)
}

// CopySettings copies the settings from the given source index into the destination one.
//
// This method can only be used with indices which belong to the same Algolia application.
func (c *Client) CopySettings(source, destination string, opts ...interface{}) (UpdateTaskRes, error) {
	opts = opt.InsertOrReplaceOption(opts, opt.Scopes("settings"))
	return c.CopyIndex(source, destination, opts...)
}

// CopySynonyms copies the synonyms from the given source index into the destination one.
//
// This method can only be used with indices which belong to the same Algolia application.
func (c *Client) CopySynonyms(source, destination string, opts ...interface{}) (UpdateTaskRes, error) {
	opts = opt.InsertOrReplaceOption(opts, opt.Scopes("synonyms"))
	return c.CopyIndex(source, destination, opts...)
}

// CopyIndex copies the full content (objects, synonyms, rules, settings) of the
// given source index into the destination one.
//
// This method can only be used with indices which belong to the same Algolia application. To
// perform the same operation on indices which belong to different Algolia applications, use
// Account.CopyIndex which is optimized for this use-case.
func (c *Client) CopyIndex(source, destination string, opts ...interface{}) (UpdateTaskRes, error) {
	return c.operation(source, destination, "copy", opts...)
}

// MoveIndex moves the full content (objects, synonyms, rules, settings) of the
// given source index into the destination one, effectively deleting the source
// index.
func (c *Client) MoveIndex(source, destination string, opts ...interface{}) (UpdateTaskRes, error) {
	return c.operation(source, destination, "move", opts...)
}

func (c *Client) operation(source, destination, op string, opts ...interface{}) (res UpdateTaskRes, err error) {
	return c.InitIndex(source).operation(destination, op, opts...)
}
