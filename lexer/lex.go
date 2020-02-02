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
        
		var len_ noarch.SsizeT = stream_read(lex[0].in, lex[0].input[0+int32(lex[0].length):], 4096)
		if len_ < noarch.SsizeT(0) {
			lex_errorf(lex, []byte("Error reading input: %s\x00"), lex[0].in[0].error_.message)
			return
		}

		lex[0].length += uint32(len_)
		lex[0].input[uint32(lex[0].length)] = byte(0)
	}
	lex[0].width = 1
	return lex[0].input[func() uint32 {
		tempVar := &lex[0].pos
		defer func() {
			*tempVar++
		}()
		return *tempVar
	}()]
}

// lex_backup - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/lex.c:88
func lex_backup(lex []lex_t) {
	lex[0].pos -= uint32(lex[0].width)
}

// lex_peek - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/lex.c:92
func lex_peek(lex []lex_t) byte {
	var c byte = lex_next(lex)
	lex_backup(lex)
	return c
}

// lex_accept - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/lex.c:98
func lex_accept(lex []lex_t, matching []byte) int32 {
	var t byte = lex_next(lex)
	var c []byte = matching
	for int32(c[0]) != int32('\x00') {
		if int32(t) == int32(c[0]) {
			return 1
		}
		func() []byte {
			tempVarUnary := c
			defer func() {
				c = c[0+1:]
			}()
			return tempVarUnary
		}()
	}
	return 0
}

// lex_accept_run - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/lex.c:108
func lex_accept_run(lex []lex_t, matching []byte) {
	for int32((lex_accept(lex, matching))) != 0 {
	}
	lex_backup(lex)
}

// lex_ignore - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/lex.c:113
func lex_ignore(lex []lex_t) {
	lex[0].start = uint32(lex[0].pos)
}

// lex_emit - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/lex.c:117
func lex_emit(lex []lex_t, it lex_item_type) {
	count_newlines(lex)
	var i lex_item_t = lex_item_new(substring(lex[0].input, uint32(lex[0].start), uint32(lex[0].pos)), it, uint32(lex[0].line), uint32(lex[0].line_pos), uint32(lex[0].start))
	lex[0].items = lex_buffer_push(lex[0].items, i)
	lex[0].start = uint32(lex[0].pos)
}

// lex_next_item - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/lex.c:128
func lex_next_item(lex []lex_t) lex_item_t {
	for uint32(lex[0].items[0].length) == 0 && len(lex_state_fn(lex[0].state)) != 0 {
		lex[0].state = lex[0].state(lex).(lex_state_fn)
	}
	return lex_buffer_next(lex[0].items)
}

// lex_free - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/lex.c:135
func lex_free(lex []lex_t) {
	stream_close(lex[0].in)
	lex_buffer_free(lex[0].items)
	_ = lex[0].filename
	_ = lex[0].input
	_ = lex
}

// substring - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/lex.c:144
func substring(input []byte, start uint32, end uint32) []byte {
	return strndup(input[0+start:], end-start)
}

// count_newlines - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/lex.c:148
func count_newlines(lex []lex_t) {
	var i uint32
	for i = uint32(lex[0].start); i < uint32(lex[0].pos); i++ {
		if int32(lex[0].input[i]) == int32('\n') {
			lex[0].line += uint32(1)
			lex[0].line_pos = i + 1
		}
	}
}

type _Bool int32

// memcpy is function from string.h.
// c function : void * memcpy( void * , const void * , size_t )
// dep pkg    : reflect
// dep func   :
func memcpy(dst, src interface{}, size uint32) interface{} {
	switch reflect.TypeOf(src).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(src)
		d := reflect.ValueOf(dst)
		if s.Len() == 0 {
			return dst
		}
		if s.Len() > 0 {
			size /= uint32(int(s.Index(0).Type().Size()))
		}
		var val reflect.Value
		for i := 0; i < int(size); i++ {
			if i < s.Len() {
				val = s.Index(i)
			}
			d.Index(i).Set(val)
		}
	}
	return dst
}

// realloc is function from stdlib.h.
// c function : void * realloc(void* , size_t )
// dep pkg    : reflect
// dep func   : memcpy
func realloc(ptr interface{}, size uint32) interface{} {
	if ptr == nil {
		return make([]byte, size)
	}
	elemType := reflect.TypeOf(ptr).Elem()
	ptrNew := reflect.MakeSlice(reflect.SliceOf(elemType), int(size), int(size)).Interface()
	// copy elements
	memcpy(ptrNew, ptr, size)
	return ptrNew
}

// strndup - add c-binding for implemention function
func strndup(arg0 []byte, arg1 uint32) []byte {
	if arg0 == nil {
		return []byte{}
	}
	return []byte(C.GoString(C.strndup((*C.char)(unsafe.Pointer(&arg0[0])), C.ulong(arg1))))
}

// vasprintf - add c-binding for implemention function
func vasprintf(arg0 [][]byte, arg1 []byte, arg2 *va_list) int32 {
	return int32(C.vasprintf(unsafe.Pointer(&arg0), (*C.char)(unsafe.Pointer(&arg1[0])), (*C.va_list)(unsafe.Pointer(&arg2))))
}

// va_list is C4GO implementation of va_list from "stdarg.h"
type va_list struct {
	position int
	Slice    []interface{}
}

func create_va_list(list []interface{}) *va_list {
	return &va_list{
		position: 0,
		Slice:    list,
	}
}

func va_start(v *va_list, count interface{}) {
	v.position = 0
}

func va_end(v *va_list) {
	// do nothing
}

func va_arg(v *va_list) interface{} {
	defer func() {
		v.position++
	}()
	value := v.Slice[v.position]
	switch value.(type) {
	case int:
		return int32(value.(int))
	default:
		return value
	}
}
