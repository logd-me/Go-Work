/*
@Title : randomID
@Description : 产生随机账号
@Author : 谭靖渝
@Update : 2021/3/29 11:34
*/
package Model

import (
	"math"
	"math/rand"
	"strconv"
	"time"
)

func Random(width int) string {
	var rands =0
	for i := 0; i < width; i++ {
		r := rand.New(rand.NewSource(time.Now().Add(time.Second * time.Duration(i)).UnixNano()))
		n := r.Intn(10)
		rands = rands+n*int(math.Pow10(width-i-1))
	}
	str := strconv.Itoa(rands)
	return str
}
