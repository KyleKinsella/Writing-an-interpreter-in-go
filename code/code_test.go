package code

import "testing"

func TestMake(t *testing.T) {
	tests := []struct {
		op Opcode
		operands []int
		expected []byte
	}{
		{OpConstant, []int{65534}, []byte{byte(OpConstant), 255, 254}},
		{OpAdd, []int{}, []byte{byte(OpAdd)}},
	}
	
	for _, tt := range tests {
		instruction := Make(tt.op, tt.operands...)
		
		if len(instruction) != len(tt.expected) {
			t.Errorf("instruction has wrong lenght. want=%d, got=%d", len(tt.expected), len(instruction))
		}
		
		for i, b := range tt.expected {
			if instruction[i] != tt.expected[i] {
				t.Errorf("wrong byte at pos %d. want=%d, got=%d", i, b, instruction[i])
			}
		}
	}
} 

func TestInstructionsString(t *testing.T) {
	instructions := []Instructions{
		//Make(OpConstant, 1),
		
		Make(OpAdd),
		Make(OpConstant, 2),
		Make(OpConstant, 65535),
	}
	
	// This here was the problem (page 52 in the book) - my test was not passing and i was not too sure why
	// but the problem was I had line 42 and 43 tabed in by 1, this was then treated as a tab (\t) and that is 
	// why my test failed. P.S. I tried to use ChatGpt for some help / guidance for this problem but it was 
	// unable to help me fix the problem, I fixed this problem on my own!
	
	expected := `0000 OpAdd
0001 OpConstant 2
0004 OpConstant 65535
`
	
	concatted := Instructions{}
	
	for _, ins := range instructions {
		concatted = append(concatted, ins...)
	}
	
	if concatted.String() != expected {				
		t.Errorf("instructions wrongly formatted.\nwant=%q\ngot=%q", expected, concatted.String())
	}
}

func TestReadOperands(t *testing.T) {
	tests := []struct{
		op Opcode
		operands []int
		bytesRead int
	} {
		{OpConstant, []int{65535}, 2},
	}
	
	for _, tt := range tests {
		instruction := Make(tt.op, tt.operands...)
		
		def, err := Lookup(byte(tt.op))
		if err != nil {
			t.Fatalf("definition not found: %q\n", err)
		}
		
		operandsRead, n := ReadOperands(def, instruction[1:])
		if n != tt.bytesRead {
			t.Fatalf("n wrong. want=%d, got=%d", tt.bytesRead, n)
		}
		
		for i, want := range tt.operands {
			if operandsRead[i] != want {
				t.Errorf("operand wrong. want=%d, got=%d", want, operandsRead[i])
			}
		}
	}
}
