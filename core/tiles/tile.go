package tiles

import (
	"fmt"
)

type Tile struct {
	address   AddressData
	neighbors *neighborsCollection
}

func NewTile(
	column int,
	row int,
) *Tile {
	return &Tile{
		address:   newAddress(column, row),
		neighbors: newEmptyNeighborsCollection(),
	}
}

func (t *Tile) CountNeighbors() int {
	return len(t.neighbors.tiles)
}

func (t *Tile) GetAddress() (column, row int) {
	return t.address.GetAddress()
}

type neighborTile struct {
	address AddressData
}

func newNeighborTile(column int, row int) *neighborTile {
	return &neighborTile{newAddress(column, row)}
}

func (n neighborTile) GetAddress() (int, int) {
	return n.address.GetAddress()
}

func (t *Tile) addNeighbor(neighbor *Tile) {
	neighborColumn, neighborRow := neighbor.GetAddress()
	t.neighbors.set(neighborColumn, neighborRow, newNeighborTile(neighborColumn, neighborRow))
}

type neighborsCollection struct {
	tiles map[AddressData]*neighborTile
}

func newEmptyNeighborsCollection() *neighborsCollection {
	return &neighborsCollection{
		tiles: map[AddressData]*neighborTile{},
	}
}

func (collection *neighborsCollection) has(column int, row int) bool {
	if _, ok := collection.tiles[newAddress(column, row)]; ok {
		return true
	}

	return false
}

func (collection *neighborsCollection) get(column int, row int) (*neighborTile, error) {
	if tile, ok := collection.tiles[newAddress(column, row)]; ok {
		return tile, nil
	}

	return nil, fmt.Errorf("can't find tile by AddressData Column [%d] and Row [%d]", column, row)
}

func (collection *neighborsCollection) set(column int, row int, tile *neighborTile) {
	collection.tiles[newAddress(column, row)] = tile
}

func (collection *neighborsCollection) asList() []*neighborTile {
	neighbors := make([]*neighborTile, 0, len(collection.tiles))

	for _, tile := range collection.tiles {
		neighbors = append(neighbors, tile)
	}

	return neighbors
}
