package strings

import (
	"bytes"
	"unicode/utf8"
)

// 截取字符串 如果超过 指定长度 则截取不超过长度的字符串 并且以 指定的mark结束
// str 原始字符串
// length 字符串的长度
// mark 最后一个结束的字符
func SplitZhString(str string, length int, sliceString []string) []string {
	// 小于 指定长度 直接返回
	if  utf8.RuneCountInString(str) <= length {
		return append(sliceString, str)
	}

	// 转成中文 []rune
	strZh := []rune(str)
	// 获取指定长度
	split := string(strZh[:length])
	var markIndex int = length
	marks := []string{"。", "\n", "，", "," ,"、", "；", ":", "：", "?", "？", ";", "!","！"}
	for _, mark := range marks{
		middle := bytes.LastIndex([]byte(split), []byte(mark))
		if middle != -1 {
			markIndex = middle + len(mark)
		}
	}

	// 通过句号获取最后的结束
	newZhStr := []byte(split)[:markIndex]
	// 剩余自字符串
	oldZhStr := []byte(str)[markIndex:]
	newSliceString := append(sliceString, string(newZhStr))
	return SplitZhString(string(oldZhStr), length, newSliceString)
}

