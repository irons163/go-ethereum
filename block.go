package main

type Block struct {
	transactions []*Transaction
}

func NewBlock(/* TODO use raw data */transactions []*Transaction) *Block {
	block := &Block {
		transactions: transactions,
	}

	return block
}

func (block *Block) Update() {

}