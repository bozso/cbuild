//
//	Package - transpiled by c4go
//
//	If you have found any issues, please raise an issue at:
//	https://github.com/Konstantin8105/c4go/
//

package parser

// init_options - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/identifier.c:19
func init_options() {
	options = kh_init_ptr()
	hash_set(options, []byte("enum\x00"), 1)
	hash_set(options, []byte("union\x00"), 2)
	hash_set(options, []byte("struct\x00"), 3)
}

// token - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/identifier.c:27
func token(p []parser_t, s []lex_item_stack_t) lex_item_t {
	var item lex_item_t = parser_next(p)
	for uint32(int32((item.type_))) == uint32(int32((item_whitespace))) {
		lex_item_stack_push(s, item)
		item = parser_next(p)
	}
	return item
}

// cleanup - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/identifier.c:36
func cleanup(p []parser_t, s []lex_item_stack_t) lex_item_t {
	for uint32(s[0].length) > 1 {
		parser_backup(p, lex_item_stack_pop(s))
	}
	var out lex_item_t = lex_item_stack_pop(s)
	lex_item_stack_free(s)
	return out
}

// rewind_until - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/identifier.c:45
func rewind_until(p []parser_t, s []lex_item_stack_t, item lex_item_t) {
	var i lex_item_t
	for uint32(s[0].length) > 0 && uint32(int32(((func() lex_item_t {
		i = lex_item_stack_pop(s)
		return i
	}()).type_))) != 0 && noarch.Not(lex_item_equals(i, item)) {
		parser_backup(p, i)
	}
}

// emit - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/identifier.c:52
func emit(p []parser_t, s []lex_item_stack_t) lex_item_t {
	var i int32
	for i = 0; uint32(i) < uint32(s[0].length)-1; i++ {
		package_emit(p[0].pkg, s[0].items[i].value)
	}
	var out lex_item_t = lex_item_stack_pop(s)
	lex_item_stack_free(s)
	return out
}

// parse_type - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/identifier.c:65
func parse_type(p []parser_t, s []lex_item_stack_t, item lex_item_t, type_ []lex_item_t) lex_item_t {
	if len(options) == 0 {
		init_options()
	}
	if noarch.Not(hash_has(options, item.value)) {
		return item
	}
	type_[0] = item
	lex_item_stack_push(s, item)
	return token(p, s)
}

// parse_import - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/identifier.c:74
func parse_import(p []parser_t, s []lex_item_stack_t, item lex_item_t, pkg []lex_item_t, name []lex_item_t) lex_item_t {
	var temp lex_item_t
	lex_item_stack_push(s, item)
	if uint32(int32((item.type_))) != uint32(int32((item_id))) {
		return lex_item_empty
	}
	temp = item
	item = parser_next(p)
	lex_item_stack_push(s, item)
	if uint32(int32((item.type_))) != uint32(int32((item_symbol))) || int32(item.value[0]) != int32('.') {
		return temp
	}
	item = parser_next(p)
	lex_item_stack_push(s, item)
	if uint32(int32((item.type_))) != uint32(int32((item_id))) {
		return lex_item_empty
	}
	pkg[0] = temp
	name[0] = item
	return item
}

// parse_symbol - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/identifier.c:96
func parse_symbol(p []parser_t, s []lex_item_stack_t, type_ lex_item_t, item lex_item_t) lex_item_t {
	var symbol []package_export_t = hash_get(p[0].pkg[0].symbols, item.value).([]package_export_t)
	if len(symbol) == 0 {
		return cleanup(p, s)
	}
	var has_type int32 = int32((noarch.BoolToInt(uint32(int32((type_.type_))) != 0)))
	rewind_until(p, s, item)
	var symbol_name []byte
	asprintf((*[1000000][]byte)(unsafe.Pointer(&symbol_name))[:], []byte("%s%s%s\x00"), func() []byte {
		if int32((has_type)) != 0 {
			return type_.value
		}
		return []byte("\x00")
	}(), func() []byte {
		if int32((has_type)) != 0 {
			return []byte(" \x00")
		}
		return []byte("\x00")
	}(), symbol[0].symbol)
	item = lex_item_replace_value(item, symbol_name)
	lex_item_stack_free(s)
	return item
}

// parser_identifier_parse_typed - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/identifier.c:112
func parser_identifier_parse_typed(p []parser_t, type_ lex_item_t, item lex_item_t, is_export int32) lex_item_t {
	var ident lex_item_t = lex_item_t{lex_item_type((0)), nil, 0, 0, 0, 0, 0}
	var from lex_item_t = lex_item_t{lex_item_type((0)), nil, 0, 0, 0, 0, 0}
	var name lex_item_t = lex_item_t{lex_item_type((0)), nil, 0, 0, 0, 0, 0}
	var s []lex_item_stack_t = lex_item_stack_new(8)
	item = parse_import(p, s, item, (*[1000000]lex_item_t)(unsafe.Pointer(&from))[:], (*[1000000]lex_item_t)(unsafe.Pointer(&name))[:])
	if uint32(int32((item.type_))) == 0 {
		return cleanup(p, s)
	}
	if uint32(int32((name.type_))) == 0 {
		return parse_symbol(p, s, lex_item_empty, item)
	}
	if noarch.Strcmp(from.value, []byte("global\x00")) == 0 {
		lex_item_stack_free(s)
		return name
	}
	var imp []package_import_t = hash_get(p[0].pkg[0].deps, from.value).([]package_import_t)
	if len(imp) == 0 || len(imp[0].pkg) == 0 {
		return emit(p, s)
	}
	var exp []package_export_t = hash_get(imp[0].pkg[0].exports, name.value).([]package_export_t)
	if len(exp) == 0 {
		parser_errorf(p, name, []byte("\x00"), []byte("Package '%s' does not export the symbol '%s'\x00"), from.value, name.value)
		return cleanup(p, s)
	}
	if int32((is_export)) != 0 {
		package_export_export_headers(p[0].pkg, imp[0].pkg)
	}
	var typed_name []byte
	asprintf((*[1000000][]byte)(unsafe.Pointer(&typed_name))[:], []byte("%s %s\x00"), type_.value, exp[0].symbol)
	ident = lex_item_replace_value(type_, typed_name)
	lex_item_stack_free(s)
	return ident
}

// parser_identifier_parse - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/identifier.c:149
func parser_identifier_parse(p []parser_t, item lex_item_t, is_export int32) lex_item_t {
	var ident lex_item_t = lex_item_t{lex_item_type((0)), nil, 0, 0, 0, 0, 0}
	var type_ lex_item_t = lex_item_t{lex_item_type((0)), nil, 0, 0, 0, 0, 0}
	var from lex_item_t = lex_item_t{lex_item_type((0)), nil, 0, 0, 0, 0, 0}
	var name lex_item_t = lex_item_t{lex_item_type((0)), nil, 0, 0, 0, 0, 0}
	var s []lex_item_stack_t = lex_item_stack_new(8)
	item = parse_type(p, s, item, (*[1000000]lex_item_t)(unsafe.Pointer(&type_))[:])
	item = parse_import(p, s, item, (*[1000000]lex_item_t)(unsafe.Pointer(&from))[:], (*[1000000]lex_item_t)(unsafe.Pointer(&name))[:])
	if uint32(int32((item.type_))) == 0 {
		return cleanup(p, s)
	}
	if uint32(int32((name.type_))) == 0 {
		return parse_symbol(p, s, type_, item)
	}
	if noarch.Strcmp(from.value, []byte("global\x00")) == 0 {
		name = lex_item_dup(name)
		lex_item_stack_free(s)
		return name
	}
	var imp []package_import_t = hash_get(p[0].pkg[0].deps, from.value).([]package_import_t)
	if len(imp) == 0 || len(imp[0].pkg) == 0 {
		return emit(p, s)
	}
	var exp []package_export_t = hash_get(imp[0].pkg[0].exports, name.value).([]package_export_t)
	if len(exp) == 0 {
		parser_errorf(p, name, []byte("\x00"), []byte("Package '%s' does not export the symbol '%s'\x00"), from.value, name.value)
		lex_item_stack_free(s)
		return cleanup(p, s)
	}
	var has_type int32 = int32((noarch.BoolToInt(uint32(int32((type_.type_))) != 0)))
	if int32((is_export)) != 0 {
		package_export_export_headers(p[0].pkg, imp[0].pkg)
	}
	var symbol_name []byte
	asprintf((*[1000000][]byte)(unsafe.Pointer(&symbol_name))[:], []byte("%s%s%s\x00"), func() []byte {
		if int32((has_type)) != 0 {
			return type_.value
		}
		return []byte("\x00")
	}(), func() []byte {
		if int32((has_type)) != 0 {
			return []byte(" \x00")
		}
		return []byte("\x00")
	}(), exp[0].symbol)
	var start lex_item_t = func() lex_item_t {
		if int32((has_type)) != 0 {
			return type_
		}
		return from
	}()
	ident = lex_item_new(symbol_name, start.type_, uint32(start.line), uint32(start.line_pos), uint32(start.start))
	lex_item_stack_free(s)
	return ident
}
