package uulang

import (
	"github.com/huichen/sego"
	"strings"
)

var segmenter sego.Segmenter

func init() {
	//segmenter.LoadDictionary("./data/dictionary.txt")
}

type SegWord struct {
	Id      int
	Word    string `json:"word"`
	Pos     string `json:"pos"`
	Pinyin  string `json:"pinyin"`
	PinyinN string `json:"pinyin_n"`
}

type SegWords []*SegWord

func Cut(text string) *SegWords {
	btext := []byte(text)
	segs := segmenter.Segment(btext)

	var segWords SegWords
	for k, seg := range segs {
		_text := seg.Token().Text()
		_pos := seg.Token().Pos()

		segWords = append(segWords, &SegWord{
			Id:   k,
			Word: _text,
			Pos:  _pos,
		})
	}

	return &segWords
}

func CutWithPinyin(text string) *SegWords {
	btext := []byte(text)
	segs := segmenter.Segment(btext)

	var segWords SegWords
	{
		for k, seg := range segs {
			_text := seg.Token().Text()
			_pos := seg.Token().Pos()
			_py := strings.Join(LazyPinyinTone(_text), " ")
			_pyn := strings.Join(LazyPinyinTone2(_text), " ")

			segWords = append(segWords, &SegWord{
				Id:      k,
				Word:    _text,
				Pos:     _pos,
				Pinyin:  _py,
				PinyinN: _pyn,
			})
		}
	}

	return &segWords
}
