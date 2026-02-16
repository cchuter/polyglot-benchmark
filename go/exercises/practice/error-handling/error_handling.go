package erratum

func Use(opener ResourceOpener, input string) (err error) {
	// 1. Open the resource, retrying on TransientError
	var r Resource
	for {
		r, err = opener()
		if err == nil {
			break
		}
		if _, ok := err.(TransientError); !ok {
			return err
		}
	}

	// 2. Ensure Close is called exactly once after successful open
	defer r.Close()

	// 3. Recover from panics in Frob
	defer func() {
		if x := recover(); x != nil {
			if frobErr, ok := x.(FrobError); ok {
				r.Defrob(frobErr.defrobTag)
			}
			err = x.(error)
		}
	}()

	// 4. Call Frob
	r.Frob(input)

	return nil
}
