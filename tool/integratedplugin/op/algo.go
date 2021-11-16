// 算法

package op

func (com *Opsoft) AStarFindPath(mapWidth, mapHeight int, disablePoints string, beginX, beginY, endX, endY int) int {
	ret, _ := com.iDispatch.CallMethod("AStarFindPath", mapWidth, mapHeight, disablePoints, beginX, beginY, endX, endY)
	return int(ret.Val)
}
