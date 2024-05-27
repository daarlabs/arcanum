package form

import (
	"fmt"
	"testing"
	
	"github.com/stretchr/testify/assert"
	
	"github.com/daarlabs/arcanum/gox"
)

func TestCsrf(t *testing.T) {
	t.Run(
		"hidden inputs", func(t *testing.T) {
			name, token := "test", "123456789"
			assert.Equal(
				t,
				fmt.Sprintf(
					`<input type="hidden" name="%s" value="%s" /><input type="hidden" name="%s" value="%s" />`,
					CsrfName, name,
					CsrfToken, token,
				),
				gox.Render(Csrf(name, token)),
			)
		},
	)
}
