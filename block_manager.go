package main

type BlockManager struct {
	vm *Vm
}

func NewBlockManager() *BlockManager {
	bm := &BlockManager{vm: NewVm()}

	return bm
}

// Process a block.
func (bm *BlockManager) ProcessBlock(block *Block) error {
	txCount := len(block.transactions)
	lockChan := make(chan bool, txCount)

	for _, tx := range block.transactions {
		go bm.ProcessTransaction(tx, lockChan)
	}

	// Wait for all Tx to finish processing
	for i := 0; i < txCount; i++ {
		<- lockChan
	}

	return nil
}

func (bm *BlockManager) ProcessTransaction(tx *Transaction, lockChan chan bool) {
	if tx.recipient == 0x0 {
		bm.vm.RunTransaction(tx, func(opType OpType) bool {
			// TODO calculate fees

			return true // Continue
		})
	}

	// Broadcast we're done
	lockChan <- true
}