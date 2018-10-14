package main

import "fmt"

//定义结构

type Block struct {
	//1,前区块哈希
	PrevHash []byte
	//2,当前区块哈希
	Hash []byte
	//3,数据
	Data []byte
}

//2,创建区块
func NewBlock(data string,prevBlockHash []byte) *Block {
	block:=Block{
		PrevHash:prevBlockHash,//先填空，后面再计算 //TODO
		Hash: []byte{},
		Data: []byte(data),
	}
	return &block

}

func main()  {
	block :=NewBlock("Jack转给Rose一枚比特币!",[]byte{})
	fmt.Printf("前区块哈希值:%x\n",block.PrevHash)
	fmt.Printf("当前区块哈希值： %x\n", block.Hash)
	fmt.Printf("区块数据 :%s\n", block.Data)
	fmt.Println("hello")

}

