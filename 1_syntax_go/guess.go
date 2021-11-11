package main

//猜数字小游戏
import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
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
