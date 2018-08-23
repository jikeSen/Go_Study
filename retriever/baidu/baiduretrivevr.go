package baidu

type Retrivevr struct {
	Contents string
}

func (r Retrivevr) Get(url string) string {
	return r.Contents
}


