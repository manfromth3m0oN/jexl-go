#include <stdarg.h>
#include <stdbool.h>
#include <stdint.h>
#include <stdlib.h>

typedef struct JexlEngine JexlEngine;

struct JexlEngine *new_engine(const char *context_ptr, const char *script_ptr);

void free_engine(struct JexlEngine *ptr);

const char *run_engine(struct JexlEngine *ptr);

const char *eval(const char *script);
