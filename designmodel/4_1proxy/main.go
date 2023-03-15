package main

import "fmt"

// 抽象主题
type BeautyWoman interface {
	MakeEyesWithMan()
	HappyWithMan()
}

// 具体主题
type PanJinLian struct{}

func (p *PanJinLian) MakeEyesWithMan() {
	fmt.Println("Pan make eyes with someone")
}

func (p *PanJinLian) HappyWithMan() {
	fmt.Println("Pan happy with someone")
}

// 代理中介
type WangPo struct {
	wowan BeautyWoman
}

func NewProxy(w BeautyWoman) BeautyWoman {
	return &WangPo{wowan: w}
}

func (p *WangPo) MakeEyesWithMan() {
	p.wowan.MakeEyesWithMan()
}

func (p *WangPo) HappyWithMan() {
	p.wowan.HappyWithMan()
}

func main() {
	wangpo := NewProxy(new(PanJinLian))
	wangpo.MakeEyesWithMan()
	wangpo.HappyWithMan()
}
