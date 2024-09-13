package models

type Activity_SubTypes string

const (
	Left        Activity_SubTypes = "left_breast"
	Right       Activity_SubTypes = "right_breast"
	Formula     Activity_SubTypes = "formula"
	Breast_Milk Activity_SubTypes = "breast_milk"
	Goats_Milk  Activity_SubTypes = "goats_milk"
	Cows_Milk   Activity_SubTypes = "cows_milk"
)
