package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func testMap() {
	//定义map,map属于引用类型，var 变量名 map[key的类型]值的类型
	var a map[string]int
	fmt.Println(a)
	// a["stu1"]=100 引用类型未初始化就进行复制会报错，未初始化时值为nil
	// 使用make进行初始化，
	a = make(map[string]int, 10) //map类型，容量为10
	//key相同时value被覆盖
	a["stu1"] = 100
	a["stu2"] = 100
	a["stu3"] = 100
	fmt.Println(a)
	fmt.Printf("%#v\n", a) //%#v显示原来数据类型

	// 声明时初始化
	var b map[string]int = map[string]int{
		"stu1": 100,
		"stu2": 200,
		"stu3": 300,
	}
	fmt.Println("b=", b)
	c := map[string]string{
		"1": "555",
		"2": "666",
	}
	fmt.Println(c)

	//访问map
	key := "1"
	fmt.Println(c[key], c["2"])

	// 判断一个key是否存在于map里面，访问不存在的key不会报错，会返回一个默认值
	var restul string
	var ok bool
	restul, ok = c["2"]
	if ok == false {
		fmt.Println("key is not exist,ok:", ok)
	} else {
		fmt.Println("key is exist,value :", restul, "ok:", ok)
	}

	//循环map
	for k, v := range b {
		fmt.Println(k, v)
	}
	delete(b, "stu3") //删除某个key及其对应的值
	fmt.Println(b)
	fmt.Println(len(b)) //长度
}

func testMap2() {
	//map属于引用类型，赋值，当作参数传递会修改原来的数据
	a := map[string]int{
		"1": 100,
		"2": 200,
	}
	b := a
	b["1"] = 200
	fmt.Println(a, b)
}

func sort_map() {
	//排序
	rand.Seed(time.Now().UnixNano())
	var a map[string]int = make(map[string]int, 1024) //使用make定义一个空的map

	for i := 0; i < 128; i++ {
		key := fmt.Sprintf("stu%d", i)
		val := rand.Intn(1000) //在[0,1000)内随机生成整数
		a[key] = val
	}

	//对map的key排序
	var keys []string = make([]string, 0, 128)
	for k, _ := range a {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Println(keys)
	for _, v := range keys {
		fmt.Printf("key:%s,val:%d\n", v, a[v])
	}
}

func slice_map()  {
	//切片类型的map，切片里面的元素类型为map
	var s []map[string]int
	s=make([]map[string]int,5,10)
	//初始化切片之后，里面map也需要进行初始化，不能直接赋值
	s[0]=make(map[string]int,10)
	s[0]["stu1"]=100
	s[0]["stu3"]=200
	s[0]["stu3"]=300
	for i,v:=range s{
		fmt.Printf("slice:%d=%v\n",i,v)
	}
	fmt.Println(s)

}

func map_slice()  {
	//map类型的切片，map里面的值的类型为切片
	var m map[string][]int
	m=make(map[string][]int,10)
	//先判断key是否存在，不存在就初始化里面的切片，再插入值
	key:="stu1"
	val,ok:=m[key]
	if !ok {
		m[key]=make([]int,0,10)
		val=m[key]
	}
	val=append(val,100)
	val=append(val,200)
	val=append(val,300)
	m[key]=val
	fmt.Println(m)
}

func main() {
	// testMap()
	// testMap2()
	// sort_map()
	// slice_map()
	map_slice()

}
