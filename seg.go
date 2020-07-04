package uulang

import (
	"github.com/huichen/sego"
	"regexp"
	"strconv"
)

var segmenter sego.Segmenter

func init() {
	//segmenter.LoadDictionary("./data/dictionary.txt")
}

type SegWord struct {
	Id      int
	Word    string   `json:"word"`
	Pos     string   `json:"pos"`
	Pinyin  []string `json:"pinyin"`
	PinyinN []string `json:"pinyin_n"`
	PinyinT []int    `json:"pinyin_t"`
}

type SegWords []*SegWord

func LoadSegDictionary(file string) {
	segmenter.LoadDictionary(file)
}

func GetSegmenter() *sego.Segmenter {
	return &segmenter
}

func Cut(text string) SegWords {
	btext := []byte(text)
	segs := segmenter.Segment(btext)
	re := regexp.MustCompile(`(\d)+`)

	var segWords SegWords
	for k, seg := range segs {
		_text := seg.Token().Text()
		_pos := seg.Token().Pos()
		_py := LazyPinyinTone(_text)
		_pyn := LazyPinyinTone3(_text)
		_pyt := []int{}

		for _, _tmp := range _pyn {
			m := re.FindString(_tmp)
			i, _ := strconv.Atoi(m)
			_pyt = append(_pyt, i)
		}

		segWords = append(segWords, &SegWord{
			Id:      k,
			Word:    _text,
			Pos:     _pos,
			Pinyin:  _py,
			PinyinN: _pyn,
			PinyinT: _pyt,
		})
	}

	return segWords
}

func Cut4Search(text string) SegWords {
	btext := []byte(text)
	segs := segmenter.Segment(btext)
	re := regexp.MustCompile(`(\d)+`)

	words := sego.SegmentsToSlice(segs, true)

	var segWords SegWords
	for k, w := range words {
		_text := w
		_py := LazyPinyinTone(_text)
		_pyn := LazyPinyinTone3(_text)
		_pyt := []int{}

		for _, _tmp := range _pyn {
			m := re.FindString(_tmp)
			i, _ := strconv.Atoi(m)
			_pyt = append(_pyt, i)
		}

		segWords = append(segWords, &SegWord{
			Id:      k,
			Word:    _text,
			Pinyin:  _py,
			PinyinN: _pyn,
			PinyinT: _pyt,
		})
	}

	return segWords
}

func CutWithPinyin(text string) SegWords {
	btext := []byte(text)
	segs := segmenter.Segment(btext)
	re := regexp.MustCompile(`(\d)+`)

	var segWords SegWords
	{
		for k, seg := range segs {
			_text := seg.Token().Text()
			_pos := seg.Token().Pos()
			//_py := strings.Join(LazyPinyinTone(_text), " ")
			//_pyn := strings.Join(LazyPinyinTone2(_text), " ")
			_py := LazyPinyinTone(_text)
			_pyn := LazyPinyinTone3(_text)
			_pyt := []int{}

			// set pos = "unknow" if pinyin != ""
			if _pos == "x" && len(_py) > 0 {
				_pos = "unknown"
			}

			for _, _tmp := range _pyn {
				m := re.FindString(_tmp)
				i, _ := strconv.Atoi(m)
				_pyt = append(_pyt, i)
			}

			segWords = append(segWords, &SegWord{
				Id:      k,
				Word:    _text,
				Pos:     _pos,
				Pinyin:  _py,
				PinyinN: _pyn,
				PinyinT: _pyt,
			})
		}
	}

	return segWords
}
