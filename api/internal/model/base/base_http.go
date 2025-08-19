package basemodel

// BaseRequest strcut
type BaseRequest[TDto any] struct {
	Data BaseRequestData[TDto] `json:"data" binding:"required"`
}

type BaseRequestData[TDto any] struct {
	Attributes TDto `json:"attributes" binding:"required"`
}

// BaseResponse struct
type BaseResponse struct {
	// Cờ đánh dấu request có thành công hay không?
	Status bool `json:"status"`
	// Thông tin lỗi
	Error BaseResponseError `json:"error,omitzero"`
	// Dữ liệu khi thành công
	Data any `json:"data,omitempty"`
}

type BaseResponseError struct {
	// Mã lỗi
	Code string `json:"code,omitempty"`
	// Thống báo
	Message string `json:"message,omitempty"`
}

const CookieAccessToken = "x-sora-access-token"
