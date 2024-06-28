package mirage

func defaultDynamicHandler(c Ctx) error {
	return c.Response().
		Status(c.Response().Intercept().Status()).
		Error(c.Response().Intercept().Error())
}
