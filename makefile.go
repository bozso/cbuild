//
//	Package - transpiled by c4go
//
//	If you have found any issues, please raise an issue at:
//	https://github.com/Konstantin8105/c4go/
//

package main

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
