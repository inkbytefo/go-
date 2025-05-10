source_filename = "goplus_module"

define i32 @anonymous_func() {
entry:
	%x = alloca i32
	store i32 10, i32* %x
	%y = alloca i32
	store i32 20, i32* %y
	ret i32 0
}

define i32 @main() {
entry:
	ret i32 0
}
