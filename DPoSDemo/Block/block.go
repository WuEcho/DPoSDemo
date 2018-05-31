package Block

import (
	"time"
	"strconv"
	"crypto/sha256"
	"encoding/hex"
)

type Block struct {
	BMP int

	PrefHash string

	HashCode string

	TimeStamp string

	Index int

	//区块验证者
	Validator string
}

func GenerateNextBlock(bmp int,oldBlock Block,validator string) Block {

	var newBlock Block
	newBlock.PrefHash = oldBlock.HashCode
	newBlock.TimeStamp = time.Now().String()
	newBlock.Index = oldBlock.Index + 1
	newBlock.Validator = validator
	newBlock.BMP = bmp
	newBlock.HashCode = SetHash(newBlock)
    return newBlock
}

func SetHash(b Block) string {

    hashCode := []byte(b.Validator+strconv.Itoa(b.Index)+b.TimeStamp+b.PrefHash+strconv.Itoa(b.BMP))

    sha := sha256.New()

    sha.Write(hashCode)

    hash := sha.Sum(nil)

    //fmt.Print(hex.EncodeToString(hash))

    return hex.EncodeToString(hash)
}


