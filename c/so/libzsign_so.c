#include "libzsign_so.h"
#include <dlfcn.h>

int do_so_func(char *argv[])
{
    void* handle;
    typedef int (*FPTR)(char *[]);

    handle = dlopen("./libzsign.so", 1);
    FPTR fptr = (FPTR)dlsym(handle, "start");

    int result = (*fptr)(argv);
    return result;
}