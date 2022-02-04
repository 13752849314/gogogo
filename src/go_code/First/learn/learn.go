package learn

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Books struct {
	title   string
	author  string
	subject string
	bookId  int
}

type show interface {
	print()
}

func (book Books) print() {
	fmt.Println(book.title)
}
func Main() {
	//变量声明
	var n = 123
	fmt.Printf("%d\n", n)
	var str = "hello,world"
	fmt.Printf("%s\n", str)
	var n1 int = 123
	fmt.Printf("%d\n", n1)
	str1 := "321"
	fmt.Println(str1)

	const Path string = "Address" //常量
	fmt.Println(Path)
	fmt.Println(&n) //取地址

	//数组
	var array = []float32{1, 2, 3, 4, 5}
	fmt.Println(array)

	var array2 [][]float32
	array2 = append(array2, array)
	array2 = append(array2, array)
	fmt.Println(array2)
	fmt.Println(array2[0][2])

	for i := 0; i < len(array2); i++ {
		for j := 0; j < len(array2[i]); j++ {
			fmt.Print(array2[i][j], " ")
		}
		fmt.Println()
	}
	for i := range array2 {
		fmt.Println(array2[i])
	}

	var ptr [5]*float32
	for i := 0; i < len(array); i++ {
		ptr[i] = &array[i]
	}
	for i := 0; i < len(ptr); i++ {
		fmt.Print(*ptr[i])
	}

	//切片
	fmt.Println(array[1:3])

	//结构体
	//type Books struct {
	//	title   string
	//	author  string
	//	subject string
	//	bookId  int
	//}

	var book Books
	book = Books{"数据结构", "严蔚敏", "计算机", 654}

	p1 := new(Books)
	p1.author = "123"
	p1.bookId = 345
	p1.subject = "math"
	p1.title = "gs"
	fmt.Println(book)
	fmt.Println(p1)
	p1.print()
	//集合
	var hg map[int]string
	hg = make(map[int]string)
	hg[20] = "a" //插入
	hg[1] = "23"
	hg[12] = "273"
	hg[13] = "243"
	for i, s := range hg {
		fmt.Println(i, "->", s)
	}
	delete(hg, 1) //删除
	for i, s := range hg {
		fmt.Println(i, "->", s)
	}

	//type show interface {
	//	print()
	//}

	var file, _ = os.OpenFile("D:\\go\\poj\\src\\go_code\\First\\learn\\hg.text", os.O_RDWR|os.O_CREATE, 0775)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	_, err := file.WriteString(time.Now().Local().String() + "\n")
	if err != nil {
		return
	}
	defer file.Close()
}
