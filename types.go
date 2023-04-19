package witherrors

type Priority int

const (
	// constant for enum Priority types
	VERY_LOW Priority = iota
	LOW
	MEDIUM
	HIGH
	VERY_HIGH

	// constant in string type for priority types
	veryLow  = "VERY_LOW"
	low      = "LOW"
	medium   = "MEDIUM"
	high     = "HIGH"
	veryHigh = "VERY_HIGH"

	// constant for default err
	emptyCodeStr          = "code is empty"
	emptyDependencyStr    = "dependency is empty"
	emptyErrStr           = "err message is empty"
	emptyCustomMessageStr = "custom message is empty"
	emptyPriorityStr      = "priority is empty"
)
