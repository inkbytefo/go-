source_filename = "goplus_module"

define i32 @anonymous_func() {
entry:
	%x = alloca i32
	store i32 10, i32* %x
	%0 = load i32, i32* %x
	%1 = icmp sgt i32 %0, 5
	br i1 %1, label %if.then, label %if.else

if.then:
	%2 = load i32, i32* %x
	%3 = load i32, i32* %x
	%4 = add i32 %3, 1
	store i32 %4, i32* %x
	br label %if.end

if.else:
	%5 = load i32, i32* %x
	%6 = load i32, i32* %x
	%7 = sub i32 %6, 1
	store i32 %7, i32* %x
	br label %if.end

if.end:
	ret i32 0
}

define i32 @main() {
entry:
	ret i32 0
}
