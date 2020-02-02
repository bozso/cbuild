//
//	Package - transpiled by c4go
//
//	If you have found any issues, please raise an issue at:
//	https://github.com/Konstantin8105/c4go/
//

/* AST Error :
cannot parse line: `UnusedAttr 0x6a2ede0 <line:15:451> Inherited unused`. could not match regexp with string
^(?P<address>[0-9a-fx]+) <(?P<position>.*)>(?P<unused> unused)?[\s]*$
0x6a2ede0 <line:15:451> Inherited unused

*/

package main

// #include </usr/include/libgen.h>
// #include </usr/include/unistd.h>
import "C"

import "reflect"
import "unsafe"
import "github.com/Konstantin8105/c4go/noarch"

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

// khint32_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../package/../deps/hash/khash.h:135
type khint32_t = uint32

// khint64_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../package/../deps/hash/khash.h:141
type khint64_t = uint32

// khint_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../package/../deps/hash/khash.h:162
type khint_t = khint32_t

// khiter_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../package/../deps/hash/khash.h:163
type khiter_t = khint_t

// __ac_HASH_UPPER - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../package/../deps/hash/khash.h:192
var __ac_HASH_UPPER float64 = 0.77

// kh_cstr_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../package/../deps/hash/khash.h:611
type kh_cstr_t = []byte

// kh_ptr_s - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../package/../deps/hash/hash.h:15
type kh_ptr_s struct {
	n_buckets   khint_t
	size        khint_t
	n_occupied  khint_t
	upper_bound khint_t
	flags       []khint32_t
	keys        []kh_cstr_t
	vals        []interface{}
}

// kh_ptr_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../package/../deps/hash/hash.h:15
type kh_ptr_t = kh_ptr_s

// hash_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../package/../deps/hash/hash.h:21
type hash_t = kh_ptr_t

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

// parser_parse_fn - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/parser.c:18
type parser_parse_fn = func([]parser_parser_s) interface{}

// parser_parser_s - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/parser.c:20
type parser_parser_s struct {
	lexer  []lex_t
	state  parser_parse_fn
	items  []lex_item_stack_t
	pkg    []package_t
	errors int32
}

// parser_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/parser.c:20
type parser_t = parser_parser_s

// __ac_X31_hash_string - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../package/../deps/hash/khash.h:395
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

// __ac_Wang_hash - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../package/../deps/hash/khash.h:412
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

// kh_init_ptr - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../package/../deps/hash/hash.h:15
func kh_init_ptr() []kh_ptr_t {
	//
	// hash.h
	//
	// Copyright (c) 2012 TJ Holowaychuk <tj@vision-media.ca>
	//
	// pointer hash
	return make([]kh_ptr_t, 1)
}

// kh_destroy_ptr - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../package/../deps/hash/hash.h:15
func kh_destroy_ptr(h []kh_ptr_t) {
	if h != nil {
		_ = h[0].keys
		_ = h[0].flags
		_ = h[0].vals
		_ = h
	}
}

// kh_clear_ptr - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../package/../deps/hash/hash.h:15
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

// kh_get_ptr - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../package/../deps/hash/hash.h:15
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

// kh_resize_ptr - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../package/../deps/hash/hash.h:15
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

// kh_put_ptr - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../package/../deps/hash/hash.h:15
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

// kh_del_ptr - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/../package/../deps/hash/hash.h:15
func kh_del_ptr(h []kh_ptr_t, x khint_t) {
	if x != khint_t(h[0].n_buckets) && noarch.Not(h[0].flags[x>>uint64(4)]>>uint64(uint32((khint32_t((x & khint_t((khint32_t((uint32(15))))) << uint64(1))))))&khint32_t((3))) {
		h[0].flags[x>>uint64(4)] |= khint32_t((uint32(1 << uint64(uint32((khint32_t((x & khint_t((khint32_t((uint32(15))))) << uint64(1)))))))))
		h[0].size -= khint_t((khint32_t((uint32(1)))))
	}
}

// parser_parse - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/parser.c:28
func parser_parse(lexer []lex_t, start parser_parse_fn, pkg []package_t) int32 {
	var p []parser_t = make([]parser_t, 1)
	p[0].lexer = lexer
	p[0].state = start
	p[0].items = lex_item_stack_new(1)
	p[0].pkg = pkg
	p[0].errors = 0
	var cwd []byte = getcwd(nil, 0)
	var directory []byte = noarch.Strdup(lexer[0].filename)
	chdir(dirname(directory))
	_ = directory
	for len(parser_parse_fn(p[0].state)) != 0 {
		p[0].state = p[0].state(p).(parser_parse_fn)
	}
	lex_free(lexer)
	lex_item_stack_free(p[0].items)
	chdir(cwd)
	_ = cwd
	var errors int32 = p[0].errors
	_ = p
	if errors > 0 {
		noarch.Fprintf(noarch.Stderr, []byte("%d error%s generated\n\x00"), errors, func() []byte {
			if errors == 1 {
				return []byte("\x00")
			}
			return []byte("s\x00")
		}())
	}
	return errors
}

// parser_next - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/parser.c:56
func parser_next(p []parser_t) lex_item_t {
	var item lex_item_t
	if uint32(p[0].items[0].length) != 0 {
		item = lex_item_stack_pop(p[0].items)
	} else {
		item = lex_next_item(p[0].lexer)
	}
	return item
}

// parser_backup - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/parser.c:66
func parser_backup(p []parser_t, item lex_item_t) {
	p[0].items = lex_item_stack_push(p[0].items, item)
}

// parser_skip - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/parser.c:70
func parser_skip(p []parser_t, c4goArgs ...interface{}) lex_item_t {
	var args = create_va_list(c4goArgs)
	var len_ int32
	var i int32
	var types []lex_item_type = make([]lex_item_type, 14)
	va_start(args, p)
	for {
		types[len_] = va_arg(args).(lex_item_type)
		len_++
		if !(uint32(int32((types[len_-1]))) != 0) {
			break
		}
	}
	len_--
	var item lex_item_t = lex_item_t{lex_item_type((0)), nil, 0, 0, 0, 0, 0}
	for {
		var do_skip int32
		lex_item_free(item)
		item = parser_next(p)
		for i = 0; i < len_; i++ {
			if uint32(int32((types[i]))) == uint32(int32((item.type_))) {
				do_skip = 1
				break
			}
		}
		if do_skip != 0 {
			continue
		}
		return item
		if !(uint32(int32((item.type_))) != uint32(int32((item_eof))) && uint32(int32((item.type_))) != uint32(int32((item_error)))) {
			break
		}
	}
	return item
}

// parser_collect - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/parser.c:102
func parser_collect(p []parser_t, buf [][]byte, blen []noarch.SsizeT, c4goArgs ...interface{}) lex_item_t {
	var args = create_va_list(c4goArgs)
	var len_ int32
	var i int32
	var text []byte = func() []byte {
		if len(buf) == 0 {
			return nil
		}
		return buf[0]
	}()
	var text_len uint32 = uint32(func() int32 {
		if len(blen) == 0 {
			return 0
		}
		return int32(blen[0])
	}())
	var types []lex_item_type = make([]lex_item_type, 14)
	va_start(args, blen)
	for {
		types[len_] = va_arg(args).(lex_item_type)
		len_++
		if !(uint32(int32((types[len_-1]))) != 0) {
			break
		}
	}
	len_--
	var item lex_item_t
	for {
		var do_skip int32
		item = parser_next(p)
		for i = 0; i < len_; i++ {
			if uint32(int32((types[i]))) == uint32(int32((item.type_))) {
				do_skip = 1
				text = realloc(text, text_len+uint32(item.length)+1).([]byte)
				noarch.Strcpy(text[0+text_len:], item.value)
				text_len += uint32(item.length)
				break
			}
		}
		if do_skip != 0 {
			continue
		}
		break
		if !(uint32(int32((item.type_))) != uint32(int32((item_eof))) && uint32(int32((item.type_))) != uint32(int32((item_error)))) {
			break
		}
	}
	buf[0] = text
	blen[0] = noarch.SsizeT(text_len)
	return item
}

// parser_verrorf - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/parser.c:141
func parser_verrorf(p []parser_t, item lex_item_t, context []byte, fmt_ []byte, args *va_list) {
	var start uint32 = uint32(item.start)
	var line uint32 = uint32(item.line)
	var line_pos uint32 = uint32(item.line_pos)
	var col int32 = int32(start - line_pos)
	p[0].errors++
	noarch.Fprintf(noarch.Stderr, []byte("\x1b[1m%s:%ld:%d: \x1b[0m\x1b[1m\x1b[31merror: \x1b[0m\x1b[1m%s\x00"), p[0].lexer[0].filename, line+1, col+1, context)
	if uint32(int32((item.type_))) == uint32(int32((item_error))) {
		noarch.Fprintf(noarch.Stderr, []byte("%s\x00"), lex_item_to_string(item))
	} else {
		noarch.Vfprintf(noarch.Stderr, fmt_, args)
	}
	if len(p[0].lexer[0].input) != 0 && uint32(p[0].lexer[0].length) > line_pos {
		var line_start []byte = p[0].lexer[0].input[0+line_pos:]
		var line_end []byte = noarch.Strchr(line_start, int32('\n'))
		var line_length int32 = int32(func() uint32 {
			if len(line_end) == 0 {
				return uint32(col) + uint32(item.length)
			}
			return uint32(int32((int64(uintptr(unsafe.Pointer(&line_end[0])))/int64(1) - int64(uintptr(unsafe.Pointer(&line_start[0])))/int64(1))))
		}())
		noarch.Fprintf(noarch.Stderr, []byte("\x1b[0m\n%.*s\n\x00"), line_length, line_start)
		noarch.Fprintf(noarch.Stderr, []byte("\x1b[32m%*.*s^\n\x1b[0m\x00"), col, col, []byte(" \x00"))
	} else {
		noarch.Fprintf(noarch.Stderr, []byte("\x1b[0m\n\x00"))
	}
}

// parser_errorf - transpiled function from  /home/istvan/packages/downloaded/cbuild/parser/parser.c:174
func parser_errorf(p []parser_t, item lex_item_t, context []byte, fmt_ []byte, c4goArgs ...interface{}) {
	var args = create_va_list(c4goArgs)
	va_start(args, fmt_)
	parser_verrorf(p, item, context, fmt_, args)
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

// chdir - add c-binding for implemention function
func chdir(arg0 []byte) int32 {
	return int32(C.chdir((*C.char)(unsafe.Pointer(&arg0[0]))))
}

// dirname - add c-binding for implemention function
func dirname(arg0 []byte) []byte {
	if arg0 == nil {
		return []byte{}
	}
	return []byte(C.GoString(C.dirname((*C.char)(unsafe.Pointer(&arg0[0])))))
}

// getcwd - add c-binding for implemention function
func getcwd(arg0 []byte, arg1 uint32) []byte {
	if arg0 == nil {
		return []byte{}
	}
	return []byte(C.GoString(C.getcwd((*C.char)(unsafe.Pointer(&arg0[0])), C.ulong(arg1))))
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
