//
//	Package - transpiled by c4go
//
//	If you have found any issues, please raise an issue at:
//	https://github.com/Konstantin8105/c4go/
//

/* AST Error :
cannot parse line: `UnusedAttr 0x6383910 <line:15:451> Inherited unused`. could not match regexp with string
^(?P<address>[0-9a-fx]+) <(?P<position>.*)>(?P<unused> unused)?[\s]*$
0x6383910 <line:15:451> Inherited unused

*/
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/makefile.c:119 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = n_buckets
// #include </usr/include/stdio.h>
// #include </usr/include/unistd.h>
// #include </usr/include/libgen.h>
import "C"



// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/makefile.c:119 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = flags
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/makefile.c:119 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = vals
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/makefile.c:129 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = n_buckets
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/makefile.c:129 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = flags
// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/makefile.c:129 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = vals

package main

import "reflect"
import "unsafe"
import "github.com/Konstantin8105/c4go/noarch"
// khint32_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/deps/hash/khash.h:135
/* AST Error :
cannot parse line: `UnusedAttr 0x6383910 <line:15:451> Inherited unused`. could not match regexp with string
^(?P<address>[0-9a-fx]+) <(?P<position>.*)>(?P<unused> unused)?[\s]*$
0x6383910 <line:15:451> Inherited unused

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
// package_var_type - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/package.h:8
type package_var_type = int32

const (
	build_var_set         package_var_type = 0
	build_var_set_default                  = 1
	build_var_append                       = 2
)
// package_var_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/package.h:14
type package_var_t struct {
	name      []byte
	value     []byte
	operation package_var_type
}
// stream_error_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/../deps/stream/stream.h:7
type stream_error_t struct {
	message []byte
	code    int32
}
// stream_read_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/../deps/stream/stream.h:12
type stream_read_t = func(interface{} , interface{} , uint32 , []stream_error_t)(noarch.SsizeT)
// stream_write_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/../deps/stream/stream.h:13
type stream_write_t = func(interface{} , interface{} , uint32 , []stream_error_t)(noarch.SsizeT)
// stream_pipe_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/../deps/stream/stream.h:14
type stream_pipe_t = func(interface{} , interface{} , []stream_error_t)(noarch.SsizeT)
// stream_close_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/../deps/stream/stream.h:15
type stream_close_t = func(interface{} , []stream_error_t)(noarch.SsizeT)
// stream_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/../deps/stream/stream.h:17
type stream_t struct {
	ctx    interface{}
	read   stream_read_t
	write  stream_write_t
	pipe   stream_pipe_t
	close  stream_close_t
	error_ stream_error_t
	type_  int32
}
// package_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/package.h:22
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
// package_export_type - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/export.h:4
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
// package_export_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/export.h:14
type package_export_t struct {
	local_name  []byte
	export_name []byte
	declaration []byte
	symbol      []byte
	type_       package_export_type
}
// package_import_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/package/import.h:8
type package_import_t struct {
	alias    []byte
	filename []byte
	c_file   int32
	pkg      []package_t
}
// ops - transpiled function from  /home/istvan/packages/downloaded/cbuild/makefile.c:19
var ops [][]byte = [][]byte{[]byte(":=\x00"), []byte("?=\x00"), []byte("+=\x00")}
// makevars - transpiled function from  /home/istvan/packages/downloaded/cbuild/makefile.c:25
type makevars struct {
	makefile      []byte
	makefile_dir  []byte
	makefile_base []byte
	target        []byte
	cwd           []byte
}
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
			h = h<<uint64(5)-h+khint_t(s[0])
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
	key ^= key>>uint64(10)
	key += key<<uint64(3)
	key ^= key>>uint64(6)
	key += ^(key << uint64(11))
	key ^= key>>uint64(16)
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
		mask = khint_t(h[0].n_buckets)-khint_t((khint32_t((1))))
		k = __ac_X31_hash_string([]byte((key)))
		i = k&mask
		last = i
		for noarch.Not(h[0].flags[i>>uint64(4)]>>uint64(uint32((khint32_t((i & khint_t((khint32_t((uint32(15))))) << uint64(1))))))&khint32_t((2))) && (uint32((h[0].flags[i>>uint64(4)]>>uint64(uint32((khint32_t((i & khint_t((khint32_t((uint32(15))))) << uint64(1))))))&khint32_t((1)))) != 0 || !(noarch.Strcmp([]byte((h[0].keys[i])), []byte((key))) == 0)) {
			i = (i+func() khint_t {
				step ++
				return step
			}())&mask
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
		new_n_buckets |= new_n_buckets>>uint64(1)
		new_n_buckets |= new_n_buckets>>uint64(2)
		new_n_buckets |= new_n_buckets>>uint64(4)
		new_n_buckets |= new_n_buckets>>uint64(8)
		new_n_buckets |= new_n_buckets>>uint64(16)
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
				new_mask = new_n_buckets-khint_t((khint32_t((1))))
				if 1 != 0 {
					val = h[0].vals[j]
				}
				h[0].flags[j>>uint64(4)] |= khint32_t((uint32(1 << uint64(uint32((khint32_t((j & khint_t((khint32_t((uint32(15))))) << uint64(1)))))))))
				for 1 != 0 {
					var k khint_t
					var i khint_t
					var step khint_t
					k = __ac_X31_hash_string([]byte((key)))
					i = k&new_mask
					for noarch.Not(new_flags[i>>uint64(4)] >> uint64(uint32((khint32_t((i & khint_t((khint32_t((uint32(15))))) << uint64(1)))))) & khint32_t((2))) {
						i = (i+func() khint_t {
							step ++
							return step
						}())&new_mask
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
		h[0].upper_bound = khint_t(float64(uint32((khint32_t((khint_t(h[0].n_buckets))))))*__ac_HASH_UPPER+0.5)
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
		i = k&mask
		if uint32((h[0].flags[i>>uint64(4)] >> uint64(uint32((khint32_t((i & khint_t((khint32_t((uint32(15))))) << uint64(1)))))) & khint32_t((2)))) != 0 {
			x = i
		} else {
			last = i
			for noarch.Not(h[0].flags[i>>uint64(4)]>>uint64(uint32((khint32_t((i & khint_t((khint32_t((uint32(15))))) << uint64(1))))))&khint32_t((2))) && (uint32((h[0].flags[i>>uint64(4)]>>uint64(uint32((khint32_t((i & khint_t((khint32_t((uint32(15))))) << uint64(1))))))&khint32_t((1)))) != 0 || !(noarch.Strcmp([]byte((h[0].keys[i])), []byte((key))) == 0)) {
				if uint32((h[0].flags[i>>uint64(4)] >> uint64(uint32((khint32_t((i & khint_t((khint32_t((uint32(15))))) << uint64(1)))))) & khint32_t((1)))) != 0 {
					site = i
				}
				i = (i+func() khint_t {
					step ++
					return step
				}())&mask
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
// get_makevars - transpiled function from  /home/istvan/packages/downloaded/cbuild/makefile.c:33
func get_makevars(pkg []package_t, makefile []byte) makevars {
	var v makevars = makevars{makefile, noarch.Strdup(dirname(makefile)), noarch.Strdup(__xpg_basename(makefile)), noarch.Strdup(__xpg_basename(pkg[0].generated)), getcwd(nil, 0)}
	var ext []byte = noarch.Strrchr(v.target, int32('.'))
	if noarch.Strcmp(pkg[0].name, []byte("main\x00")) == 0 {
		ext[0] = byte(0)
	} else {
		ext[1] = 'a'
	}
	chdir(v.makefile_dir)
	return v
}
// clear_makevars - transpiled function from  /home/istvan/packages/downloaded/cbuild/makefile.c:53
func clear_makevars(v makevars, result int32, cmd []byte) int32 {
	chdir(v.cwd)
	_ = v.cwd
	_ = v.target
	_ = v.makefile_dir
	_ = v.makefile_base
	_ = v.makefile
	_ = cmd
	if result == 0 || result == -1 {
		return result
	}
	if result == 127 {
		return -1
	}
	return result & 65280 >> uint64(8)
}
// makefile_make - transpiled function from  /home/istvan/packages/downloaded/cbuild/makefile.c:68
func makefile_make(pkg []package_t, makefile []byte) int32 {
	if len(pkg) == 0 {
		return -1
	}
	var v makevars = get_makevars(pkg, makefile)
	var cmd []byte
	asprintf((*[1000000][]byte)(unsafe.Pointer(&cmd))[:], []byte("make -f %s %s\x00"), v.makefile_base, v.target)
	return clear_makevars(v, noarch.System(cmd), cmd)
}
// makefile_clean - transpiled function from  /home/istvan/packages/downloaded/cbuild/makefile.c:78
func makefile_clean(pkg []package_t, makefile []byte) int32 {
	if len(pkg) == 0 {
		return -1
	}
	var v makevars = get_makevars(pkg, makefile)
	var cmd []byte
	asprintf((*[1000000][]byte)(unsafe.Pointer(&cmd))[:], []byte("make -f %s CLEAN_%s\x00"), v.makefile_base, v.target)
	return clear_makevars(v, noarch.System(cmd), cmd)
}
// write_deps - transpiled function from  /home/istvan/packages/downloaded/cbuild/makefile.c:88
func write_deps(pkg []package_t, root []package_t, out []stream_t, deps []byte) []byte {
	if len(pkg) == 0 || int32((pkg[0].exported)) != 0 {
		return deps
	}
	pkg[0].exported = 1
	var source []byte = utils_relative(root[0].generated, pkg[0].generated)
	var object []byte = noarch.Strdup(source)
	object[noarch.Strlen(source)-int32(1)] = 'o'
	var len_ uint32 = func() uint32 {
		if len(deps) == 0 {
			return 0
		}
		return uint32(noarch.Strlen(deps))
	}()
	deps = realloc(deps, len_+uint32(noarch.Strlen(object))+2).([]byte)
	noarch.Sprintf(deps[0+len_:], []byte(" %s\x00"), object)
	var package_rel []byte = utils_relative(root[0].source_abs, pkg[0].generated)
	stream_printf(out, []byte("#dependencies for package '%s'\n\x00"), package_rel)
	_ = package_rel
	var i int32
	for i = 0; uint32(i) < uint32(pkg[0].n_variables); i++ {
		var v package_var_t = pkg[0].variables[i]
		stream_printf(out, []byte("%s %s %s\n\x00"), v.name, ops[v.operation], v.value)
	}
	stream_printf(out, []byte("%s: %s\x00"), object, source)
	_ = source
	_ = object
	if len(pkg[0].deps) == 0 {
		stream_printf(out, []byte("\n\n\x00"))
		return deps
	}
	{
		var val interface{}
		{
			var k khiter_t = khiter_t((khint_t(0)))
			for ; k < khiter_t(((pkg[0].deps)[0].n_buckets)); k++ {
				if !noarch.Not((pkg[0].deps)[0].flags[k>>uint64(4)] >> uint64(uint32((khint32_t((khint_t((k & khiter_t((khint_t((khint32_t((uint32(15))))))) << uint64(1)))))))) & khint32_t((3))) {
					// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/makefile.c:119 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = n_buckets
					// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/makefile.c:119 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = flags
					continue
				}
				// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/makefile.c:119 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = vals
				val = (pkg[0].deps)[0].vals[k]
				{
					var dep []package_import_t = val.([]package_import_t)
					if len(dep) == 0 && len(dep[0].pkg) == 0 && dep[0].pkg[0].header != nil {
						var path []byte = utils_relative(root[0].source_abs, dep[0].pkg[0].header)
						stream_printf(out, []byte(" %s\x00"), path)
						_ = path
					}
				}
			}
		}
	}
	stream_printf(out, []byte("\n\n\x00"))
	{
		var val interface{}
		{
			var k khiter_t = khiter_t((khint_t(0)))
			for ; k < khiter_t(((pkg[0].deps)[0].n_buckets)); k++ {
				if !noarch.Not((pkg[0].deps)[0].flags[k>>uint64(4)] >> uint64(uint32((khint32_t((khint_t((k & khiter_t((khint_t((khint32_t((uint32(15))))))) << uint64(1)))))))) & khint32_t((3))) {
					// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/makefile.c:129 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = n_buckets
					// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/makefile.c:129 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = flags
					continue
				}
				// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/makefile.c:129 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = vals
				val = (pkg[0].deps)[0].vals[k]
				{
					var dep []package_import_t = val.([]package_import_t)
					deps = write_deps(dep[0].pkg, root, out, deps)
				}
			}
		}
	}
	return deps
}
// get_makefile_name - transpiled function from  /home/istvan/packages/downloaded/cbuild/makefile.c:137
func get_makefile_name(path []byte) []byte {
	var name []byte = noarch.Strdup(path)
	name[noarch.Strlen(name)-noarch.Strlen([]byte("module.c\x00"))] = byte(0)
	noarch.Strcat(name, []byte("mk\x00"))
	return name
}
// makefile_write - transpiled function from  /home/istvan/packages/downloaded/cbuild/makefile.c:144
func makefile_write(pkg []package_t, name []byte) []byte {
	var target []byte
	var mkfile_name []byte = get_makefile_name(name)
	var mkfile []stream_t = atomic_stream_open(mkfile_name)
	var deps []byte = write_deps(pkg, pkg, mkfile, nil)
	if noarch.Strcmp(pkg[0].name, []byte("main\x00")) == 0 {
		var buf []byte = noarch.Strdup(pkg[0].generated)
		var base []byte = __xpg_basename(buf)
		asprintf((*[1000000][]byte)(unsafe.Pointer(&target))[:], []byte("%.*s\x00"), noarch.Strlen(base)-2, base)
		_ = buf
		stream_printf(mkfile, []byte("%s:%s\n\x00"), target, deps)
		stream_printf(mkfile, []byte("\t$(CC) $(CFLAGS) $(LDFLAGS) %s -o %s $(LDLIBS)\n\n\x00"), deps, target)
	} else {
		target = make([]byte, noarch.Strlen(pkg[0].name)+int32(3))
		noarch.Strcpy(target, pkg[0].name)
		noarch.Strcat(target, []byte(".a\x00"))
		stream_printf(mkfile, []byte("%s:%s\n\x00"), target, deps)
		stream_printf(mkfile, []byte("\tar rcs $@ $^\n\n\x00"), deps, target)
	}
	stream_printf(mkfile, []byte("CLEAN_%s:\n\x00"), target)
	stream_printf(mkfile, []byte("\trm -rf %s%s\n\x00"), target, deps)
	_ = target
	stream_close(mkfile)
	_ = deps
	return mkfile_name
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



// __xpg_basename - add c-binding for implemention function
func __xpg_basename(arg0 []byte) []byte {
	if arg0 == nil {
		return []byte{}
	}
	return []byte(C.GoString(C.__xpg_basename((*C.char)(unsafe.Pointer(&arg0[0])))))
}


// asprintf - add c-binding for implemention function
func asprintf(arg0 [][]byte, arg1 ...[]byte) int32 {
	return int32(C.asprintf(unsafe.Pointer(&arg0), (*C.char)(unsafe.Pointer(&arg1[0][0]))))
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

