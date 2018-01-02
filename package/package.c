

#include "../deps/hash/hash.h"
#include <stdlib.h>
#include <stdbool.h>

#include <stdio.h>

#include "../../stream/stream.h"

typedef struct {
  hash_t   * deps;
  hash_t   * exports;
  hash_t   * symbols;
  void    ** ordered;
  size_t     n_exports;
  char     * name;
  char     * source_abs;
  char     * source_rel;
  char     * generated;
  char     * header_abs;
  bool       exported;
  bool       c_file;
  stream_t * out;
} package_t;



hash_t * package_path_cache = NULL;
hash_t * package_id_cache = NULL;

/* TODO: fix the circrular dependency issue */

package_t * (*package_new)(const char * relative_path, char ** error) = NULL;

void package_emit(package_t * pkg, char * value) {
  if (pkg->out) stream_write(pkg->out, value, strlen(value));
}

package_t * package_c_file(char * abs_path, char ** error) {
  package_t * cached = hash_get(package_path_cache, abs_path);
  if (cached != NULL) return cached;

  package_t * pkg = calloc(1, sizeof(package_t));

  pkg->name       = abs_path;
  pkg->source_abs = abs_path;
  pkg->generated  = abs_path;
  pkg->c_file     = true;

  hash_set(package_path_cache, abs_path, pkg);
  return pkg;
}

void package_free(package_t * pkg) {
  if (pkg == NULL) return;

  hash_clear(pkg->deps);
  hash_clear(pkg->exports);
  hash_clear(pkg->symbols);

  free(pkg->name);
  free(pkg->source_abs);
  free(pkg->source_rel);
  free(pkg->generated);
  free(pkg->header_abs);
  free(pkg);
}