package libs

type TileMap [][]int

func (t *TileMap) AddRow(row []int) {
	*t = append(*t, row)
}

func (t TileMap) NumRows() int {
	return len(t)
}

type TileMapSet []TileMap

func (tms *TileMapSet) AddTileMap(tileMap TileMap) {
	*tms = append(*tms, tileMap)
}

func (tms TileMapSet) GetTileMap(index int) (TileMap, bool) {
	if index < 0 || index >= len(tms) {
		return nil, false
	}
	return tms[index], true
}

func (tms *TileMapSet) RemoveTileMap(index int) bool {
	if index < 0 || index >= len(*tms) {
		return false
	}
	*tms = append((*tms)[:index], (*tms)[index+1:]...)
	return true
}
