package algoliasearch

type PartialUpdateOp struct {
	Op    string      `json:"_operation"`
	Value interface{} `json:"value"`
}

func IncrementOp(value int) PartialUpdateOp {
	return PartialUpdateOp{
		Op:    "Increment",
		Value: value,
	}
}

func DecrementOp(value int) PartialUpdateOp {
	return PartialUpdateOp{
		Op:    "Decrement",
		Value: value,
	}
}

func AddOp(value interface{}) PartialUpdateOp {
	return PartialUpdateOp{
		Op:    "Add",
		Value: value,
	}
}

func RemoveOp(value interface{}) PartialUpdateOp {
	return PartialUpdateOp{
		Op:    "Remove",
		Value: value,
	}
}

func AddUniqueOp(value interface{}) PartialUpdateOp {
	return PartialUpdateOp{
		Op:    "AddUnique",
		Value: value,
	}
}
