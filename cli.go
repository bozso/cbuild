//
//	Package - transpiled by c4go
//
//	If you have found any issues, please raise an issue at:
//	https://github.com/Konstantin8105/c4go/
//

/* AST Error :
cannot parse line: `UnusedAttr 0x7b3be80 <line:15:451> Inherited unused`. could not match regexp with string
^(?P<address>[0-9a-fx]+) <(?P<position>.*)>(?P<unused> unused)?[\s]*$
0x7b3be80 <line:15:451> Inherited unused

*/
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:148 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = n_buckets
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:148 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = flags
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:148 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = keys
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:155 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = n_buckets
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:155 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = flags
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:155 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = keys
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:155 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = vals
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:168 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = n_buckets
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:168 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = flags
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:168 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = keys
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:172 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = n_buckets
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:172 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = flags
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:172 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = keys
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:172 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = vals
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:242 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = n_buckets
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:242 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = flags
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:242 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = keys
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:242 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = vals
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:249 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = n_buckets
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:249 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = flags
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:249 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = vals

package main

import "reflect"
import "github.com/Konstantin8105/c4go/noarch"
import "unsafe"

// khint32_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/deps/hash/khash.h:135
/* AST Error :
cannot parse line: `UnusedAttr 0x7b3be80 <line:15:451> Inherited unused`. could not match regexp with string
^(?P<address>[0-9a-fx]+) <(?P<position>.*)>(?P<unused> unused)?[\s]*$
0x7b3be80 <line:15:451> Inherited unused

*/type khint32_t = uint32

// khint64_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/deps/hash/khash.h:141
type khint64_t = uint32

// khint_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/deps/hash/khash.h:162
type khint_t = khint32_t

// khiter_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/deps/hash/khash.h:163
type khiter_t = khint_t

// __ac_HASH_UPPER - transpiled function from  /home/istvan/packages/downloaded/cbuild/deps/hash/khash.h:192
var __ac_HASH_UPPER float64 = 0.77

// kh_cstr_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/deps/hash/khash.h:611
type kh_cstr_t = []byte

// kh_ptr_s - transpiled function from  /home/istvan/packages/downloaded/cbuild/deps/hash/hash.h:15
type kh_ptr_s struct {
	n_buckets   khint_t
	size        khint_t
	n_occupied  khint_t
	upper_bound khint_t
	flags       []khint32_t
	keys        []kh_cstr_t
	vals        []interface{}
}

// kh_ptr_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/deps/hash/hash.h:15
type kh_ptr_t = kh_ptr_s

// hash_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/deps/hash/hash.h:21
type hash_t = kh_ptr_t

// cli_flag_options - transpiled function from  /home/istvan/packages/downloaded/cbuild/cli.c:10
type cli_flag_options struct {
	description []byte
	long_name   []byte
	short_name  []byte
	required    int32
}

// flag_type - transpiled function from  /home/istvan/packages/downloaded/cbuild/cli.c:17
type flag_type = int32

const (
	flag_type_bool   flag_type = 0
	flag_type_int              = 1
	flag_type_string           = 2
)

// flag_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/cli.c:23
type flag_t struct {
	dest           interface{}
	type_          flag_type
	description    []byte
	long_name      []byte
	short_name     []byte
	short_name_buf [3]byte
	required       int32
	has_value      int32
}

// cli_cmd_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/cli.c:39
type cli_cmd_t struct {
	cb          interface{}
	ctx         interface{}
	name        []byte
	description []byte
}

// cli_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/cli.c:51
type cli_t struct {
	flags           []hash_t
	commands        []hash_t
	default_command []cli_cmd_t
	has_commands    int32
	usage           []byte
	name            []byte
	argc            int32
	argv            [][]byte
}

// cli_cmd_cb - transpiled function from  /home/istvan/packages/downloaded/cbuild/cli.c:65
type cli_cmd_cb = func([]cli_t, []byte, interface{}) int32

// __ac_X31_hash_string - transpiled function from  /home/istvan/packages/downloaded/cbuild/deps/hash/khash.h:395
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

// __ac_Wang_hash - transpiled function from  /home/istvan/packages/downloaded/cbuild/deps/hash/khash.h:412
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

// kh_init_ptr - transpiled function from  /home/istvan/packages/downloaded/cbuild/deps/hash/hash.h:15
func kh_init_ptr() []kh_ptr_t {
	//
	// hash.h
	//
	// Copyright (c) 2012 TJ Holowaychuk <tj@vision-media.ca>
	//
	// pointer hash
	return make([]kh_ptr_t, 1)
}

// kh_destroy_ptr - transpiled function from  /home/istvan/packages/downloaded/cbuild/deps/hash/hash.h:15
func kh_destroy_ptr(h []kh_ptr_t) {
	if h != nil {
		_ = h[0].keys
		_ = h[0].flags
		_ = h[0].vals
		_ = h
	}
}

// kh_clear_ptr - transpiled function from  /home/istvan/packages/downloaded/cbuild/deps/hash/hash.h:15
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

// kh_get_ptr - transpiled function from  /home/istvan/packages/downloaded/cbuild/deps/hash/hash.h:15
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

// kh_resize_ptr - transpiled function from  /home/istvan/packages/downloaded/cbuild/deps/hash/hash.h:15
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

// kh_put_ptr - transpiled function from  /home/istvan/packages/downloaded/cbuild/deps/hash/hash.h:15
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

// kh_del_ptr - transpiled function from  /home/istvan/packages/downloaded/cbuild/deps/hash/hash.h:15
func kh_del_ptr(h []kh_ptr_t, x khint_t) {
	if x != khint_t(h[0].n_buckets) && noarch.Not(h[0].flags[x>>uint64(4)]>>uint64(uint32((khint32_t((x & khint_t((khint32_t((uint32(15))))) << uint64(1))))))&khint32_t((3))) {
		h[0].flags[x>>uint64(4)] |= khint32_t((uint32(1 << uint64(uint32((khint32_t((x & khint_t((khint32_t((uint32(15))))) << uint64(1)))))))))
		h[0].size -= khint_t((khint32_t((uint32(1)))))
	}
}

// flag_free - transpiled function from  /home/istvan/packages/downloaded/cbuild/cli.c:34
func flag_free(f []flag_t) {
	_ = f[0].long_name
	_ = f
}

// cmd_free - transpiled function from  /home/istvan/packages/downloaded/cbuild/cli.c:46
func cmd_free(c []cli_cmd_t) {
	_ = c[0].name
	_ = c
}

// cli_new - transpiled function from  /home/istvan/packages/downloaded/cbuild/cli.c:67
func cli_new(usage []byte) []cli_t {
	var cli []cli_t = make([]cli_t, 1)
	cli[0].flags = kh_init_ptr()
	cli[0].commands = kh_init_ptr()
	cli[0].usage = usage
	return cli
}

// flag - transpiled function from  /home/istvan/packages/downloaded/cbuild/cli.c:76
func flag(cli []cli_t, dest interface{}, type_ flag_type, options cli_flag_options) {
	var f []flag_t = make([]flag_t, 1)
	f[0].type_ = type_
	f[0].dest = dest
	f[0].required = options.required
	f[0].description = options.description
	f[0].long_name = make([]byte, noarch.Strlen(options.long_name)+int32(3))
	noarch.Strcpy(f[0].long_name, []byte("--\x00"))
	noarch.Strcat(f[0].long_name, options.long_name)
	hash_set(cli[0].flags, f[0].long_name, f)
	if len(options.short_name) == 0 {
		f[0].short_name = nil
	} else {
		f[0].short_name_buf[:][0] = '-'
		f[0].short_name_buf[:][1] = options.short_name[0]
		f[0].short_name_buf[:][2] = byte(0)
		f[0].short_name = f[0].short_name_buf[:]
		hash_set(cli[0].flags, f[0].short_name, f)
	}
}

// cli_flag_bool - transpiled function from  /home/istvan/packages/downloaded/cbuild/cli.c:102
func cli_flag_bool(cli []cli_t, out []int32, options cli_flag_options) {
	options.required = 0
	flag(cli, out, flag_type_bool, options)
}

// cli_flag_int - transpiled function from  /home/istvan/packages/downloaded/cbuild/cli.c:107
func cli_flag_int(cli []cli_t, out []int32, options cli_flag_options) {
	flag(cli, out, flag_type_string, options)
}

// cli_flag_string - transpiled function from  /home/istvan/packages/downloaded/cbuild/cli.c:111
func cli_flag_string(cli []cli_t, out [][]byte, options cli_flag_options) {
	flag(cli, out, flag_type_string, options)
}

// cli_command - transpiled function from  /home/istvan/packages/downloaded/cbuild/cli.c:115
func cli_command(cli []cli_t, name []byte, cb cli_cmd_cb, description []byte, is_default int32, ctx interface{}) {
	cli[0].has_commands = 1
	var cmd []cli_cmd_t = make([]cli_cmd_t, 1)
	cmd[0].name = noarch.Strdup(name)
	cmd[0].cb = cli_cmd_cb(cb)
	cmd[0].ctx = ctx
	cmd[0].description = description
	if int32((is_default)) != 0 {
		cli[0].default_command = cmd
	}
	hash_set(cli[0].commands, cmd[0].name, cmd)
}

// parse_flag - transpiled function from  /home/istvan/packages/downloaded/cbuild/cli.c:128
func parse_flag(flag []flag_t, arg []byte) {
	switch uint32(int32((flag[0].type_))) {
	case uint32(int32((flag_type_bool))):
		if flag[0].dest != nil {
			(flag[0].dest.([]int32))[0] = 1
		}
	case uint32(int32((flag_type_int))):
		if flag[0].dest != nil {
			(flag[0].dest.([]int32))[0] = noarch.Strtol(arg, nil, 0)
		}
	case uint32(int32((flag_type_string))):
		if flag[0].dest != nil {
			(flag[0].dest.([][]byte))[0] = arg
		}
	}
	flag[0].has_value = 1
}

// cli_usage - transpiled function from  /home/istvan/packages/downloaded/cbuild/cli.c:142
func cli_usage(cli []cli_t) int32 {
	noarch.Fprintf(noarch.Stderr, []byte("usage %s: [options]%s %s\n\n\x00"), cli[0].name, func() []byte {
		if int32((cli[0].has_commands)) != 0 {
			return []byte(" command\x00")
		}
		return []byte("\x00")
	}(), cli[0].usage)
	{
		// flags
		noarch.Fprintf(noarch.Stderr, []byte("FLAGS: \n\n\x00"))
		var width int32 = noarch.Strlen([]byte("--help\x00"))
		{
			var key []byte
			{
				var k khiter_t = khiter_t((khint_t(0)))
				for ; k < khiter_t(((cli[0].flags)[0].n_buckets)); k++ {
					if !noarch.Not((cli[0].flags)[0].flags[k>>uint64(4)] >> uint64(uint32((khint32_t((khint_t((k & khiter_t((khint_t((khint32_t((uint32(15))))))) << uint64(1)))))))) & khint32_t((3))) {
						// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:148 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = n_buckets
						// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:148 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = flags
						continue
					}
					// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:148 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = keys
					key = []byte(((cli[0].flags)[0].keys[k]))
					{
						width = int32(func() uint32 {
							if noarch.Strlen(key) > int32(uint32(width)) {
								return uint32(noarch.Strlen(key))
							}
							return uint32(width)
						}())
					}
				}
			}
		}
		var w int32 = int32(uint32(width) - uint32(noarch.Strlen([]byte("--help\x00"))))
		noarch.Fprintf(noarch.Stderr, []byte("          %s%*.*s %s\t%s\n\x00"), []byte("--help\x00"), w, w, []byte(" \x00"), []byte("-h\x00"), []byte("prints this message\x00"))
		{
			var key []byte
			var val interface{}
			{
				var k khiter_t = khiter_t((khint_t(0)))
				for ; k < khiter_t(((cli[0].flags)[0].n_buckets)); k++ {
					if !noarch.Not((cli[0].flags)[0].flags[k>>uint64(4)] >> uint64(uint32((khint32_t((khint_t((k & khiter_t((khint_t((khint32_t((uint32(15))))))) << uint64(1)))))))) & khint32_t((3))) {
						// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:155 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = n_buckets
						// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:155 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = flags
						continue
					}
					// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:155 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = keys
					key = []byte(((cli[0].flags)[0].keys[k]))
					// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:155 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = vals
					val = (cli[0].flags)[0].vals[k]
					{
						if int32(key[1]) != int32('-') {
							continue
						}
						var flag []flag_t = val.([]flag_t)
						var w int32 = int32(uint32(width) - uint32(noarch.Strlen(key)))
						noarch.Fprintf(noarch.Stderr, []byte("          %s%*.*s %s\t%s\n\x00"), key, w, w, []byte(" \x00"), func() []byte {
							if flag[0].short_name != nil {
								return flag[0].short_name
							}
							return []byte("  \x00")
						}(), flag[0].description)
					}
				}
			}
		}
		noarch.Fprintf(noarch.Stderr, []byte("\n\x00"))
	}
	if int32((cli[0].has_commands)) != 0 {
		noarch.Fprintf(noarch.Stderr, []byte("COMMANDS: \n\n\x00"))
		var width int32
		{
			var key []byte
			{
				var k khiter_t = khiter_t((khint_t(0)))
				for ; k < khiter_t(((cli[0].commands)[0].n_buckets)); k++ {
					if !noarch.Not((cli[0].commands)[0].flags[k>>uint64(4)] >> uint64(uint32((khint32_t((khint_t((k & khiter_t((khint_t((khint32_t((uint32(15))))))) << uint64(1)))))))) & khint32_t((3))) {
						// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:168 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = n_buckets
						// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:168 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = flags
						continue
					}
					// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:168 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = keys
					key = []byte(((cli[0].commands)[0].keys[k]))
					{
						width = int32(func() uint32 {
							if noarch.Strlen(key) > int32(uint32(width)) {
								return uint32(noarch.Strlen(key))
							}
							return uint32(width)
						}())
					}
				}
			}
		}
		{
			var key []byte
			var val interface{}
			{
				var k khiter_t = khiter_t((khint_t(0)))
				for ; k < khiter_t(((cli[0].commands)[0].n_buckets)); k++ {
					if !noarch.Not((cli[0].commands)[0].flags[k>>uint64(4)] >> uint64(uint32((khint32_t((khint_t((k & khiter_t((khint_t((khint32_t((uint32(15))))))) << uint64(1)))))))) & khint32_t((3))) {
						// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:172 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = n_buckets
						// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:172 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = flags
						continue
					}
					// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:172 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = keys
					key = []byte(((cli[0].commands)[0].keys[k]))
					// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:172 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = vals
					val = (cli[0].commands)[0].vals[k]
					{
						var cmd []cli_cmd_t = val.([]cli_cmd_t)
						var w int32 = int32(uint32(width) - uint32(noarch.Strlen(key)))
						noarch.Fprintf(noarch.Stderr, []byte("%s%s%*.*s  %s\n\x00"), func() []byte {
							if (int64(uintptr(unsafe.Pointer(&cmd[0])))/int64(64) - int64(uintptr(unsafe.Pointer(&cli[0].default_command[0])))/int64(64)) == 0 {
								return []byte("(default) \x00")
							}
							return []byte("          \x00")
						}(), cmd[0].name, w, w, []byte(" \x00"), cmd[0].description)
					}
				}
			}
		}
		noarch.Fprintf(noarch.Stderr, []byte("\n\x00"))
	}
	return -1
}

// cli_parse - transpiled function from  /home/istvan/packages/downloaded/cbuild/cli.c:185
func cli_parse(cli []cli_t, argc int32, argv [][]byte) int32 {
	cli[0].argv = make([][]byte, 1*int32(uint32(argc-1)))
	cli[0].name = argv[0]
	var cmd []cli_cmd_t
	var i int32
	{
		// collect args, commands, flags
		for i = 1; i < argc; i++ {
			var arg []byte = argv[i]
			if noarch.Strcmp([]byte("--\x00"), arg) == 0 {
				break
			}
			if noarch.Strcmp([]byte("-h\x00"), arg) == 0 || noarch.Strcmp([]byte("--help\x00"), arg) == 0 {
				return cli_usage(cli)
			}
			var arg_allocated []byte = noarch.Strdup(arg)
			var flag []flag_t = hash_get(cli[0].flags, arg_allocated).([]flag_t)
			if len(flag) != 0 {
				parse_flag(flag, arg)
				_ = arg_allocated
				continue
			}
			if int32(arg[0]) == int32('-') {
				noarch.Fprintf(noarch.Stderr, []byte("Unrecognized flag '%s'\n\n\x00"), arg)
				_ = arg_allocated
				return cli_usage(cli)
			}
			if len(cmd) == 0 {
				cmd = hash_get(cli[0].commands, arg_allocated).([]cli_cmd_t)
				if cmd != nil {
					_ = arg_allocated
					continue
				}
			}
			cli[0].argv[cli[0].argc] = arg
			cli[0].argc++
			_ = arg_allocated
		}
	}
	for ; i < argc; i++ {
		// collect args after a --
		cli[0].argv[cli[0].argc] = argv[i]
		cli[0].argc++
	}
	if len(cmd) == 0 {
		cmd = cli[0].default_command
	}
	if len(cmd) == 0 {
		return func() int32 {
			if int32((cli[0].has_commands)) != 0 {
				return cli_usage(cli)
			}
			return 0
		}()
	}
	var cb cli_cmd_cb = cmd[0].cb.(cli_cmd_cb)
	var result int32 = cb(cli, cmd[0].name, cmd[0].ctx)
	if result == 0 {
		return result
	}
	return cli_usage(cli)
}

// cli_free - transpiled function from  /home/istvan/packages/downloaded/cbuild/cli.c:241
func cli_free(cli []cli_t) {
	{
		var key []byte
		var val interface{}
		{
			var k khiter_t = khiter_t((khint_t(0)))
			for ; k < khiter_t(((cli[0].flags)[0].n_buckets)); k++ {
				if !noarch.Not((cli[0].flags)[0].flags[k>>uint64(4)] >> uint64(uint32((khint32_t((khint_t((k & khiter_t((khint_t((khint32_t((uint32(15))))))) << uint64(1)))))))) & khint32_t((3))) {
					// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:242 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = n_buckets
					// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:242 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = flags
					continue
				}
				// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:242 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = keys
				key = []byte(((cli[0].flags)[0].keys[k]))
				// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:242 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = vals
				val = (cli[0].flags)[0].vals[k]
				{
					if int32(key[0]) == int32('-') && int32(key[1]) == int32('-') {
						flag_free(val.([]flag_t))
					}
				}
			}
		}
	}
	kh_destroy_ptr(cli[0].flags)
	{
		var val interface{}
		{
			var k khiter_t = khiter_t((khint_t(0)))
			for ; k < khiter_t(((cli[0].commands)[0].n_buckets)); k++ {
				if !noarch.Not((cli[0].commands)[0].flags[k>>uint64(4)] >> uint64(uint32((khint32_t((khint_t((k & khiter_t((khint_t((khint32_t((uint32(15))))))) << uint64(1)))))))) & khint32_t((3))) {
					// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:249 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = n_buckets
					// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:249 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = flags
					continue
				}
				// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cli.c:249 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = vals
				val = (cli[0].commands)[0].vals[k]
				{
					cmd_free(val.([]cli_cmd_t))
				}
			}
		}
	}
	kh_destroy_ptr(cli[0].commands)
	_ = cli[0].argv
	_ = cli
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
