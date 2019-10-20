package typedefs

type HttpResponse struct {
	Url      string
	Response []byte
	Err      error
}
