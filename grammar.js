/**
 * @file A Tree-sitter parser that maps out your hierarchical structures using pipeline (|) characters. It's giving peak organization and clean data reads.
 * @author Hanvim <marifatalhanif@gmail.com>
 * @license MIT
 */

/// <reference types="tree-sitter-cli/dsl" />
// @ts-check

module.exports = grammar({
  name: "tspipe",

  rules: {
    // TODO: add the actual grammar rules
    source_file: $ => "hello"
  }
});
