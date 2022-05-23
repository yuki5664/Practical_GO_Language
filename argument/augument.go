package argument

import "time"

type Portion int

const (
	Reguller Portion = iota
	Small
	Large
)

type Udon struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

// インスタンス作成
func NewUdon(p Portion, aburaage bool, ebiten uint) *Udon {
	return &Udon{
		men:      p,
		aburaage: aburaage,
		ebiten:   ebiten,
	}
}

var tempuraUdon = NewUdon(Large, false, 2)

// 別名の関数によるオプション引数
func NewKakeUdon(p Portion) *Udon {
	return &Udon{
		men:      p,
		aburaage: false,
		ebiten:   0,
	}
}

func NewKitusneUdon(p Portion) *Udon {
	return &Udon{
		men:      p,
		aburaage: true,
		ebiten:   0,
	}
}

func NewTempuraUdon(p Portion) *Udon {
	return &Udon{
		men:      p,
		aburaage: false,
		ebiten:   3,
	}
}

// 構造体を利用したオプション引数
type Option struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

func NewUdon(opt Option) *Udon {
	if opt.ebiten == 0 && time.Now().Hour() < 10 {
		opt.ebiten = 1
	}
	return &Udon{
		men:      opt.men,
		aburaage: opt.aburaage,
		ebiten:   opt.ebiten,
	}
}
