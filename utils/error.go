package utils

func PanicIfError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func ErrReturnErr(err error) error {
	if err != nil {
		return err
	}

	return nil
}
