package lexer

import (
    "io"
    "github.com/bozso/cbuild/lexer/item"
)

type (
    Lexer struct {
        in io.Reader
        input, filename string
        
        
        length, start, pos, width, line, line_pos int
        
        items []item.T;
        state StateFn
    }

    StateFn func(*Lexer) 
)

func new(start StateFn, in io.Reader, filename string) (l Lexer) {
	l.in = in
    l.filename = filename
    l.items = make([]item.Item, 2)
    l.state = start
    
	return;
}

func (l *Lexer) Errorf(string fmt, args ...interface{}) StateFn {
    msg = fmt.Sprintf(fmt, args...)
    
    err := item.New(msg, item.Error, l.line, l.linePos, l.pos)
    l.items = append(l.items, err)
    
    return nil
}

func (l *Lexer) Next() byte {
	if (l.pos + 1 > l.length) {
		if (lex->in->error.code != 0) {
			l.Errorf("Error reading input: %s", lex->in->error.message)
			return 0
		}

		nextLength := l.length + 4096 + 1
		l.input = realloc(l.input, nextLength)
        
		Len := stream.read(l.in, l.input + l.length, 4096)

		if (Len < 0) {
			l.Errorf("Error reading input: %s", lex->in->error.message);
			return 0
		}

		if (Len == 0) return 0

		l.length += Len
		l.input[l.length] = 0;
	}

	l.width = 1;
	return l.input[l.pos++]
}

func (l *Lexer) Backup() {
	l.pos -= l.width
}

func (l *Lexer) Peek() (c byte) {
	c = l.Next()
	l.Backup()
    
	return
}

func (l *Lexer) Accept(matching string) bool {
	t := l.Next()
    
    // possible better implementation
    // return strings.Contains(matching, l.Next())
    
    ll := len(matching)
    
    for ii := 0; ii < ll; ii++ {
        if t == matching[ii] {
            return true
        }
    }
	return false;
}

func (l *Lexer) AcceptRun(matching string) {
	for l.Accept(matching) {
        l.Backup()
    }
}

func (l *Lexer) Ignore() {
	l.start  = l.pos
}

func (l *Lexer) Emit(it item.T) {
	l.countNewlines()
	
    i := item.New(l.input[l.start : l.pos] it, l.line, l.linePos,
        l.start)

	l.items = append(l.items, i);
	l.start = l.pos
}

func (l *Lexer) NextItem() item.T {
	while(lex->items->length == 0 && lex->state != NULL) {
		lex->state = (state_fn) lex->state(lex);
	}
    
	return buffer.next(lex->items);
}


func (l *Lexer) countNewlines() {
    for ii := lex.start; ii < lex.pos; ii++ {
        if l.input[ii] == '\n' {
            l.line++
            l.linePos = ii + 1
        }
    }
}
