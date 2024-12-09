package code

import "testing"

func TestMake(t *testing.T) {
	tests := []struct {
		op       Opcode
		operands []int
		expected []byte
	}{
		{OpConstant, []int{65534}, []byte{byte(OpConstant), 255, 254}},
	}

	for _, tt := range tests {
		instructions := Make(tt.op, tt.operands...)
		if len(instructions) != len(tt.expected) {
			t.Errorf("instructions length wrong. want=%d, got=%d", len(tt.expected), len(instructions))
		}
		for i, b := range tt.expected {
			if instructions[i] != b {
				t.Errorf("instruction wrong at %d. want=%d, got=%d", i, b, instructions[i])
			}
		}
	}
}
