package game

import (
    "math/rand"
)


func RandomNumber(min, max int) int {
    return rand.Intn(max - min) + min
}