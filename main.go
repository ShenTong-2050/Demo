package main

import (
	"demo/go-design/factory/abstract_factory"
	"fmt"
	"reflect"
	"sort"
	"strings"
	//"os"
	//"io/ioutil"
)

func init() {
	/*ch := make(chan struct{})
	go func() {
		defer close(ch)
		fmt.Println("init:",time.Now())
		time.Sleep(time.Second * 2)
	}()
	<-ch*/
}

func main() {
	// test2.Build()

	//immok.ImmokMain()
	//testA()

	//file,err := os.Open("maze.in")
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()

	//body,_ := ioutil.ReadAll(file)
	//fmt.Println(string(body))

	//testA()

	//fmt.Println(1)

	//interfaces.MethodMain()
	//interfaces.EnforcementMain()
	//interfaces.ChangeMain()

	//goroutine.GoroutineMain()
	//goroutine.StorageMain()
	//goroutine.GoExitMain()

	//passpackage.GetPrivatePkg()
	//fmt.Println("wait init done",time.Now())

	//lissajous.MakeLissajous()

	factory := abstract_factory.GetMySports("nike")
	factory.MakeShoes().SetLogo("my nike")
	logo := factory.MakeShoes().GetLogo()
	fmt.Println(logo)
}

func copyReturn() {
	sep := "--"
	ts := make([]byte,100)
	sa := []string{"1","2","3","4","5"}
	bp := copy(ts,sa[0])
	for i,v := range sa[1:] {
		bp += copy(ts[bp:],sep)
		fmt.Println(i,bp)
		bp += copy(ts[bp:],v)
		fmt.Println(i,bp)
	}
	fmt.Println(string(ts))
}

func CheckPermutation(s1 string, s2 string) bool {
	s1Len,s2Len := len(s1),len(s2)
	m1,m2 := map[string]int{},map[string]int{}
	if s1Len != s2Len {
		return false
	}
	for _,v := range s1 {
		s2Index := strings.Index(s2,string(v))
		m1[string(v)]++
		m2[string(v)]++
		if s2Index < 0 {
			return false
		}
	}
	return true
}

func none() {
	sl := make([]int,0,10)
	appendFn := func(s []int) {
		// 值传递,s 并不是 sl
		// 但数组是引用类型,可以修改底层数组
		fmt.Printf("s ptr(old): %p -> %v\n",s,s)
		s = append(s,10,20,30)
		fmt.Printf("s ptr(new): %p -> %v\n",s,s)
	}

	fmt.Printf("ptr(sl): %p -> %v\n",sl,sl)
	appendFn(sl)
	fmt.Println(sl)

	fmt.Println(sl[:10])
}

func testA() {
	mg := make([]int,3,6)
	mg = []int{1,2,3}
	b := []int{4,5,6,7,8}
	fmt.Printf("the mg is:%d, len(mg):%d, cap(mg):%d\n",mg,len(mg),cap(mg))
	mg = append(mg,b...)
	fmt.Printf("appended mg is:%d, len(mg):%d, cap(mg):%d\n",mg,len(mg),cap(mg))
	fmt.Println(b[2:])
}

func capSlice() {
	arr := [4]int{10,20,30,40}
	slice := arr[0:2]
	fmt.Println("slice=>",slice)
	testSlice1 := slice
	//testSlice2 := append(append(append(slice,1),2),3)
	testSlice2 := fnAppend(slice)
	fmt.Println("testSlice1=>",testSlice1)
	fmt.Println("testSlice2=>",testSlice2)
	slice[0] = 11
	fmt.Println("edited testSlice1=>",testSlice1)
	fmt.Println("edited testSlice2=>",testSlice2)
	fmt.Println("testSlice1[0]=>",testSlice1[0])
	fmt.Println("testSlice2[0]=>",testSlice2[0])
}

func fnAppend(s []int) []int {
	return append(append(append(s,1),2),3)
}

func app() func(string) string {
	t := "Hi"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	fmt.Println(reflect.TypeOf(t))
	fmt.Println(reflect.TypeOf(c))
	return c
}

func test(a,b []int) (diffA, diffB []int) {
	sort.Ints(a)
	sort.Ints(b)
	i,n := 0, len(a)
	j,m := 0, len(b)
	for {
		if i == n {
			diffB = append(diffB,b[j:]...)
			return
		}
		if j == m {
			diffA = append(diffA,a[i:]...)
			return
		}
		x , y := a[i], b[j]
		if x < y {
			diffA = append(diffA,x)
			i++
		} else if x > y {
			diffB = append(diffB,y)
			j++
		} else {
			i++
			j++
		}
	}
}

