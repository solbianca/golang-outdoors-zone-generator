package tiles

import (
	"fmt"
	"log"
)

type Collection struct {
	columns, rows int
	t             map[AddressData]*Tile

	tilesListCache []*Tile
}

func NewEmptyTileCollection(columns, rows int) *Collection {
	tileCollection := &Collection{
		columns: columns,
		rows:    rows,
		t:       map[AddressData]*Tile{},
	}

	for column := 0; column < columns; column++ {
		for row := 0; row < rows; row++ {
			tileCollection.Set(NewTile(column, row))
		}
	}

	for column := 0; column < columns; column++ {
		for row := 0; row < rows; row++ {
			tile, _ := tileCollection.Get(column, row)

			tile.neighbors = findNeighborsForTile(column, row, tileCollection)
		}
	}

	return tileCollection
}

func (c *Collection) AsList() []*Tile {
	if c.tilesListCache != nil {
		return c.tilesListCache
	}

	tiles := make([]*Tile, 0, len(c.t))

	for column := 0; column < c.columns; column++ {
		for row := 0; row < c.rows; row++ {
			if tile, err := c.Get(column, row); err == nil {
				tiles = append(tiles, tile)
			}
		}
	}

	if c.tilesListCache == nil {
		c.tilesListCache = tiles
	}

	return tiles
}

func (c *Collection) Has(column int, row int) bool {
	if _, ok := c.t[newAddress(column, row)]; ok {
		return true
	}

	return false
}

func (c *Collection) Get(column, row int) (*Tile, error) {
	if tile, ok := c.t[newAddress(column, row)]; ok {
		return tile, nil
	}

	return nil, fmt.Errorf("can't find tile by AddressData Column [%d] and Row [%d]", column, row)
}

func (c *Collection) Set(tile *Tile) {
	c.t[tile.address] = tile
}

func findNeighborsForTile(column, row int, collection *Collection) *neighborsCollection {
	if !collection.Has(column, row) {
		log.Print(fmt.Sprintf(
			"Try to find neighbors for tile that not exists in tile collection. Tile [Column: %d Row: %d].",
			column,
			row,
		))

		return newEmptyNeighborsCollection()
	}

	neighbors := newEmptyNeighborsCollection()

	if collection.Has(column+1, row) {
		neighbors.set(column+1, row, newNeighborTile(column+1, row))
	}
	if collection.Has(column-1, row) {
		neighbors.set(column-1, row, newNeighborTile(column-1, row))
	}
	if collection.Has(column, row+1) {
		neighbors.set(column, row+1, newNeighborTile(column, row+1))
	}
	if collection.Has(column, row-1) {
		neighbors.set(column, row-1, newNeighborTile(column, row-1))
	}
	if collection.Has(column+1, row+1) {
		neighbors.set(column+1, row+1, newNeighborTile(column+1, row+1))
	}
	if collection.Has(column+1, row-1) {
		neighbors.set(column+1, row-1, newNeighborTile(column+1, row-1))
	}
	if collection.Has(column-1, row-1) {
		neighbors.set(column-1, row-1, newNeighborTile(column-1, row-1))
	}
	if collection.Has(column-1, row+1) {
		neighbors.set(column-1, row+1, newNeighborTile(column-1, row+1))
	}

	return neighbors
}
