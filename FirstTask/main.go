package FirstTask

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	err := Call()
	if errors.Cause(err) == sql.ErrNoRows {
		fmt.Printf("data not found,%v\n", err)
		fmt.Printf("%+v\n", err)
		return
	}
}

func getSql() error {
	return errors.Wrap(sql.ErrNoRows, "getSql failed")
}
func Call() error {
	return errors.WithMessage(getSql(), "call failed")
}
