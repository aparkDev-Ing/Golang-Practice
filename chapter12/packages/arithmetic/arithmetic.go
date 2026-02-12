package arithmetic

type Numbers struct {
	X int
	Y int
}

// plus function 10+10 for example
func (c *Numbers) Plus() int {

	return c.X + c.Y
}

func (c *Numbers) Minus() int {

	return c.X - c.Y
}

func (c *Numbers) Multiply() int {

	return c.X * c.Y
}

func (c *Numbers) Divide() int {

	return c.X / c.Y
}
