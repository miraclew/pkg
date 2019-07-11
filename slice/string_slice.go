package types

type StringSlice []string

func (s StringSlice) Has(item string) bool {
	for _, v := range s {
		if v == item {
			return true
		}
	}

	return false
}

func (s StringSlice) Sub(items StringSlice) StringSlice {
	var res StringSlice
	for _, v := range s {
		if !items.Has(v) {
			res = append(res, v)
		}
	}

	return res
}

func (s StringSlice) Equals(s2 StringSlice) bool {
	if len(s) != len(s2) {
		return false
	}

	for i, v := range s {
		if v != s2[i] {
			return false
		}
	}
	return true
}
