package selector

type strsArr [][]string

func (a strsArr) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a strsArr) Less(i, j int) bool {
	return a[i][0] < a[j][0]
}

func (a strsArr) Len() int {
	return len(a)
}

type strs []string

func (a strs) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a strs) Less(i, j int) bool {
	return a[i] > a[j]
}

func (a strs) Len() int {
	return len(a)
}

type itemArr []item

func (a itemArr) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a itemArr) Less(i, j int) bool {
	return a[i].line < a[j].line
}

func (a itemArr) Len() int {
	return len(a)
}
