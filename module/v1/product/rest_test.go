package product

import (
	"go-simple/config"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestModule_productList(t *testing.T) {
	type fields struct {
		Config config.Configuration
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Module{
				Config: tt.fields.Config,
			}
			if err := m.productList(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Module.productList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
