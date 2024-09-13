package models

type Activity_Types string

const (
	Undefined Activity_SubTypes = "undefined"
	Breast    Activity_Types    = "breast"
	Bottle    Activity_Types    = "bottle"
	Pump      Activity_Types    = "pump"
	Nappy     Activity_Types    = "nappy"
	Sleep     Activity_Types    = "sleep"
	Solids    Activity_Types    = "solids"
	Tummy     Activity_Types    = "tummy"
	Custom    Activity_Types    = "custom"
)
