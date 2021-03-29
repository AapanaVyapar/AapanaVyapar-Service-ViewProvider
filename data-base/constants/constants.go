package constants

type Status uint8

const (
	PENDING Status = iota
	CONFORM
	DISPATCHED
	DELIVERED
)

type Ratings uint8

const (
	VERY_BAD Ratings = iota
	BAD
	OK
	GOOD
	VERY_GOOD
)

type Categories uint8

const (
	SSPORTS_AND_FITNESS Categories = iota
	ELECTRIC
	DEVOTIONAL
	AGRICULTURAL
	WONENS_CLOTHING
	WONENS_ACCSSORIES
	MENS_CLOTHING
	MENS_ACCSSORIES
	HOUSE_HOME_GADGETS
	TOYS
	ELECTRONIC
)
