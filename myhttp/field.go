package myhttp

import "regexp"

type Field struct {
	Name string
	Value string
}

type ContentType  Field

func (c  ContentType ) GetBoundary() string {
	return regexp.MustCompile(`boundary=(.*)$`).FindStringSubmatch(c.Value)[1]
}

func (c  ContentType) GetMediaType() string {
	r1 := regexp.MustCompile(`(.+/.+);`)
	r2 := regexp.MustCompile(`(.+/.+)`)

	m1 := r1.FindStringSubmatch(c.Value)
	m2 := r2.FindStringSubmatch(c.Value)

	if len(m1)>0 {
		return m1[1]
	}

	if len(m2) >0 {
		return m2[1]
	}

	return ""

}

