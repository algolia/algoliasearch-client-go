package call

type Kind int

const (
	Read Kind = iota
	Write
	Analytics
)

func IsRead(k Kind) bool      { return k == Read }
func IsWrite(k Kind) bool     { return k == Write }
func IsAnalytics(k Kind) bool { return k == Analytics }
func IsReadWrite(k Kind) bool { return IsRead(k) || IsWrite(k) }
