package bitsdecoder

import (
	"fmt"
	"os"

	"github.com/CurlyQuokka/AdventOfCode2021/Day22/pkg/utils"
)

const (
	versionLength      = 3
	versionTypeLength  = 6
	literalValueNibble = 5
	typeZeroLength     = 15
	typeOneLength      = 11

	zeroTypeOperator            = "0"
	literalValueLastNibbleToken = "0"

	sumOp         = 0
	productOp     = 1
	minimumOp     = 2
	maximumOp     = 3
	literal       = 4
	greaterThanOp = 5
	lessThanOp    = 6
	equalOp       = 7

	notEvaluatedValue = -1
)

type generalPacket struct {
	pVersion uint64
	pType    uint64
}

func getVerType(data string, n *int) (uint64, uint64) {
	v := utils.BinStringToDec(data[*n : *n+versionLength])
	t := utils.BinStringToDec(data[*n+versionLength : *n+versionTypeLength])
	*n += versionTypeLength
	return v, t
}

type literalValuePacket struct {
	pGeneral generalPacket
	value    uint64
}

func newLiteralValuePacket(v, t uint64) literalValuePacket {
	p := literalValuePacket{}
	p.pGeneral.pVersion = v
	p.pGeneral.pType = t
	return p
}

func (lvp *literalValuePacket) Process(data string, n *int) {
	value := ""
	for {
		value += data[*n+1 : *n+literalValueNibble]
		*n += literalValueNibble
		if string(data[*n-literalValueNibble]) == literalValueLastNibbleToken {
			break
		}

	}
	lvp.value = utils.BinStringToDec(value)
}

type operatorPacket struct {
	pGeneral   generalPacket
	lengthType string
	len        int
	packets    []interface{}
	value      int
}

func newOperatorPacket(v, t uint64) operatorPacket {
	p := operatorPacket{}
	p.pGeneral.pVersion = v
	p.pGeneral.pType = t
	p.value = notEvaluatedValue
	return p
}

func (op *operatorPacket) Process(data string, n *int) {
	op.lengthType = string(data[*n])
	if op.lengthType == zeroTypeOperator {
		op.len = utils.ConvertBinToDec(data[*n+1 : *n+typeZeroLength+1])

		*n += typeZeroLength + 1
		oldN := *n
		for {
			if *n-oldN >= op.len {
				break
			}
			p := createPacket(data, n)
			op.packets = append(op.packets, p)
		}

	} else {
		op.len = utils.ConvertBinToDec(data[*n+1 : *n+typeOneLength+1])
		*n += typeOneLength + 1
		for i := 0; i < op.len; i++ {
			p := createPacket(data, n)
			op.packets = append(op.packets, p)
		}
	}
}

func (op *operatorPacket) Evaluate() {
	switch op.pGeneral.pType {
	case sumOp:
		op.evaluateSum()
	case productOp:
		op.evaluateProduct()
	case minimumOp:
		op.evaluateMinimum()
	case maximumOp:
		op.evaluateMaximum()
	case greaterThanOp:
		op.evaluateGreaterThan()
	case lessThanOp:
		op.evaluateLessThan()
	case equalOp:
		op.evaluateEqual()
	}
}

type BitsDecoder struct {
	data    []string
	binary  string
	packets []interface{}
}

func NewBitsDecoder(path string) *BitsDecoder {
	bd := &BitsDecoder{
		data: utils.LoadData(path),
	}
	bd.binary = utils.HexStringToBin(&bd.data[0])
	return bd
}

func createPacket(binary string, n *int) interface{} {
	v, t := getVerType(binary, n)
	if t == literal {
		p := newLiteralValuePacket(v, t)
		p.Process(binary, n)
		return &p
	} else {
		p := newOperatorPacket(v, t)
		p.Process(binary, n)
		return &p
	}
}

func (bd *BitsDecoder) Decode() {
	n := 0
	p := createPacket(bd.binary, &n)
	bd.packets = append(bd.packets, p)
}

func (bd *BitsDecoder) SumVersions() {
	sum := sumVersions(bd.packets)
	fmt.Printf("Version sum: %d\n", sum)
}

func sumVersions(packets []interface{}) uint64 {
	var sum uint64
	for _, p := range packets {
		if casted, ok := p.(*operatorPacket); ok {
			sum += casted.pGeneral.pVersion + sumVersions(casted.packets)
		}
		if casted, ok := p.(*literalValuePacket); ok {
			sum += casted.pGeneral.pVersion
		}
	}
	return sum
}

func (bd *BitsDecoder) EvaluateBits() {
	if casted, ok := bd.packets[0].(*operatorPacket); ok {
		casted.Evaluate()
		fmt.Printf("Evaluated value: %d\n", casted.value)
	} else {
		fmt.Printf("Unable to cast and evaluate")
		os.Exit(999)
	}
}
