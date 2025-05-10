source_filename = "goplus_module"

define i32 @anonymous_func() {
entry:
	%x = alloca i32
	store i32 10, i32* %x
	%y = alloca i32
	store i32 20, i32* %y
	%sum = alloca i32
	%0 = load i32, i32* %x
	%1 = load i32, i32* %y
	%2 = add i32 %0, %1
	store i32 %2, i32* %sum
	%diff = alloca i32
	%3 = load i32, i32* %x
	%4 = load i32, i32* %y
	%5 = sub i32 %3, %4
	store i32 %5, i32* %diff
	%prod = alloca i32
	%6 = load i32, i32* %x
	%7 = load i32, i32* %y
	%8 = mul i32 %6, %7
	store i32 %8, i32* %prod
	%quot = alloca i32
	%9 = load i32, i32* %x
	%10 = load i32, i32* %y
	%11 = sdiv i32 %9, %10
	store i32 %11, i32* %quot
	ret i32 0
}

define i32 @main() {
entry:
	ret i32 0
}
