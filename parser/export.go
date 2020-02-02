//
//	Package - transpiled by c4go
//
//	If you have found any issues, please raise an issue at:
//	https://github.com/Konstantin8105/c4go/
//

/* AST Error :
cannot parse line: `UnusedAttr 0x6f518d0 <line:15:451> Inherited unused`. could not match regexp with string
^(?P<address>[0-9a-fx]+) <(?P<position>.*)>(?P<unused> unused)?[\s]*$
0x6f518d0 <line:15:451> Inherited unused

*/

package main

// #include </usr/include/stdio.h>
import "C"

import "reflect"
import "unsafe"
import "github.com/Konstantin8105/c4go/noarch"

// khint32_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../deps/hash/khash.h:135
/* AST Error :
cannot parse line: `UnusedAttr 0x6f518d0 <line:15:451> Inherited unused`. could not match regexp with string
^(?P<address>[0-9a-fx]+) <(?P<position>.*)>(?P<unused> unused)?[\s]*$
0x6f518d0 <line:15:451> Inherited unused

*/type khint32_t = uint32

// khint64_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../deps/hash/khash.h:141
type khint64_t = uint32

// khint_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../deps/hash/khash.h:162
type khint_t = khint32_t

// khiter_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../deps/hash/khash.h:163
type khiter_t = khint_t

// __ac_HASH_UPPER - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../deps/hash/khash.h:192
var __ac_HASH_UPPER float64 = 0.77

// kh_cstr_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../deps/hash/khash.h:611
type kh_cstr_t = []byte

// kh_ptr_s - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../deps/hash/hash.h:15
type kh_ptr_s struct {
	n_buckets   khint_t
	size        khint_t
	n_occupied  khint_t
	upper_bound khint_t
	flags       []khint32_t
	keys        []kh_cstr_t
	vals        []interface{}
}

// kh_ptr_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../deps/hash/hash.h:15
type kh_ptr_t = kh_ptr_s

// hash_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../deps/hash/hash.h:21
type hash_t = kh_ptr_t

// parser_parse_fn - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/parser.h:9
type parser_parse_fn = func([]parser_parser_s) interface{}

// lex_state_fn - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../lexer/lex.h:9
type lex_state_fn = func([]lex_lexer_s) interface{}

// stream_error_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../lexer/../deps/stream/stream.h:7
type stream_error_t struct {
	message []byte
	code    int32
}

// stream_read_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../lexer/../deps/stream/stream.h:12
type stream_read_t = func(interface{}, interface{}, uint32, []stream_error_t) noarch.SsizeT

// stream_write_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../lexer/../deps/stream/stream.h:13
type stream_write_t = func(interface{}, interface{}, uint32, []stream_error_t) noarch.SsizeT

// stream_pipe_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../lexer/../deps/stream/stream.h:14
type stream_pipe_t = func(interface{}, interface{}, []stream_error_t) noarch.SsizeT

// stream_close_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../lexer/../deps/stream/stream.h:15
type stream_close_t = func(interface{}, []stream_error_t) noarch.SsizeT

// stream_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../lexer/../deps/stream/stream.h:17
type stream_t struct {
	ctx    interface{}
	read   stream_read_t
	write  stream_write_t
	pipe   stream_pipe_t
	close  stream_close_t
	error_ stream_error_t
	type_  int32
}

// lex_item_type - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../lexer/item.h:7
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

// lex_item_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../lexer/item.h:30
type lex_item_t struct {
	type_    lex_item_type
	value    []byte
	length   uint32
	line     uint32
	line_pos uint32
	start    uint32
	index    uint32
}

// lex_buffer_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../lexer/buffer.h:8
type lex_buffer_t struct {
	items    []lex_item_t
	capacity uint32
	length   uint32
	cursor   uint32
}

// lex_lexer_s - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../lexer/lex.h:14
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

// lex_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../lexer/lex.h:14
type lex_t = lex_lexer_s

// lex_item_stack_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../lexer/stack.h:8
type lex_item_stack_t struct {
	items    []lex_item_t
	capacity uint32
	length   uint32
}

// package_var_type - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../package/package.h:8
type package_var_type = int32

const (
	build_var_set         package_var_type = 0
	build_var_set_default                  = 1
	build_var_append                       = 2
)

// package_var_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../package/package.h:14
type package_var_t struct {
	name      []byte
	value     []byte
	operation package_var_type
}

// package_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../package/package.h:22
type package_t struct {
	deps        []hash_t
	exports     []hash_t
	symbols     []hash_t
	ordered     []interface{}
	n_exports   uint32
	variables   []package_var_t
	n_variables uint32
	name        []byte
	source_abs  []byte
	generated   []byte
	header      []byte
	errors      uint32
	exported    int32
	c_file      int32
	force       int32
	silent      int32
	out         []stream_t
}

// parser_parser_s - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/parser.h:15
type parser_parser_s struct {
	lexer  []lex_t
	state  parser_parse_fn
	items  []lex_item_stack_t
	pkg    []package_t
	errors int32
}

// parser_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/parser.h:15
type parser_t = parser_parser_s

// package_export_type - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../package/export.h:4
type package_export_type = int32

const (
	type_block    package_export_type = 0
	type_type                         = 1
	type_function                     = 2
	type_enum                         = 3
	type_union                        = 4
	type_struct                       = 5
	type_header                       = 6
)

// package_export_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../package/export.h:14
type package_export_t struct {
	local_name  []byte
	export_name []byte
	declaration []byte
	symbol      []byte
	type_       package_export_type
}

// package_import_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../package/import.h:8
type package_import_t struct {
	alias    []byte
	filename []byte
	c_file   int32
	pkg      []package_t
}

// enum_i - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:17
var enum_i lex_item_t = lex_item_t{item_id, []byte("enum\x00"), 4, 0, 0, 0, 0}

// union_i - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:18
var union_i lex_item_t = lex_item_t{item_id, []byte("union\x00"), 5, 0, 0, 0, 0}

// struct_i - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:19
var struct_i lex_item_t = lex_item_t{item_id, []byte("struct\x00"), 6, 0, 0, 0, 0}

// decl_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:21
type decl_t struct {
	items  []lex_item_t
	length uint32
	error_ int32
}

// export_types - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:167
var export_types []hash_t

// export_fn - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:177
type export_fn = func([]parser_t, []decl_t) lex_item_t

// __ac_X31_hash_string - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../deps/hash/khash.h:395
func __ac_X31_hash_string(s []byte) khint_t {
	var h khint_t = khint_t(s[0])
	if uint32((khint32_t((h)))) != 0 {
		// The MIT License
		//
		//   Copyright (c) 2008, 2009, 2011 by Attractive Chaos <attractor@live.co.uk>
		//
		//   Permission is hereby granted, free of charge, to any person obtaining
		//   a copy of this software and associated documentation files (the
		//   "Software"), to deal in the Software without restriction, including
		//   without limitation the rights to use, copy, modify, merge, publish,
		//   distribute, sublicense, and/or sell copies of the Software, and to
		//   permit persons to whom the Software is furnished to do so, subject to
		//   the following conditions:
		//
		//   The above copyright notice and this permission notice shall be
		//   included in all copies or substantial portions of the Software.
		//
		//   THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
		//   EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
		//   MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
		//   NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS
		//   BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN
		//   ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
		//   CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
		//   SOFTWARE.
		//
		//
		//  An example:
		//
		//#include "khash.h"
		//KHASH_MAP_INIT_INT(32, char)
		//int main() {
		// int ret, is_missing;
		// khiter_t k;
		// khash_t(32) *h = kh_init(32);
		// k = kh_put(32, h, 5, &ret);
		// kh_value(h, k) = 10;
		// k = kh_get(32, h, 10);
		// is_missing = (k == kh_end(h));
		// k = kh_get(32, h, 5);
		// kh_del(32, h, k);
		// for (k = kh_begin(h); k != kh_end(h); ++k)
		//  if (kh_exist(h, k)) kh_value(h, k) = 1;
		// kh_destroy(32, h);
		// return 0;
		//}
		//
		//
		//  2013-05-02 (0.2.8):
		//
		// * Use quadratic probing. When the capacity is power of 2, stepping function
		//   i*(i+1)/2 guarantees to traverse each bucket. It is better than double
		//   hashing on cache performance and is more robust than linear probing.
		//
		//   In theory, double hashing should be more robust than quadratic probing.
		//   However, my implementation is probably not for large hash tables, because
		//   the second hash function is closely tied to the first hash function,
		//   which reduce the effectiveness of double hashing.
		//
		// Reference: http://research.cs.vt.edu/AVresearch/hashing/quadratic.php
		//
		//  2011-12-29 (0.2.7):
		//
		//    * Minor code clean up; no actual effect.
		//
		//  2011-09-16 (0.2.6):
		//
		// * The capacity is a power of 2. This seems to dramatically improve the
		//   speed for simple keys. Thank Zilong Tan for the suggestion. Reference:
		//
		//    - http://code.google.com/p/ulib/
		//    - http://nothings.org/computer/judy/
		//
		// * Allow to optionally use linear probing which usually has better
		//   performance for random input. Double hashing is still the default as it
		//   is more robust to certain non-random input.
		//
		// * Added Wang's integer hash function (not used by default). This hash
		//   function is more robust to certain non-random input.
		//
		//  2011-02-14 (0.2.5):
		//
		//    * Allow to declare global functions.
		//
		//  2009-09-26 (0.2.4):
		//
		//    * Improve portability
		//
		//  2008-09-19 (0.2.3):
		//
		// * Corrected the example
		// * Improved interfaces
		//
		//  2008-09-11 (0.2.2):
		//
		// * Improved speed a little in kh_put()
		//
		//  2008-09-10 (0.2.1):
		//
		// * Added kh_clear()
		// * Fixed a compiling error
		//
		//  2008-09-02 (0.2.0):
		//
		// * Changed to token concatenation which increases flexibility.
		//
		//  2008-08-31 (0.1.2):
		//
		// * Fixed a bug in kh_get(), which has not been tested previously.
		//
		//  2008-08-31 (0.1.1):
		//
		// * Added destructor
		//
		//!
		//  @header
		//
		//  Generic hash table library.
		//
		// compiler specific configuration
		// --- BEGIN OF HASH FUNCTIONS ---
		//! @function
		//  @abstract     Integer hash function
		//  @param  key   The integer [khint32_t]
		//  @return       The hash value [khint_t]
		//
		//! @function
		//  @abstract     Integer comparison function
		//
		//! @function
		//  @abstract     64-bit integer hash function
		//  @param  key   The integer [khint64_t]
		//  @return       The hash value [khint_t]
		//
		//! @function
		//  @abstract     64-bit integer comparison function
		//
		//! @function
		//  @abstract     const char* hash function
		//  @param  s     Pointer to a null terminated string
		//  @return       The hash value
		//
		for func() []byte {
			tempVarUnary := s
			defer func() {
				s = s[0+1:]
			}()
			return tempVarUnary
		}(); s[0] != 0; func() []byte {
			tempVarUnary := s
			defer func() {
				s = s[0+1:]
			}()
			return tempVarUnary
		}() {
			h = h<<uint64(5) - h + khint_t(s[0])
		}
	}
	return h
}

// __ac_Wang_hash - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../deps/hash/khash.h:412
func __ac_Wang_hash(key khint_t) khint_t {
	//! @function
	//  @abstract     Another interface to const char* hash function
	//  @param  key   Pointer to a null terminated string [const char*]
	//  @return       The hash value [khint_t]
	//
	//! @function
	//  @abstract     Const char* comparison function
	//
	key += ^(key << uint64(15))
	key ^= key >> uint64(10)
	key += key << uint64(3)
	key ^= key >> uint64(6)
	key += ^(key << uint64(11))
	key ^= key >> uint64(16)
	return key
}

// kh_init_ptr - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../deps/hash/hash.h:15
func kh_init_ptr() []kh_ptr_t {
	//
	// hash.h
	//
	// Copyright (c) 2012 TJ Holowaychuk <tj@vision-media.ca>
	//
	// pointer hash
	return make([]kh_ptr_t, 1)
}

// kh_destroy_ptr - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../deps/hash/hash.h:15
func kh_destroy_ptr(h []kh_ptr_t) {
	if h != nil {
		_ = h[0].keys
		_ = h[0].flags
		_ = h[0].vals
		_ = h
	}
}

// kh_clear_ptr - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../deps/hash/hash.h:15
func kh_clear_ptr(h []kh_ptr_t) {
	if len(h) == 0 && len(h[0].flags) == 0 {
		noarch.Memset((*[1000000]byte)(unsafe.Pointer(uintptr(int64(uintptr(unsafe.Pointer(&h[0].flags[0]))) / int64(1))))[:], byte(170), func() uint32 {
			if h[0].n_buckets < khint_t((khint32_t((16)))) {
				return 1
			}
			return uint32((khint32_t((h[0].n_buckets >> uint64(4)))))
		}()*4)
		h[0].n_occupied = 0
		h[0].size = h[0].n_occupied
	}
}

// kh_get_ptr - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../deps/hash/hash.h:15
func kh_get_ptr(h []kh_ptr_t, key kh_cstr_t) (c4goDefaultReturn khint_t) {
	if uint32((khint32_t((khint_t(h[0].n_buckets))))) != 0 {
		var k khint_t
		var i khint_t
		var last khint_t
		var mask khint_t
		var step khint_t
		mask = khint_t(h[0].n_buckets) - khint_t((khint32_t((1))))
		k = __ac_X31_hash_string([]byte((key)))
		i = k & mask
		last = i
		for noarch.Not(h[0].flags[i>>uint64(4)]>>uint64(uint32((khint32_t((i & khint_t((khint32_t((uint32(15))))) << uint64(1))))))&khint32_t((2))) && (uint32((h[0].flags[i>>uint64(4)]>>uint64(uint32((khint32_t((i & khint_t((khint32_t((uint32(15))))) << uint64(1))))))&khint32_t((1)))) != 0 || !(noarch.Strcmp([]byte((h[0].keys[i])), []byte((key))) == 0)) {
			i = (i + func() khint_t {
				step++
				return step
			}()) & mask
			if i == last {
				return khint_t(h[0].n_buckets)
			}
		}
		return khint_t((khint32_t((func() uint32 {
			if uint32((h[0].flags[i>>uint64(4)] >> uint64(uint32((khint32_t((i & khint_t((khint32_t((uint32(15))))) << uint64(1)))))) & khint32_t((3)))) != 0 {
				return uint32((khint32_t((khint_t(h[0].n_buckets)))))
			}
			return uint32((khint32_t((i))))
		}()))))
	} else {
		return 0
	}
	return
}

// kh_resize_ptr - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../deps/hash/hash.h:15
func kh_resize_ptr(h []kh_ptr_t, new_n_buckets khint_t) int32 {
	var new_flags []khint32_t
	var j khint_t = 1
	{
		new_n_buckets -= khint_t((khint32_t((uint32(1)))))
		new_n_buckets |= new_n_buckets >> uint64(1)
		new_n_buckets |= new_n_buckets >> uint64(2)
		new_n_buckets |= new_n_buckets >> uint64(4)
		new_n_buckets |= new_n_buckets >> uint64(8)
		new_n_buckets |= new_n_buckets >> uint64(16)
		new_n_buckets += khint_t((khint32_t((uint32(1)))))
		if new_n_buckets < khint_t((khint32_t((4)))) {
			new_n_buckets = 4
		}
		if khint_t(h[0].size) >= khint_t(float64(uint32((khint32_t((new_n_buckets)))))*__ac_HASH_UPPER+0.5) {
			j = 0
		} else {
			new_flags = make([]khint32_t, func() uint32 {
				if new_n_buckets < khint_t((khint32_t((16)))) {
					return 1
				}
				return uint32((khint32_t((new_n_buckets >> uint64(4)))))
			}()*uint32(1))
			if new_flags == nil {
				return -1
			}
			noarch.Memset((*[1000000]byte)(unsafe.Pointer(uintptr(int64(uintptr(unsafe.Pointer(&new_flags[0]))) / int64(1))))[:], byte(170), func() uint32 {
				if new_n_buckets < khint_t((khint32_t((16)))) {
					return 1
				}
				return uint32((khint32_t((new_n_buckets >> uint64(4)))))
			}()*4)
			if khint_t(h[0].n_buckets) < new_n_buckets {
				var new_keys []kh_cstr_t = realloc(h[0].keys, uint32((khint32_t((new_n_buckets))))*8).([]kh_cstr_t)
				if new_keys == nil {
					_ = new_flags
					return -1
				}
				h[0].keys = new_keys
				if 1 != 0 {
					var new_vals []interface{} = realloc(h[0].vals, uint32((khint32_t((new_n_buckets))))*8).([]interface{})
					if new_vals == nil {
						_ = new_flags
						return -1
					}
					h[0].vals = new_vals
				}
			}
		}
	}
	if uint32((khint32_t((j)))) != 0 {
		for j = 0; j != khint_t(h[0].n_buckets); j++ {
			if h[0].flags[j>>uint64(4)]>>uint64(uint32((khint32_t((j & khint_t((khint32_t((uint32(15))))) << uint64(1))))))&khint32_t((3)) == khint32_t((0)) {
				var key kh_cstr_t = h[0].keys[j]
				var val interface{}
				var new_mask khint_t
				new_mask = new_n_buckets - khint_t((khint32_t((1))))
				if 1 != 0 {
					val = h[0].vals[j]
				}
				h[0].flags[j>>uint64(4)] |= khint32_t((uint32(1 << uint64(uint32((khint32_t((j & khint_t((khint32_t((uint32(15))))) << uint64(1)))))))))
				for 1 != 0 {
					var k khint_t
					var i khint_t
					var step khint_t
					k = __ac_X31_hash_string([]byte((key)))
					i = k & new_mask
					for noarch.Not(new_flags[i>>uint64(4)] >> uint64(uint32((khint32_t((i & khint_t((khint32_t((uint32(15))))) << uint64(1)))))) & khint32_t((2))) {
						i = (i + func() khint_t {
							step++
							return step
						}()) & new_mask
					}
					new_flags[i>>uint64(4)] &= khint32_t((^(uint32(2 << uint64(uint32((khint32_t((i & khint_t((khint32_t((uint32(15))))) << uint64(1))))))))))
					if i < khint_t(h[0].n_buckets) && h[0].flags[i>>uint64(4)]>>uint64(uint32((khint32_t((i&khint_t((khint32_t((uint32(15)))))<<uint64(1))))))&khint32_t((3)) == khint32_t((0)) {
						{
							var tmp kh_cstr_t = h[0].keys[i]
							h[0].keys[i] = key
							key = tmp
						}
						if 1 != 0 {
							var tmp interface{} = h[0].vals[i]
							h[0].vals[i] = val
							val = tmp
						}
						h[0].flags[i>>uint64(4)] |= khint32_t((uint32(1 << uint64(uint32((khint32_t((i & khint_t((khint32_t((uint32(15))))) << uint64(1)))))))))
					} else {
						h[0].keys[i] = key
						if 1 != 0 {
							h[0].vals[i] = val
						}
						break
					}
				}
			}
		}
		if khint_t(h[0].n_buckets) > new_n_buckets {
			h[0].keys = realloc(h[0].keys, uint32((khint32_t((new_n_buckets))))*8).([]kh_cstr_t)
			if 1 != 0 {
				h[0].vals = realloc(h[0].vals, uint32((khint32_t((new_n_buckets))))*8).([]interface{})
			}
		}
		_ = h[0].flags
		h[0].flags = new_flags
		h[0].n_buckets = new_n_buckets
		h[0].n_occupied = khint_t(h[0].size)
		h[0].upper_bound = khint_t(float64(uint32((khint32_t((khint_t(h[0].n_buckets))))))*__ac_HASH_UPPER + 0.5)
	}
	return 0
}

// kh_put_ptr - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../deps/hash/hash.h:15
func kh_put_ptr(h []kh_ptr_t, key kh_cstr_t, ret []int32) khint_t {
	var x khint_t
	if khint_t(h[0].n_occupied) >= khint_t(h[0].upper_bound) {
		if khint_t(h[0].n_buckets) > khint_t(h[0].size)<<uint64(1) {
			if kh_resize_ptr(h, khint_t(h[0].n_buckets)-khint_t((khint32_t((1))))) < 0 {
				ret[0] = -1
				return khint_t(h[0].n_buckets)
			}
		} else if kh_resize_ptr(h, khint_t(h[0].n_buckets)+khint_t((khint32_t((1))))) < 0 {
			ret[0] = -1
			return khint_t(h[0].n_buckets)
		}
	}
	{
		var k khint_t
		var i khint_t
		var site khint_t
		var last khint_t
		var mask khint_t = khint_t(h[0].n_buckets) - khint_t((khint32_t((1))))
		var step khint_t
		site = khint_t(h[0].n_buckets)
		x = site
		k = __ac_X31_hash_string([]byte((key)))
		i = k & mask
		if uint32((h[0].flags[i>>uint64(4)] >> uint64(uint32((khint32_t((i & khint_t((khint32_t((uint32(15))))) << uint64(1)))))) & khint32_t((2)))) != 0 {
			x = i
		} else {
			last = i
			for noarch.Not(h[0].flags[i>>uint64(4)]>>uint64(uint32((khint32_t((i & khint_t((khint32_t((uint32(15))))) << uint64(1))))))&khint32_t((2))) && (uint32((h[0].flags[i>>uint64(4)]>>uint64(uint32((khint32_t((i & khint_t((khint32_t((uint32(15))))) << uint64(1))))))&khint32_t((1)))) != 0 || !(noarch.Strcmp([]byte((h[0].keys[i])), []byte((key))) == 0)) {
				if uint32((h[0].flags[i>>uint64(4)] >> uint64(uint32((khint32_t((i & khint_t((khint32_t((uint32(15))))) << uint64(1)))))) & khint32_t((1)))) != 0 {
					site = i
				}
				i = (i + func() khint_t {
					step++
					return step
				}()) & mask
				if i == last {
					x = site
					break
				}
			}
			if x == khint_t(h[0].n_buckets) {
				if uint32((h[0].flags[i>>uint64(4)]>>uint64(uint32((khint32_t((i & khint_t((khint32_t((uint32(15))))) << uint64(1))))))&khint32_t((2)))) != 0 && site != khint_t(h[0].n_buckets) {
					x = site
				} else {
					x = i
				}
			}
		}
	}
	if uint32((h[0].flags[x>>uint64(4)] >> uint64(uint32((khint32_t((x & khint_t((khint32_t((uint32(15))))) << uint64(1)))))) & khint32_t((2)))) != 0 {
		h[0].keys[x] = key
		h[0].flags[x>>uint64(4)] &= khint32_t((^(uint32(3 << uint64(uint32((khint32_t((x & khint_t((khint32_t((uint32(15))))) << uint64(1))))))))))
		h[0].size += khint_t((khint32_t((uint32(1)))))
		h[0].n_occupied += khint_t((khint32_t((uint32(1)))))
		ret[0] = 1
	} else if uint32((h[0].flags[x>>uint64(4)] >> uint64(uint32((khint32_t((x & khint_t((khint32_t((uint32(15))))) << uint64(1)))))) & khint32_t((1)))) != 0 {
		h[0].keys[x] = key
		h[0].flags[x>>uint64(4)] &= khint32_t((^(uint32(3 << uint64(uint32((khint32_t((x & khint_t((khint32_t((uint32(15))))) << uint64(1))))))))))
		h[0].size += khint_t((khint32_t((uint32(1)))))
		ret[0] = 2
	} else {
		ret[0] = 0
	}
	return x
}

// kh_del_ptr - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../deps/hash/hash.h:15
func kh_del_ptr(h []kh_ptr_t, x khint_t) {
	if x != khint_t(h[0].n_buckets) && noarch.Not(h[0].flags[x>>uint64(4)]>>uint64(uint32((khint32_t((x & khint_t((khint32_t((uint32(15))))) << uint64(1))))))&khint32_t((3))) {
		h[0].flags[x>>uint64(4)] |= khint32_t((uint32(1 << uint64(uint32((khint32_t((x & khint_t((khint32_t((uint32(15))))) << uint64(1)))))))))
		h[0].size -= khint_t((khint32_t((uint32(1)))))
	}
}

// errorf - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:27
func errorf(p []parser_t, item lex_item_t, decl []decl_t, fmt_ []byte, c4goArgs ...interface{}) lex_item_t {
	var args = create_va_list(c4goArgs)
	va_start(args, fmt_)
	parser_verrorf(p, item, []byte("Invalid export syntax: \x00"), fmt_, args)
	decl[0].error_ = 1
	return lex_item_empty
}

// append_ - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:36
func append_(decl []decl_t, value lex_item_t) {
	decl[0].items = realloc(decl[0].items, 40*(uint32(decl[0].length)+1)).([]lex_item_t)
	decl[0].items[uint32(decl[0].length)] = value
	decl[0].length += uint32(1)
}

// rewind_until - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:42
func rewind_until(p []parser_t, decl []decl_t, item lex_item_t) {
	var last noarch.SsizeT = noarch.SsizeT(uint32(decl[0].length) - 1)
	for last >= noarch.SsizeT(0) && noarch.Not(lex_item_equals(item, decl[0].items[last])) {
		parser_backup(p, decl[0].items[last])
		last--
	}
	parser_backup(p, decl[0].items[last])
	decl[0].length = uint32(last)
}

// rewind_whitespace - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:54
func rewind_whitespace(p []parser_t, decl []decl_t, item lex_item_t) {
	parser_backup(p, item)
	var last noarch.SsizeT = noarch.SsizeT(uint32(decl[0].length) - 1)
	for last >= noarch.SsizeT(0) && uint32(int32((decl[0].items[last].type_))) == uint32(int32((item_whitespace))) {
		parser_backup(p, decl[0].items[last])
		last--
	}
	decl[0].length = uint32(last + noarch.SsizeT(1))
}

// symbol_rename - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:66
func symbol_rename(p []parser_t, decl []decl_t, name lex_item_t, alias lex_item_t) lex_item_t {
	var i int32
	for i = 0; uint32(i) < uint32(decl[0].length); i++ {
		var item lex_item_t = decl[0].items[i]
		if int32((lex_item_equals(name, item))) != 0 {
			var symbol_name []byte
			var export_name []byte = func() []byte {
				if uint32(int32((alias.type_))) == 0 {
					return name.value
				}
				return alias.value
			}()
			asprintf((*[1000000][]byte)(unsafe.Pointer(&symbol_name))[:], []byte("%s_%s\x00"), p[0].pkg[0].name, export_name)
			item = lex_item_replace_value(item, symbol_name)
			decl[0].items[i] = item
			return item
		}
	}
	return lex_item_empty
}

// free_decl - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:84
func free_decl(decl []decl_t) {
	var i int32
	for i = 0; uint32(i) < uint32(decl[0].length); i++ {
		lex_item_free(decl[0].items[i])
	}
	_ = decl[0].items
}

// emit - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:92
func emit(p []parser_t, decl []decl_t, is_function int32, is_extern int32) []byte {
	var length uint32
	var start int32
	var end int32 = int32(uint32(decl[0].length))
	var i int32
	for start < end && uint32(int32((decl[0].items[start].type_))) == uint32(int32((item_whitespace))) {
		// trim leading/trailing whitespace
		start++
	}
	for end > start && uint32(int32((decl[0].items[end-1].type_))) == uint32(int32((item_whitespace))) {
		end--
	}
	for i = 0; uint32(i) < uint32(decl[0].length); i++ {
		var item lex_item_t = decl[0].items[i]
		if noarch.Not(is_extern) {
			package_emit(p[0].pkg, item.value)
		}
		if i >= start && i < end {
			length += uint32(item.length)
		}
	}
	if int32((is_function)) != 0 {
		length++
	}
	var output []byte = make([]byte, length+1)
	length = 0
	for i = start; i < end; i++ {
		var item lex_item_t = decl[0].items[i]
		memcpy(output[0+length:], item.value, uint32(item.length))
		length += uint32(item.length)
	}
	if int32((is_function)) != 0 && noarch.Not(is_extern) {
		output[func() uint32 {
			defer func() {
				length++
			}()
			return length
		}()] = ';'
	}
	output[length] = byte(0)
	return output
}

// collect_newlines - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:124
func collect_newlines(p []parser_t, decl []decl_t) lex_item_t {
	var line uint32 = uint32(p[0].lexer[0].line)
	var item lex_item_t = parser_next(p)
	for uint32(int32((item.type_))) == uint32(int32((item_whitespace))) || uint32(int32((item.type_))) == uint32(int32((item_comment))) {
		if uint32(p[0].lexer[0].line) != line {
			append_(decl, item)
		} else {
			lex_item_free(item)
		}
		line = uint32(p[0].lexer[0].line)
		item = parser_next(p)
	}
	return item
}

// collect - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:142
func collect(p []parser_t, decl []decl_t) lex_item_t {
	var item lex_item_t = parser_next(p)
	for uint32(int32((item.type_))) == uint32(int32((item_whitespace))) || uint32(int32((item.type_))) == uint32(int32((item_comment))) {
		append_(decl, item)
		item = parser_next(p)
	}
	return item
}

// init_export_types - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:168
func init_export_types() {
	export_types = kh_init_ptr()
	hash_set(export_types, []byte("typedef\x00"), parse_typedef)
	hash_set(export_types, []byte("struct\x00"), parse_struct)
	hash_set(export_types, []byte("enum\x00"), parse_enum)
	hash_set(export_types, []byte("union\x00"), parse_union)
}

// parse_semicolon - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:179
func parse_semicolon(p []parser_t, decl []decl_t) {
	if int32((decl[0].error_)) != 0 {
		return
	}
	var item lex_item_t = collect_newlines(p, decl)
	if uint32(int32((item.type_))) != uint32(int32((item_symbol))) || int32(item.value[0]) != int32(';') {
		errorf(p, item, decl, []byte("expecting ';' but got %s\x00"), lex_item_to_string(item))
		return
	}
	append_(decl, item)
}

// export_parse - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:191
func export_parse(p []parser_t) int32 {
	if len(export_types) == 0 {
		init_export_types()
	}
	var decl decl_t = decl_t{nil, 0, 0}
	var fn export_fn
	var name lex_item_t = lex_item_t{lex_item_type((0)), nil, 0, 0, 0, 0, 0}
	var alias lex_item_t = lex_item_t{lex_item_type((0)), nil, 0, 0, 0, 0, 0}
	var has_semicolon int32
	var is_extern int32
	var t int32
	var type_ lex_item_t = collect_newlines(p, (*[1000000]decl_t)(unsafe.Pointer(&decl))[:])
	switch uint32(int32((type_.type_))) {
	case uint32(int32((item_id))):
		if noarch.Strcmp([]byte("extern\x00"), type_.value) == 0 {
			append_((*[1000000]decl_t)(unsafe.Pointer(&decl))[:], type_)
			is_extern = 1
			type_ = collect(p, (*[1000000]decl_t)(unsafe.Pointer(&decl))[:])
		}
		fn = hash_get(export_types, type_.value).(export_fn)
		has_semicolon = int32((noarch.BoolToInt(int32((is_extern)) != 0 || len(fn) != 0)))
		t = 1
	case uint32(int32((item_symbol))):
		if int32(type_.value[0]) == int32('*') {
			lex_item_free(type_)
			return parse_passthrough(p)
		}
	case uint32(int32((item_open_symbol))):
		if int32(type_.value[0]) == int32('{') {
			fn = parse_export_block
			lex_item_free(type_)
			type_.value = []byte("{\x00")
			t = 0
		}
	default:
		fn = nil
		break
	}
	if len(fn) == 0 {
		parser_backup(p, type_)
		type_.type_ = item_symbol
		type_.value = []byte("variable\x00")
		type_.length = strings_len([]byte("variable\x00"))
		name = parse_variable(p, (*[1000000]decl_t)(unsafe.Pointer(&decl))[:])
		t = 2
	} else {
		if t == 1 {
			append_((*[1000000]decl_t)(unsafe.Pointer(&decl))[:], type_)
		}
		name = fn(p, (*[1000000]decl_t)(unsafe.Pointer(&decl))[:])
	}
	alias = parse_as(p, (*[1000000]decl_t)(unsafe.Pointer(&decl))[:])
	if int32((has_semicolon)) != 0 {
		parse_semicolon(p, (*[1000000]decl_t)(unsafe.Pointer(&decl))[:])
	}
	if int32((decl.error_)) != 0 {
		free_decl((*[1000000]decl_t)(unsafe.Pointer(&decl))[:])
		return -1
	}
	if uint32(int32((name.type_))) == 0 {
		errorf(p, type_, (*[1000000]decl_t)(unsafe.Pointer(&decl))[:], []byte("can't export anonymous symbols\x00"))
		free_decl((*[1000000]decl_t)(unsafe.Pointer(&decl))[:])
		return -1
	}
	var original_name lex_item_t = lex_item_dup(name)
	var symbol lex_item_t = symbol_rename(p, (*[1000000]decl_t)(unsafe.Pointer(&decl))[:], name, alias)
	if uint32(int32((symbol.type_))) == 0 && (int64(uintptr(unsafe.Pointer(&fn[0])))/int64(1)-int64(uintptr(unsafe.Pointer(&parse_export_block[0])))/int64(1)) != 0 {
		errorf(p, type_, (*[1000000]decl_t)(unsafe.Pointer(&decl))[:], []byte("unable to export symbol '%s'\x00"), name.value)
		free_decl((*[1000000]decl_t)(unsafe.Pointer(&decl))[:])
		return -1
	}
	package_export_add(strings_dup(original_name.value), strings_dup(alias.value), strings_dup(symbol.value), type_.value, emit(p, (*[1000000]decl_t)(unsafe.Pointer(&decl))[:], int32((noarch.BoolToInt(t == 2))), is_extern), p[0].pkg)
	lex_item_free(original_name)
	lex_item_free(alias)
	free_decl((*[1000000]decl_t)(unsafe.Pointer(&decl))[:])
	return 1
}

// parse_passthrough - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:281
func parse_passthrough(p []parser_t) int32 {
	var decl decl_t = decl_t{nil, 0, 0}
	var from lex_item_t = collect(p, (*[1000000]decl_t)(unsafe.Pointer(&decl))[:])
	if uint32(int32((from.type_))) != uint32(int32((item_id))) || noarch.Strcmp(from.value, []byte("from\x00")) != 0 {
		parser_errorf(p, from, []byte("Exporting passthrough: \x00"), []byte("expected 'from', but got %s\x00"), lex_item_to_string(from))
		free_decl((*[1000000]decl_t)(unsafe.Pointer(&decl))[:])
		return -1
	}
	lex_item_free(from)
	var filename lex_item_t = collect(p, (*[1000000]decl_t)(unsafe.Pointer(&decl))[:])
	if uint32(int32((filename.type_))) != uint32(int32((item_quoted_string))) {
		parser_errorf(p, filename, []byte("Exporting passthrough: \x00"), []byte("expected filename, but got %s\x00"), lex_item_to_string(filename))
		free_decl((*[1000000]decl_t)(unsafe.Pointer(&decl))[:])
		return -1
	}
	parse_semicolon(p, (*[1000000]decl_t)(unsafe.Pointer(&decl))[:])
	if int32((decl.error_)) != 0 {
		free_decl((*[1000000]decl_t)(unsafe.Pointer(&decl))[:])
		return -1
	}
	var error_ []byte
	var imp []package_import_t = package_import_passthrough(p[0].pkg, strings_dup(string_parse(filename.value)), (*[1000000][]byte)(unsafe.Pointer(&error_))[:])
	if len(imp) == 0 {
		parser_errorf(p, filename, []byte("Exporting passthrough: \x00"), []byte("%s\x00"), error_)
		free_decl((*[1000000]decl_t)(unsafe.Pointer(&decl))[:])
		return -1
	}
	lex_item_free(filename)
	var header []byte
	var rel []byte = utils_relative(p[0].pkg[0].source_abs, imp[0].pkg[0].header)
	asprintf((*[1000000][]byte)(unsafe.Pointer(&header))[:], []byte("#include \"%s\"\x00"), rel)
	package_emit(p[0].pkg, header)
	free_decl((*[1000000]decl_t)(unsafe.Pointer(&decl))[:])
	_ = rel
	_ = header
	return 1
}

// parse_as - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:324
func parse_as(p []parser_t, decl []decl_t) lex_item_t {
	if int32((decl[0].error_)) != 0 {
		return lex_item_empty
	}
	var start uint32 = uint32(decl[0].length)
	var as lex_item_t = collect(p, decl)
	if uint32(int32((as.type_))) != uint32(int32((item_id))) || noarch.Strcmp([]byte("as\x00"), as.value) != 0 {
		parser_backup(p, as)
		for uint32(decl[0].length) > start {
			decl[0].length -= uint32(1)
			parser_backup(p, decl[0].items[uint32(decl[0].length)])
		}
		return lex_item_empty
	}
	lex_item_free(as)
	for uint32(decl[0].length) > start {
		decl[0].length -= uint32(1)
		parser_backup(p, decl[0].items[uint32(decl[0].length)])
	}
	var alias lex_item_t = collect_newlines(p, decl)
	if uint32(int32((alias.type_))) != uint32(int32((item_id))) {
		errorf(p, alias, decl, []byte("expecting identifier but got %s\x00"), lex_item_to_string(alias))
		return lex_item_empty
	}
	return alias
}

// parse_function_ptr - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:355
func parse_function_ptr(p []parser_t, decl []decl_t) lex_item_t {
	var star lex_item_t = collect(p, decl)
	if uint32(int32((star.type_))) != uint32(int32((item_symbol))) || int32(star.value[0]) != int32('*') {
		return errorf(p, star, decl, []byte("function pointer: expecting '*' but found '%s'\x00"), star.value)
	}
	append_(decl, star)
	var name lex_item_t = collect(p, decl)
	if uint32(int32((name.type_))) != uint32(int32((item_id))) {
		return errorf(p, name, decl, []byte("function pointer: expecting identifier but found '%s'\x00"), name.value)
	}
	append_(decl, name)
	var item lex_item_t
	item = collect(p, decl)
	if uint32(int32((item.type_))) != uint32(int32((item_close_symbol))) || int32(item.value[0]) != int32(')') {
		return errorf(p, item, decl, []byte("function pointer: expecting ')' but found '%s'\x00"), item.value)
	}
	append_(decl, item)
	item = collect(p, decl)
	if uint32(int32((item.type_))) != uint32(int32((item_open_symbol))) || int32(item.value[0]) != int32('(') {
		return errorf(p, item, decl, []byte("function pointer: expecting '(' but found '%s'\x00"), item.value)
	}
	append_(decl, item)
	parse_function_args(p, decl)
	if int32((decl[0].error_)) != 0 {
		return lex_item_empty
	}
	return name
}

// parse_declaration - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:387
func parse_declaration(p []parser_t, decl []decl_t) lex_item_t {
	var type_ lex_item_t = collect(p, decl)
	if uint32(int32((type_.type_))) != uint32(int32((item_id))) {
		if uint32(int32((type_.type_))) == uint32(int32((item_close_symbol))) && int32(type_.value[0]) == int32('}') {
			return type_
		}
		return errorf(p, type_, decl, []byte("in declaration: expecting identifier but got %s\x00"), lex_item_to_string(type_))
	}
	if noarch.Strcmp([]byte("as\x00"), type_.value) == 0 {
		return type_
	}
	type_ = parser_identifier_parse(p, type_, 1)
	var fn export_fn = hash_get(export_types, type_.value).(export_fn)
	if (int64(uintptr(unsafe.Pointer(&fn[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&parse_typedef[0])))/int64(1)) == 0 {
		return errorf(p, type_, decl, []byte("in declaration: unexpected identifier 'typedef'\x00"))
	}
	if len(fn) == 0 {
		fn = parse_variable
	}
	append_(decl, type_)
	var item lex_item_t = fn(p, decl)
	if int32((decl[0].error_)) != 0 {
		return lex_item_empty
	}
	if (int64(uintptr(unsafe.Pointer(&fn[0])))/int64(1)-int64(uintptr(unsafe.Pointer(&parse_enum[0])))/int64(1)) == 0 || (int64(uintptr(unsafe.Pointer(&fn[0])))/int64(1)-int64(uintptr(unsafe.Pointer(&parse_union[0])))/int64(1)) == 0 || (int64(uintptr(unsafe.Pointer(&fn[0])))/int64(1)-int64(uintptr(unsafe.Pointer(&parse_struct[0])))/int64(1)) == 0 {
		item = collect(p, decl)
		if uint32(int32((item.type_))) != uint32(int32((item_id))) {
			parser_backup(p, item)
		} else {
			append_(decl, item)
		}
	}
	parse_semicolon(p, decl)
	if int32((decl[0].error_)) != 0 {
		return lex_item_empty
	}
	return item
}

// parse_typedef - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:425
func parse_typedef(p []parser_t, decl []decl_t) lex_item_t {
	var type_ lex_item_t = collect(p, decl)
	if uint32(int32((type_.type_))) != uint32(int32((item_id))) {
		return errorf(p, type_, decl, []byte("in typedef: expected identifier\x00"))
	}
	var fn export_fn = hash_get(export_types, type_.value).(export_fn)
	//if (fn != parse_struct && fn != parse_enum && fn != parse_union) append(decl, type);
	append_(decl, type_)
	if len(fn) != 0 {
		fn(p, decl)
		if int32((decl[0].error_)) != 0 {
			return lex_item_empty
		}
	}
	var item lex_item_t
	var name lex_item_t = lex_item_t{lex_item_type((0)), nil, 0, 0, 0, 0, 0}
	var function_ptr int32
	var as int32
	for {
		item = collect(p, decl)
		switch uint32(int32((item.type_))) {
		case uint32(int32((item_eof))):
			return errorf(p, item, decl, []byte("in typedef: expected identifier or '('\x00"))
		case uint32(int32((item_id))):
			if noarch.Strcmp([]byte("as\x00"), item.value) == 0 {
				rewind_whitespace(p, decl, item)
				as = 1
				continue
			}
			if noarch.Not(function_ptr) {
				name = item
			}
		case uint32(int32((item_symbol))):
			if int32(item.value[0]) == int32(';') {
				continue
			}
			fallthrough
		case uint32(int32((item_open_symbol))):
			if noarch.Not(function_ptr) && int32(item.value[0]) == int32('(') {
				function_ptr = 1
				append_(decl, item)
				name = parse_function_ptr(p, decl)
				if int32((decl[0].error_)) != 0 {
					return name
				}
				continue
			}
			fallthrough
		default:
			break
		}
		append_(decl, item)
		if !(uint32(int32((item.type_))) != uint32(int32((item_error))) && !(uint32(int32((item.type_))) == uint32(int32((item_symbol))) && int32(item.value[0]) == int32(';')) && noarch.Not(as)) {
			break
		}
	}
	if uint32(int32((item.type_))) == uint32(int32((item_error))) {
		decl[0].error_ = 1
		return item
	}
	if noarch.Not(as) {
		parser_backup(p, item)
	}
	return name
}

// parse_struct - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:489
func parse_struct(p []parser_t, decl []decl_t) lex_item_t {
	var name lex_item_t = lex_item_t{lex_item_type((0)), nil, 0, 0, 0, 0, 0}
	var item lex_item_t = collect(p, decl)
	if uint32(int32((item.type_))) == uint32(int32((item_id))) {
		name = parser_identifier_parse_typed(p, struct_i, item, 1)
		append_(decl, name)
		item = collect(p, decl)
	}
	if uint32(int32((item.type_))) != uint32(int32((item_open_symbol))) || int32(item.value[0]) != int32('{') {
		parser_backup(p, item)
		return name
	}
	append_(decl, item)
	parse_struct_block(p, decl, item)
	return func() lex_item_t {
		if int32((decl[0].error_)) != 0 {
			return lex_item_empty
		}
		return name
	}()
}

// parse_union - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:509
func parse_union(p []parser_t, decl []decl_t) lex_item_t {
	var name lex_item_t = lex_item_t{lex_item_type((0)), nil, 0, 0, 0, 0, 0}
	var item lex_item_t = collect(p, decl)
	if uint32(int32((item.type_))) == uint32(int32((item_id))) {
		name = parser_identifier_parse_typed(p, union_i, item, 1)
		append_(decl, name)
		item = collect(p, decl)
	}
	if uint32(int32((item.type_))) != uint32(int32((item_open_symbol))) || int32(item.value[0]) != int32('{') {
		return errorf(p, item, decl, []byte("in union: expecting '{' but got %s\x00"), lex_item_to_string(item))
	}
	append_(decl, item)
	parse_struct_block(p, decl, item)
	return func() lex_item_t {
		if int32((decl[0].error_)) != 0 {
			return lex_item_empty
		}
		return name
	}()
}

// parse_enum - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:527
func parse_enum(p []parser_t, decl []decl_t) lex_item_t {
	var name lex_item_t = lex_item_t{lex_item_type((0)), nil, 0, 0, 0, 0, 0}
	var item lex_item_t = collect(p, decl)
	if uint32(int32((item.type_))) == uint32(int32((item_id))) {
		name = parser_identifier_parse_typed(p, enum_i, item, 1)
		append_(decl, name)
		item = collect(p, decl)
	}
	if uint32(int32((item.type_))) != uint32(int32((item_open_symbol))) || int32(item.value[0]) != int32('{') {
		return errorf(p, item, decl, []byte("in enum: expecting '{' but got %s\x00"), lex_item_to_string(item))
	}
	append_(decl, item)
	parse_enum_block(p, decl, item)
	return func() lex_item_t {
		if int32((decl[0].error_)) != 0 {
			return lex_item_empty
		}
		return name
	}()
}

// parse_export_block - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:545
func parse_export_block(p []parser_t, decl []decl_t) (c4goDefaultReturn lex_item_t) {
	var item lex_item_t
	var escaped_id int32
	for 1 != 0 {
		var record int32
		item = collect(p, decl)
		switch uint32(int32((item.type_))) {
		case uint32(int32((item_c_code))):
			fallthrough
		case uint32(int32((item_quoted_string))):
			fallthrough
		case uint32(int32((item_comment))):
			fallthrough
		case uint32(int32((item_symbol))):
			fallthrough
		case uint32(int32((item_preprocessor))):
			record = 1
		case uint32(int32((item_arrow))):
			escaped_id = 1
			record = 1
		case uint32(int32((item_id))):
			if escaped_id == 0 {
				item = parser_identifier_parse(p, item, 1)
			}
			record = 1
			escaped_id = 0
		case uint32(int32((item_close_symbol))):
			if int32(item.value[0]) == int32('}') {
				lex_item_free(item)
				item.value = []byte("block\x00")
				item.length = strings_len([]byte("block\x00"))
				return item
			}
		case uint32(int32((item_eof))):
			return errorf(p, item, decl, []byte("in block: unmatched '{'\x00"))
		case uint32(int32((item_error))):
			parser_backup(p, item)
			decl[0].error_ = 1
			return lex_item_empty
		default:
			return errorf(p, item, decl, []byte("in block: unexpected input '%s'\x00"), lex_item_to_string(item))
		}
		if record != 0 {
			append_(decl, item)
		}
	}
	return
}

// parse_struct_block - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:596
func parse_struct_block(p []parser_t, decl []decl_t, start lex_item_t) lex_item_t {
	var item lex_item_t
	for {
		item = parse_declaration(p, decl)
		if uint32(int32((item.type_))) == uint32(int32((item_eof))) {
			return errorf(p, start, decl, []byte("in struct block: missing matching '}' before end of file\x00"))
		}
		if uint32(int32((item.type_))) == uint32(int32((item_error))) {
			return item
		}
		if !!(uint32(int32((item.type_))) == uint32(int32((item_close_symbol))) && int32(item.value[0]) == int32('}')) {
			break
		}
	}
	append_(decl, item)
	return item
}

// parse_enum_declaration - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:609
func parse_enum_declaration(p []parser_t, decl []decl_t) lex_item_t {
	var item lex_item_t = collect(p, decl)
	if uint32(int32((item.type_))) != uint32(int32((item_id))) {
		if uint32(int32((item.type_))) == uint32(int32((item_close_symbol))) && int32(item.value[0]) == int32('}') {
			return item
		}
		return errorf(p, item, decl, []byte("in enum: expecting identifier but got %s\x00"), lex_item_to_string(item))
	}
	if noarch.Strcmp([]byte("as\x00"), item.value) == 0 {
		return item
	}
	append_(decl, item)
	item = collect(p, decl)
	if !(uint32(int32((item.type_))) == uint32(int32((item_symbol))) && (int32(item.value[0]) != int32(',') || int32(item.value[0]) != int32('='))) {
		return item
	}
	append_(decl, item)
	if int32(item.value[0]) == int32('=') {
		item = collect(p, decl)
		if uint32(int32((item.type_))) != uint32(int32((item_number))) {
			return errorf(p, item, decl, []byte("in enum: expecting a number but got %s\x00"), lex_item_to_string(item))
		}
		append_(decl, item)
		item = collect(p, decl)
		if uint32(int32((item.type_))) != uint32(int32((item_symbol))) || int32(item.value[0]) != int32(',') {
			return item
		}
		append_(decl, item)
	}
	return item
}

// parse_enum_block - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:643
func parse_enum_block(p []parser_t, decl []decl_t, start lex_item_t) lex_item_t {
	var item lex_item_t
	for {
		item = parse_enum_declaration(p, decl)
		if uint32(int32((item.type_))) == uint32(int32((item_eof))) {
			return errorf(p, start, decl, []byte("in struct block: missing matching '}' before end of file\x00"))
		}
		if uint32(int32((item.type_))) == uint32(int32((item_error))) {
			return item
		}
		if !!(uint32(int32((item.type_))) == uint32(int32((item_close_symbol))) && int32(item.value[0]) == int32('}')) {
			break
		}
	}
	append_(decl, item)
	return item
}

// parse_type - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:656
func parse_type(p []parser_t, decl []decl_t) lex_item_t {
	var item lex_item_t
	var name lex_item_t = lex_item_t{lex_item_type((0)), nil, 0, 0, 0, 0, 0}
	var escaped_id int32
	for {
		item = collect(p, decl)
		switch uint32(int32((item.type_))) {
		case uint32(int32((item_arrow))):
			escaped_id = 1
			append_(decl, item)
			continue
			fallthrough
		case uint32(int32((item_id))):
			if escaped_id == 0 {
				item = parser_identifier_parse(p, item, 1)
			}
			name = item
		case uint32(int32((item_symbol))):
			if int32(item.value[0]) == int32('*') {
				break
			}
			fallthrough
		case uint32(int32((item_open_symbol))):
			if int32(item.value[0]) == int32('(') {
				append_(decl, item)
				return name
			}
			fallthrough
		default:
			return errorf(p, item, decl, []byte("in type declaration: expecting identifier or '('\x00"))
		}
		append_(decl, item)
		if !(uint32(int32((item.type_))) != uint32(int32((item_open_symbol))) || int32(item.value[0]) != int32('(')) {
			break
		}
	}
	return lex_item_empty
}

// parse_variable - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:686
func parse_variable(p []parser_t, decl []decl_t) lex_item_t {
	var item lex_item_t
	var name lex_item_t = lex_item_t{lex_item_type((0)), nil, 0, 0, 0, 0, 0}
	var escaped_id int32
	var escape_till_semicolon int32
	for {
		item = collect(p, decl)
		if int32((escape_till_semicolon)) != 0 && (uint32(int32((item.type_))) != uint32(int32((item_symbol))) || int32(item.value[0]) != int32(';')) {
			append_(decl, item)
			continue
		}
		switch uint32(int32((item.type_))) {
		case uint32(int32((item_arrow))):
			escaped_id = 1
			append_(decl, item)
			continue
			fallthrough
		case uint32(int32((item_id))):
			if noarch.Strcmp(item.value, []byte("as\x00")) == 0 {
				rewind_whitespace(p, decl, item)
				return name
			}
			if int32((escaped_id)) == 0 {
				item = parser_identifier_parse(p, item, 1)
			}
			name = item
		case uint32(int32((item_symbol))):
			if int32(item.value[0]) == int32('*') {
				break
			}
			if int32(item.value[0]) == int32('=') {
				escape_till_semicolon = 1
				break
			}
			if int32(item.value[0]) == int32(';') {
				parser_backup(p, item)
				return name
			}
			fallthrough
		case uint32(int32((item_open_symbol))):
			if int32(item.value[0]) == int32('(') {
				var star lex_item_t = parser_next(p)
				parser_backup(p, star)
				if uint32(int32((star.type_))) == uint32(int32((item_symbol))) && int32(star.value[0]) == int32('*') {
					append_(decl, item)
					return parse_function_ptr(p, decl)
				}
				parser_backup(p, item)
				rewind_until(p, decl, name)
				return parse_function(p, decl)
			}
			if int32(item.value[0]) == int32('[') {
				append_(decl, item)
				item = parser_next(p)
				if uint32(int32((item.type_))) == uint32(int32((item_number))) {
					append_(decl, item)
					item = parser_next(p)
				}
				if uint32(int32((item.type_))) != uint32(int32((item_close_symbol))) || int32(item.value[0]) != int32(']') {
					return errorf(p, item, decl, []byte("in type declaration: expecting terminating ']' but got %s\x00"), lex_item_to_string(item))
				}
				break
			}
			fallthrough
		default:
			return errorf(p, item, decl, []byte("in type declaration: expecting identifier or '(' but got %s\x00"), lex_item_to_string(item))
		}
		append_(decl, item)
		if !(uint32(int32((item.type_))) != uint32(int32((item_eof)))) {
			break
		}
	}
	return name
}

// parse_function_args - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:755
func parse_function_args(p []parser_t, decl []decl_t) lex_item_t {
	var level int32 = 1
	var escaped_id int32
	var item lex_item_t
	for {
		item = collect(p, decl)
		switch uint32(int32((item.type_))) {
		case uint32(int32((item_arrow))):
			escaped_id = 1
			append_(decl, item)
			continue
			fallthrough
		case uint32(int32((item_id))):
			if escaped_id == 0 {
				item = parser_identifier_parse(p, item, 1)
			}
		case uint32(int32((item_open_symbol))):
			if int32(item.value[0]) == int32('(') {
				level++
			}
		case uint32(int32((item_close_symbol))):
			if int32(item.value[0]) == int32(')') {
				level--
			}
		default:
			break
		}
		escaped_id = 0
		append_(decl, item)
		if !(level > 0 && uint32(int32((item.type_))) != uint32(int32((item_eof))) && uint32(int32((item.type_))) != uint32(int32((item_error)))) {
			break
		}
	}
	return item
}

// parse_function - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/export.c:787
func parse_function(p []parser_t, decl []decl_t) lex_item_t {
	var name lex_item_t = parse_type(p, decl)
	if int32((decl[0].error_)) != 0 {
		emit(p, decl, 1, 0)
		decl[0].items = nil
		decl[0].length = 0
		return lex_item_empty
	}
	parse_function_args(p, decl)
	if int32((decl[0].error_)) != 0 {
		return lex_item_empty
	}
	return name
}

type _Bool int32

// --- END OF HASH FUNCTIONS ---
// Other convenient macros...
//!
//  @abstract Type of the hash table.
//  @param  name  Name of the hash table [symbol]
//
//! @function
//  @abstract     Initiate a hash table.
//  @param  name  Name of the hash table [symbol]
//  @return       Pointer to the hash table [khash_t(name)*]
//
//! @function
//  @abstract     Destroy a hash table.
//  @param  name  Name of the hash table [symbol]
//  @param  h     Pointer to the hash table [khash_t(name)*]
//
//! @function
//  @abstract     Reset a hash table without deallocating memory.
//  @param  name  Name of the hash table [symbol]
//  @param  h     Pointer to the hash table [khash_t(name)*]
//
//! @function
//  @abstract     Resize a hash table.
//  @param  name  Name of the hash table [symbol]
//  @param  h     Pointer to the hash table [khash_t(name)*]
//  @param  s     New size [khint_t]
//
//! @function
//  @abstract     Insert a key to the hash table.
//  @param  name  Name of the hash table [symbol]
//  @param  h     Pointer to the hash table [khash_t(name)*]
//  @param  k     Key [type of keys]
//  @param  r     Extra return code: -1 if the operation failed;
//                0 if the key is present in the hash table;
//                1 if the bucket is empty (never used); 2 if the element in
//				the bucket has been deleted [int*]
//  @return       Iterator to the inserted element [khint_t]
//
//! @function
//  @abstract     Retrieve a key from the hash table.
//  @param  name  Name of the hash table [symbol]
//  @param  h     Pointer to the hash table [khash_t(name)*]
//  @param  k     Key [type of keys]
//  @return       Iterator to the found element, or kh_end(h) if the element is absent [khint_t]
//
//! @function
//  @abstract     Remove a key from the hash table.
//  @param  name  Name of the hash table [symbol]
//  @param  h     Pointer to the hash table [khash_t(name)*]
//  @param  k     Iterator to the element to be deleted [khint_t]
//
//! @function
//  @abstract     Test whether a bucket contains data.
//  @param  h     Pointer to the hash table [khash_t(name)*]
//  @param  x     Iterator to the bucket [khint_t]
//  @return       1 if containing data; 0 otherwise [int]
//
//! @function
//  @abstract     Get key given an iterator
//  @param  h     Pointer to the hash table [khash_t(name)*]
//  @param  x     Iterator to the bucket [khint_t]
//  @return       Key [type of keys]
//
//! @function
//  @abstract     Get value given an iterator
//  @param  h     Pointer to the hash table [khash_t(name)*]
//  @param  x     Iterator to the bucket [khint_t]
//  @return       Value [type of values]
//  @discussion   For hash sets, calling this results in segfault.
//
//! @function
//  @abstract     Alias of kh_val()
//
//! @function
//  @abstract     Get the start iterator
//  @param  h     Pointer to the hash table [khash_t(name)*]
//  @return       The start iterator [khint_t]
//
//! @function
//  @abstract     Get the end iterator
//  @param  h     Pointer to the hash table [khash_t(name)*]
//  @return       The end iterator [khint_t]
//
//! @function
//  @abstract     Get the number of elements in the hash table
//  @param  h     Pointer to the hash table [khash_t(name)*]
//  @return       Number of elements in the hash table [khint_t]
//
//! @function
//  @abstract     Get the number of buckets in the hash table
//  @param  h     Pointer to the hash table [khash_t(name)*]
//  @return       Number of buckets in the hash table [khint_t]
//
//! @function
//  @abstract     Iterate over the entries in the hash table
//  @param  h     Pointer to the hash table [khash_t(name)*]
//  @param  kvar  Variable to which key will be assigned
//  @param  vvar  Variable to which value will be assigned
//  @param  code  Block of code to execute
//
//! @function
//  @abstract     Iterate over the values in the hash table
//  @param  h     Pointer to the hash table [khash_t(name)*]
//  @param  vvar  Variable to which value will be assigned
//  @param  code  Block of code to execute
//
// More conenient interfaces
//! @function
//  @abstract     Instantiate a hash set containing integer keys
//  @param  name  Name of the hash table [symbol]
//
//! @function
//  @abstract     Instantiate a hash map containing integer keys
//  @param  name  Name of the hash table [symbol]
//  @param  khval_t  Type of values [type]
//
//! @function
//  @abstract     Instantiate a hash map containing 64-bit integer keys
//  @param  name  Name of the hash table [symbol]
//
//! @function
//  @abstract     Instantiate a hash map containing 64-bit integer keys
//  @param  name  Name of the hash table [symbol]
//  @param  khval_t  Type of values [type]
//
//! @function
//  @abstract     Instantiate a hash map containing const char* keys
//  @param  name  Name of the hash table [symbol]
//
//! @function
//  @abstract     Instantiate a hash map containing const char* keys
//  @param  name  Name of the hash table [symbol]
//  @param  khval_t  Type of values [type]
//
//
// * Hash type.
//
//
// * Allocate a new hash.
//
//
// * Destroy the hash.
//
//
// * Hash size.
//
//
// * Remove all pairs in the hash.
//
//
// * Iterate hash keys and ptrs, populating
// * `key` and `val`.
//
//
// * Iterate hash keys, populating `key`.
//
//
// * Iterate hash ptrs, populating `val`.
//
// protos

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

// asprintf - add c-binding for implemention function
func asprintf(arg0 [][]byte, arg1 ...[]byte) int32 {
	return int32(C.asprintf(unsafe.Pointer(&arg0), (*C.char)(unsafe.Pointer(&arg1[0][0]))))
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
