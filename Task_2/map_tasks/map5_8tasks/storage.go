package main

type Keeper interface {
	GetPriceByName(string) float64
	AppendPr(string, float64) error
	UpdatePr(string, float64) error
}

type StorageIds map[string]uint // [товар]id
type Products map[uint]float64  // [id_товара]цена
// строка заказа вида ["|id:кол-во|id:кол-во..."]суммаЗаказа
type Orders map[string]float64

type StorageBase struct {
	Ids StorageIds
	Prd Products
	Ord Orders
}

// Keeper methods
func (strBase *StorageBase) GetPriceByName(prdName string) float64 {
	return strBase.Prd[strBase.Ids[prdName]]
}

func (strBase *StorageBase) AppendPr(product string, cost float64) error {
	if _, ok := strBase.Ids[product]; product == "" || cost <= 0 || ok {
		return AppendErr("Unable to append! Possible problems: cost < 0, empty product str, prd. already in base")
	}

	// получить новое id
	idNew := (strBase.Ids).GetUniqId()
	// записать новые продукт и id в карту [продукты]id
	strBase.Ids[product] = idNew
	// записать новые id и центу в карту [id]цена
	strBase.Prd[idNew] = cost

	return nil
}

// получить id на единицу большее максимального
func (ids *StorageIds) GetUniqId() uint {
	var max uint = 0
	for _, val := range *ids {
		if val > max {
			max = val
		}
	}
	return max + 1
}

// обновить товар
func (strBase *StorageBase) UpdatePr(product string, cost float64) error {
	if _, ok := strBase.Ids[product]; product == "" || cost <= 0 || !ok {
		return AppendErr("Unable to append! Possible problems: cost < 0, empty product str, prd. not in base")
	}
	id := strBase.Ids[product]
	strBase.Prd[id] = cost
	return nil
}
