package main

import (
	"DPoSDemo/Block"
	"math/rand"
	"time"
	"fmt"
)

var blockChain []Block.Block

//存放代理人的地址
var delegate = [] string{"aaa","bbb","ccc","ddd"}


//随机调换
func randeDelegate() {
	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(delegate))
	index1 := rand.Intn(len(delegate))

	if index1 == index {
	  Label :
		index1 = rand.Intn(len(delegate))
		if index1 == index{
			goto Label
		}
	}

	tempDelegate := delegate[index]
	delegate[index] = delegate[index1]
	delegate[index1] = tempDelegate
}


func main()  {

	var firstBlock Block.Block

	blockChain = append(blockChain,firstBlock)

	//通过变量n按顺序让代理人作为矿工
	var n  = 0

	var temp = 0

	for {

		randeDelegate()
		//每隔3秒产生一个新区快
		time.Sleep(3 * time.Second)

		var nextBlock = Block.GenerateNextBlock(temp,blockChain[temp],delegate[n])

        blockChain = append(blockChain,nextBlock)

        n ++

		temp ++

		n = n %len(delegate)

		fmt.Println(blockChain)
	}

}
