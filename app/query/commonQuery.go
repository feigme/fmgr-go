package query

type CommonQuery struct {
	PageNo   int
	PageSize int
}

func (c *CommonQuery) GetPageNo() int {
	if c.PageNo < 1 {
		c.PageNo = 1
	}
	return c.PageNo
}

func (c *CommonQuery) GetPageSize() int {
	if c.PageSize < 1 {
		c.PageSize = 20
	}
	return c.PageSize
}

func (c *CommonQuery) GetOffset() int {
	return (c.GetPageNo() - 1) * c.GetPageSize()
}
