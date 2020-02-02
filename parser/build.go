//
//	Package - transpiled by c4go
//
//	If you have found any issues, please raise an issue at:
//	https://github.com/Konstantin8105/c4go/
//

/* AST Error :
cannot parse line: `UnusedAttr 0x6c03330 <line:15:451> Inherited unused`. could not match regexp with string
^(?P<address>[0-9a-fx]+) <(?P<position>.*)>(?P<unused> unused)?[\s]*$
0x6c03330 <line:15:451> Inherited unused

*/

package parser

type parser_parse_fn = func([]parser_parser_s) interface{}

// options - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/build.c:14
var options []hash_t

// parse_fn - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/build.c:16
type parse_fn = func([]parser_t) int32

// init_options - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/build.c:21
func init_options() {
	options = kh_init_ptr()
	hash_set(options, []byte("depends\x00"), parse_depends)
	hash_set(options, []byte("set\x00"), parse_set)
	hash_set(options, []byte("append\x00"), parse_append)
}

// errorf - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/build.c:29
func errorf(p []parser_t, item lex_item_t, fmt_ []byte, c4goArgs ...interface{}) int32 {
	var args = create_va_list(c4goArgs)
	va_start(args, fmt_)
	parser_verrorf(p, item, []byte("Invalid build syntax: \x00"), fmt_, args)
	return -1
}

// build_parse - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/build.c:46
func build_parse(p []parser_t) int32 {
	if len(options) == 0 {
		//
		// * Build options manipulate the resulting makefile.
		// *
		// * Parses the following syntaxes:
		// * - build depends      "<filename>";          # add "<filenme>" to the dependency list for the current file
		// * - build set          <variable> "<value>";  # set the value of a makefile valiable         (:=)
		// * - build set default  <variable> "<value>";  # set the default value of a makefile variable (?=)
		// * - build append       <variable> "<value>";  # append to a makefile variable                (+=)
		//
		init_options()
	}
	var item lex_item_t = parser_skip(p, item_whitespace, 0)
	var fn parse_fn = hash_get(options, item.value).(parse_fn)
	lex_item_free(item)
	if len(fn) != 0 {
		return fn(p)
	}
	return errorf(p, item, []byte("Expecting one of \n\t'depends', 'set', 'set default' or 'append', but got %s\x00"), lex_item_to_string(item))
}

// parse_depends - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/build.c:60
func parse_depends(p []parser_t) int32 {
	var filename lex_item_t = parser_skip(p, item_whitespace, 0)
	if uint32(int32((filename.type_))) != uint32(int32((item_quoted_string))) {
		return errorf(p, filename, []byte("Expecting a filename but got '%s'\x00"), filename.value)
	}
	var semicolon lex_item_t = parser_skip(p, item_whitespace, 0)
	if uint32(int32((semicolon.type_))) != uint32(int32((item_symbol))) && int32(semicolon.value[0]) != int32(';') {
		return errorf(p, semicolon, []byte("Expecting ';' but got '%s'\x00"), lex_item_to_string(semicolon))
	}
	lex_item_free(semicolon)
	var error_ []byte
	var imp []package_import_t = package_import_add_c_file(p[0].pkg, strings_dup(string_parse(filename.value)), (*[1000000][]byte)(unsafe.Pointer(&error_))[:])
	if len(imp) == 0 {
		parser_errorf(p, filename, []byte("\x00"), []byte("Error adding dependency: %s\x00"), error_)
		return -1
	}
	lex_item_free(filename)
	return 1
}

// parse_set - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/build.c:86
func parse_set(p []parser_t) int32 {
	var is_default int32
	var have_name int32
	var name lex_item_t
	for {
		name = parser_skip(p, item_whitespace, 0)
		if uint32(int32((name.type_))) != uint32(int32((item_id))) {
			return errorf(p, name, []byte("Expecting a variable name, but got %s\x00"), name.value)
		}
		if noarch.Strcmp(name.value, []byte("default\x00")) == 0 {
			is_default = 1
			lex_item_free(name)
		} else {
			have_name = 1
		}
		if !(int32((have_name)) == 0) {
			break
		}
	}
	var value lex_item_t = parser_skip(p, item_whitespace, 0)
	if uint32(int32((value.type_))) != uint32(int32((item_quoted_string))) {
		return errorf(p, value, []byte("Experting a quoted string, but got '%s'\x00"), value.value)
	}
	var semicolon lex_item_t = parser_skip(p, item_whitespace, 0)
	if uint32(int32((semicolon.type_))) != uint32(int32((item_symbol))) && int32(semicolon.value[0]) != int32(';') {
		return errorf(p, semicolon, []byte("Expecting ';' but got '%s'\x00"), lex_item_to_string(semicolon))
	}
	lex_item_free(semicolon)
	var v package_var_t = package_var_t{nil, nil, 0}
	v.name = strings_dup(name.value)
	v.value = strings_dup(string_parse(value.value))
	v.operation = package_var_type((func() int32 {
		if int32((is_default)) != 0 {
			return int32((build_var_set_default))
		}
		return int32((build_var_set))
	}()))
	p[0].pkg[0].variables = realloc(p[0].pkg[0].variables, 40*(uint32(p[0].pkg[0].n_variables)+1)).([]package_var_t)
	p[0].pkg[0].variables[uint32(p[0].pkg[0].n_variables)] = v
	p[0].pkg[0].n_variables += uint32(1)
	lex_item_free(name)
	lex_item_free(value)
	return 1
}

// parse_append - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/build.c:132
func parse_append(p []parser_t) int32 {
	var name lex_item_t = parser_skip(p, item_whitespace, 0)
	if uint32(int32((name.type_))) != uint32(int32((item_id))) {
		return errorf(p, name, []byte("Expecting a variable name, but got %s\x00"), name.value)
	}
	var value lex_item_t = parser_skip(p, item_whitespace, 0)
	if uint32(int32((value.type_))) != uint32(int32((item_quoted_string))) {
		return errorf(p, value, []byte("Experting a quoted string, but got '%s'\x00"), value.value)
	}
	var semicolon lex_item_t = parser_skip(p, item_whitespace, 0)
	if uint32(int32((semicolon.type_))) != uint32(int32((item_symbol))) && int32(semicolon.value[0]) != int32(';') {
		return errorf(p, semicolon, []byte("Expecting ';' but got '%s'\x00"), lex_item_to_string(semicolon))
	}
	lex_item_free(semicolon)
	var v package_var_t = package_var_t{nil, nil, 0}
	v.name = strings_dup(name.value)
	v.value = strings_dup(string_parse(value.value))
	v.operation = build_var_append
	p[0].pkg[0].variables = realloc(p[0].pkg[0].variables, 40*(uint32(p[0].pkg[0].n_variables)+1)).([]package_var_t)
	p[0].pkg[0].variables[uint32(p[0].pkg[0].n_variables)] = v
	p[0].pkg[0].n_variables += uint32(1)
	lex_item_free(name)
	lex_item_free(value)
	return 1
}
