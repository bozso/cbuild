//
//	Package - transpiled by c4go
//
//	If you have found any issues, please raise an issue at:
//	https://github.com/Konstantin8105/c4go/
//

package lexer

import (
    "io"
    "github.com/bozso/cbuild/lexer/item"
    "github.com/bozso/cbuild/lexer/buffer"
)

type stateFn = func(*Lexer) interface{}

// Lexer - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/lex.c:17
type Lexer struct {
	in                                        io.Reader
	input, filename                           string
	length, start, pos, width, line, line_pos uint32
    items                                     item.Buffer
	state                                     stateFn
}

// lex_new - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/lex.c:35
func New(start stateFn, in io.Reader, filename string) (l Lexer) {
	l.in = in
	l.filename = filename
	l.input = nil
	l.length = 0
	l.items = item.NewBuffer()
	l.state = start
	l.line = 0
	
    return
}

func (l *Lexer) Errorf(format string, args ...interface{}) stateFn {
    msg := fmt.Sprintf(format, args...)
    
	e := item.New(message, item.Error, l.line, l.line_pos, l.pos)
	l.items.Push(e)
	return nil
}

func (l *Lexer) Next() (b byte) {
    if l.pos + 1 > l.length {
		if l.in.error_.code != 0 {
			l.Errorf("Error reading input: %s\x00", l.in.error_.message)
			return
		}
        
        // TODO: properly figure out the logic here
        // do prevoius reads into l.input need to be preserved?
		nextLength := l.length + 4096
        
        l.input = make([]byte, nextLength)
        
        _, e = lex.in.Read(lex.input)
        
        if e != nil {
            return
        }
        
		var len_ noarch.SsizeT = stream_read(lex[0].in, lex[0].input[0+int32(lex[0].length):], 4096)
		
        if len_ < noarch.SsizeT(0) {
			lex_errorf(lex, []byte("Error reading input: %s\x00"), lex[0].in[0].error_.message)
			return
		}

		l.length += len_
		l.input[l.length] = byte(0)
	}
    
	l.width = 1
    
    b = l.input[l.pos++]
    return
}

func (l *Lexer) Backup() {
	l.pos -= l.width
}

func (l* Lexer) Peek() (b byte) {
	b = l.Next()
	l.Backup()
	
    return
}

func (l* Lexer) Accept(matching string) bool {
    return strings.Contains(matching, l.Next())
}

func (l *Lexer) AcceptRun(matching string) {
	for l.Accept(matching) {}
    
	l.Backup()
}

func (l *Lexer) Ignore() {
	l.start = l.pos
}

func (l *Lexer) Emit(it item.Type) {
	l.CountNewlines()
	ii := item.New(l.input[l.start:l.pos], it, l.line, l.line_pos,
        l.start)
    
	l.items.Push(ii)
	l.start = l.pos
}

// lex_next_item - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/lex.c:128
func (l *Lexer) NextItem() item.Item {
    for l.items.length == 0 && l.state != nil {
		l.state = l.state(l).(stateFn)
	}
    
	return l.items.Next()
}

// count_newlines - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/lex.c:148
func (l *Lexer) CountNewlines() {
	for ii := l.start; ii < l.pos; ii++ {
		if l.input[ii] == '\n' {
			l.line += 1
			l.line_pos =i i + 1
		}
	}
}
