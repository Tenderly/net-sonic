package difftool

import (
	"testing"

	"github.com/andrecronje/lachesis/src/common"
)

func TestExample(t *testing.T) {
	logger := common.NewTestLogger(t)

	nodes := NewNodeList(3, logger)

	stop := nodes.StartRandTxStream()
	nodes.WaitForBlock(5)
	stop()

	diff_result := Compare(nodes.Values()...)

	if !diff_result.IsEmpty() {
		t.Fatal("\n" + diff_result.ToString())
	}
}
