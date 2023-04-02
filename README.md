# GoLang-C-Bindings
Interpolation of C in GoLang - CGO - GoLang C Bindings

#### Step 1: Create C Binary
- 1.a (create object file)
  ```
  gcc -c example.c -o example.o
  ```
- 1.b (create binary from object)
  ```
  ar rcs libexample.a example.o
  ```
  
  #### Directory Structure
  ```
  main.go
  -lib
  --example.a
  -include
  --example.h
  ```
#### Step 2: (include in main.go file)
```
/*
#cgo CFLAGS: -I./include
#cgo LDFLAGS: -L./lib -lexample
#include "example.h"
#include <stdlib.h>
*/
```
**Note: In GoLang the above commented code is not a usual comment. it is called preambles.**

And All Done.
