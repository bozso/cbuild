package lexer

import (
    "github.com/bozso/cbuild/lexer/item"
)

type ItemBuffer struct {
	item.Item  []items;
    cursor int
}

func New(int count) (ib ItemBuffer) {
    ib.items = make([]items, count)
    ib.cursor = 0
    
    return
}

// appending to slices is already implemented in go
//export item_buffer_t * resize(item_buffer_t * b, size_t count) {
	//if (b->capacity >= count) {
		//return b;
	//}

	//lex_item.t * items = malloc(count * sizeof(lex_item.t));


	//int i;
	//if (count > b->length) {
		//for (i = 0; i < b->length; i++) {
			//items[i] = b->items[(i + b->cursor) % b->capacity];
		//}
	//} else {
		//// on shrink drop items from the end
		//for (i = 0; i < count; i++) {
			//items[i] = b->items[(i + b->cursor) % b->capacity];
		//}
	//}

	//free(b->items);
	//b->items = items;
	//b->capacity = count;
	//b->cursor = 0;

	//return b;
//}

//export item_buffer_t * push(item_buffer_t * b, lex_item.t item) {
	//b = resize(b, b->length + 1);
	//size_t last = (b->length + b->cursor) % b->capacity;

	//b->items[last] = item;
	//b->length++;

	//return b;
//}
