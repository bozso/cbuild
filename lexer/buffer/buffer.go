//
//	Package - transpiled by c4go
//
//	If you have found any issues, please raise an issue at:
//	https://github.com/Konstantin8105/c4go/
//

package buffer

type T struct {
	items    []item.T
	cursor uint32
}

// lex_buffer_new - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/buffer.c:16
func New(count uint32) (b T) {
	b.items = make([]item.T, count)
	b.cursor = 0

	return
}

// lex_buffer_push - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/buffer.c:55
func (b *T) Push(it item.T) *T {
	b = lex_buffer_resize(b, uint32(b[0].length)+1)
	var last uint32 = (uint32(b[0].length) + uint32(b[0].cursor)) % uint32(b[0].capacity)
	b.items[last] = item
	b.length += uint32(1)
	return b
}

// lex_buffer_next - transpiled function from  /home/istvan/packages/downloaded/cbuild/lexer/buffer.c:65
func (b *T) Next() item.T {
	if uint32(b.length) == 0 {
		return lex_item_new([]byte("No more items\x00"), item_error, 0, 0, 0)
	}
	
    it := b.items[uint32(b.cursor)]
	
    b.cursor = func() uint32 {
		if uint32(b.cursor)+1 < uint32(b.capacity) {
			return uint32(b.cursor) + 1
		}
		return 0
	}()
    
	b.length -= uint32(1)
	return it
}
