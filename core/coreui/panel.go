package coreui

//type PanelInterface interface {
//	createNoiseMap()
//	Draw(screen *ebiten.Image)
//	Update()
//}
//
//type Panel struct {
//	windowWidth, windowHeight float64
//	GlobalX, GlobalY          float64
//	Width, Height             int
//}
//
//func NewPanel(windowWidth, windowHeight float64, width, height int, globalX, globalY float64) *Panel {
//	return &Panel{
//		windowWidth:  windowWidth,
//		windowHeight: windowHeight,
//		GlobalX:      globalX,
//		GlobalY:      globalY,
//		Width:        width,
//		Height:       height,
//	}
//}
//
//func (p Panel) GlobalToLocalCoordinates(globalX, globalY float64) (bool, float64, float64) {
//	return false, 0.0, 0.0
//}
//
//func (p Panel) LocalToGlobalCoordinates(localX, localY float64) (float64, float64) {
//	if localX > p.windowWidth {
//		panic(fmt.Sprintf("Local value of [X:%f] is more than window Columns.", localX))
//	}
//
//	if localY > p.windowHeight {
//		panic(fmt.Sprintf("Local value of [Y:%f] is more than window Rows.", localY))
//	}
//
//	if localX < 0.0 || localY < 0.0 {
//		panic(fmt.Sprintf("Local value of [X:%f] or [Y:%f] can't be less that zero.", localX, localY))
//	}
//
//	return p.GlobalX + localX, p.GlobalY + localY
//}
