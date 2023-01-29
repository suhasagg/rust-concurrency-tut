type arena struct {
	buf []byte
	off int
}

func newArena(size int) *arena {
	return &arena{
		buf: make([]byte, size),
	}
}

func (a *arena) allocate(size int) []byte {
	if a.off+size > len(a.buf) {
		return nil
	}
	b := a.buf[a.off : a.off+size]
	a.off += size
	return b
}

func (a *arena) reset() {
	a.off = 0
}

type ArenaPool struct {
	arenas []*arena
	size   int
}

func NewArenaPool(blockSize, maxBlocks int) *ArenaPool {
	return &ArenaPool{
		size: blockSize,
	}
}

func (p *ArenaPool) Allocate(size int) []byte {
	for _, a := range p.arenas {
		b := a.allocate(size)
		if b != nil {
			return b
		}
	}
	a := newArena(p.size)
	p.arenas = append(p.arenas, a)
	return a.allocate(size)
}

func (p *ArenaPool) Reset() {
	for _, a := range p.arenas {
		a.reset()
	}
}
