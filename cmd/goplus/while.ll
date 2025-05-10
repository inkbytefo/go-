source_filename = "goplus_module"

define i32 @anonymous_func() {
entry:
	%i = alloca i32
	store i32 0, i32* %i
	%sum = alloca i32
	store i32 0, i32* %sum
	br label %while.cond

while.cond:
	%0 = load i32, i32* %i
	%1 = icmp slt i32 %0, 10
	br i1 %1, label %while.body, label %while.end

while.body:
	%2 = load i32, i32* %sum
	%3 = load i32, i32* %sum
	%4 = load i32, i32* %i
	%5 = add i32 %3, %4
	store i32 %5, i32* %sum
	%6 = load i32, i32* %i
	%7 = load i32, i32* %i
	%8 = add i32 %7, 1
	store i32 %8, i32* %i
	br label %while.cond

while.end:
	ret i32 0
}

define i32 @main() {
entry:
	ret i32 0
}
