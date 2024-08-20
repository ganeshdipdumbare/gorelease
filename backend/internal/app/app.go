package app

import (
	"backend/internal/gocode"
	"backend/internal/gocodeplayer"
)

type CodePlayer interface {
	Play(code *gocode.GoCode) (*gocode.GoCodePlayOutput, error)
}

type CodePlayerI struct {
	codePlayer gocodeplayer.CodePlayer
}

func NewCodePlayer(codePlayer gocodeplayer.CodePlayer) *CodePlayerI {
	return &CodePlayerI{
		codePlayer: codePlayer,
	}
}

func (cp *CodePlayerI) Play(code *gocode.GoCode) (*gocode.GoCodePlayOutput, error) {
	return cp.codePlayer.Play(code)
}
