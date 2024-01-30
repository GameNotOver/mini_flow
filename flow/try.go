package flow

type TryCatchHandle struct {
	err   interface{}
	catch bool
}

func Try(fn func()) (ret *TryCatchHandle) {
	ret = &TryCatchHandle{}
	defer func() {
		if r := recover(); r != nil {
			ret.err = r
			ret.catch = true
			//_ = fmt.Errorf("err: %v", r)
		}
	}()
	fn()
	return
}

func (t *TryCatchHandle) Error() ierr.IBizError {
	switch v := t.err.(type) {
	case nil:
		return nil
	case ierr.IBizError:
		return v
	}
	return nil
}

func (t *TryCatchHandle) CatchBizError(fn func(bizError ierr.IBizError)) *TryCatchHandle {
	if t.err == nil {
		return t
	}
	if v, ok := t.err.(ierr.IBizError); ok {
		fn(v)
	}
	return t
}

func (t *TryCatchHandle) Catch(fn func(err interface{})) *TryCatchHandle {
	if t.err == nil {
		return t
	}
	fn(t.err)
	return t
}
