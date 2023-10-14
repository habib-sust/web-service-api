package constant

type ResponseStatus int
type Headers int
type General int

//Constatn Api

const (
	Success ResponseStatus = iota + 1
	DataNotFound
	UnknownError
	InvalidRequest
	Unauthorized
)

func (r ResponseStatus) GetResponseStatue() string {
	return [...]string{
		"SUCCESS",
		"DATA_NOT-FOUND",
		"UNKNOWN_ERROR",
		"INVALID_REQUEST",
		"UNAUTHORIZED",
	}[r-1]
}

func (r ResponseStatus) GetResponseMessage() string {
	return [...]string{
		"Success",
		"Data Not Found",
		"Unknown Error",
		"Invalid Request",
		"Unauthorized",
	}[r-1]
}
