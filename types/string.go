package types

import (
	"strconv"
	"strings"
)

type String struct {
	string
}
type StringArray []String

func NewString(str string) String {
	return String{str}
}

func NewStringFromInt(i int) String {
	return String{strconv.Itoa(i)}
}

func (str *String) String() string {
	return str.string
}

func (str *String) ToInt() int {
	i, err := strconv.Atoi(str.string)
	if err != nil {
		panic(err)
	}

	return i
}

func (str *String) ToLower() string {
	return strings.ToLower(str.string)
}

func (str *String) ToUpper() string {
	return strings.ToUpper(str.string)
}

func (str *String) Join(elems []string) string {
	return strings.Join(elems, str.string)
}

func (str *String) Contains(substr string) bool {
	return strings.Contains(str.string, substr)
}

func (str *String) Count(substr string) int {
	return strings.Count(str.string, substr)
}

func (str *String) HasPrefix(prefix string) bool {
	return strings.HasPrefix(str.string, prefix)
}

func (str *String) StartsWith(prefix string) bool {
	return strings.HasPrefix(str.string, prefix)
}

func (str *String) HasSuffix(suffix string) bool {
	return strings.HasSuffix(str.string, suffix)
}

func (str *String) EndsWith(suffix string) bool {
	return strings.HasSuffix(str.string, suffix)
}

func (str *String) Index(substr string) int {
	return strings.Index(str.string, substr)
}

func (str *String) LastIndex(substr string) int {
	return strings.LastIndex(str.string, substr)
}

func (str *String) Replace(old string, new string) string {
	return strings.Replace(str.string, old, new, -1)
}

func (str *String) ReplaceWithLimit(old string, new string, limit int) string {
	return strings.Replace(str.string, old, new, limit)
}

func (str *String) Repeat(count int) string {
	return strings.Repeat(str.string, count)
}

func (str *String) Split(sep string) []string {
	return strings.Split(str.string, sep)
}
