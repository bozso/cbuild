//
//	Package - transpiled by c4go
//
//	If you have found any issues, please raise an issue at:
//	https://github.com/Konstantin8105/c4go/
//

package main

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

// lex_buffer_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/buffer.c:9
type lex_buffer_t struct {
	items    []lex_item_t
	capacity uint32
	length   uint32
	cursor   uint32
}

// lex_buffer_new - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/buffer.c:16
func lex_buffer_new(count uint32) []lex_buffer_t {
	var b []lex_buffer_t = make([]lex_buffer_t, 1)
	b[0].items = make([]lex_item_t, count*uint32(1))
	b[0].capacity = count
	b[0].length = 0
	b[0].cursor = 0
	return b
}

// lex_buffer_resize - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/buffer.c:27
func lex_buffer_resize(b []lex_buffer_t, count uint32) []lex_buffer_t {
	if uint32(b[0].capacity) >= count {
		return b
	}
	var items []lex_item_t = make([]lex_item_t, count*uint32(1))
	var i int32
	if count > uint32(b[0].length) {
		for i = 0; uint32(i) < uint32(b[0].length); i++ {
			items[i] = b[0].items[(uint32(i)+uint32(b[0].cursor))%uint32(b[0].capacity)]
		}
	} else {
		{
			// on shrink drop items from the end
			for i = 0; uint32(i) < count; i++ {
				items[i] = b[0].items[(uint32(i)+uint32(b[0].cursor))%uint32(b[0].capacity)]
			}
		}
	}
	_ = b[0].items
	b[0].items = items
	b[0].capacity = count
	b[0].cursor = 0
	return b
}

// lex_buffer_push - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/buffer.c:55
func lex_buffer_push(b []lex_buffer_t, item lex_item_t) []lex_buffer_t {
	b = lex_buffer_resize(b, uint32(b[0].length)+1)
	var last uint32 = (uint32(b[0].length) + uint32(b[0].cursor)) % uint32(b[0].capacity)
	b[0].items[last] = item
	b[0].length += uint32(1)
	return b
}

// lex_buffer_next - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/buffer.c:65
func lex_buffer_next(b []lex_buffer_t) lex_item_t {
	if uint32(b[0].length) == 0 {
		return lex_item_new([]byte("No more items\x00"), item_error, 0, 0, 0)
	}
	var i lex_item_t = b[0].items[uint32(b[0].cursor)]
	b[0].cursor = func() uint32 {
		if uint32(b[0].cursor)+1 < uint32(b[0].capacity) {
			return uint32(b[0].cursor) + 1
		}
		return 0
	}()
	b[0].length -= uint32(1)
	return i
}

// lex_buffer_free - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/buffer.c:77
func lex_buffer_free(b []lex_buffer_t) {
	for uint32(b[0].length) != 0 {
		lex_item_free(lex_buffer_next(b))
	}
	_ = b[0].items
	_ = b
}

type _Bool int32
