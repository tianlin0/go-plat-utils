package crypto

type Encoder func(origin []byte) string
type Decoder func(originString string) ([]byte, error)
