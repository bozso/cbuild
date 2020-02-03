//
//	Package - transpiled by c4go
//
//	If you have found any issues, please raise an issue at:
//	https://github.com/Konstantin8105/c4go/
//

package lexer

func NewSyntax(input io.Reader, filename string) Lexer {
	return New(lex_c, input, filename)
}

func (l *Lexer) emit_c_code(fn stateFn) interface{} {
	l.Backup()
	
    if l.pos > l.start {
		l.Emit(item.CCode)
	}
	
    if fn == nil {
		l.Next()
	}
	
    return fn
}

func (l *Lexer) Eof() interface{} {
	l.Backup()
	
    if l.pos > l.start {
		l.Emit(item.CCode)
	}
	
    l.Emit(item.Eof)
	
    return nil
}

func lex_c(l *Lexer) interface{} {
	var next byte
	
    for {
		c := l.Next()
		
        if int32(c) == 0 {
			// lexes standard C code looking for things that might be modular c
			return l.Eof()
		}
        
        // translate this to idiomatic go - probably using the unicode package
		if c == 0      return eof(lex);
		if isspace(c)  return emit_c_code(lex, lex_whitespace);
		if isalpha(c)  return emit_c_code(lex, lex_id);
		if isdigit(c) return emit_c_code(lex, lex_number);
		
        //if int32(((__ctype_b_loc())[0])[int32(c)])&int32(uint16(noarch.ISspace)) != 0 {
			//return emit_c_code(lex, lex_whitespace)
		//}
		
        //if int32(((__ctype_b_loc())[0])[int32(c)])&int32(uint16(noarch.ISalpha)) != 0 {
			//return emit_c_code(lex, lex_id)
		//}
		
        //if int32(((__ctype_b_loc())[0])[int32(c)])&int32(uint16(noarch.ISdigit)) != 0 {
			//return emit_c_code(lex, lex_number)
		//}
        
		switch int32(c) {

		case '\'':
			return l.emit_c_code(lex_squote)

		case '"':
			return l.emit_c_code(lex_quote)

		case ';', '.', ',', '=', '*':
			l.emit_c_code(nil)
			l.Emit(item.Symbol)

		case '-':
			if int32(l.Peek()) == int32('>') {
				l.Next()
				l.Emit(item.Arrow)
			}
		case '#':
			if l.pos > 2 && l.input[l.pos - 2] == '\n' {
				return l.emit_c_code(lex_preprocessor)
			}
		case '_':
			return l.emit_c_code(lex_id)
		case '/':
			next = l.Peek()
			if next == '/' || next == '*' {
				return l.emit_c_code(lex_comment)
			}
		case '(', '[', '{':
			l.emit_c_code(nil)
			l.Emit(item.OpenSymbol)

		case ')', ']', '}':
			l.emit_c_code(nil)
			l.Emit(item.CloseSymbol)
		}
	}
	return
}

func lex_whitespace(l *Lexer) interface{} {
	for (c := l.Next()) != 0 && isspace(c) {}

	l.Backup()
	l.Emit(item.Whitespace)
    
	if int32(c) == 0 {
		return l.Eof()
	}
    
	return lex_c
}

func lex_oneline_comment(l *Lexer) interface{} {
	for (c := l.Next()) != 0 && c != '\n' {}

	l.Emit(item.Comment)
    
	if int32(c) == 0 {
		return l.Eof()
	}
    
	return lex_c
}

func lex_multiline_comment(l *Lexer) interface{} {
	var c byte
    
    for {
		for (c = l.Next()) != 0 && c != '*' {}
        
		if int32(c) == 0 {
			l.Errorf("Unterminated multiline comment\n %*s\x00",
                l.input[l.start:])
            
			return l.Eof()
		}
        
		if !(int32(l.Peek()) != int32('/')) {
			break
		}
	}
    
	l.Next()
	l.Emit(item.Comment)
    
	return lex_c
}

func lex_comment(l *Lexer) interface{} {
	l.pos += 1
	c := l.Next()
    
	if int32(c) == int32('/') {
		return lex_oneline_comment(l)
	} else {
		return lex_multiline_comment(l)
	}
    
	return
}

func lex_number(l *Lexer) interface{} {
	var c byte
	for (c = l.Next()) != 0 && isdigit(c) {}

	l.Backup()
	l.Emit(item.Number)
    
	if int32(c) == 0 {
		return l.Eof()
	}
    
	return lex_c
}

func lex_id(l *Lexer) interface{} {
	var c byte
	for (c = l.Next()) != 0 && (isalnum(c) || c == '_') {}
	
    l.Backup()
	l.Emit(item.Id)
	
    if int32(c) == 0 {
		return l.Eof()
	}
    
	return lex_c
}

func lex_quote(l *Lexer) interface{} {
	l.pos += 1
    
	var c byte
	for {
		c := l.Next()
        
		if int32(c) == int32('\\') && int32(l.Next()) != 0 {
			c = l.Next()
			continue
		}
        
		if !(int32(c) != 0 && int32(c) != int32('"') && int32(c) != int32('\n')) {
			break
		}
	}
    
	if int32(c) == 0 || int32(c) == int32('\n') {
		return l.Errorf("Missing terminating '\"' character\n")
	}
    
	l.Emit(item.QuotedString)
	
    return lex_c
}

func lex_squote(l *) interface{} {
	l.pos += 1
	
    var c byte
	for {
		c := l.Next()
        
		if int32(c) == int32('\\') && int32(l.Next()) != 0 {
			c = l.Next()
			continue
		}
        
		if !(int32(c) != 0 && int32(c) != int32('\'')) {
			break
		}
	}
    
	if int32(c) == 0 {
		length := l.pos - l.start
        
		if length > 10 {
			return l.Errorf("Missing terminating ' character\n%.*s...",
                10, l.input[l.start:])
		} else {
			return l.Errorf("Missing terminating ' character\n%s"),
                l.input[l.start:])
		}
	}
    
	l.Emit(item.CharLiteral)
	return lex_c
}

func lex_preprocessor(l *Lexer) interface{} {
	c := l.Next()
    
	for int32(c) != int32('\n') && int32(c) != 0 {
		c = l.Next()
	}
    
	l.Backup()
	l.Emit(item.Preprocessor)
    
	return lex_c
}
