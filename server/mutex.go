package main

/**
	mutex 互斥锁
	互斥锁是传统并发程序对共享资源进行访问的主要手段。是sync包中的Mutex结构体
	   type Mutex struct{}
	该结构体包括两个方法，可以说是非常简单使用的
	   func (m *Mutex) Lock() {}
	   func (m *Mutex) UnLock() {}
	我们来通过一个简单的例子来说明他的用法
 */
