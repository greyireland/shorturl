package dao

func (d *dao) getCode(raw string) (code string, err error) {
	res, ok := d.lru.Get(raw)
	if !ok {
		return
	}
	return res.(string), nil
}
func (d *dao) addCode(code, raw string) (err error) {
	d.lru.Add(raw, code)
	return nil
}
