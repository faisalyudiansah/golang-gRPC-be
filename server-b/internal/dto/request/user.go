package dtorequest

type RequestGetUserByID struct {
	ID int64 `json:"id" uri:"id" binding:"required"`
}
