# Learning go with tests

[Reference URL](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/iteration)

## Run tests

~~The package should be named as `SOMETHING_test`~~. Just ignore that

- `go test -v ./folder`

## Modules

- Hello world
- Integers
- Iteration
- Arrays and Slices
- Structs, methods and interfaces

## Quick Notes (bc i have short memory 🙃)

- `t.Helper()` is used in helper test functions so that if the test fails
  the line number reported will be in our function call rather than inside
  our test helper.
- `slices` package includes the `Equal` method already.
- `func x(numbers ...[]int)` is called a _Variadic Function_ and represents
  a scenario when you can pass _n_ number of parameters of the same type
- `%#v` prints the struct values in a `printf` format
- In Go **interface resolution is implicit**. If the type you pass in
  matches what the interface is asking for, it will compile.
- The structs can be defined with his properties names so that the reading
  can be easier: `{shape: Rectangle{Width: 12, Height: 6}, want: 72.0}`
