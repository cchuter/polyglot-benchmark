package erratum

func Use(opener ResourceOpener, input string) (err error) {
	var res Resource
	for {
		res, err = opener()
		if err == nil {
			break
		}
		if _, ok := err.(TransientError); !ok {
			return err
		}
	}
	defer res.Close()
	defer func() {
		if r := recover(); r != nil {
			if fe, ok := r.(FrobError); ok {
				res.Defrob(fe.defrobTag)
				err = fe
			} else if e, ok := r.(error); ok {
				err = e
			}
		}
	}()
	res.Frob(input)
	return nil
}
