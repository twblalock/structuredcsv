package structuredcsv

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestEmpty(t *testing.T) {
	csv := ""

	_, err := NewReader(strings.NewReader(csv))
	if err != io.EOF {
		t.Error(err)
	}
}

func TestOnlyHeaders(t *testing.T) {
	csv := "A,B,C"

	reader, err := NewReader(strings.NewReader(csv))
	if err != nil {
		t.Error(err)
	}

	_, err = reader.Read()
	if err != io.EOF {
		t.Error(err)
	}
}

func TestSimple(t *testing.T) {
	csv := `A,B,C
1,2,3
4,5,6`

	reader, err := NewReader(strings.NewReader(csv))
	if err != nil {
		t.Error(err)
	}

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			t.Error(err)
		}

		for _, col := range row.Columns {
			fmt.Println("header:", col.Header, "value:", col.Value)
		}
	}
}

func TestForEach(t *testing.T) {
	csv := `A,B,C
1,2,3
4,5,6`

	reader, err := NewReader(strings.NewReader(csv))
	if err != nil {
		t.Errorf("error: %s", err)
	}

	err = reader.ForEach(func(row *Row) {
		for _, col := range row.Columns {
			fmt.Println("header:", col.Header, "value:", col.Value)
		}
	})

	if err != nil {
		t.Error(err)
	}
}

func TestGet(t *testing.T) {
	csv := `A,B,C
1,2,3
4,5,6`

	reader, err := NewReader(strings.NewReader(csv))
	if err != nil {
		t.Errorf("error: %s", err)
	}

	err = reader.ForEach(func(row *Row) {
		fmt.Println("A is", row.Get("A"))
	})

	if err != nil {
		t.Error(err)
	}
}

func TestSet(t *testing.T) {
	csv := `A,B,C
1,2,3
4,5,6`

	reader, err := NewReader(strings.NewReader(csv))
	if err != nil {
		t.Errorf("error: %s", err)
	}

	err = reader.ForEach(func(row *Row) {
		row.Set("A", "expected")
		if row.Get("A") != "expected" {
			t.Error("Failed to set column value")
		}
	})

	if err != nil {
		t.Error(err)
	}
}