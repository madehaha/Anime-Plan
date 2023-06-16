package model

type CollectionType = uint8

const (
	Wish CollectionType = iota
	Watched
	Doing
	OnHold
	Dropped
)

const (
	TextWish    = "想看"
	TextWatched = "看过"
	TextDoing   = "在看"
	TextOnHold  = "搁置"
	TextDropped = "抛弃"
)

func CollectionTypeString(s uint8) string {
	switch s {
	case Wish:
		return TextWish
	case OnHold:
		return TextOnHold
	case Doing:
		return TextDoing
	case Dropped:
		return TextDropped
	case Watched:
		return TextWatched
	default:
		return "unknown type"
	}
}
