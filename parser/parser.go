//
//	Package - transpiled by c4go
//
//	If you have found any issues, please raise an issue at:
//	https://github.com/Konstantin8105/c4go/
//

package parser

// parser_parse_fn - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/parser.c:18
type parser_parse_fn = func([]parser_parser_s) interface{}

// parser_parser_s - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/parser.c:20
type parser_parser_s struct {
	lexer  []lex_t
	state  parser_parse_fn
	items  []lex_item_stack_t
	pkg    []package_t
	errors int32
}

// parser_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/parser.c:20
type parser_t = parser_parser_s


// parser_parse - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/parser.c:28
func parser_parse(lexer []lex_t, start parser_parse_fn, pkg []package_t) int32 {
	var p []parser_t = make([]parser_t, 1)
	p[0].lexer = lexer
	p[0].state = start
	p[0].items = lex_item_stack_new(1)
	p[0].pkg = pkg
	p[0].errors = 0
	var cwd []byte = getcwd(nil, 0)
	var directory []byte = noarch.Strdup(lexer[0].filename)
	chdir(dirname(directory))
	_ = directory
	for len(parser_parse_fn(p[0].state)) != 0 {
		p[0].state = p[0].state(p).(parser_parse_fn)
	}
	lex_free(lexer)
	lex_item_stack_free(p[0].items)
	chdir(cwd)
	_ = cwd
	var errors int32 = p[0].errors
	_ = p
	if errors > 0 {
		noarch.Fprintf(noarch.Stderr, []byte("%d error%s generated\n\x00"), errors, func() []byte {
			if errors == 1 {
				return []byte("\x00")
			}
			return []byte("s\x00")
		}())
	}
	return errors
}

// parser_next - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/parser.c:56
func parser_next(p []parser_t) lex_item_t {
	var item lex_item_t
	if uint32(p[0].items[0].length) != 0 {
		item = lex_item_stack_pop(p[0].items)
	} else {
		item = lex_next_item(p[0].lexer)
	}
	return item
}

// parser_backup - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/parser.c:66
func parser_backup(p []parser_t, item lex_item_t) {
	p[0].items = lex_item_stack_push(p[0].items, item)
}

// parser_skip - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/parser.c:70
func parser_skip(p []parser_t, c4goArgs ...interface{}) lex_item_t {
	var args = create_va_list(c4goArgs)
	var len_ int32
	var i int32
	var types []lex_item_type = make([]lex_item_type, 14)
	va_start(args, p)
	for {
		types[len_] = va_arg(args).(lex_item_type)
		len_++
		if !(uint32(int32((types[len_-1]))) != 0) {
			break
		}
	}
	len_--
	var item lex_item_t = lex_item_t{lex_item_type((0)), nil, 0, 0, 0, 0, 0}
	for {
		var do_skip int32
		lex_item_free(item)
		item = parser_next(p)
		for i = 0; i < len_; i++ {
			if uint32(int32((types[i]))) == uint32(int32((item.type_))) {
				do_skip = 1
				break
			}
		}
		if do_skip != 0 {
			continue
		}
		return item
		if !(uint32(int32((item.type_))) != uint32(int32((item_eof))) && uint32(int32((item.type_))) != uint32(int32((item_error)))) {
			break
		}
	}
	return item
}

// parser_collect - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/parser.c:102
func parser_collect(p []parser_t, buf [][]byte, blen []noarch.SsizeT, c4goArgs ...interface{}) lex_item_t {
	var args = create_va_list(c4goArgs)
	var len_ int32
	var i int32
	var text []byte = func() []byte {
		if len(buf) == 0 {
			return nil
		}
		return buf[0]
	}()
	var text_len uint32 = uint32(func() int32 {
		if len(blen) == 0 {
			return 0
		}
		return int32(blen[0])
	}())
	var types []lex_item_type = make([]lex_item_type, 14)
	va_start(args, blen)
	for {
		types[len_] = va_arg(args).(lex_item_type)
		len_++
		if !(uint32(int32((types[len_-1]))) != 0) {
			break
		}
	}
	len_--
	var item lex_item_t
	for {
		var do_skip int32
		item = parser_next(p)
		for i = 0; i < len_; i++ {
			if uint32(int32((types[i]))) == uint32(int32((item.type_))) {
				do_skip = 1
				text = realloc(text, text_len+uint32(item.length)+1).([]byte)
				noarch.Strcpy(text[0+text_len:], item.value)
				text_len += uint32(item.length)
				break
			}
		}
		if do_skip != 0 {
			continue
		}
		break
		if !(uint32(int32((item.type_))) != uint32(int32((item_eof))) && uint32(int32((item.type_))) != uint32(int32((item_error)))) {
			break
		}
	}
	buf[0] = text
	blen[0] = noarch.SsizeT(text_len)
	return item
}

// parser_verrorf - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/parser.c:141
func parser_verrorf(p []parser_t, item lex_item_t, context []byte, fmt_ []byte, args *va_list) {
	var start uint32 = uint32(item.start)
	var line uint32 = uint32(item.line)
	var line_pos uint32 = uint32(item.line_pos)
	var col int32 = int32(start - line_pos)
	p[0].errors++
	noarch.Fprintf(noarch.Stderr, []byte("\x1b[1m%s:%ld:%d: \x1b[0m\x1b[1m\x1b[31merror: \x1b[0m\x1b[1m%s\x00"), p[0].lexer[0].filename, line+1, col+1, context)
	if uint32(int32((item.type_))) == uint32(int32((item_error))) {
		noarch.Fprintf(noarch.Stderr, []byte("%s\x00"), lex_item_to_string(item))
	} else {
		noarch.Vfprintf(noarch.Stderr, fmt_, args)
	}
	if len(p[0].lexer[0].input) != 0 && uint32(p[0].lexer[0].length) > line_pos {
		var line_start []byte = p[0].lexer[0].input[0+line_pos:]
		var line_end []byte = noarch.Strchr(line_start, int32('\n'))
		var line_length int32 = int32(func() uint32 {
			if len(line_end) == 0 {
				return uint32(col) + uint32(item.length)
			}
			return uint32(int32((int64(uintptr(unsafe.Pointer(&line_end[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&line_start[0])))/int64(1))))
		}())
		noarch.Fprintf(noarch.Stderr, []byte("\x1b[0m\n%.*s\n\x00"), line_length, line_start)
		noarch.Fprintf(noarch.Stderr, []byte("\x1b[32m%*.*s^\n\x1b[0m\x00"), col, col, []byte(" \x00"))
	} else {
		noarch.Fprintf(noarch.Stderr, []byte("\x1b[0m\n\x00"))
	}
}

// parser_errorf - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/parser.c:174
func parser_errorf(p []parser_t, item lex_item_t, context []byte, fmt_ []byte, c4goArgs ...interface{}) {
	var args = create_va_list(c4goArgs)
	va_start(args, fmt_)
	parser_verrorf(p, item, context, fmt_, args)
}
