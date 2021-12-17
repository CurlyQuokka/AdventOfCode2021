package bitsdecoder

func (op *operatorPacket) evaluateSum() {
	sum := 0
	for _, p := range op.packets {
		if casted, ok := p.(*operatorPacket); ok {
			if casted.value == notEvaluatedValue {
				casted.Evaluate()
				sum += casted.value
			}
		}
		if casted, ok := p.(*literalValuePacket); ok {
			sum += int(casted.value)
		}
	}
	op.value = sum
}

func (op *operatorPacket) evaluateProduct() {
	product := 1
	for _, p := range op.packets {
		if casted, ok := p.(*operatorPacket); ok {
			if casted.value == notEvaluatedValue {
				casted.Evaluate()
				product *= casted.value
			}
		}
		if casted, ok := p.(*literalValuePacket); ok {
			product *= int(casted.value)
		}
	}
	op.value = product
}

func (op *operatorPacket) evaluateMinimum() {
	minimum := -1
	for _, p := range op.packets {
		if casted, ok := p.(*operatorPacket); ok {
			if casted.value == notEvaluatedValue {
				casted.Evaluate()
				if minimum <= -1 || casted.value < minimum {
					minimum = casted.value
				}

			}
		}
		if casted, ok := p.(*literalValuePacket); ok {
			if minimum <= -1 || int(casted.value) < minimum {
				minimum = int(casted.value)
			}
		}
	}
	op.value = minimum
}

func (op *operatorPacket) evaluateMaximum() {
	maximum := -1
	for _, p := range op.packets {
		if casted, ok := p.(*operatorPacket); ok {
			if casted.value == notEvaluatedValue {
				casted.Evaluate()
				if casted.value > maximum {
					maximum = casted.value
				}

			}
		}
		if casted, ok := p.(*literalValuePacket); ok {
			if int(casted.value) > maximum {
				maximum = int(casted.value)
			}
		}
	}
	op.value = maximum
}

func (op *operatorPacket) evaluateGreaterThan() {
	values := []int{}
	for _, p := range op.packets {
		if casted, ok := p.(*operatorPacket); ok {
			if casted.value == notEvaluatedValue {
				casted.Evaluate()
				values = append(values, casted.value)
			}
		}
		if casted, ok := p.(*literalValuePacket); ok {
			values = append(values, int(casted.value))
		}
	}
	if values[0] > values[1] {
		op.value = 1
	} else {
		op.value = 0
	}
}

func (op *operatorPacket) evaluateLessThan() {
	values := []int{}
	for _, p := range op.packets {
		if casted, ok := p.(*operatorPacket); ok {
			if casted.value == notEvaluatedValue {
				casted.Evaluate()
				values = append(values, casted.value)
			}
		}
		if casted, ok := p.(*literalValuePacket); ok {
			values = append(values, int(casted.value))
		}
	}
	if values[0] < values[1] {
		op.value = 1
	} else {
		op.value = 0
	}
}

func (op *operatorPacket) evaluateEqual() {
	values := []int{}
	for _, p := range op.packets {
		if casted, ok := p.(*operatorPacket); ok {
			if casted.value == notEvaluatedValue {
				casted.Evaluate()
				values = append(values, casted.value)
			}
		}
		if casted, ok := p.(*literalValuePacket); ok {
			values = append(values, int(casted.value))
		}
	}
	if values[0] == values[1] {
		op.value = 1
	} else {
		op.value = 0
	}
}
