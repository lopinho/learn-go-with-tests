package main

const REPEAT_COUNT = 5

func Repeat(str string) string {
	retorno := ""
	for i := 0; i < REPEAT_COUNT; i++ {
		retorno += str
	}
	return retorno
}
