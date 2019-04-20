package uulang

import (
	"github.com/mozillazg/go-pinyin"
)

type (
	pySlice [][]string
	lzSlice []string
)

func Pinyin(hans string) pySlice {
	return PinyinTone(hans)
}

func PinyinTone(hans string) pySlice {
	a := pinyin.NewArgs()
	a.Style = pinyin.Tone
	a.Heteronym = true
	return pinyin.Pinyin(hans, a)
}

func PinyinTone2(hans string) pySlice {
	a := pinyin.NewArgs()
	a.Style = pinyin.Tone2
	a.Heteronym = true
	return pinyin.Pinyin(hans, a)
}

func PinyinTone3(hans string) pySlice {
	a := pinyin.NewArgs()
	a.Style = pinyin.Tone3
	a.Heteronym = true
	return pinyin.Pinyin(hans, a)
}

func LazyPinyinTone(hans string) lzSlice {
	a := pinyin.NewArgs()
	a.Style = pinyin.Tone
	return pinyin.LazyPinyin(hans, a)
}

func LazyPinyinTone2(hans string) lzSlice {
	a := pinyin.NewArgs()
	a.Style = pinyin.Tone2
	return pinyin.LazyPinyin(hans, a)
}

func LazyPinyinTone3(hans string) lzSlice {
	a := pinyin.NewArgs()
	a.Style = pinyin.Tone3
	return pinyin.LazyPinyin(hans, a)
}
