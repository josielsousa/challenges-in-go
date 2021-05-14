package main

import (
	"encoding/json"
	"fmt"
)

type CommonCondition struct {
	Type               string `json:"__type"`
	LogicalOperator    string `json:"LogicalOperator"`
	ComparisonOperator string `json:"ComparisonOperator"`
	Field              string `json:"Field"`
	Value              string `json:"Value"`
}

type GroupedConditionList struct {
	CommonCondition
}

type Condition struct {
	GroupedConditionList []GroupedConditionList `json:"GroupedConditionList,omitempty"`
	CommonCondition
}

func main() {
	// Example Initialize group outside variables
	//groupCondition := GroupedConditionList {}
	//groupCondition.CommonCondition.Type = "type teste"
	//groupCondition.CommonCondition.LogicalOperator = "+"
	//groupCondition.CommonCondition.ComparisonOperator = `>=`
	//groupCondition.CommonCondition.Field = "amount"
	//groupCondition.CommonCondition.Value = "33"
	//
	//texto, _ := json.Marshal(groupCondition)
	//fmt.Printf("Group %+v", string(texto))

	// Example Initialize group from JSON variables
	//strTxt := `{"__type":"type teste","LogicalOperator":"+","ComparisonOperator":"\u003e=","Field":"amount","Value":"33"}`
	//var groupCondition GroupedConditionList
	//json.Unmarshal([]byte(strTxt), &groupCondition)

	// Example Initialize condition inside composition
	// Use nested composite type to initialize a value in a single expression:
	groupCondition := GroupedConditionList{
		CommonCondition: CommonCondition{
			Type:               "type teste group",
			LogicalOperator:    "+",
			ComparisonOperator: `>=`,
			Field:              "amount",
			Value:              "33",
		},
	}

	condition := Condition{
		CommonCondition: CommonCondition{
			Type:               "type teste condition",
			LogicalOperator:    "+",
			ComparisonOperator: `>=`,
			Field:              "amount",
			Value:              "33",
		},
		GroupedConditionList: []GroupedConditionList{groupCondition},
	}

	fmt.Printf("Group %+v\n", condition)
	fmt.Printf("Group %+v\n", condition.Type)
	fmt.Printf("Group %+v\n", condition.GroupedConditionList[0].Type)

	texto, _ := json.Marshal(condition)
	fmt.Printf("JSON %+v", string(texto))
}
