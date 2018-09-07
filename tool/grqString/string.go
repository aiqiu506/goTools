package grqString

import "errors"

/**
 字符串取子串
@params matherString 被取串
@params start 开始位置  为负数时，表示从右往左数位置
@params lens 长度 为0时，表示截取到末尾

*/
func SubStr(matherString string, start, lens int) (string, error) {
	matherStr := []rune(matherString)
	strLen := len(matherStr)
	//位置错误
	if start >= strLen || start < (-strLen) {
		return "", errors.New("start position error")
	}
	if lens < 0 {
		return "", errors.New("length error")
	}
	//开始位置为负数，从后往前取
	if start < 0 {
		//不是所指定位置到结束
		if -start < lens {
			return "", errors.New("start position or length error")
		}
		if lens > 0 {
			return string(matherStr[strLen+start : strLen+start+lens]), nil
		}
		if lens == 0 {
			return string(matherStr[strLen+start:]), nil
		}

		return "", errors.New("start position or length error")
	}
	//从指定位置到结束
	if lens == 0 {
		return string(matherStr[start:]), nil
	}
	return string(matherStr[start : lens+start]), nil
}
