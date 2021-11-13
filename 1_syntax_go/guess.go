package main

//猜数字小游戏
import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	//printTest()
	
	/* root,err := squareRoot(45)
	if err != nil {
		fmt.Println(err);
	}else {
		fmt.Printf("%.6f",root)
	} */
	
	var number int = 5
	Pointer(&number)
	fmt.Println(&number)
	fmt.Println(number)
}

func printTest()  {
	fmt.Println(1.0/3)
	fmt.Printf("%0.2f\n",1.0/3)
	result := fmt.Sprint(1.0/3)
    fmt.Println(result)
	result = fmt.Sprintf("%0.4f",1.0/3)
	fmt.Println(result)
}

func guessGame()  {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100)
	for i := 1; i <= 10; i++ {
		print("请输入你猜的数字！")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if target < guess {
			println("猜大了！还有", 10-i, "次机会")
		} else if target > guess {
			println("猜小了!还有", 10-i, "次机会")
		} else {
			fmt.Println("恭喜你，猜对了！ game over！")
			return
		}
	}
	fmt.Println("很遗憾，你失败了！目标数是：", target)
}

func squareRoot(number float64) (float64,error) {
	if number < 0 {
		return 0,fmt.Errorf("负数不能计算平方根！")
	}
	return math.Sqrt(number),nil
}
func Pointer(number *int) int{
	fmt.Println("number = ",number)
	*number = *number * 2
	fmt.Println("*number =",*number)
	return *number
}