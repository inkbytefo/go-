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
	%3 = load i32, i32* %z
	%4 = icmp sgt i32 %3, 25
	br i1 %4, label %if.then, label %if.else

if.then:
	%5 = load i32, i32* %z
	%6 = load i32, i32* %z
	%7 = sub i32 %6, 5
	store i32 %7, i32* %z
	br label %if.end

if.else:
	%8 = load i32, i32* %z
	%9 = load i32, i32* %z
	%10 = add i32 %9, 5
	store i32 %10, i32* %z
	br label %if.end

if.end:
	%i = alloca i32
	store i32 0, i32* %i
	br label %while.cond

while.cond:
	%11 = load i32, i32* %i
	%12 = icmp slt i32 %11, 5
	br i1 %12, label %while.body, label %while.end

while.body:
	%13 = load i32, i32* %z
	%14 = load i32, i32* %z
	%15 = add i32 %14, 1
	store i32 %15, i32* %z
	%16 = load i32, i32* %i
	%17 = load i32, i32* %i
	%18 = add i32 %17, 1
	store i32 %18, i32* %i
	br label %while.cond

while.end:
	ret i32 0
}

declare i32 @add()

define i32 @main() {
entry:
	ret i32 0
}
