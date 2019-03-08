package orderService

import (
	"fmt"
	"strings"
	"testing"
)

func TestOrderQuery(t *testing.T) {
	s := "aa"
	c := "vv"
	v := ""

	patterns := []string{}
	values := []string{}
	fields := []interface{}{}
	pattern := "%v LIKE ?"
	comma := " OR "
	if len(s) > 0 {
		patterns = append(patterns, pattern)
		fields = append(fields, "services.service_name")
		values = append(values, "%"+s+"%")
	}
	if len(c) > 0 {
		patterns = append(patterns, pattern)
		fields = append(fields, "customers.username")
		values = append(values, "%"+c+"%")
	}
	if len(v) > 0 {
		patterns = append(patterns, pattern)
		fields = append(fields, "vendors.username")
		values = append(values, "%"+v+"%")
	}
	whereFormat := strings.Join(patterns, comma)
	whereStr := fmt.Sprintf(whereFormat, fields...)

	fmt.Println(whereStr)
}
