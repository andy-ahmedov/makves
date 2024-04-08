package postgres

import (
	"context"
	"reflect"

	"github.com/andy-ahmedov/makves/internal/domain"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Makves struct {
	db *sqlx.DB
}

func NewMakvesRepository(db *sqlx.DB) *Makves {
	return &Makves{
		db: db,
	}
}

func (m *Makves) GetOne(ctx context.Context, id int64) (domain.MakvesPtr, error) {
	var Makves domain.MakvesPtr

	request := `SELECT * FROM makves_data WHERE id=$1`

	err := m.db.Get(&Makves, request, id)

	return Makves, err
}

func (t *Makves) Get(ctx context.Context, ids []int64) ([]domain.Makves, error) {
	var makveses []domain.MakvesPtr

	query, args, err := sqlx.In(`SELECT * FROM makves_data WHERE id IN (?)`, ids)
	if err != nil {
		return nil, err
	}

	query = t.db.Rebind(query)
	err = t.db.Select(&makveses, query, args...)
	if err != nil {
		return nil, err
	}
	results := make([]domain.Makves, len(makveses))

	for i, data := range makveses {
		results[i] = ConvertMakvesPtrToMakves(&data)
	}

	return results, err
}

func ConvertMakvesPtrToMakves(mPtr *domain.MakvesPtr) domain.Makves {
	result := domain.Makves{}
	valPtr := reflect.ValueOf(mPtr).Elem()
	val := reflect.ValueOf(&result).Elem()

	for i := 0; i < valPtr.NumField(); i++ {
		fieldPtr := valPtr.Field(i)
		field := val.Field(i)

		if fieldPtr.Kind() == reflect.Ptr {
			if fieldPtr.IsNil() {
				switch field.Kind() {
				case reflect.String:
					field.SetString("")
				case reflect.Int:
					field.SetInt(0)
				}
			} else {
				field.Set(reflect.ValueOf(fieldPtr.Elem().Interface()))
			}
		} else {
			field.Set(fieldPtr)
		}
	}

	return result
}
