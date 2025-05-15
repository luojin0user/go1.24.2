package demo

var sink interface{}
var g1, g2 *int

type Data struct {
	value int
}

/*
func escape() *int {
	v := 42 // 局部变量
	// 启动goroutine访问变量（闭包捕获导致逃逸）
	go func() {
		v = 100 // 在goroutine中修改局部变量
	}()
	return &v // 返回局部变量的地址
}

func main() {
	x := escape()
	// 这里x指向的变量已经逃逸到堆上
	// 因为goroutine可能在任何时间修改它
	sink = x
}
*/
// 处理数据的函数
/*
func processData(d *int) {

	*d = *d + 2
	_ = d
}

func main2() {
	// 创建一个共享的数据对象
	data := 42

	// 启动多个goroutine，这些goroutine会共享同一个数据对象
	go processData(&data) // 启动goroutine并传递共享数据

	go func(d *int) {
		*d = *d + 2
		_ = d
	}(&data)
}
*/
/*
func LEAppendUint32(b []byte, v uint32) []byte {
	return append(b,
		byte(v),
		byte(v>>8),
		byte(v>>16),
		byte(v>>24),
	)
}
*/
/*
func return2() string {
	i := 2
	k := "hello"
	return string(k[i:])
}

func return3() *int {
	i := 3
	t := &i
	return t
}
*/

// GetOSVersion returns OS version, kernel and bitness
//
//	func GetOSVersion(b []byte) (k string) {
//		k = string(b)
//		return
//	}
type Address struct {
	Name    string // Proper name; may be empty.
	Address string // user@domain
}

func hhh() []*Address {
	return []*Address{{
		Name:    "111",
		Address: "spec",
	}}
}

// Uitoa converts val to a decimal string.
/*
func Uitoa(val uint) string {
	if val == 0 { // avoid string allocation
		return "0"
	}
	var buf [20]byte // big enough for 64bit value base 10
	i := len(buf) - 1
	for val >= 10 {
		q := val / 10
		buf[i] = byte('0' + val - q*10)
		i--
		val = q
	}
	// val < 10
	buf[i] = byte('0' + val)
	return string(buf[i:])
}
*/

/*
// Generate 返回一个闭包，闭包中捕获了外部局部变量 a
func Generate() func() int {
	a := 0              // 局部变量 a
	return func() int { // 闭包捕获 a
		a++ // 对 a 进行读写
		return a
	}
}

func main() {
	fn := Generate() // 此时 a 已被判定逃逸到堆
	_ = fn
}
*/
/*
func param14a(x [4]*int) interface{} { // ERROR "leaking param: x$"

	return x // ERROR "x escapes to heap"
}

func param14a2(kint int) interface{} { // ERROR "leaking param: x$"

	return kint // ERROR "x escapes to heap"
}

func param14a3(pint [4]*int) [4]*int { // ERROR "leaking param: x$"

	return pint // ERROR "x escapes to heap"
}

func param14a4(pint2 *int) *int { // ERROR "leaking param: x$"

	return pint2 // ERROR "x escapes to heap"
}

func param9(p ***int) **int { // ERROR "leaking param: p to result ~r0 level=1"

	return *p
}
*/
/*
func caller9a() {
	i := 0
	p := &i
	p2 := &p
	_ = param9(&p2)
}
*/
/*
func returnAddress() *int {
	i_testAddr := 10
	y := &i_testAddr
	return y // 局部变量 i 的地址被返回，i 会逃逸到堆上
}

func returnAddress2() *int {
	i_testAddr := 10
	return &i_testAddr // 局部变量 i 的地址被返回，i 会逃逸到堆上
}

func returnAddress3() **int {
	// i_testAddr := 10
	// p = &i_testAddr
	//return y // 局部变量 i 的地址被返回，i 会逃逸到堆上
	i := 10
	p1 := &i
	p2 := &p1
	return p2 // 局部变量 i 的地址被返回，i 会逃逸到堆上
}

func outerLoopReference() {
	var outerRef *int

	for i := 0; i < 10; i++ {
		i_testOuter := i
		outerRef = &i_testOuter // 循环内部的变量 num 被外部引用，num 会逃逸到堆上
	}

	_ = outerRef
}
*/
/*
func bigVariable() {
	// 定义一个非常大的数组，可能会导致栈溢出，因此会逃逸到堆上
	bigArray := [1000000]int{}
	_ = bigArray
}

func variableSize() {
	x := 10
	// 切片的大小在运行时确定，可能会逃逸到堆上
	slice := make([]int, x)
	_ = slice
}
*/
/*
func paramPointee(p *int) {
	//sink = p // 参数 p 指向的变量被全局变量引用，该变量会逃逸到堆上
	i := 10
	//k := 20
	var k *int
	g1, k = p, &i
	sink = k
	//*g1 += i

}

func paramCall() {
	i_testParam := 20
	paramPointee(&i_testParam)
}
*/
/*
func referencedByGlobal() {
	i_testGlobal := 30
	b := &i_testGlobal // 局部变量 i 被全局变量引用，i 会逃逸到堆上
	sink = &b
	//sink = &i_testGlobal // 局部变量 i 被全局变量引用，i 会逃逸到堆上
}
*/
/*
func level0() ***int {

	i := 0    // ERROR "moved to heap: i"
	p0 := &i  // ERROR "moved to heap: p0"
	p1 := &p0 // ERROR "moved to heap: p1"
	p2 := &p1 // ERROR "moved to heap: p2"
	p3 := p2
	// p1 -> p2 -> p3 -> sink
	//    -1     0    0
	sink = p3
	return p2
	//g1, g2 = p0, *p1

	//return p2
}
*/
/*
func referencedByGlobal2() {
	//i_testGlobal := 30
	//b := 20 // 局部变量 i 被全局变量引用，i 会逃逸到堆上
	var b *int
	for i := 0; i < 1; i++ {
		q := 10

		b = &q // 局部变量 i 被全局变量引用，i 会逃逸到堆上
	}
	//q := &b
	//sink = &p
	//sink = &i_testGlobal // 局部变量 i 被全局变量引用，i 会逃逸到堆上
	//sink, p = &q, &b
	sink = &b
}
*/
/*
var gg **int

type X struct {
	i *int
	j int
}

func indirect() {
	//i_testAddr := 10
	//p := &i_testAddr
	// sink = &p // 局部变量 i 的地址被返回，i 会逃逸到堆上
	//var p **int
	// // var i int
	// *p = &i
	q := 10
	x := []X{}
	x[0].i = &q

	i := 10
	*gg = &i
}

type P struct {
	X *int
	Y int
}

var p P

func test() {

	// i := 20
	// sink = &i

	// i := 20
	// p1 := i
	// var p2 *int
	// *p2 = p1

	// i := 20
	// p.X = &i // p.X 是合法左值

	// var k = [2]*int{}
	// i := 20
	// k[0] = &i

	// i := 20
	// m := make(map[string]*int)
	// m["Go"] = &i // m["Go"] 是合法左值

	// i := 20
	// (sink) = &i

}
*/
