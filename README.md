goscript
------------------------------------------

This is a simple script language.

Example：

~~~
#! /usr/bin/env goscript

Print("hello")

a=1+1

Print(a)

if 1>2 {
    Print("a")
} else {
    Print("b")
}

for i=1;i<10;i++ {
   Print(i)
}

func A(a) {
 Print(a)
}

A(1)
~~~

