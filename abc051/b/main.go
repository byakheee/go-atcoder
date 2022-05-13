package main

import (
	"bufio"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriterSize(os.Stdout, 1024*1024)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)

	sc.Scan()
	k, _ := strconv.Atoi(sc.Text())

	sc.Scan()
	s, _ := strconv.Atoi(sc.Text())

	// 0 <= X,Y,Z <= K を満たす整数の値
	// X+Y+Z=S を満たす組み合わせの数
	n := 0
	for x := 0; x <= k; x++ {
		for y := 0; y <= k; y++ {
			z := s - (x + y)
			if 0 <= z && z <= k {
				n++
			}
		}
	}

	wr.WriteString(strconv.Itoa(n) + "\n")
}
