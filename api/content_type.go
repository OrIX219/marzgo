package api

type ContentType int

const (
	ContentTypeNone ContentType = iota
	ContentTypeUrlEncoded
	ContentTypeJson
)

func (ct ContentType) String() string {
	return [...]string{
		"",
		"application/x-www-form-urlencoded",
		"application/json",
	}[ct]
}
