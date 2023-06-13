package model

type SubjectType = uint8

const (
	SubjectTypeAll   SubjectType = 0
	SubjectTypeBook  SubjectType = 1
	SubjectTypeAnime SubjectType = 2
	SubjectTypeMusic SubjectType = 3
	SubjectTypeGame  SubjectType = 4
)

const (
	textSubjectBook  = "书籍"
	textSubjectAnime = "动画"
	textSubjectMusic = "音乐"
	textSubjectGame  = "游戏"
)

func SubjectTypeString(s uint8) string {
	switch s {
	case SubjectTypeBook:
		return textSubjectBook
	case SubjectTypeAnime:
		return textSubjectAnime
	case SubjectTypeMusic:
		return textSubjectMusic
	case SubjectTypeGame:
		return textSubjectGame
	default:
		return "unknown type"
	}
}
