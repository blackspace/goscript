goscript
------------------------------------------

This is a simple script language.

Example：

~~~
#! /usr/bin/env goscript

print("hello")

a=1+1

print(a)

if 1>2 {
    print("a")
} else {
    print("b")
}

for i=1;i<10;i++ {
   print(i)
}

func A(a) {
 print(a)
}

A(1)
~~~

