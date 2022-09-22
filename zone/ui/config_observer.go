package ui

import (
	"zone/zone/noise"
	"zone/zone/zone/layers"
)

type ConfigObserver struct {
	columns, rows int
	Config        *ConfigPanel
	Zone          *ZonePanel
}

func NewConfigObserver(columns, rows int) *ConfigObserver {
	return &ConfigObserver{columns: columns, rows: rows}
}

func (c *ConfigObserver) Notify(alias string) {
	noiseConfig := noise.NewNoiseMapConfig(
		c.columns,
		c.rows,
		c.Config.layers[alias].noiseConfig.pointSize,
		c.Config.layers[alias].noiseConfig.n,
		c.Config.layers[alias].noiseConfig.seed,
		c.Config.layers[alias].noiseConfig.frequency,
		c.Config.layers[alias].noiseConfig.alpha,
		c.Config.layers[alias].noiseConfig.beta,
	)
	reliefLayer, _ := c.Zone.zone.GetLayer(layers.ReliefLayer)
	plantsLayer, _ := c.Zone.zone.GetLayer(layers.TreeLayer)

	switch alias {
	case relief:
		reliefLayer.SetNoise(noise.NewNoiseMap(noiseConfig))
		plantsLayer.(*layers.Plants).SetReliefSigns(reliefLayer.GetSigns())
	case plants:
		plantsLayer.SetNoise(noise.NewNoiseMap(noiseConfig))
	}
	//reliefLayer, _ := c.Zone.zone.GetLayer(layers.ReliefLayer)
	//reliefLayer.SetNoise(noise.NewNoiseMap(noiseConfig))

	//plantsLayer, _ := c.Zone.zone.GetLayer(layers.TreeLayer)
	//plantsLayer.(*layers.Plants).SetReliefSigns(reliefLayer.GetSigns())
}
