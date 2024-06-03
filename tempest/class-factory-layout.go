package tempest

import "fmt"

func containerClass(breakpoints, container map[string]string) string {
	return createContainerCenter() + " " +
		createBreakpoint(breakpoints, None, createContainer(container[None])) + " " +
		createBreakpoint(breakpoints, SizeSm, createContainer(container[SizeSm])) + " " +
		createBreakpoint(breakpoints, SizeMd, createContainer(container[SizeMd])) + " " +
		createBreakpoint(breakpoints, SizeLg, createContainer(container[SizeLg])) + " " +
		createBreakpoint(breakpoints, SizeXl, createContainer(container[SizeXl])) + " " +
		createBreakpoint(breakpoints, SizeXxl, createContainer(container[SizeXxl]))
}

func overflowClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{overflow: %s;}`,
		selector,
		value,
	)
}

func overflowXAxisClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{overflow-x: %s;}`,
		selector,
		value,
	)
}

func overflowYAxisClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{overflow-y: %s;}`,
		selector,
		value,
	)
}

func positionClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{position: %s;}`,
		selector,
		value,
	)
}

func topClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{top: %s;}`,
		selector,
		value,
	)
}

func rightClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{top: %s;}`,
		selector,
		value,
	)
}

func bottomClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{top: %s;}`,
		selector,
		value,
	)
}

func leftClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{top: %s;}`,
		selector,
		value,
	)
}

func insetClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{top: %[2]s;right: %[2]s;bottom: %[2]s;left: %[2]s;}`,
		selector,
		value,
	)
}

func insetXAxisClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{right: %[2]s;left: %[2]s;}`,
		selector,
		value,
	)
}

func insetYAxisClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{top: %[2]s;bottom: %[2]s;}`,
		selector,
		value,
	)
}
