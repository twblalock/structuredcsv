package structuredcsv

import (
	"testing"
	"strings"
	"io"
	"fmt"
)

func TestSimple(t *testing.T) {
	var csv = `A,B,C
1,2,3
4,5,6`

	reader, err := NewReader(strings.NewReader(csv))
	if err != nil {
		t.Errorf("error: %s", err)
	}
	
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			t.Errorf("error: %s", err)
		}

		// TODO figure out how to do assertions and throw errors without Errorf
		
		for _, col := range row.Columns {
			fmt.Println("header:", col.Header, "value:", col.Value)
		}
	}
  
}
