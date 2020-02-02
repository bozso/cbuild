//
//	Package - transpiled by c4go
//
//	If you have found any issues, please raise an issue at:
//	https://github.com/Konstantin8105/c4go/
//

package main

import "os"
import "fmt"


// options_t - transpiled function from  /home/istvan/packages/downloaded/cbuild/cbuild.c:35
type options_t struct {
	force int32
}

// generate - transpiled function from  /home/istvan/packages/downloaded/cbuild/cbuild.c:21
func generate(filename []byte, force int32, no_output int32) []package_t {
	var error_ []byte
	var pkg []package_t = index_new(filename, (*[1000000][]byte)(unsafe.Pointer(&error_))[:], force, no_output)
	lex_item_unfreed()
	if len(error_) != 0 {
		//if (pkg == NULL || pkg->errors > 0) return NULL;
		noarch.Fprintf(noarch.Stderr, []byte("%s\n\x00"), error_)
		return nil
	}
	return pkg
}

// do_generate - transpiled function from  /home/istvan/packages/downloaded/cbuild/cbuild.c:39
func do_generate(cli []cli_t, cmd []byte, arg interface{}) int32 {
	var opts []options_t = arg.([]options_t)
	if cli[0].argc < 1 {
		noarch.Fprintf(noarch.Stderr, []byte("no root module specified\n\x00"))
		return -1
	}
	if int32((opts[0].force)) != 0 {
		fmt.Printf("FORCED REBUILD\n")
	}
	var root []package_t = generate(cli[0].argv[0], opts[0].force, 0)
	if len(root) == 0 {
		noarch.Exit(-1)
	}
	return 0
}

// do_build - transpiled function from  /home/istvan/packages/downloaded/cbuild/cbuild.c:53
func do_build(cli []cli_t, cmd []byte, arg interface{}) int32 {
	var opts []options_t = arg.([]options_t)
	if cli[0].argc < 1 {
		noarch.Fprintf(noarch.Stderr, []byte("no root module specified\n\x00"))
		return -1
	}
	var root []package_t = generate(cli[0].argv[0], opts[0].force, 0)
	if len(root) == 0 {
		noarch.Exit(-1)
	}
	var mkfile_name []byte = makefile_write(root, cli[0].argv[0])
	var result int32 = makefile_make(root, mkfile_name)
	if result != 0 {
		noarch.Exit(result)
	}
	return 0
}

// clean_generated - transpiled function from  /home/istvan/packages/downloaded/cbuild/cbuild.c:69
func clean_generated(pkg []package_t) {
	if len(pkg) == 0 || int32((pkg[0].c_file)) != 0 || int32((pkg[0].exported)) == 0 {
		return
	}
	pkg[0].exported = 0
	if pkg[0].generated != nil {
		noarch.Printf([]byte("unlink: %s\n\x00"), pkg[0].generated)
		noarch.Unlink(pkg[0].generated)
	}
	if pkg[0].header != nil {
		noarch.Printf([]byte("unlink: %s\n\x00"), pkg[0].header)
		noarch.Unlink(pkg[0].header)
	}
	{
		var val interface{}
		{
			var k khiter_t = khiter_t((khint_t(0)))
			for ; k < khiter_t(((pkg[0].deps)[0].n_buckets)); k++ {
				if !noarch.Not((pkg[0].deps)[0].flags[k>>uint64(4)] >> uint64(uint32((khint32_t((khint_t((k & khiter_t((khint_t((khint32_t((uint32(15))))))) << uint64(1)))))))) & khint32_t((3))) {
					// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cbuild.c:82 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = n_buckets
					// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cbuild.c:82 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = flags
					continue
				}
				// Warning (*ast.MemberExpr):  /home/istvan/packages/downloaded/cbuild/cbuild.c:82 :cannot determine type for LHS '[hash_t * hash_t *]', will use 'void *' for all fields. Is lvalue = true. n.Name = vals
				val = (pkg[0].deps)[0].vals[k]
				{
					var imp []package_import_t = val.([]package_import_t)
					clean_generated(imp[0].pkg)
				}
			}
		}
	}
}

// do_clean - transpiled function from  /home/istvan/packages/downloaded/cbuild/cbuild.c:88
func do_clean(cli []cli_t, cmd []byte, arg interface{}) int32 {
	var opts []options_t = arg.([]options_t)
	if cli[0].argc < 1 {
		noarch.Fprintf(noarch.Stderr, []byte("no root module specified\n\x00"))
		return -1
	}
	var root []package_t = generate(cli[0].argv[0], opts[0].force, 1)
	if len(root) == 0 {
		noarch.Exit(-1)
	}
	var mkfile_name []byte = makefile_write(root, cli[0].argv[0])
	var result int32 = makefile_clean(root, mkfile_name)
	clean_generated(root)
	if result != 0 {
		noarch.Exit(result)
	}
	return 0
}

// main - transpiled function from  /home/istvan/packages/downloaded/cbuild/cbuild.c:106
func main() {
	argc := int32(len(os.Args))
	argv := [][]byte{}
	for _, argvSingle := range os.Args {
		argv = append(argv, []byte(argvSingle))
	}
	defer noarch.AtexitRun()
	var options options_t = options_t{0}
	var c []cli_t = cli_new([]byte("<root module>\x00"))
	cli_flag_bool(c, (*[1000000]int32)(unsafe.Pointer(&options.force))[:], cli_flag_options{[]byte("force rebuilding assets\x00"), []byte("force\x00"), []byte("f\x00"), 0})
	cli_command(c, []byte("build\x00"), do_build, []byte("generate code and build\x00"), 1, (*[1000000]options_t)(unsafe.Pointer(&options))[:])
	cli_command(c, []byte("generate\x00"), do_generate, []byte("generate .c .h and .mk files\x00"), 0, (*[1000000]options_t)(unsafe.Pointer(&options))[:])
	cli_command(c, []byte("clean\x00"), do_clean, []byte("clean generated files\x00"), 0, (*[1000000]options_t)(unsafe.Pointer(&options))[:])
	var result int32 = cli_parse(c, argc, argv)
	cli_free(c)
	noarch.Exit(int32(result))
}
