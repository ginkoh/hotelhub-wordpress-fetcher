package typedefs

type HttpResponse struct {
	Url      string
	Response []byte
	err      error
}
