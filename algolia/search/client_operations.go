package search

import (
	iopt "github.com/algolia/algoliasearch-client-go/algolia/internal/opt"
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

func (c *Client) CopyRules(source, destination string, opts ...interface{}) (UpdateTaskRes, error) {
	opts = iopt.InsertOrReplaceOption(opts, opt.Scopes("rules"))
	return c.CopyIndex(source, destination, opts...)
}

func (c *Client) CopySettings(source, destination string, opts ...interface{}) (UpdateTaskRes, error) {
	opts = iopt.InsertOrReplaceOption(opts, opt.Scopes("settings"))
	return c.CopyIndex(source, destination, opts...)
}

func (c *Client) CopySynonyms(source, destination string, opts ...interface{}) (UpdateTaskRes, error) {
	opts = iopt.InsertOrReplaceOption(opts, opt.Scopes("synonyms"))
	return c.CopyIndex(source, destination, opts...)
}

func (c *Client) CopyIndex(source, destination string, opts ...interface{}) (UpdateTaskRes, error) {
	return c.operation(source, destination, "copy", opts...)
}

func (c *Client) MoveRules(source, destination string, opts ...interface{}) (UpdateTaskRes, error) {
	opts = iopt.InsertOrReplaceOption(opts, opt.Scopes("rules"))
	return c.MoveIndex(source, destination, opts...)
}

func (c *Client) MoveSettings(source, destination string, opts ...interface{}) (UpdateTaskRes, error) {
	opts = iopt.InsertOrReplaceOption(opts, opt.Scopes("settings"))
	return c.MoveIndex(source, destination, opts...)
}

func (c *Client) MoveSynonyms(source, destination string, opts ...interface{}) (UpdateTaskRes, error) {
	opts = iopt.InsertOrReplaceOption(opts, opt.Scopes("synonyms"))
	return c.MoveIndex(source, destination, opts...)
}

func (c *Client) MoveIndex(source, destination string, opts ...interface{}) (UpdateTaskRes, error) {
	return c.operation(source, destination, "move", opts...)
}

func (c *Client) operation(source, destination, op string, opts ...interface{}) (res UpdateTaskRes, err error) {
	return c.InitIndex(source).operation(destination, op, opts...)
}
