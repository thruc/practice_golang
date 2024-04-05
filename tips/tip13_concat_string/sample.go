package main

import "strings"

func badConcat(values []string) string {
	s := ""
	for _, v := range values {
		s += v
	}
	return s
}

func goodConcat(values []string) string {
	sb := strings.Builder{}
	for _, v := range values {
		sb.WriteString(v)
	}
	return sb.String()
}

func bestConcat(values []string) string {
	total := 0
	// ↓ 合計バイト数を求めるために各文字列を反復処理する
	for i := 0; i < len(values); i++ {
		total += len(values[i])
	}
	sb := strings.Builder{}
	sb.Grow(total) // ← total で Grow を呼び出す
	for _, value := range values {
		_, _ = sb.WriteString(value)
	}
	return sb.String()
}
