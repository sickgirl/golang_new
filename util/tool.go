package util

/**
工具类 放默认函数
*/
import (
	"math/rand"
	"time"
)

func RandomString(n int) string {
	var letters = []byte("qwertyuiopasdfghjklzxcvbnnmQWERTYUIOPASDFGHJKLZXCVBNNM")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		letters[i] = letters[rand.Intn(len(letters))]
	}
	return string(letters)
}
