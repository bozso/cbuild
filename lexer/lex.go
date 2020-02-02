//
//	Package - transpiled by c4go
//
//	If you have found any issues, please raise an issue at:
//	https://github.com/Konstantin8105/c4go/
//

package main

// #include </usr/include/string.h>
// #include </usr/include/stdio.h>
import "C"

import "reflect"
import "unsafe"
import "github.com/Konstantin8105/c4go/noarch"

// stream_error_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/../deps/stream/stream.h:7
type stream_error_t struct {
	message []byte
	code    int32
}

// stream_read_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/../deps/stream/stream.h:12
type stream_read_t = func(interface{}, interface{}, uint32, []stream_error_t) noarch.SsizeT

// stream_write_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/../deps/stream/stream.h:13
type stream_write_t = func(interface{}, interface{}, uint32, []stream_error_t) noarch.SsizeT

// stream_pipe_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/../deps/stream/stream.h:14
type stream_pipe_t = func(interface{}, interface{}, []stream_error_t) noarch.SsizeT

// stream_close_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/../deps/stream/stream.h:15
type stream_close_t = func(interface{}, []stream_error_t) noarch.SsizeT

// stream_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/../deps/stream/stream.h:17
type stream_t struct {
	ctx    interface{}
	read   stream_read_t
	write  stream_write_t
	pipe   stream_pipe_t
	close  stream_close_t
	error_ stream_error_t
	type_  int32
}

// lex_item_type - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/item.h:7
type lex_item_type = int32

const (
	item_error         lex_item_type = 0
	item_eof                         = 1
	item_whitespace                  = 2
	item_c_code                      = 3
	item_id                          = 4
	item_number                      = 5
	item_char_literal                = 6
	item_quoted_string               = 7
	item_preprocessor                = 8
	item_comment                     = 9
	item_symbol                      = 10
	item_open_symbol                 = 11
	item_close_symbol                = 12
	item_arrow                       = 13
	item_total_symbols               = 14
)

// lex_item_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/item.h:30
type lex_item_t struct {
	type_    lex_item_type
	value    []byte
	length   uint32
	line     uint32
	line_pos uint32
	start    uint32
	index    uint32
}

// lex_buffer_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/buffer.h:8
type lex_buffer_t struct {
	items    []lex_item_t
	capacity uint32
	length   uint32
	cursor   uint32
}

// lex_state_fn - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/lex.c:15
type lex_state_fn = func([]lex_lexer_s) interface{}

// lex_lexer_s - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/lex.c:17
type lex_lexer_s struct {
	in       []stream_t
	input    []byte
	filename []byte
	length   uint32
	start    uint32
	pos      uint32
	width    uint32
	line     uint32
	line_pos uint32
	items    []lex_buffer_t
	state    lex_state_fn
}

// lex_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/lex.c:17
type lex_t = lex_lexer_s

// lex_new - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/lex.c:35
func lex_new(start lex_state_fn, in []stream_t, filename []byte) []lex_t {
	var lex []lex_t = make([]lex_t, 1)
	lex[0].in = in
	lex[0].filename = noarch.Strdup(filename)
	lex[0].input = nil
	lex[0].length = 0
	lex[0].items = lex_buffer_new(2)
	lex[0].state = start
	lex[0].line = 0
	return lex
}

// lex_errorf - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/lex.c:49
func lex_errorf(lex []lex_t, fmt_ []byte, c4goArgs ...interface{}) lex_state_fn {
	var args = create_va_list(c4goArgs)
	va_start(args, fmt_)
	var message []byte
	vasprintf((*[1000000][]byte)(unsafe.Pointer(&message))[:], fmt_, args)
	var error_ lex_item_t = lex_item_new(message, item_error, uint32(lex[0].line), uint32(lex[0].line_pos), uint32(lex[0].pos))
	lex[0].items = lex_buffer_push(lex[0].items, error_)
	va_end(args)
	return nil
}

// lex_next - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/lex.c:62
func lex_next(lex []lex_t) byte {
	if uint32(lex[0].pos)+1 > uint32(lex[0].length) {
		if lex[0].in[0].error_.code != 0 {
			lex_errorf(lex, []byte("Error reading input: %s\x00"), lex[0].in[0].error_.message)
			return byte(0)
		}
		var next_length uint32 = uint32(lex[0].length) + 4096 + 1
		lex[0].input = realloc(lex[0].input, next_length).([]byte)
		var len_ noarch.SsizeT = stream_read(lex[0].in, lex[0].input[0+int32(lex[0].length):], 4096)
		if len_ < noarch.SsizeT(0) {
			lex_errorf(lex, []byte("Error reading input: %s\x00"), lex[0].in[0].error_.message)
			return byte(0)
		}
		if len_ == noarch.SsizeT(0) {
			return byte(0)
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
