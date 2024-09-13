package models

type Activity struct {
	Type           Activity_Types
	Start_DateTime string
	End_DateTime   string
	SubType        Activity_SubTypes
}

func (activity Activity) getSubTypes() []Activity_SubTypes {
	var activity_sub_type_map = map[Activity_Types][]Activity_SubTypes{
		Breast: {Undefined},
		Bottle: {Formula, Breast_Milk, Goats_Milk, Cows_Milk},
	}

	return activity_sub_type_map[activity.Type]
}
