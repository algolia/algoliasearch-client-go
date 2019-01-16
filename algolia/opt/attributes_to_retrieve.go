package opt

type attributesToRetrieve struct{ value []string }

func AttributesToRetrieve(v []string) attributesToRetrieve {
	return attributesToRetrieve{v}
}

func ExtractAttributesToRetrieve(opts ...interface{}) []string {
	var (
		uniq = make(map[string]bool)
		res  []string
	)

	for _, opt := range opts {
		v, ok := opt.(attributesToRetrieve)
		if ok {
			for _, attr := range v.value {
				uniq[attr] = true
			}
		}
	}

	for attr := range uniq {
		res = append(res, attr)
	}
	return res
}
