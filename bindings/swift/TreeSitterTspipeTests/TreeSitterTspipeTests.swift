import XCTest
import SwiftTreeSitter
import TreeSitterTspipe

final class TreeSitterTspipeTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_tspipe())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading tree-sitter-pipe grammar")
    }
}
