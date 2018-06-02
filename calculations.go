package main

func (rgb modifiedrgba) max() uint8 {
	max := rgb.r
	if rgb.g > max {
		max = rgb.g
	}
	if rgb.b > max {
		max = rgb.b
	}
	return max
}

func (rgb modifiedrgba) min() uint8 {
	min := rgb.r
	if rgb.g < min {
		min = rgb.g
	}
	if rgb.b < min {
		min = rgb.b
	}
	return min
}

func toUint8(x float64) uint8 {
	var i uint8
	if x <= 0 {
		i = 0
	} else if x >= 255 {
		i = 255
	} else {
		i = uint8(x)
	}
	return i
}

func (rgb modifiedrgba) modd(minD float64, maxD float64) modifiedrgba {
	var x modifiedrgba
	x.r = uint8(((float64(rgb.r) - minD) / (maxD - minD)) * (255.0 - 0.0))
	x.g = uint8(((float64(rgb.g) - minD) / (maxD - minD)) * (255.0 - 0.0))
	x.b = uint8(((float64(rgb.b) - minD) / (maxD - minD)) * (255.0 - 0.0))
	return x
}
