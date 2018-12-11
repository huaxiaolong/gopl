package weightconv

import "fmt"

type Kg float64
type Pound float64

func (k Kg) String() string {
	return fmt.Sprintf("%gkg", k)
}

func (p Pound) String() string {
	return fmt.Sprintf("%glb", p)
}

func KToP(k Kg) Pound {
	return Pound(k * 0.45359)
}

func PToK(p Pound) Kg {
	return Kg(p / 0.45359)
}
