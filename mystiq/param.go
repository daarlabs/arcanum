package mystiq

import (
	"fmt"
	"strings"
	
	"github.com/daarlabs/arcanum/mirage"
	"github.com/daarlabs/arcanum/quirk"
	"github.com/daarlabs/arcanum/util"
)

type Param struct {
	Fulltext string
	Offset   int
	Limit    int
	Order    []string
	Fields   Fields
}

type OrderParam struct {
	Field     string
	Direction string
}

type Fields struct {
	Fulltext []string
	Order    map[string]string
}

func (p Param) Parse(c mirage.Ctx) Param {
	c.Parse().MustQuery(Fulltext, &p.Fulltext)
	c.Parse().MustQuery(Offset, &p.Offset)
	c.Parse().MustQuery(Limit, &p.Limit)
	c.Parse().Multiple().MustQuery(Order, &p.Order)
	return p
}

func (p Param) Use(q *quirk.Quirk) {
	useFulltextParam(q, p.Fulltext, p.Fields.Fulltext)
	useOrderParam(q, p.Order, p.Fields.Order)
	useOffsetParam(q, p.Offset)
	useLimitParam(q, p.Limit)
}

func useFulltextParam(q *quirk.Quirk, fulltext string, columns []string) {
	if len(fulltext) == 0 || len(columns) == 0 {
		return
	}
	startWord := "WHERE"
	if q.WhereExists() {
		startWord = "AND"
	}
	conditions := make([]string, len(columns))
	args := make(quirk.Map)
	for i := range conditions {
		name := fmt.Sprintf("fulltext%d", i+1)
		args[name] = quirk.CreateTsQuery(fulltext)
		conditions[i] = columns[i] + " @@ to_tsquery(@" + name + ")"
	}
	q.Q(startWord+` (`+strings.Join(conditions, " OR ")+`)`, args)
}

func useOrderParam(q *quirk.Quirk, order []string, columns map[string]string) {
	if len(order) == 0 || len(columns) == 0 {
		return
	}
	r := make([]string, 0)
	for _, o := range order {
		if !strings.Contains(o, ":") {
			continue
		}
		parts := strings.Split(o, ":")
		if len(parts) < 2 || parts[1] == "" {
			continue
		}
		column, ok := columns[parts[0]]
		if !ok {
			continue
		}
		r = append(r, util.EscapeString(column)+" "+util.EscapeString(strings.ToUpper(parts[1])))
	}
	if len(r) == 0 {
		return
	}
	q.Q(`ORDER BY ` + strings.Join(r, ","))
}

func useOffsetParam(q *quirk.Quirk, offset int) {
	q.Q(`OFFSET @offset`, quirk.Map{Offset: offset})
}

func useLimitParam(q *quirk.Quirk, limit int) {
	if limit == -1 {
		return
	}
	if limit == 0 {
		limit = DefaultLimit
	}
	q.Q(`LIMIT @limit`, quirk.Map{Limit: limit})
}
