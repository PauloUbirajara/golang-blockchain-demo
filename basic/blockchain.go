package basic

type Blockchain struct {
	Blocks []Block
}

func (bc *Blockchain) NewBlock(content string, difficulty int) {
	DEFAULT_PREVIOUS_HASH := "secret"

	var lastBlock Block
	lastBlock.CurrentHash = DEFAULT_PREVIOUS_HASH

	if len(bc.Blocks) > 0 {
		lastBlock = bc.Blocks[len(bc.Blocks)-1]
	}

	newBlock := Block{
		Index:        lastBlock.Index + 1,
		Content:      content,
		PreviousHash: lastBlock.CurrentHash,
	}

	newBlock.SearchHash(difficulty)

	bc.Blocks = append(bc.Blocks, newBlock)
}

func (bc *Blockchain) PrintBlocks() {
	for _, b := range bc.Blocks {
		b.Print()
	}
}
