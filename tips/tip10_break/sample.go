package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", i)
		switch i {
		case 2:
			// breakしてもswitch文を抜けるだけで、for文は抜けない
			break
		default:
		}

	}

	fmt.Println("-----")
	// labelを使ってfor文を抜ける
Loop:
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", i)
		switch i {
		case 2:
			break Loop
		default:
		}
	}
}

// deferの呼び出しはループの反復処理の間ではなく、readFiles関数の終了時に行われる
func readFiles(ch <-chan string) error {
	for path := range ch {
		// ←チャネルを反復処理する
		file, err := os.Open(path) // ←ファイルをオープンする
		if err != nil {
			return err
		}
		defer file.Close() // ← file.Close()の呼び出しを遅延させる
		// ファイルで何らかの処理を行う
	}
	return nil
}
