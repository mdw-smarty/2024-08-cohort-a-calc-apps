package handlers

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

type CSVHandler struct {
	reader      *csv.Reader
	writer      *csv.Writer
	calculators map[string]Calculator
}

func NewCSVHandler(reader io.Reader, writer io.Writer, calculators map[string]Calculator) *CSVHandler {
	csvReader := csv.NewReader(reader)
	csvReader.FieldsPerRecord = -1
	return &CSVHandler{
		reader:      csvReader,
		writer:      csv.NewWriter(writer),
		calculators: calculators,
	}
}

func (this *CSVHandler) Handle() error {
	for {
		record, err := this.reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("read error: %w", err)
		}
		if len(record) < 3 {
			continue
		}
		calculator, ok := this.calculators[record[1]]
		if !ok {
			continue
		}
		operand1, err := strconv.Atoi(record[0])
		if err != nil {
			continue
		}
		operand2, err := strconv.Atoi(record[2])
		if err != nil {
			continue
		}
		result := calculator.Calculate(operand1, operand2)
		err = this.writer.Write(append(record, strconv.Itoa(result)))
		if err != nil {
			break
		}
	}
	this.writer.Flush()
	return this.writer.Error()
}
