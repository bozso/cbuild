package lexer

import (
    "fmt"
    "github.com/bozso/cbuild/lexer/item"
)

func New(input stream, string filename) Lexer {
    return new(lexC, input, filename) 
}

func (l *Lexer) emitCCode(fn StateFn) StateFn {
    l.Backup()
    
    if l.pos > l.start {
        l.Emit(item.CCode)
    }
    
    if fn == nil {
        l.Next()
    }

	return fn
}

func (l * Lexer) eof() {
	l.Backup()

	if (l.pos > l.start) {
        l.Emit(item.CCode)
	}
    
	l.Emit(item.Eof);

	return nil;
}

/* lexes standard C code looking for things that might be modular c */

func lex_c(l *Lexer) {
    var next byte
    
	for {
		c := l.Next()

		if c == 0 {
            return l.eof()
        }
        
		if isspace(c) {
            return l.emitCCode(lex_whitespace)
        }
        
		if isalpha(c) {
            return l.emitCCode(lex_id)
        }
        
		if isdigit(c) {
            return l.EmitCCode(lex_number)
        }

		switch c {
			case '\'':
				return l.emitCCode(lex_squote);

			case '"':
				return l.emitCCode(lex_quote);

			case ';', '.', ',', '=', '*':
				l.emitCCode(nil)
				l.Emit(item.Symbol)

			case '-':
				if (l.Peek() == '>') {
					l.Next();
					l.Emit(item.Arrow)
				}
			case '#':
				if (l.pos > 2 && l.input[l.pos - 2] == '\n') {
					return l.emitCCode(lex_preprocessor)
                }
			case '_':
				return l.emitCCode(lex_id)

			case '/':
				next = l.Peek();
				if next == '/' || next == '*' {
					return l.emitCCode(lex_comment)
                }
			
            case '(', '[', '{':
				l.emitCCode(nil)
				l.Emit(item.OpenSymbol)
            
			case ')', ']', '}':
				l.emitCCode(nil)
				l.emit(itemCloseSymbol)
		}
	}
}

func lex_whitespace(l *Lexer) StateFn {
	var c byte
    
    for (c = l.Next()) != 0 && isspace(c) {}
	
    l.Backup()
	l.Emit(item.Whitespace)

	if (c == 0) {
        return l.eof()
    }
    
	return lex_c;
}

func lex_oneline_comment(l *Lexer) StateFn {
	var c byte
    
    for (c = l.Next()) != 0 && c != '\n' {}
	
    l.Emit(item.Comment)
    
	if (c == 0) {
        return l.eof()
    }
    
	return lex_c
}

func lex_multiline_comment(l *Lexer) StateFn {
	var c byte
	
    for l.Peek() != '/' {
		for (c = l.Next()) != 0 && c != '*' {}

		if (c == 0) {
			l.Errorf("Unterminated multiline comment\n %*s",
                l.input + l.start)
            
			return l.eof()
		}

	}

	l.Next();
	l.Emit(item.Comment)
    
	return lex_c
}

func lex_comment(l *Lexer) StateFn {
	l.pos++
    
	c = l.Next()

	if (c == '/') {
        return lex_oneline_comment(l)
    } else {
        return lex_multiline_comment(l)
    }
}

func lex_number(l * Lexer) StateFn {
	var c byte
    
	for (c = l.Next()) != 0 && isdigit(c) {}
	
    l.Backup()
	l.Emit(item.Number)

	if (c == 0) {
        return l.eof()
    }
    
	return lex_c
}

func lex_id(l *Lexer) StateFn {
	var c byte
    
	for (c = l.Next()) != 0 && (isalnum(c) || c == '_') {}
    
	l.Backup();
	l.Emit(item.Id)

	if (c == 0) {
        return l.eof()
    }
    
	return lex_c
}


func lex_quote(l *Lexer) StateFn {
	l.pos++

	c = byte('\'')
    
	for c != 0 && c != '"' && c != '\n' {
		c = l.Next()
        
		if c == '\\' && l.Next() != 0 {
			c = l.Next()
			continue
		}
	}

	if c == 0 || c == '\n' {
		return l.Errorf("Missing terminating '\"' character\n");
	}

	l.Emit(item.QuotedString)
    
	return lex_c
}

func lex_squote(l *Lexer) StateFn {
	l.pos++

	c := byte('"')
    
	for c != 0 && c != '\'' {
		c = lexer.next(lex);
		if c == '\\' && l.Next() != 0 {
			c = l.Next()
			continue
		}
	}

	if (c == 0) {
		length := lex.pos - l.start;
		if (length > 10) {
			return l.Errorf("Missing terminating ' character\n%.*s...",
                10, l.input + l.start)
		} else {
			return l.Errorf("Missing terminating ' character\n%s",
                l.input + l.start)
		}
	}

	l.Emit(item.CharLiteral)
    
	return lex_c
}

func lex_preprocessor(lexer.t * lex) StateFn {
	c := l.Next()
    
	for c != '\n' && c != 0 {
		c = l.Next();
	}
    
	l.Backup();
	l.Emit(item.Preprocessor)
    
	return lex_c
}
