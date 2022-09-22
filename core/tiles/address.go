package tiles

type Address interface {
	GetAddress() (int, int)
}

type AddressData struct {
	column, row int
}

func (a AddressData) GetAddress() (int, int) {
	return a.column, a.row
}

func newAddress(column int, row int) AddressData {
	return AddressData{column: column, row: row}
}
