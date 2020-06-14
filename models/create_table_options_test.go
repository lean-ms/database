package models

import (
	"fmt"
	"testing"

	"github.com/go-pg/pg/orm"
)

var conversionTestCases = []*CreateTableOptions{
	nil,
	{},
	{Temp: true},
}

func TestCreateTableOptionsConversion(t *testing.T) {
	for _, conversionTestCase := range conversionTestCases {
		x := (*orm.CreateTableOptions)(conversionTestCase)
		fmt.Println(x)
	}
}
