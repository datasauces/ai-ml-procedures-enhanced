package ml

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"

	"github.com/ornerylawn/linear"
)

func LoadMatrixCSV(filename string) (linear.Matrix, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(f)
	recordLength := -1
	values := []float64{}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if recordLength == -1 {
			recordLength = len(record)
			continue // first record is header
		} else if len(record) != recordLength {
			return nil, fmt.Errorf("expected record of length %d but found length %d", recordLength, len(record))
		}
		for _, v := range record {
			x, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return nil, err
			}
			values = append(values, x)
		}
	}
	if len(values) == 0 {
		return nil, errors.New("no values found in csv")
	}
	ins := recordLength
	outs := len(values) / recordLength
	A := linear.NewArrayMatrix(ins, outs)
	for o := 0; o < outs; o++ {
		for i := 0; i < ins; i++ {
			A.Set(