package rtcsrtpkeyparam

type RTCSrtpKeyParam struct {
	keyMethod *string
	keySalt   *string
	lifetime  *string
	mkiLength *uint8
	mkiValue  *uint8
}
