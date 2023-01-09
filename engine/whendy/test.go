package whendy

func Test1() string {
	print(9)
	s := Store{}
	return s.GetSes("BUILD").(string)
}
