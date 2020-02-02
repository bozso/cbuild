//
//	Package - transpiled by c4go
//
//	If you have found any issues, please raise an issue at:
//	https://github.com/Konstantin8105/c4go/
//

package main

import "reflect"

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

// lex_item_stack_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/stack.c:8
type lex_item_stack_t struct {
	items    []lex_item_t
	capacity uint32
	length   uint32
}

// lex_item_stack_new - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/stack.c:14
func lex_item_stack_new(count uint32) []lex_item_stack_t {
	var s []lex_item_stack_t = make([]lex_item_stack_t, 1)
	s[0].items = make([]lex_item_t, count*uint32(1))
	s[0].capacity = count
	s[0].length = 0
	return s
}

// lex_item_stack_resize - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/stack.c:24
func lex_item_stack_resize(s []lex_item_stack_t, count uint32) []lex_item_stack_t {
	if uint32(s[0].capacity) >= count {
		return s
	}
	s[0].items = realloc(s[0].items, count*40).([]lex_item_t)
	s[0].capacity = count
	return s
}

// lex_item_stack_push - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/stack.c:35
func lex_item_stack_push(s []lex_item_stack_t, item lex_item_t) []lex_item_stack_t {
	s = lex_item_stack_resize(s, uint32(s[0].length)+1)
	s[0].items[uint32(s[0].length)] = item
	s[0].length += uint32(1)
	return s
}

// lex_item_stack_pop - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/stack.c:44
func lex_item_stack_pop(s []lex_item_stack_t) lex_item_t {
	if uint32(s[0].length) == 0 {
		return lex_item_new([]byte("No more items\x00"), item_error, 0, 0, 0)
	}
	s[0].length -= uint32(1)
	return s[0].items[uint32(s[0].length)]
}

// lex_item_stack_free - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/stack.c:53
func lex_item_stack_free(s []lex_item_stack_t) {
	var i int32
	for i = 0; uint32(i) < uint32(s[0].length); i++ {
		lex_item_free(s[0].items[i])
	}
	s[0].length = 0
	_ = s[0].items
	s[0].items = nil
	_ = s
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
