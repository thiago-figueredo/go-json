package test

import (
	"math/rand"
	"time"
)

func RandomElementFromArray(arr []any) any {
	rand.New(rand.NewSource(time.Now().UnixNano()))
  randomIndex := rand.Intn(len(arr))

  return arr[randomIndex]
}
