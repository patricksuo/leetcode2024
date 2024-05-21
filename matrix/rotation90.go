package matrix

func RotateNightDegree(matrix [][]int) {
	/*
	   观察发现，N*N的矩阵旋转90度，最外层的元素顺着矩阵边缘流转N-1格子
	   矩阵内部是一个子问题
	*/

	rotateN(matrix, len(matrix))
}

func rotateN(matrix [][]int, n int) {
	if n <= 1 {
		return
	}

	// 计算出起点 x0, y0
	// 最多移动 n-1 步
	// 先像右移动 ，如果超过矩阵范围，则改为向下移动
	// 如果超出范围则改向左移动
	// 如果超出范围则向上移动

	// 有非常多犯错的点:
	// 1) 一轮交换多少个元素？
	// 2) 一共需要交换多少轮
	// 3) x,y 坐标的计算维护 array2d[x][y] 向右移动需要先增大y

	originSize := len(matrix)
	minX, minY := (originSize-n)/2, (originSize-n)/2
	maxX := minX + n - 1
	maxY := minY + n - 1

	// 一轮交换4个元素，交换 N-1次
	for i := 0; i < n-1; i++ {
		x1, y1 := minX, minY+i

		x2, y2 := add(x1, y1, minX, minY, maxX, maxY, n-1)
		x3, y3 := add(x2, y2, minX, minY, maxX, maxY, n-1)
		x4, y4 := add(x3, y3, minX, minY, maxX, maxY, n-1)

		a, b, c, d := matrix[x1][y1], matrix[x2][y2], matrix[x3][y3], matrix[x4][y4]
		matrix[x1][y1], matrix[x2][y2], matrix[x3][y3], matrix[x4][y4] = d, a, b, c

	}

	rotateN(matrix, n-2)
}

func add(x, y, minX, minY, maxX, maxY, step int) (int, int) {
	if x == minX {
		// move right move down
		y += step
		if y > maxY {
			step = y - maxY
			y = maxY
			x += step
		}

		return x, y
	}

	if y == maxY {
		// move down move left
		x += step
		if x > maxX {
			step = x - maxX
			x = maxX
			y -= step
		}
		return x, y
	}

	if x == maxX {
		// move left move up
		y -= step
		if y < minY {
			step = minY - y
			y = minY
			x -= step
		}

		return x, y
	}

	// move up move right

	x -= step
	if x <= minX {
		step = minX - x
		x = minX
		y += step
	}

	return x, y
}
