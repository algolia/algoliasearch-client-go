package gen

//go:generate go run gen_option.go        options.go functions.go
//go:generate go run gen_extract.go       options.go functions.go
//go:generate go run gen_settings.go      options.go functions.go
//go:generate go run gen_search_params.go options.go functions.go
//go:generate go run gen_iterator.go      options.go functions.go
