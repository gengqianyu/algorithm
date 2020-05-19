//内存管理之程序装入
//静态装入：编译时决定程序代码，机器指令，变量常量等的物理绝对地址。
//	  缺点：地址空间连续固定，可移植性差，不灵活。（使用最早的单道程序）

//可重定位装入(静态重定位)：编译时决定程序的逻辑地址。装入内存时进行地址转换（基址《不可变》+逻辑地址=物理地址）。
//	  因为程序装入是从基址开始，往后依次连续装入，占用连续的存储空间，所以只要一个重定位基址就可以。
//	  缺点：地址空间连续，程序执行之前就要分配所需全部地址空间。而且程序执行期间内存不能移动（地址空间写死）。

//(动态运行时装入)动态重定位：编译决定程序逻辑地址；装入内存时进行地址转换，利用重定位(基址)寄存器存储进程的起始基址。
//装入的时候，将编译好的二进制指令代码段，切分成内存页的大小的块，离散的装入内存中不同的页框中。因此基址寄存器中需要记录多个指向不同页框的重定位基址。
//cpu访问内存地址=重定位基址(可变)+程序逻辑地址；最厉害地方在于程序运行时，不用全部分配内存地址，可以用到时分块按需加载分配。
//优点：由于按需加载，可以将程序地址空间分配到不是连续的存储区。那么运行时我们想让程序发生移动就简单了。只要改变重定位(基址)寄存器里的地址就行，大大节省了内存。

// 静态链接。程序在运行前编译的最后进行模块链接。
// 动态链接：程序在运行时动态链接，按需加载模块。

//程序从产生=>运行的过程
//1,程序员用ide编写程序(a.go defined main func，b.go defined b func,c.go defined c func)
//2,经过编译这几个源代码文件，会被编译成几个与之对应的目标模块文件。并且目标模块文件已经包含了代码所对应的机器指令了。
// 这些指令的编址都是逻辑地址。每一个模块编址都是从逻辑地址0开始的(因为重定向寄存器基址可变，所以最后的物理地址还是不一样的)。就这样编译形成了机器语言。
//3,链接，把目标模块链接形成装入模块(静态链接)。除了自己写的程序机器指令还要链接库函数模块，比如fmt.printf 。(在windows中也就是.exe文件)
//注意：链接也可以在第4步装入时进行链接=》动态装入时链接。
//	   链接还可以在第5步运行时进行链接=》动态运行时链接。只有我们需要那个模块才把它调入内存。
//4,将装入模块，按照动态重定位的方式，将装入模块装入内存。
//5,cpu读取内存(重定位基址(可变)+指令逻辑地址)，取址执行。
package main

func main() {

}