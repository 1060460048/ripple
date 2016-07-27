package ripple

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"reflect"
	"strings"
)

// Model facilitate database interactions, support mysql
type Model struct {
	models   map[string]reflect.Value
	isOpenDB bool
	*gorm.DB
}

// NewModel returns a new Model without opening database connection
func NewModel() *Model {
	return &Model{
		models: make(map[string]reflect.Value),
	}
}

// NewModelWithConfig creates a new model, and opens database connection based on cfg settings
func NewModelWithConfig() (*Model, error) {
	m := NewModel()
	if err := m.OpenWithConfig(); err != nil {
		return nil, err
	}
	return m, nil
}

// OpenWithConfig opens database connection with the settings found in cfg
func (m *Model) OpenWithConfig() error {
	db, err := GetDbWithGorm()
	if err != nil {
		return err
	}
	m.DB = db
	m.isOpenDB = true
	return nil
}

// IsOpenDB returns true if the Model has already established connection
// to the database
func (m *Model) IsOpenDB() bool {
	return m.isOpenDB
}

// AddModels add the values to the models registry
func (m *Model) AddModels(values ...interface{}) error {
	// do not work on them.models first, this is like an insurance policy
	// whenever we encounter any error in the values nothing goes into the registry
	models := make(map[string]reflect.Value)
	if len(values) > 0 {
		for _, val := range values {
			rVal := reflect.ValueOf(val)
			if rVal.Kind() == reflect.Ptr {
				rVal = rVal.Elem()
			}
			switch rVal.Kind() {
			case reflect.Struct:
				models[getTypName(rVal.Type())] = reflect.New(rVal.Type())
			default:
				return errors.New("ripple: model must be struct type")
			}
		}
	}
	for k, v := range models {
		m.models[k] = v
	}
	return nil
}

// AutoMigrateAll runs migrations for all the registered models
func (m *Model) AutoMigrateAll() {
	for _, v := range m.models {
		m.AutoMigrate(v.Interface())
	}
}

// getTypName returns a string representing the name of the object typ.
// if the name is defined then it is used, otherwise, the name is derived from the
// Stringer interface.
//
// the stringer returns something like *somepkg.MyStruct, so skip
// the *somepkg and return MyStruct
func getTypName(typ reflect.Type) string {
	if typ.Name() != "" {
		return typ.Name()
	}
	split := strings.Split(typ.String(), ".")
	return split[len(split)-1]
}
