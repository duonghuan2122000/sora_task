package basemodel

type LogicError struct {
	// Mã lỗi
	Code string
	// mô tả mã lỗi
	Message string
}

func (e LogicError) Error() string {
	return "Có lỗi logic: " + e.Message
}
