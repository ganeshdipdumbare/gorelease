package gocodeplayer

import "backend/internal/gocode"

type CodePlayer interface {
	Play(code *gocode.GoCode) (*gocode.GoCodePlayOutput, error)
}
