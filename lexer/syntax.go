//
//	Package - transpiled by c4go
//
//	If you have found any issues, please raise an issue at:
//	https://github.com/Konstantin8105/c4go/
//

package lexer


// lex_state_fn - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/lex.h:9
type lex_state_fn = func([]lex_lexer_s) interface{}

// lex_syntax_new - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/syntax.c:22
func lex_syntax_new(input []stream_t, filename []byte, error_ [][]byte) []lex_t {
	// declaration for state functions
	//static void * lex_export(lexer.t * lex);
	return lex_new(lex_c, input, filename)
}

// emit_c_code - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/syntax.c:26
func emit_c_code(lex []lex_t, fn lex_state_fn) interface{} {
	lex_backup(lex)
	if uint32(lex[0].pos) > uint32(lex[0].start) {
		lex_emit(lex, item_c_code)
	}
	if len(fn) == 0 {
		lex_next(lex)
	}
	return lex_state_fn(fn)
}

// eof - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/syntax.c:38
func eof(lex []lex_t) interface{} {
	lex_backup(lex)
	if uint32(lex[0].pos) > uint32(lex[0].start) {
		lex_emit(lex, item_c_code)
	}
	lex_emit(lex, item_eof)
	return nil
}

// lex_c - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/syntax.c:50
func lex_c(lex []lex_t) (c4goDefaultReturn interface{}) {
	var next byte
	for 1 != 0 {
		var c byte = lex_next(lex)
		if int32(c) == 0 {
			// lexes standard C code looking for things that might be modular c
			return eof(lex)
		}
		if int32(((__ctype_b_loc())[0])[int32(c)])&int32(uint16(noarch.ISspace)) != 0 {
			return emit_c_code(lex, lex_whitespace)
		}
		if int32(((__ctype_b_loc())[0])[int32(c)])&int32(uint16(noarch.ISalpha)) != 0 {
			return emit_c_code(lex, lex_id)
		}
		if int32(((__ctype_b_loc())[0])[int32(c)])&int32(uint16(noarch.ISdigit)) != 0 {
			return emit_c_code(lex, lex_number)
		}
		switch int32(c) {
		case '\'':
			return emit_c_code(lex, lex_squote)
		case '"':
			return emit_c_code(lex, lex_quote)
		case ';':
			fallthrough
		case '.':
			fallthrough
		case ',':
			fallthrough
		case '=':
			fallthrough
		case '*':
			emit_c_code(lex, nil)
			lex_emit(lex, item_symbol)
		case '-':
			if int32(lex_peek(lex)) == int32('>') {
				lex_next(lex)
				lex_emit(lex, item_arrow)
			}
		case '#':
			if uint32(lex[0].pos) > 2 && int32(lex[0].input[uint32(lex[0].pos)-2]) == int32('\n') {
				return emit_c_code(lex, lex_preprocessor)
			}
		case '_':
			return emit_c_code(lex, lex_id)
		case '/':
			next = lex_peek(lex)
			if int32(next) == int32('/') || int32(next) == int32('*') {
				return emit_c_code(lex, lex_comment)
			}
		case '(':
			fallthrough
		case '[':
			fallthrough
		case '{':
			emit_c_code(lex, nil)
			lex_emit(lex, item_open_symbol)
		case ')':
			fallthrough
		case ']':
			fallthrough
		case '}':
			emit_c_code(lex, nil)
			lex_emit(lex, item_close_symbol)
			break
		}
	}
	return
}

// lex_whitespace - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/syntax.c:117
func lex_whitespace(lex []lex_t) interface{} {
	var c byte
	for int32((func() byte {
		c = lex_next(lex)
		return c
	}())) != 0 && int32(((__ctype_b_loc())[0])[int32(c)])&int32(uint16(noarch.ISspace)) != 0 {
	}
	lex_backup(lex)
	lex_emit(lex, item_whitespace)
	if int32(c) == 0 {
		return eof(lex)
	}
	return lex_c
}

// lex_oneline_comment - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/syntax.c:127
func lex_oneline_comment(lex []lex_t) interface{} {
	var c byte
	for int32((func() byte {
		c = lex_next(lex)
		return c
	}())) != 0 && int32(c) != int32('\n') {
	}
	lex_emit(lex, item_comment)
	if int32(c) == 0 {
		return eof(lex)
	}
	return lex_c
}

// lex_multiline_comment - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/syntax.c:135
func lex_multiline_comment(lex []lex_t) interface{} {
	var c byte
	for {
		for int32((func() byte {
			c = lex_next(lex)
			return c
		}())) != 0 && int32(c) != int32('*') {
		}
		if int32(c) == 0 {
			lex_errorf(lex, []byte("Unterminated multiline comment\n %*s\x00"), lex[0].input[0+int32(lex[0].start):])
			return eof(lex)
		}
		if !(int32(lex_peek(lex)) != int32('/')) {
			break
		}
	}
	lex_next(lex)
	lex_emit(lex, item_comment)
	return lex_c
}

// lex_comment - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/syntax.c:152
func lex_comment(lex []lex_t) (c4goDefaultReturn interface{}) {
	lex[0].pos += uint32(1)
	var c byte = lex_next(lex)
	if int32(c) == int32('/') {
		return lex_oneline_comment(lex)
	} else {
		return lex_multiline_comment(lex)
	}
	return
}

// lex_number - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/syntax.c:160
func lex_number(lex []lex_t) interface{} {
	var c byte
	for int32((func() byte {
		c = lex_next(lex)
		return c
	}())) != 0 && int32(((__ctype_b_loc())[0])[int32(c)])&int32(uint16(noarch.ISdigit)) != 0 {
	}
	lex_backup(lex)
	lex_emit(lex, item_number)
	if int32(c) == 0 {
		return eof(lex)
	}
	return lex_c
}

// lex_id - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/syntax.c:170
func lex_id(lex []lex_t) interface{} {
	var c byte
	for int32((func() byte {
		c = lex_next(lex)
		return c
	}())) != 0 && (int32(((__ctype_b_loc())[0])[int32(c)])&int32(uint16(noarch.ISalnum)) != 0 || int32(c) == int32('_')) {
	}
	lex_backup(lex)
	lex_emit(lex, item_id)
	if int32(c) == 0 {
		return eof(lex)
	}
	return lex_c
}

// lex_quote - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/syntax.c:181
func lex_quote(lex []lex_t) interface{} {
	lex[0].pos += uint32(1)
	var c byte
	for {
		c = lex_next(lex)
		if int32(c) == int32('\\') && int32(lex_next(lex)) != 0 {
			c = lex_next(lex)
			continue
		}
		if !(int32(c) != 0 && int32(c) != int32('"') && int32(c) != int32('\n')) {
			break
		}
	}
	if int32(c) == 0 || int32(c) == int32('\n') {
		return lex_errorf(lex, []byte("Missing terminating '\"' character\n\x00"))
	}
	lex_emit(lex, item_quoted_string)
	return lex_c
}

// lex_squote - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/syntax.c:201
func lex_squote(lex []lex_t) interface{} {
	lex[0].pos += uint32(1)
	var c byte
	for {
		c = lex_next(lex)
		if int32(c) == int32('\\') && int32(lex_next(lex)) != 0 {
			c = lex_next(lex)
			continue
		}
		if !(int32(c) != 0 && int32(c) != int32('\'')) {
			break
		}
	}
	if int32(c) == 0 {
		var length uint32 = uint32(lex[0].pos) - uint32(lex[0].start)
		if length > 10 {
			return lex_errorf(lex, []byte("Missing terminating ' character\n%.*s...\x00"), 10, lex[0].input[0+int32(lex[0].start):])
		} else {
			return lex_errorf(lex, []byte("Missing terminating ' character\n%s\x00"), lex[0].input[0+int32(lex[0].start):])
		}
	}
	lex_emit(lex, item_char_literal)
	return lex_c
}

// lex_preprocessor - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/syntax.c:226
func lex_preprocessor(lex []lex_t) interface{} {
	var c byte = lex_next(lex)
	for int32(c) != int32('\n') && int32(c) != 0 {
		c = lex_next(lex)
	}
	lex_backup(lex)
	lex_emit(lex, item_preprocessor)
	return lex_c
}

// __ctype_b_loc from ctype.h
// c function : const unsigned short int** __ctype_b_loc()
// dep pkg    : unicode
// dep func   :
func __ctype_b_loc() [][]uint16 {
	var characterTable []uint16

	for i := 0; i < 255; i++ {
		var c uint16

		// Each of the bitwise expressions below were copied from the enum
		// values, like _ISupper, etc.

		if unicode.IsUpper(rune(i)) {
			c |= ((1 << (0)) << 8)
		}

		if unicode.IsLower(rune(i)) {
			c |= ((1 << (1)) << 8)
		}

		if unicode.IsLetter(rune(i)) {
			c |= ((1 << (2)) << 8)
		}

		if unicode.IsDigit(rune(i)) {
			c |= ((1 << (3)) << 8)
		}

		if unicode.IsDigit(rune(i)) ||
			(i >= 'a' && i <= 'f') ||
			(i >= 'A' && i <= 'F') {
			// IsXDigit. This is the same implementation as the Mac version.
			// There may be a better way to do this.
			c |= ((1 << (4)) << 8)
		}

		if unicode.IsSpace(rune(i)) {
			c |= ((1 << (5)) << 8)
		}

		if unicode.IsPrint(rune(i)) {
			c |= ((1 << (6)) << 8)
		}

		// The IsSpace check is required because Go treats spaces as graphic
		// characters, which C does not.
		if unicode.IsGraphic(rune(i)) && !unicode.IsSpace(rune(i)) {
			c |= ((1 << (7)) << 8)
		}

		// http://www.cplusplus.com/reference/cctype/isblank/
		// The standard "C" locale considers blank characters the tab
		// character ('\t') and the space character (' ').
		if i == int('\t') || i == int(' ') {
			c |= ((1 << (8)) >> 8)
		}

		if unicode.IsControl(rune(i)) {
			c |= ((1 << (9)) >> 8)
		}

		if unicode.IsPunct(rune(i)) {
			c |= ((1 << (10)) >> 8)
		}

		if unicode.IsLetter(rune(i)) || unicode.IsDigit(rune(i)) {
			c |= ((1 << (11)) >> 8)
		}

		// Yes, I know this is a hideously slow way to do it but I just want to
		// test if this works right now.
		characterTable = append(characterTable, c)
	}
	return [][]uint16{characterTable}
}
