package structuredcsv

import (
	"encoding/csv"
	"io"
)

type Row struct {
	Columns []*Column
}

type Column struct {
	Header, Value string
}

type StructuredReader struct {
	reader  *csv.Reader
	headers []string
}

func (r StructuredReader) Read() (Row, error) {
	var row Row

	values, err := r.reader.Read()
	if err != nil {
		return row, err
	}

	cols := make([]*Column, len(values))
	for i, v := range values {
		var header string
		if i < len(r.headers) {
			header = r.headers[i]
		}
		cols[i] = &Column{Header: header, Value: v}
	}

	return Row{Columns: cols}, nil
}

func (r StructuredReader) ReadAll() ([]Row, error) {
	records := make([]Row, 0)

	for {
		record, err := r.Read()
		if err == io.EOF {
			return records, nil
		}
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}

func (r StructuredReader) ForEach(f func(*Row)) error {
	for {
		row, err := r.Read()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		f(&row)
	}

	return nil
}

func (r Row) Get(header string) string {
	for _, c := range r.Columns {
		if c.Header == header {
			return c.Value
		}
	}

	return ""
}

func (r Row) Set(header string, value string) {
	for _, c := range r.Columns {
		if c.Header == header {
			c.Value = value
			return
		}
	}
}

func NewReader(r io.Reader) (*StructuredReader, error) {
	csvReader := csv.NewReader(r)

	headers, err := csvReader.Read()
	if err != nil {
		return nil, err
	}

	return &StructuredReader{reader: csvReader, headers: headers}, nil
}
