package myhttp

type StatusLine struct {
	HTTPVersion string
	StatusCode string
	ReasonPhrase string
}

func (sl StatusLine)ToString() string {
	return sl.HTTPVersion+" "+sl.StatusCode+" "+sl.ReasonPhrase

}


type Response struct {
	StatusLine StatusLine
	Header  Header
	Body string
}


func (r Response)ToString() string {
	return r.StatusLine.ToString()+"\r\n"+r.Header.ToString()+"\r\n"+r.Body
}

