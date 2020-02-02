//
//	Package - transpiled by c4go
//
//	If you have found any issues, please raise an issue at:
//	https://github.com/Konstantin8105/c4go/
//

package parser


type parser_parse_fn = func([]parser_parser_s) interface{}

// errorf - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/package.c:10
func errorf(p []parser_t, item lex_item_t, fmt_ []byte, c4goArgs ...interface{}) int32 {
	var args = create_va_list(c4goArgs)
	va_start(args, fmt_)
	parser_verrorf(p, item, []byte("Invalid package syntax: \x00"), fmt_, args)
	return -1
}

// parser_package_parse - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/package.c:18
func parser_package_parse(p []parser_t) int32 {
	var name lex_item_t = parser_skip(p, item_whitespace, 0)
	if uint32(int32((name.type_))) != uint32(int32((item_quoted_string))) {
		return errorf(p, name, []byte("Expecting a quoted string, but got %s\x00"), lex_item_to_string(name))
	}
	var semicolon lex_item_t = parser_skip(p, item_whitespace, 0)
	if uint32(int32((semicolon.type_))) != uint32(int32((item_symbol))) && int32(semicolon.value[0]) != int32(';') {
		return errorf(p, semicolon, []byte("Expecting ';' got %s\x00"), lex_item_to_string(semicolon))
	}
	lex_item_free(semicolon)
	_ = p[0].pkg[0].name
	p[0].pkg[0].name = strings_dup(string_parse(name.value))
	lex_item_free(name)
	return 1
}
