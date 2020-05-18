/*
单例模式：
1 单例类只能有一个实例
2 单例类必须自己创建自己的唯一实例
3 单例类必须给所有其他对象提供这一实例

应用：数据库连接池使用单例模式，降低打开或者关闭数据库连接带来的效率损耗；配置信息也可以使用单例模式

写法：懒汉式单例、饿汉式单例

*/
//懒汉式，线程不安全
//懒汉式，线程安全，加锁保证单例
//饿汉式，线程安全
package design

import "sync"

// type Singleton struct{}

// var singleton *Singleton
// var once sync.Once

// func GetInstance() *Singleton {
// 	once.Do(func() {
// 		singleton = &Singleton{}
// 	})
// 	return singleton
// }

type Singleton struct{}

var singleton *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		singleton = new(Singleton)
	})
	return singleton
}
