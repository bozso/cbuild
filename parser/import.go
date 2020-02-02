//
//	Package - transpiled by c4go
//
//	If you have found any issues, please raise an issue at:
//	https://github.com/Konstantin8105/c4go/
//

package parser

type parser_parse_fn = func([]parser_parser_s) interface{}


// package_import_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../package/import.h:8
type package_import_t struct {
	alias    []byte
	filename []byte
	c_file   int32
	pkg      []package_t
}

// errorf - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/import.c:13
func errorf(p []parser_t, item lex_item_t, fmt_ []byte, c4goArgs ...interface{}) int32 {
	var args = create_va_list(c4goArgs)
	va_start(args, fmt_)
	parser_verrorf(p, item, []byte("Invalid import syntax: \x00"), fmt_, args)
	return -1
}

// import_parse - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/import.c:22
func import_parse(p []parser_t) int32 {
	var alias lex_item_t = parser_skip(p, item_whitespace, 0)
	if uint32(int32((alias.type_))) != uint32(int32((item_id))) {
		return errorf(p, alias, []byte("Expecting identifier, but got %s\x00"), lex_item_to_string(alias))
	}
	var from lex_item_t = parser_skip(p, item_whitespace, 0)
	if uint32(int32((from.type_))) != uint32(int32((item_id))) || noarch.Strcmp(from.value, []byte("from\x00")) != 0 {
		return errorf(p, from, []byte("Expecting 'from', but got %s\x00"), lex_item_to_string(from))
	}
	lex_item_free(from)
	var filename lex_item_t = parser_skip(p, item_whitespace, 0)
	if uint32(int32((filename.type_))) != uint32(int32((item_quoted_string))) {
		return errorf(p, filename, []byte("Expecting filename, but got %s\x00"), lex_item_to_string(filename))
	}
	var semicolon lex_item_t = parser_skip(p, item_whitespace, 0)
	if uint32(int32((semicolon.type_))) != uint32(int32((item_symbol))) && int32(semicolon.value[0]) != int32(';') {
		return errorf(p, semicolon, []byte("Expecting ';', but got %s\x00"), lex_item_to_string(semicolon))
	}
	lex_item_free(semicolon)
	var error_ []byte
	var imp []package_import_t = package_import_add(strings_dup(alias.value), strings_dup(string_parse(filename.value)), p[0].pkg, (*[1000000][]byte)(unsafe.Pointer(&error_))[:])
	lex_item_free(alias)
	lex_item_free(filename)
	if len(imp) == 0 {
		return -1
	}
	if len(error_) != 0 {
		return errorf(p, filename, error_)
	}
	var include []byte
	var rel []byte = utils_relative(p[0].pkg[0].source_abs, imp[0].pkg[0].header)
	asprintf((*[1000000][]byte)(unsafe.Pointer(&include))[:], []byte("#include \"%s\"\x00"), rel)
	package_emit(p[0].pkg, include)
	_ = include
	_ = rel
	return 1
}
