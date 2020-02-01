import (
    "io"
    "strings"
    "path/filepath"
    "github.com/bozso/cbuild/parser"    
    "github.com/bozso/cbuild/utils"    
)

var idCache = make(map[string])

func pkgName(relPath string) string {
	buffer := relPath
	filename := filepath.Base(buffer)

    ll := len(filename)
    // trim the .c extension - implement this on proper go way
	filename[ll - 2] = 0 

	for ii := 0; ii < ll; ii++ {
		c := &filename[ii]
        
        switch *c {
			case '.', '-', ' ':
				*c = '_'
			default:
		}
	}

	name = filename
	return name
}

var suffixLen = len(".module.c")

func assertName(relativePath string, err *string) (err error) {
	err = nil
    ll := len(relativePath)

	if ll < suffix_l || ".module.c" == relative_path[ll - suffixL:ll] {
        err = fmt.Errorf("Unsupported input filename '%s', Expecting '<file>.module.c'",
                relativePath)
	}
	return
}

func GeneratedName(path string) string {
	return strings.Replace(path, ".module", "")
}

func Parse(
    input io.Reader
    out io.Writer
    rel, key, generated string
    force, silent bool
) (p T, e error) {
	p.deps       = hash_new();
	p.exports    = hash_new();
	p.symbols    = hash_new();
	p.ordered    = NULL;
	p.sourceAbs = key
	p.generated  = generated
	p.out        = out
	p.name       = pkgName(generated)
	p.force      = force
	p.silent     = silent

    pathCache[key] = p

	e = parser.Parse(input, rel, p, e)
	return
}

func new_(relativePath string, force, silent bool) (p T, e error) {
    e = assert_name(relativePath)
    if e != nil { return }
	
    key, e := filepath.Abs(relativePath)
    
    if e != nil { return }

	p, ok := pathCache[key]
    
    if ok {
        return
    }

	input, err := os.Open(relativePath)
    
    if err != nil { return }

	generated := generatedName(key)
    
    var out io.Writer
	if force || (!silent && utils.Newer(relativePath, generated)) {
		out, err = atomic.Open(generated)
        
        if err != nil { return }
		}
	}

	p, e = Parse(input, out, relativePath, key, generated, force, silent)
    
    if e != nil { return }
    
	if (p == NULL || *error != NULL) {
		free(key);
		if (out) atomic.abort(out);
		return NULL;
	}

	if (out) stream.close(out);
	return
}

func (p *T) Release() {
    delete(pathCache, p.sourceAbs)
}

