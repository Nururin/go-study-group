package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

// go-cutコマンドを実装しよう
func main() {

	d := flag.String("d", "\t", "delimiter")

	// TODO: 範囲指定
	// TODO: カンマ区切り指定
	f := flag.Int("f", 0, "field")

	flag.Parse()

	file, err := os.Open(flag.Arg(0))
	if err != nil {
		fmt.Fprintln(os.Stderr, "ファイルを開けません。")
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}
	defer file.Close()

	s := bufio.NewScanner(file)

	if *f <= 0 {
		fmt.Fprintln(os.Stderr, "fieldには1以上の値を設定してください。")
		return
	}

	i := *f - 1
	for s.Scan() {
		stxt := s.Text()
		sl := strings.Split(stxt, *d)
		if len(sl) <= i {
			fmt.Fprintln(os.Stdout, "")
		} else {
			fmt.Fprintln(os.Stdout, sl[i])
		}
	}

	if s.Err() != nil {
		fmt.Fprintln(os.Stderr, "ファイル読み込み中にエラーが発生しました。")
		fmt.Fprintf(os.Stderr, "%+v\n", s.Err())
		return
	}
}
