package mystiq

import (
	"cmp"
	"encoding/json"
	"fmt"
	"reflect"
	"slices"
	"strings"
	
	"github.com/daarlabs/arcanum/quirk"
	"github.com/daarlabs/arcanum/util"
)

type Mystiq interface {
	DB(db *quirk.DB, query Query) Mystiq
	Data(data []map[string]any) Mystiq
	
	GetAllFunc(fn func(param Param, t any) error) Mystiq
	GetOneFunc(fn func(name string, v any, t any) error) Mystiq
	GetManyFunc(fn func(name string, v any, t any) error) Mystiq
	
	GetAll(param Param, t any) error
	GetOne(name string, v any, t any) error
	GetMany(name string, v any, t any) error
	
	MustGetAll(param Param, t any)
	MustGetOne(name string, v any, t any)
	MustGetMany(name string, v any, t any)
}

type mystiq struct {
	db          *quirk.DB
	data        []map[string]any
	query       Query
	getAllFunc  func(param Param, t any) error
	getOneFunc  func(name string, v any, t any) error
	getManyFunc func(name string, v any, t any) error
}

func New() Mystiq {
	return &mystiq{
		data: make([]map[string]any, 0),
	}
}

func (m *mystiq) DB(db *quirk.DB, query Query) Mystiq {
	m.db = db
	m.query = query
	return m
}

func (m *mystiq) Data(data []map[string]any) Mystiq {
	m.data = data
	return m
}

func (m *mystiq) GetAllFunc(fn func(param Param, t any) error) Mystiq {
	m.getAllFunc = fn
	return m
}

func (m *mystiq) GetOneFunc(fn func(name string, v any, t any) error) Mystiq {
	m.getOneFunc = fn
	return m
}

func (m *mystiq) GetManyFunc(fn func(name string, v any, t any) error) Mystiq {
	m.getManyFunc = fn
	return m
}

func (m *mystiq) GetAll(param Param, t any) error {
	if param.Limit == 0 {
		param.Limit = DefaultLimit
	}
	if m.getAllFunc != nil {
		return m.getAllFunc(param, t)
	}
	if m.shouldUseDb() {
		return m.getAllWithDb(param, t)
	}
	return m.getAllWithData(param, t)
}

func (m *mystiq) GetOne(name string, v any, t any) error {
	if m.getOneFunc != nil {
		return m.getOneFunc(name, v, t)
	}
	if m.shouldUseDb() {
		return m.getOneWithDb(name, v, t)
	}
	return m.getOneWithData(name, v, t)
}

func (m *mystiq) GetMany(name string, v any, t any) error {
	if m.getManyFunc != nil {
		return m.getManyFunc(name, v, t)
	}
	if m.shouldUseDb() {
		return m.getManyWithDb(name, v, t)
	}
	return m.getManyWithData(name, v, t)
}

func (m *mystiq) MustGetAll(param Param, t any) {
	if err := m.GetAll(param, t); err != nil {
		panic(err)
	}
}

func (m *mystiq) MustGetOne(name string, v any, t any) {
	if err := m.GetOne(name, v, t); err != nil {
		panic(err)
	}
}

func (m *mystiq) MustGetMany(name string, v any, t any) {
	if err := m.GetMany(name, v, t); err != nil {
		panic(err)
	}
}

func (m *mystiq) getAllWithDb(param Param, t any) error {
	q := m.db.Q("SELECT " + m.createFields()).
		Q(`FROM ` + m.query.Table).
		Q(`AS ` + m.query.Alias)
	param.Use(q)
	return q.Exec(t)
}

func (m *mystiq) getAllWithData(param Param, t any) error {
	result := make([]map[string]any, 0)
	for i, row := range m.data {
		if param.Offset != 0 && i < param.Offset {
			continue
		}
		if param.Fulltext != "" {
			shouldContinue := false
			for _, v := range row {
				if strings.Contains(quirk.Normalize(fmt.Sprintf("%v", v)), quirk.Normalize(param.Fulltext)) {
					shouldContinue = true
				}
			}
			if !shouldContinue {
				continue
			}
		}
		result = append(result, row)
		if param.Limit != 0 && i >= param.Limit+param.Offset-1 {
			break
		}
	}
	m.sortDataResult(param, result)
	resultBytes, err := json.Marshal(result)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(resultBytes, t); err != nil {
		return err
	}
	return nil
}

func (m *mystiq) sortDataResult(param Param, data []map[string]any) {
	slices.SortFunc(
		data, func(a, b map[string]any) int {
			if len(param.Order) == 0 {
				return 0
			}
			order := param.Order[0]
			parts := strings.Split(order, ":")
			if len(parts) < 2 {
				return 0
			}
			name := util.EscapeString(parts[0])
			direction := util.EscapeString(parts[1])
			va, oka := a[name]
			vb, okb := b[name]
			if !oka || !okb {
				return 0
			}
			if direction == Asc {
				return cmp.Compare(quirk.Normalize(fmt.Sprintf("%v", va)), quirk.Normalize(fmt.Sprintf("%v", vb)))
			}
			if direction == Desc {
				return cmp.Compare(quirk.Normalize(fmt.Sprintf("%v", vb)), quirk.Normalize(fmt.Sprintf("%v", va)))
			}
			return 0
		},
	)
}

func (m *mystiq) getOneWithDb(name string, v any, t any) error {
	fields := make([]string, 0)
	for key, alias := range m.query.Fields {
		fields = append(fields, key+" AS "+alias)
	}
	q := m.db.Q("SELECT "+m.createFields()).
		Q(`FROM `+m.query.Table).
		Q(fmt.Sprintf(`WHERE %[1]s = @%[1]s`, name), quirk.Map{name: v}).
		Q(`LIMIT 1`)
	return q.Exec(t)
}

func (m *mystiq) getOneWithData(name string, v any, t any) error {
	for _, row := range m.data {
		rv, ok := row[name]
		if !ok {
			continue
		}
		if rv != v {
			continue
		}
		rowBytes, err := json.Marshal(row)
		if err != nil {
			return err
		}
		if err := json.Unmarshal(rowBytes, t); err != nil {
			return err
		}
		break
	}
	return nil
}

func (m *mystiq) getManyWithDb(name string, v any, t any) error {
	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Slice {
		return ErrorSliceValue
	}
	q := m.db.Q("SELECT "+m.createFields()).
		Q(`FROM `+m.query.Table).
		Q(`WHERE `+name+` IN (@values)`, quirk.Map{"values": v}).
		Q(fmt.Sprintf("LIMIT %d", vv.Len()))
	return q.Exec(t)
}

func (m *mystiq) getManyWithData(name string, v any, t any) error {
	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Slice {
		return ErrorSliceValue
	}
	result := make([]map[string]any, 0)
	for _, row := range m.data {
		rv, ok := row[name]
		if !ok {
			continue
		}
		var exist bool
		for i := 0; i < vv.Len(); i++ {
			if vv.Index(i).Interface() == rv {
				exist = true
			}
		}
		if !exist {
			continue
		}
		result = append(result, row)
	}
	rowBytes, err := json.Marshal(result)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(rowBytes, t); err != nil {
		return err
	}
	return nil
}

func (m *mystiq) shouldUseDb() bool {
	return m.db != nil && m.query.Fields != nil && m.query.Table != ""
}

func (m *mystiq) createFields() string {
	fields := make([]string, 0)
	for alias, key := range m.query.Fields {
		fields = append(fields, key+" AS "+alias)
	}
	return strings.Join(fields, ", ")
}
