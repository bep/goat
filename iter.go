package goat

type canvasIterator func(width int, height int) chan Index

func upDown(width int, height int) chan Index {
	c := make(chan Index, width*height)

	go func() {
		for w := range width {
			for h := range height {
				c <- Index{w, h}
			}
		}
		close(c)
	}()

	return c
}

func leftRight(width int, height int) chan Index {
	c := make(chan Index, width*height)

	// Transpose an upDown order.
	go func() {
		for i := range upDown(height, width) {
			c <- Index{i.y, i.x}
		}

		close(c)
	}()

	return c
}

func diagDown(width int, height int) chan Index {
	c := make(chan Index, width*height)

	go func() {
		minSum := -height + 1
		maxSum := width

		for sum := minSum; sum <= maxSum; sum++ {
			for w := range width {
				for h := range height {
					if w-h == sum {
						c <- Index{w, h}
					}
				}
			}
		}
		close(c)
	}()

	return c
}

func diagUp(width int, height int) chan Index {
	c := make(chan Index, width*height)

	go func() {
		maxSum := width + height - 2

		for sum := 0; sum <= maxSum; sum++ {
			for w := range width {
				for h := range height {
					if h+w == sum {
						c <- Index{w, h}
					}
				}
			}
		}
		close(c)
	}()

	return c
}
