package opt

import "fmt"

func main() {
	for _, optName := range []string{
		"AdvancedSyntax",
	} {
		generateExtractFunction(optName)
	}
}

func generateExtractFunction(opt string) {
	fnStr := fmt.Sprintf(`
		func Extract%s(opts ...interface{}) *opt.%sOption {
			for _, o := range opts {
				if v, ok := o.(opt.%sOption); ok {
					return &v
				}
			}
			return nil
		}`, opt, opt, opt)
	_ = fnStr
}
