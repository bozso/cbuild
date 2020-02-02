//
//	Package - transpiled by c4go
//
//	If you have found any issues, please raise an issue at:
//	https://github.com/Konstantin8105/c4go/
//

package parser

// parser_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/parser.h:15
type parser_t = parser_parser_s

// package_export_type - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../package/export.h:4
type package_export_type = int32

const (
	type_block    package_export_type = 0
	type_type                         = 1
	type_function                     = 2
	type_enum                         = 3
	type_union                        = 4
	type_struct                       = 5
	type_header                       = 6
)

// package_export_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../package/export.h:14
type package_export_t struct {
	local_name  []byte
	export_name []byte
	declaration []byte
	symbol      []byte
	type_       package_export_type
}

// package_import_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../package/import.h:8
type package_import_t struct {
	alias    []byte
	filename []byte
	c_file   int32
	pkg      []package_t
}

// enum_i - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:17
var enum_i lex_item_t = lex_item_t{item_id, []byte("enum\x00"), 4, 0, 0, 0, 0}

// union_i - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:18
var union_i lex_item_t = lex_item_t{item_id, []byte("union\x00"), 5, 0, 0, 0, 0}

// struct_i - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:19
var struct_i lex_item_t = lex_item_t{item_id, []byte("struct\x00"), 6, 0, 0, 0, 0}

// decl_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:21
type decl_t struct {
	items  []lex_item_t
	length uint32
	error_ int32
}

// export_types - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:167
var export_types []hash_t

// export_fn - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:177
type export_fn = func([]parser_t, []decl_t) lex_item_t

// errorf - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:27
func errorf(p []parser_t, item lex_item_t, decl []decl_t, fmt_ []byte, c4goArgs ...interface{}) lex_item_t {
	var args = create_va_list(c4goArgs)
	va_start(args, fmt_)
	parser_verrorf(p, item, []byte("Invalid export syntax: \x00"), fmt_, args)
	decl[0].error_ = 1
	return lex_item_empty
}

// append_ - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:36
func append_(decl []decl_t, value lex_item_t) {
	decl[0].items = realloc(decl[0].items, 40*(uint32(decl[0].length)+1)).([]lex_item_t)
	decl[0].items[uint32(decl[0].length)] = value
	decl[0].length += uint32(1)
}

// rewind_until - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:42
func rewind_until(p []parser_t, decl []decl_t, item lex_item_t) {
	var last noarch.SsizeT = noarch.SsizeT(uint32(decl[0].length) - 1)
	for last >= noarch.SsizeT(0) && noarch.Not(lex_item_equals(item, decl[0].items[last])) {
		parser_backup(p, decl[0].items[last])
		last--
	}
	parser_backup(p, decl[0].items[last])
	decl[0].length = uint32(last)
}

// rewind_whitespace - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:54
func rewind_whitespace(p []parser_t, decl []decl_t, item lex_item_t) {
	parser_backup(p, item)
	var last noarch.SsizeT = noarch.SsizeT(uint32(decl[0].length) - 1)
	for last >= noarch.SsizeT(0) && uint32(int32((decl[0].items[last].type_))) == uint32(int32((item_whitespace))) {
		parser_backup(p, decl[0].items[last])
		last--
	}
	decl[0].length = uint32(last + noarch.SsizeT(1))
}

// symbol_rename - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:66
func symbol_rename(p []parser_t, decl []decl_t, name lex_item_t, alias lex_item_t) lex_item_t {
	var i int32
	for i = 0; uint32(i) < uint32(decl[0].length); i++ {
		var item lex_item_t = decl[0].items[i]
		if int32((lex_item_equals(name, item))) != 0 {
			var symbol_name []byte
			var export_name []byte = func() []byte {
				if uint32(int32((alias.type_))) == 0 {
					return name.value
				}
				return alias.value
			}()
			asprintf((*[1000000][]byte)(unsafe.Pointer(&symbol_name))[:], []byte("%s_%s\x00"), p[0].pkg[0].name, export_name)
			item = lex_item_replace_value(item, symbol_name)
			decl[0].items[i] = item
			return item
		}
	}
	return lex_item_empty
}

// free_decl - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:84
func free_decl(decl []decl_t) {
	var i int32
	for i = 0; uint32(i) < uint32(decl[0].length); i++ {
		lex_item_free(decl[0].items[i])
	}
	_ = decl[0].items
}

// emit - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:92
func emit(p []parser_t, decl []decl_t, is_function int32, is_extern int32) []byte {
	var length uint32
	var start int32
	var end int32 = int32(uint32(decl[0].length))
	var i int32
	for start < end && uint32(int32((decl[0].items[start].type_))) == uint32(int32((item_whitespace))) {
		// trim leading/trailing whitespace
		start++
	}
	for end > start && uint32(int32((decl[0].items[end-1].type_))) == uint32(int32((item_whitespace))) {
		end--
	}
	for i = 0; uint32(i) < uint32(decl[0].length); i++ {
		var item lex_item_t = decl[0].items[i]
		if noarch.Not(is_extern) {
			package_emit(p[0].pkg, item.value)
		}
		if i >= start && i < end {
			length += uint32(item.length)
		}
	}
	if int32((is_function)) != 0 {
		length++
	}
	var output []byte = make([]byte, length+1)
	length = 0
	for i = start; i < end; i++ {
		var item lex_item_t = decl[0].items[i]
		memcpy(output[0+length:], item.value, uint32(item.length))
		length += uint32(item.length)
	}
	if int32((is_function)) != 0 && noarch.Not(is_extern) {
		output[func() uint32 {
			defer func() {
				length++
			}()
			return length
		}()] = ';'
	}
	output[length] = byte(0)
	return output
}

// collect_newlines - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:124
func collect_newlines(p []parser_t, decl []decl_t) lex_item_t {
	var line uint32 = uint32(p[0].lexer[0].line)
	var item lex_item_t = parser_next(p)
	for uint32(int32((item.type_))) == uint32(int32((item_whitespace))) || uint32(int32((item.type_))) == uint32(int32((item_comment))) {
		if uint32(p[0].lexer[0].line) != line {
			append_(decl, item)
		} else {
			lex_item_free(item)
		}
		line = uint32(p[0].lexer[0].line)
		item = parser_next(p)
	}
	return item
}

// collect - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:142
func collect(p []parser_t, decl []decl_t) lex_item_t {
	var item lex_item_t = parser_next(p)
	for uint32(int32((item.type_))) == uint32(int32((item_whitespace))) || uint32(int32((item.type_))) == uint32(int32((item_comment))) {
		append_(decl, item)
		item = parser_next(p)
	}
	return item
}

// init_export_types - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:168
func init_export_types() {
	export_types = kh_init_ptr()
	hash_set(export_types, []byte("typedef\x00"), parse_typedef)
	hash_set(export_types, []byte("struct\x00"), parse_struct)
	hash_set(export_types, []byte("enum\x00"), parse_enum)
	hash_set(export_types, []byte("union\x00"), parse_union)
}

// parse_semicolon - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:179
func parse_semicolon(p []parser_t, decl []decl_t) {
	if int32((decl[0].error_)) != 0 {
		return
	}
	var item lex_item_t = collect_newlines(p, decl)
	if uint32(int32((item.type_))) != uint32(int32((item_symbol))) || int32(item.value[0]) != int32(';') {
		errorf(p, item, decl, []byte("expecting ';' but got %s\x00"), lex_item_to_string(item))
		return
	}
	append_(decl, item)
}

// export_parse - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:191
func export_parse(p []parser_t) int32 {
	if len(export_types) == 0 {
		init_export_types()
	}
	var decl decl_t = decl_t{nil, 0, 0}
	var fn export_fn
	var name lex_item_t = lex_item_t{lex_item_type((0)), nil, 0, 0, 0, 0, 0}
	var alias lex_item_t = lex_item_t{lex_item_type((0)), nil, 0, 0, 0, 0, 0}
	var has_semicolon int32
	var is_extern int32
	var t int32
	var type_ lex_item_t = collect_newlines(p, (*[1000000]decl_t)(unsafe.Pointer(&decl))[:])
	switch uint32(int32((type_.type_))) {
	case uint32(int32((item_id))):
		if noarch.Strcmp([]byte("extern\x00"), type_.value) == 0 {
			append_((*[1000000]decl_t)(unsafe.Pointer(&decl))[:], type_)
			is_extern = 1
			type_ = collect(p, (*[1000000]decl_t)(unsafe.Pointer(&decl))[:])
		}
		fn = hash_get(export_types, type_.value).(export_fn)
		has_semicolon = int32((noarch.BoolToInt(int32((is_extern)) != 0 || len(fn) != 0)))
		t = 1
	case uint32(int32((item_symbol))):
		if int32(type_.value[0]) == int32('*') {
			lex_item_free(type_)
			return parse_passthrough(p)
		}
	case uint32(int32((item_open_symbol))):
		if int32(type_.value[0]) == int32('{') {
			fn = parse_export_block
			lex_item_free(type_)
			type_.value = []byte("{\x00")
			t = 0
		}
	default:
		fn = nil
		break
	}
	if len(fn) == 0 {
		parser_backup(p, type_)
		type_.type_ = item_symbol
		type_.value = []byte("variable\x00")
		type_.length = strings_len([]byte("variable\x00"))
		name = parse_variable(p, (*[1000000]decl_t)(unsafe.Pointer(&decl))[:])
		t = 2
	} else {
		if t == 1 {
			append_((*[1000000]decl_t)(unsafe.Pointer(&decl))[:], type_)
		}
		name = fn(p, (*[1000000]decl_t)(unsafe.Pointer(&decl))[:])
	}
	alias = parse_as(p, (*[1000000]decl_t)(unsafe.Pointer(&decl))[:])
	if int32((has_semicolon)) != 0 {
		parse_semicolon(p, (*[1000000]decl_t)(unsafe.Pointer(&decl))[:])
	}
	if int32((decl.error_)) != 0 {
		free_decl((*[1000000]decl_t)(unsafe.Pointer(&decl))[:])
		return -1
	}
	if uint32(int32((name.type_))) == 0 {
		errorf(p, type_, (*[1000000]decl_t)(unsafe.Pointer(&decl))[:], []byte("can't export anonymous symbols\x00"))
		free_decl((*[1000000]decl_t)(unsafe.Pointer(&decl))[:])
		return -1
	}
	var original_name lex_item_t = lex_item_dup(name)
	var symbol lex_item_t = symbol_rename(p, (*[1000000]decl_t)(unsafe.Pointer(&decl))[:], name, alias)
	if uint32(int32((symbol.type_))) == 0 && (int64(uintptr(unsafe.Pointer(&fn[0])))/int64(1)-int64(uintptr(unsafe.Pointer(&parse_export_block[0])))/int64(1)) != 0 {
		errorf(p, type_, (*[1000000]decl_t)(unsafe.Pointer(&decl))[:], []byte("unable to export symbol '%s'\x00"), name.value)
		free_decl((*[1000000]decl_t)(unsafe.Pointer(&decl))[:])
		return -1
	}
	package_export_add(strings_dup(original_name.value), strings_dup(alias.value), strings_dup(symbol.value), type_.value, emit(p, (*[1000000]decl_t)(unsafe.Pointer(&decl))[:], int32((noarch.BoolToInt(t == 2))), is_extern), p[0].pkg)
	lex_item_free(original_name)
	lex_item_free(alias)
	free_decl((*[1000000]decl_t)(unsafe.Pointer(&decl))[:])
	return 1
}

// parse_passthrough - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:281
func parse_passthrough(p []parser_t) int32 {
	var decl decl_t = decl_t{nil, 0, 0}
	var from lex_item_t = collect(p, (*[1000000]decl_t)(unsafe.Pointer(&decl))[:])
	if uint32(int32((from.type_))) != uint32(int32((item_id))) || noarch.Strcmp(from.value, []byte("from\x00")) != 0 {
		parser_errorf(p, from, []byte("Exporting passthrough: \x00"), []byte("expected 'from', but got %s\x00"), lex_item_to_string(from))
		free_decl((*[1000000]decl_t)(unsafe.Pointer(&decl))[:])
		return -1
	}
	lex_item_free(from)
	var filename lex_item_t = collect(p, (*[1000000]decl_t)(unsafe.Pointer(&decl))[:])
	if uint32(int32((filename.type_))) != uint32(int32((item_quoted_string))) {
		parser_errorf(p, filename, []byte("Exporting passthrough: \x00"), []byte("expected filename, but got %s\x00"), lex_item_to_string(filename))
		free_decl((*[1000000]decl_t)(unsafe.Pointer(&decl))[:])
		return -1
	}
	parse_semicolon(p, (*[1000000]decl_t)(unsafe.Pointer(&decl))[:])
	if int32((decl.error_)) != 0 {
		free_decl((*[1000000]decl_t)(unsafe.Pointer(&decl))[:])
		return -1
	}
	var error_ []byte
	var imp []package_import_t = package_import_passthrough(p[0].pkg, strings_dup(string_parse(filename.value)), (*[1000000][]byte)(unsafe.Pointer(&error_))[:])
	if len(imp) == 0 {
		parser_errorf(p, filename, []byte("Exporting passthrough: \x00"), []byte("%s\x00"), error_)
		free_decl((*[1000000]decl_t)(unsafe.Pointer(&decl))[:])
		return -1
	}
	lex_item_free(filename)
	var header []byte
	var rel []byte = utils_relative(p[0].pkg[0].source_abs, imp[0].pkg[0].header)
	asprintf((*[1000000][]byte)(unsafe.Pointer(&header))[:], []byte("#include \"%s\"\x00"), rel)
	package_emit(p[0].pkg, header)
	free_decl((*[1000000]decl_t)(unsafe.Pointer(&decl))[:])
	_ = rel
	_ = header
	return 1
}

// parse_as - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:324
func parse_as(p []parser_t, decl []decl_t) lex_item_t {
	if int32((decl[0].error_)) != 0 {
		return lex_item_empty
	}
	var start uint32 = uint32(decl[0].length)
	var as lex_item_t = collect(p, decl)
	if uint32(int32((as.type_))) != uint32(int32((item_id))) || noarch.Strcmp([]byte("as\x00"), as.value) != 0 {
		parser_backup(p, as)
		for uint32(decl[0].length) > start {
			decl[0].length -= uint32(1)
			parser_backup(p, decl[0].items[uint32(decl[0].length)])
		}
		return lex_item_empty
	}
	lex_item_free(as)
	for uint32(decl[0].length) > start {
		decl[0].length -= uint32(1)
		parser_backup(p, decl[0].items[uint32(decl[0].length)])
	}
	var alias lex_item_t = collect_newlines(p, decl)
	if uint32(int32((alias.type_))) != uint32(int32((item_id))) {
		errorf(p, alias, decl, []byte("expecting identifier but got %s\x00"), lex_item_to_string(alias))
		return lex_item_empty
	}
	return alias
}

// parse_function_ptr - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:355
func parse_function_ptr(p []parser_t, decl []decl_t) lex_item_t {
	var star lex_item_t = collect(p, decl)
	if uint32(int32((star.type_))) != uint32(int32((item_symbol))) || int32(star.value[0]) != int32('*') {
		return errorf(p, star, decl, []byte("function pointer: expecting '*' but found '%s'\x00"), star.value)
	}
	append_(decl, star)
	var name lex_item_t = collect(p, decl)
	if uint32(int32((name.type_))) != uint32(int32((item_id))) {
		return errorf(p, name, decl, []byte("function pointer: expecting identifier but found '%s'\x00"), name.value)
	}
	append_(decl, name)
	var item lex_item_t
	item = collect(p, decl)
	if uint32(int32((item.type_))) != uint32(int32((item_close_symbol))) || int32(item.value[0]) != int32(')') {
		return errorf(p, item, decl, []byte("function pointer: expecting ')' but found '%s'\x00"), item.value)
	}
	append_(decl, item)
	item = collect(p, decl)
	if uint32(int32((item.type_))) != uint32(int32((item_open_symbol))) || int32(item.value[0]) != int32('(') {
		return errorf(p, item, decl, []byte("function pointer: expecting '(' but found '%s'\x00"), item.value)
	}
	append_(decl, item)
	parse_function_args(p, decl)
	if int32((decl[0].error_)) != 0 {
		return lex_item_empty
	}
	return name
}

// parse_declaration - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:387
func parse_declaration(p []parser_t, decl []decl_t) lex_item_t {
	var type_ lex_item_t = collect(p, decl)
	if uint32(int32((type_.type_))) != uint32(int32((item_id))) {
		if uint32(int32((type_.type_))) == uint32(int32((item_close_symbol))) && int32(type_.value[0]) == int32('}') {
			return type_
		}
		return errorf(p, type_, decl, []byte("in declaration: expecting identifier but got %s\x00"), lex_item_to_string(type_))
	}
	if noarch.Strcmp([]byte("as\x00"), type_.value) == 0 {
		return type_
	}
	type_ = parser_identifier_parse(p, type_, 1)
	var fn export_fn = hash_get(export_types, type_.value).(export_fn)
	if (int64(uintptr(unsafe.Pointer(&fn[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&parse_typedef[0])))/int64(1)) == 0 {
		return errorf(p, type_, decl, []byte("in declaration: unexpected identifier 'typedef'\x00"))
	}
	if len(fn) == 0 {
		fn = parse_variable
	}
	append_(decl, type_)
	var item lex_item_t = fn(p, decl)
	if int32((decl[0].error_)) != 0 {
		return lex_item_empty
	}
	if (int64(uintptr(unsafe.Pointer(&fn[0])))/int64(1)-int64(uintptr(unsafe.Pointer(&parse_enum[0])))/int64(1)) == 0 || (int64(uintptr(unsafe.Pointer(&fn[0])))/int64(1)-int64(uintptr(unsafe.Pointer(&parse_union[0])))/int64(1)) == 0 || (int64(uintptr(unsafe.Pointer(&fn[0])))/int64(1)-int64(uintptr(unsafe.Pointer(&parse_struct[0])))/int64(1)) == 0 {
		item = collect(p, decl)
		if uint32(int32((item.type_))) != uint32(int32((item_id))) {
			parser_backup(p, item)
		} else {
			append_(decl, item)
		}
	}
	parse_semicolon(p, decl)
	if int32((decl[0].error_)) != 0 {
		return lex_item_empty
	}
	return item
}

// parse_typedef - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:425
func parse_typedef(p []parser_t, decl []decl_t) lex_item_t {
	var type_ lex_item_t = collect(p, decl)
	if uint32(int32((type_.type_))) != uint32(int32((item_id))) {
		return errorf(p, type_, decl, []byte("in typedef: expected identifier\x00"))
	}
	var fn export_fn = hash_get(export_types, type_.value).(export_fn)
	//if (fn != parse_struct && fn != parse_enum && fn != parse_union) append(decl, type);
	append_(decl, type_)
	if len(fn) != 0 {
		fn(p, decl)
		if int32((decl[0].error_)) != 0 {
			return lex_item_empty
		}
	}
	var item lex_item_t
	var name lex_item_t = lex_item_t{lex_item_type((0)), nil, 0, 0, 0, 0, 0}
	var function_ptr int32
	var as int32
	for {
		item = collect(p, decl)
		switch uint32(int32((item.type_))) {
		case uint32(int32((item_eof))):
			return errorf(p, item, decl, []byte("in typedef: expected identifier or '('\x00"))
		case uint32(int32((item_id))):
			if noarch.Strcmp([]byte("as\x00"), item.value) == 0 {
				rewind_whitespace(p, decl, item)
				as = 1
				continue
			}
			if noarch.Not(function_ptr) {
				name = item
			}
		case uint32(int32((item_symbol))):
			if int32(item.value[0]) == int32(';') {
				continue
			}
			fallthrough
		case uint32(int32((item_open_symbol))):
			if noarch.Not(function_ptr) && int32(item.value[0]) == int32('(') {
				function_ptr = 1
				append_(decl, item)
				name = parse_function_ptr(p, decl)
				if int32((decl[0].error_)) != 0 {
					return name
				}
				continue
			}
			fallthrough
		default:
			break
		}
		append_(decl, item)
		if !(uint32(int32((item.type_))) != uint32(int32((item_error))) && !(uint32(int32((item.type_))) == uint32(int32((item_symbol))) && int32(item.value[0]) == int32(';')) && noarch.Not(as)) {
			break
		}
	}
	if uint32(int32((item.type_))) == uint32(int32((item_error))) {
		decl[0].error_ = 1
		return item
	}
	if noarch.Not(as) {
		parser_backup(p, item)
	}
	return name
}

// parse_struct - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:489
func parse_struct(p []parser_t, decl []decl_t) lex_item_t {
	var name lex_item_t = lex_item_t{lex_item_type((0)), nil, 0, 0, 0, 0, 0}
	var item lex_item_t = collect(p, decl)
	if uint32(int32((item.type_))) == uint32(int32((item_id))) {
		name = parser_identifier_parse_typed(p, struct_i, item, 1)
		append_(decl, name)
		item = collect(p, decl)
	}
	if uint32(int32((item.type_))) != uint32(int32((item_open_symbol))) || int32(item.value[0]) != int32('{') {
		parser_backup(p, item)
		return name
	}
	append_(decl, item)
	parse_struct_block(p, decl, item)
	return func() lex_item_t {
		if int32((decl[0].error_)) != 0 {
			return lex_item_empty
		}
		return name
	}()
}

// parse_union - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:509
func parse_union(p []parser_t, decl []decl_t) lex_item_t {
	var name lex_item_t = lex_item_t{lex_item_type((0)), nil, 0, 0, 0, 0, 0}
	var item lex_item_t = collect(p, decl)
	if uint32(int32((item.type_))) == uint32(int32((item_id))) {
		name = parser_identifier_parse_typed(p, union_i, item, 1)
		append_(decl, name)
		item = collect(p, decl)
	}
	if uint32(int32((item.type_))) != uint32(int32((item_open_symbol))) || int32(item.value[0]) != int32('{') {
		return errorf(p, item, decl, []byte("in union: expecting '{' but got %s\x00"), lex_item_to_string(item))
	}
	append_(decl, item)
	parse_struct_block(p, decl, item)
	return func() lex_item_t {
		if int32((decl[0].error_)) != 0 {
			return lex_item_empty
		}
		return name
	}()
}

// parse_enum - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:527
func parse_enum(p []parser_t, decl []decl_t) lex_item_t {
	var name lex_item_t = lex_item_t{lex_item_type((0)), nil, 0, 0, 0, 0, 0}
	var item lex_item_t = collect(p, decl)
	if uint32(int32((item.type_))) == uint32(int32((item_id))) {
		name = parser_identifier_parse_typed(p, enum_i, item, 1)
		append_(decl, name)
		item = collect(p, decl)
	}
	if uint32(int32((item.type_))) != uint32(int32((item_open_symbol))) || int32(item.value[0]) != int32('{') {
		return errorf(p, item, decl, []byte("in enum: expecting '{' but got %s\x00"), lex_item_to_string(item))
	}
	append_(decl, item)
	parse_enum_block(p, decl, item)
	return func() lex_item_t {
		if int32((decl[0].error_)) != 0 {
			return lex_item_empty
		}
		return name
	}()
}

// parse_export_block - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:545
func parse_export_block(p []parser_t, decl []decl_t) (c4goDefaultReturn lex_item_t) {
	var item lex_item_t
	var escaped_id int32
	for 1 != 0 {
		var record int32
		item = collect(p, decl)
		switch uint32(int32((item.type_))) {
		case uint32(int32((item_c_code))):
			fallthrough
		case uint32(int32((item_quoted_string))):
			fallthrough
		case uint32(int32((item_comment))):
			fallthrough
		case uint32(int32((item_symbol))):
			fallthrough
		case uint32(int32((item_preprocessor))):
			record = 1
		case uint32(int32((item_arrow))):
			escaped_id = 1
			record = 1
		case uint32(int32((item_id))):
			if escaped_id == 0 {
				item = parser_identifier_parse(p, item, 1)
			}
			record = 1
			escaped_id = 0
		case uint32(int32((item_close_symbol))):
			if int32(item.value[0]) == int32('}') {
				lex_item_free(item)
				item.value = []byte("block\x00")
				item.length = strings_len([]byte("block\x00"))
				return item
			}
		case uint32(int32((item_eof))):
			return errorf(p, item, decl, []byte("in block: unmatched '{'\x00"))
		case uint32(int32((item_error))):
			parser_backup(p, item)
			decl[0].error_ = 1
			return lex_item_empty
		default:
			return errorf(p, item, decl, []byte("in block: unexpected input '%s'\x00"), lex_item_to_string(item))
		}
		if record != 0 {
			append_(decl, item)
		}
	}
	return
}

// parse_struct_block - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:596
func parse_struct_block(p []parser_t, decl []decl_t, start lex_item_t) lex_item_t {
	var item lex_item_t
	for {
		item = parse_declaration(p, decl)
		if uint32(int32((item.type_))) == uint32(int32((item_eof))) {
			return errorf(p, start, decl, []byte("in struct block: missing matching '}' before end of file\x00"))
		}
		if uint32(int32((item.type_))) == uint32(int32((item_error))) {
			return item
		}
		if !!(uint32(int32((item.type_))) == uint32(int32((item_close_symbol))) && int32(item.value[0]) == int32('}')) {
			break
		}
	}
	append_(decl, item)
	return item
}

// parse_enum_declaration - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:609
func parse_enum_declaration(p []parser_t, decl []decl_t) lex_item_t {
	var item lex_item_t = collect(p, decl)
	if uint32(int32((item.type_))) != uint32(int32((item_id))) {
		if uint32(int32((item.type_))) == uint32(int32((item_close_symbol))) && int32(item.value[0]) == int32('}') {
			return item
		}
		return errorf(p, item, decl, []byte("in enum: expecting identifier but got %s\x00"), lex_item_to_string(item))
	}
	if noarch.Strcmp([]byte("as\x00"), item.value) == 0 {
		return item
	}
	append_(decl, item)
	item = collect(p, decl)
	if !(uint32(int32((item.type_))) == uint32(int32((item_symbol))) && (int32(item.value[0]) != int32(',') || int32(item.value[0]) != int32('='))) {
		return item
	}
	append_(decl, item)
	if int32(item.value[0]) == int32('=') {
		item = collect(p, decl)
		if uint32(int32((item.type_))) != uint32(int32((item_number))) {
			return errorf(p, item, decl, []byte("in enum: expecting a number but got %s\x00"), lex_item_to_string(item))
		}
		append_(decl, item)
		item = collect(p, decl)
		if uint32(int32((item.type_))) != uint32(int32((item_symbol))) || int32(item.value[0]) != int32(',') {
			return item
		}
		append_(decl, item)
	}
	return item
}

// parse_enum_block - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:643
func parse_enum_block(p []parser_t, decl []decl_t, start lex_item_t) lex_item_t {
	var item lex_item_t
	for {
		item = parse_enum_declaration(p, decl)
		if uint32(int32((item.type_))) == uint32(int32((item_eof))) {
			return errorf(p, start, decl, []byte("in struct block: missing matching '}' before end of file\x00"))
		}
		if uint32(int32((item.type_))) == uint32(int32((item_error))) {
			return item
		}
		if !!(uint32(int32((item.type_))) == uint32(int32((item_close_symbol))) && int32(item.value[0]) == int32('}')) {
			break
		}
	}
	append_(decl, item)
	return item
}

// parse_type - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:656
func parse_type(p []parser_t, decl []decl_t) lex_item_t {
	var item lex_item_t
	var name lex_item_t = lex_item_t{lex_item_type((0)), nil, 0, 0, 0, 0, 0}
	var escaped_id int32
	for {
		item = collect(p, decl)
		switch uint32(int32((item.type_))) {
		case uint32(int32((item_arrow))):
			escaped_id = 1
			append_(decl, item)
			continue
			fallthrough
		case uint32(int32((item_id))):
			if escaped_id == 0 {
				item = parser_identifier_parse(p, item, 1)
			}
			name = item
		case uint32(int32((item_symbol))):
			if int32(item.value[0]) == int32('*') {
				break
			}
			fallthrough
		case uint32(int32((item_open_symbol))):
			if int32(item.value[0]) == int32('(') {
				append_(decl, item)
				return name
			}
			fallthrough
		default:
			return errorf(p, item, decl, []byte("in type declaration: expecting identifier or '('\x00"))
		}
		append_(decl, item)
		if !(uint32(int32((item.type_))) != uint32(int32((item_open_symbol))) || int32(item.value[0]) != int32('(')) {
			break
		}
	}
	return lex_item_empty
}

// parse_variable - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:686
func parse_variable(p []parser_t, decl []decl_t) lex_item_t {
	var item lex_item_t
	var name lex_item_t = lex_item_t{lex_item_type((0)), nil, 0, 0, 0, 0, 0}
	var escaped_id int32
	var escape_till_semicolon int32
	for {
		item = collect(p, decl)
		if int32((escape_till_semicolon)) != 0 && (uint32(int32((item.type_))) != uint32(int32((item_symbol))) || int32(item.value[0]) != int32(';')) {
			append_(decl, item)
			continue
		}
		switch uint32(int32((item.type_))) {
		case uint32(int32((item_arrow))):
			escaped_id = 1
			append_(decl, item)
			continue
			fallthrough
		case uint32(int32((item_id))):
			if noarch.Strcmp(item.value, []byte("as\x00")) == 0 {
				rewind_whitespace(p, decl, item)
				return name
			}
			if int32((escaped_id)) == 0 {
				item = parser_identifier_parse(p, item, 1)
			}
			name = item
		case uint32(int32((item_symbol))):
			if int32(item.value[0]) == int32('*') {
				break
			}
			if int32(item.value[0]) == int32('=') {
				escape_till_semicolon = 1
				break
			}
			if int32(item.value[0]) == int32(';') {
				parser_backup(p, item)
				return name
			}
			fallthrough
		case uint32(int32((item_open_symbol))):
			if int32(item.value[0]) == int32('(') {
				var star lex_item_t = parser_next(p)
				parser_backup(p, star)
				if uint32(int32((star.type_))) == uint32(int32((item_symbol))) && int32(star.value[0]) == int32('*') {
					append_(decl, item)
					return parse_function_ptr(p, decl)
				}
				parser_backup(p, item)
				rewind_until(p, decl, name)
				return parse_function(p, decl)
			}
			if int32(item.value[0]) == int32('[') {
				append_(decl, item)
				item = parser_next(p)
				if uint32(int32((item.type_))) == uint32(int32((item_number))) {
					append_(decl, item)
					item = parser_next(p)
				}
				if uint32(int32((item.type_))) != uint32(int32((item_close_symbol))) || int32(item.value[0]) != int32(']') {
					return errorf(p, item, decl, []byte("in type declaration: expecting terminating ']' but got %s\x00"), lex_item_to_string(item))
				}
				break
			}
			fallthrough
		default:
			return errorf(p, item, decl, []byte("in type declaration: expecting identifier or '(' but got %s\x00"), lex_item_to_string(item))
		}
		append_(decl, item)
		if !(uint32(int32((item.type_))) != uint32(int32((item_eof)))) {
			break
		}
	}
	return name
}

// parse_function_args - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:755
func parse_function_args(p []parser_t, decl []decl_t) lex_item_t {
	var level int32 = 1
	var escaped_id int32
	var item lex_item_t
	for {
		item = collect(p, decl)
		switch uint32(int32((item.type_))) {
		case uint32(int32((item_arrow))):
			escaped_id = 1
			append_(decl, item)
			continue
			fallthrough
		case uint32(int32((item_id))):
			if escaped_id == 0 {
				item = parser_identifier_parse(p, item, 1)
			}
		case uint32(int32((item_open_symbol))):
			if int32(item.value[0]) == int32('(') {
				level++
			}
		case uint32(int32((item_close_symbol))):
			if int32(item.value[0]) == int32(')') {
				level--
			}
		default:
			break
		}
		escaped_id = 0
		append_(decl, item)
		if !(level > 0 && uint32(int32((item.type_))) != uint32(int32((item_eof))) && uint32(int32((item.type_))) != uint32(int32((item_error)))) {
			break
		}
	}
	return item
}

// parse_function - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:787
func parse_function(p []parser_t, decl []decl_t) lex_item_t {
	var name lex_item_t = parse_type(p, decl)
	if int32((decl[0].error_)) != 0 {
		emit(p, decl, 1, 0)
		decl[0].items = nil
		decl[0].length = 0
		return lex_item_empty
	}
	parse_function_args(p, decl)
	if int32((decl[0].error_)) != 0 {
		return lex_item_empty
	}
	return name
}
