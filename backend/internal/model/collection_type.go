package model

type CollectionType = uint8

const (
	Wish    CollectionType = 1
	OnHold  CollectionType = 2
	Doing   CollectionType = 3
	Drop    CollectionType = 4
	Watched CollectionType = 5
)

const (
	TextWish    = "想看"
	TextOnHold  = "搁置"
	TextDoing   = "在看"
	TextDrop    = "抛弃"
	TextWatched = "看过"
)

func CollectionTypeString(s uint8) string {
	switch s {
	case Wish:
		return TextWish
	case OnHold:
		return TextOnHold
	case Doing:
		return TextDoing
	case Drop:
		return TextDrop
	case Watched:
		return TextWatched
	default:
		return "unknown type"
	}
}
