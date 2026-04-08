package tree_sitter_tspipe_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_tspipe "github.com/marifat0/tree-sitter-pipe/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_tspipe.Language())
	if language == nil {
		t.Errorf("Error loading tree-sitter-pipe grammar")
	}
}
