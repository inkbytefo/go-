source_filename = "goplus_module"

define i32 @anonymous_func(i32 %a, i32 %b) {
entry:
	%a.addr = alloca i32
	store i32 %a, i32* %a.addr
	%b.addr = alloca i32
	store i32 %b, i32* %b.addr
	%0 = load i32, i32* %a.addr
	%1 = load i32, i32* %b.addr
	%2 = add i32 %0, %1
	ret i32 %2
}

define i32 @anonymous_func() {
entry:
	%x = alloca i32
	store i32 10, i32* %x
	%y = alloca i32
	store i32 20, i32* %y
	%z = alloca i32
	%0 = load i32, i32* %x
	%1 = load i32, i32* %y
	%2 = call i32 @add(i32 %0, i32 %1)
	store i32 %2, i32* %z
	ret i32 0
}

declare i32 @add()

define i32 @main() {
entry:
	ret i32 0
}
