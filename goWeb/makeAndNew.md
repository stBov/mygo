make、new操作</br>
make 用于内建类型（map 、 slice 和 channel ） 的内存分配。 new 用于各种类型的内存分配。</br>
内建函数 new 本质上说跟其它语言中的同名函数功能一样： new(T) 分配了零值填充的 T 类型的内存空</br>
间，并且返回其地址，即一个 *T 类型的值。用Go的术语说，它返回了一个指针，指向新分配的类</br>
型 T 的零值。有一点非常重要：new 返回指针。</br>
</br>
内建函数 make(T, args) 与 new(T) 有着不同的功能，make只能创建 slice 、 map 和 channel ，并</br>
且返回一个有初始值(非零)的 T 类型，而不是 *T 。本质来讲，导致这三个类型有所不同的原因是指向数</br>
据结构的引用在使用前必须被初始化。例如，一个 slice ，是一个包含指向数据（内部 array ） 的指</br>
针、长度和容量的三项描述符；在这些项目被初始化之前， slice 为 nil 。对</br>
于 slice 、 map 和 channel 来说， make 初始化了内部的数据结构，填充适当的值。</br>
make 返回初始化后的（非零） 值。</br>
下面这个图详细的解释了 new 和 make 之间的区别。</br>
</br>
图2.5 make和new对应底层的内存分配

零值</br>
关于“零值”，所指并非是空值，而是一种“变量未填充前”的默认值，通常为0。 此处罗列 部分类型 的 “零值”</br>
Go基础</br>
51int 0</br>
int8 0</br>
int32 0</br>
int64 0</br>
uint 0x0</br>
rune 0 //rune的实际类型是 int32</br>
byte 0x0 // byte的实际类型是 uint8</br>
float32 0 //长度为 4 byte</br>
float64 0 //长度为 8 byte</br>
bool false</br>
string ""</br>
