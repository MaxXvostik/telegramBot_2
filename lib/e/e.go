package e

import "fmt"

func Wrap(msq string, err error) error {
	return fmt.Errorf("%s: %w", msq, err)
}
func WrapIfErr(msq string, err error) error {
	if err == nil {
		return nil
	}
	return Wrap(msq, err)
}
