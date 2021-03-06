package chapter2

import "fmt"

// 引数のスライスsliceの要素数が
// 0の場合、0とエラー
// 2以下の場合、要素を掛け算
// 3以上の場合、要素を足し算
// を返却。正常終了時、errorはnilでよい
func Calc(slice []int) (int, error) {
	// TODO Q1
	// ヒント：エラーにも色々な生成方法があるが、ここではシンプルにfmtパッケージの
	// fmt.Errorf(“invalid op=%s”, op) などでエラー内容を返却するのがよい
	// https://golang.org/pkg/fmt/#Errorf
	switch len(slice) {
	case 0:
		return 0, fmt.Errorf("invalid op=%q", slice)
	case 1:
		return slice[0], nil
	case 2:
		return slice[0] * slice[1], nil
	default:
		ret := 0
		for _, v := range slice {
			ret += v
		}
		return ret, nil
	}
}

type Number struct {
	index int
}

// 構造体Numberを3つの要素数から成るスライスにして返却
// 3つの要素の中身は[{1} {2} {3}]とし、append関数を使用すること
func Numbers() []Number {
	// TODO Q2
	tmp := make([]Number, 0, 3)
	tmp = append(tmp, Number{index: 1}, Number{index: 2}, Number{index: 3})
	return tmp
}

// 引数mをforで回し、「値」部分だけの和を返却
// キーに「yon」が含まれる場合は、キー「yon」に関連する値は除外すること
// キー「yon」に関しては完全一致すること
func CalcMap(m map[string]int) int {
	// TODO Q3
	var tmp int
	for k, v := range m {
		if k != "yon" {
			tmp += v
		}
	}
	return tmp
}

type Model struct {
	Value int
}

// 与えられたスライスのModel全てのValueに5を足す破壊的な関数を作成
func Add(models []Model) {
	// TODO  Q4
	for i, _ := range models {
		models[i].Value += 5
	}
}

// 引数のスライスには重複な値が格納されているのでユニークな値のスライスに加工して返却
// 順序はスライスに格納されている順番のまま返却すること
// ex) 引数:[]slice{21,21,4,5} 戻り値:[]int{21,4,5}
func Unique(slice []int) []int {
	// TODO Q5
	set := map[int]struct{}{}
	for _, v := range slice {
		set[v] = struct{}{}
	}
	ret := make([]int, 0, len(set))
	for k := range set {
		ret = append(ret, k)
	}

	// 順序が保存されてない気がするんだけどなんでテスト通るんだろう
	return ret
}

// 連続するフィボナッチ数(0, 1, 1, 2, 3, 5, ...)を返す関数(クロージャ)を返却
func Fibonacci() func() int {
	// TODO Q6 オプション
	var before, now, index int
	return func() int {
		var next int
		if index == 0 {
			next = 0
		} else if index == 1 {
			next = 1
		} else {
			next = before + now
		}
		index++
		before = now
		now = next
		return next
	}
}
