//
//	Package - transpiled by c4go
//
//	If you have found any issues, please raise an issue at:
//	https://github.com/Konstantin8105/c4go/
//

package parser

// keyword_fn - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/grammer.c:21
type keyword_fn = func([]parser_t) int32

// parse_c - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/grammer.c:26
func parse_c(p []parser_t) interface{} {
	var item lex_item_t = lex_item_t{lex_item_type((0)), nil, 0, 0, 0, 0, 0}
	var last lex_item_t = lex_item_t{lex_item_type((0)), nil, 0, 0, 0, 0, 0}
	var escaped_id int32
	for {
		lex_item_free(last)
		last = item
		item = parser_next(p)
		switch uint32(int32((item.type_))) {
		case uint32(int32((item_arrow))):
			escaped_id = 1
			package_emit(p[0].pkg, item.value)
			continue
			fallthrough
		case uint32(int32((item_eof))):
		case uint32(int32((item_error))):
			parser_errorf(p, item, []byte("\x00"), []byte("\x00"))
			lex_item_free(last)
			lex_item_free(item)
			return nil
		case uint32(int32((item_id))):
			if uint32(int32((last.type_))) == 0 || uint32(last.start) < uint32(last.line_pos) {
				lex_item_free(last)
				return parse_keyword(p, item)
			}
			if escaped_id == 0 {
				lex_item_free(last)
				return parse_id(p, item)
			}
			fallthrough
		case uint32(int32((item_c_code))):
			fallthrough
		default:
			package_emit(p[0].pkg, item.value)
		}
		escaped_id = 0
		if !(uint32(int32((item.type_))) != uint32(int32((item_eof))) && uint32(int32((item.type_))) != uint32(int32((item_error)))) {
			break
		}
	}
	lex_item_free(item)
	lex_item_free(last)
	return nil
}

// init_keywords - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/grammer.c:72
func init_keywords() {
	keywords = kh_init_ptr()
	hash_set(keywords, []byte("package\x00"), parser_package_parse)
	hash_set(keywords, []byte("import\x00"), import_parse)
	hash_set(keywords, []byte("export\x00"), export_parse)
	hash_set(keywords, []byte("build\x00"), build_parse)
}

// parse_keyword - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/grammer.c:81
func parse_keyword(p []parser_t, item lex_item_t) interface{} {
	if len(keywords) == 0 {
		init_keywords()
	}
	var fn keyword_fn = hash_get(keywords, item.value).(keyword_fn)
	if len(fn) != 0 {
		lex_item_free(item)
		if fn(p) != 0 {
			return parse_c
		}
		return nil
	}
	return parse_id(p, item)
}

// parse_id - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/grammer.c:94
func parse_id(p []parser_t, item lex_item_t) interface{} {
	item = parser_identifier_parse(p, item, 0)
	package_emit(p[0].pkg, item.value)
	lex_item_free(item)
	return parse_c
}

// grammer_parse - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/grammer.c:101
func grammer_parse(in []stream_t, filename []byte, p []package_t, error_ [][]byte) int32 {
	var lexer []lex_t = lex_syntax_new(in, filename, error_)
	if len(lexer) == 0 {
		return -1
	}
	return parser_parse(lexer, parse_c, p)
}
