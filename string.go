package cool

import (
	"strconv"
	"strings"
)

type String string
type StringArray []String

func NewString(str string) String {
	return String(str)
}

func NewStringFromInt(i int) String {
	return String(strconv.Itoa(i))
}

func (str *String) String() string {
	return string(*str)
}

func (str *String) ToInt() int {
	i, err := strconv.Atoi(string(*str))
	if err != nil {
		panic(err)
	}

	return i
}

func (str *String) ToLower() string {
	return strings.ToLower(string(*str))
}

func (str *String) ToUpper() string {
	return strings.ToUpper(string(*str))
}

func (str *String) Join(elems []string) string {
	return strings.Join(elems, string(*str))
}

func (str *String) Contains(substr string) bool {
	return strings.Contains(string(*str), substr)
}

func (str *String) Count(substr string) int {
	return strings.Count(string(*str), substr)
}

func (str *String) HasPrefix(prefix string) bool {
	return strings.HasPrefix(string(*str), prefix)
}

func (str *String) StartsWith(prefix string) bool {
	return strings.HasPrefix(string(*str), prefix)
}

func (str *String) HasSuffix(suffix string) bool {
	return strings.HasSuffix(string(*str), suffix)
}

func (str *String) EndsWith(suffix string) bool {
	return strings.HasSuffix(string(*str), suffix)
}

func (str *String) Index(substr string) int {
	return strings.Index(string(*str), substr)
}

func (str *String) LastIndex(substr string) int {
	return strings.LastIndex(string(*str), substr)
}

func (str *String) Replace(old string, new string) string {
	return strings.Replace(string(*str), old, new, -1)
}

func (str *String) Repeat(count int) string {
	return strings.Repeat(string(*str), count)
}

func (str *String) Split(sep string) []string {
	return strings.Split(string(*str), sep)
}
