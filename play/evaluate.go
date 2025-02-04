package play

import "snakehero/models"

//type payoff struct {
//	snakeId string
//	payoff  float64
//}

// changes payoff in combination round
func evaluateRound(board models.MyBoard, moves []models.SnakeMove, heroId string) {
	//type x struct {
	//	snakeId string
	//	lenght  int
	//}
	//tab := []x{}
	//for i, m := range moves {
	//	tab = append(tab, x{snakeId: m.SnakeId, lenght: int(board.getBattlesnake(moves[i].SnakeId).Length)})
	//}
	//sort.SliceStable(tab, func(i, j int) bool {
	//	return tab[i].lenght < tab[j].lenght
	//})
	//
	//for i, s := range tab {
	//
	//}

	for i := range moves {
		p := float64(board.GetBattlesnake(moves[i].SnakeId).Length)
		moves[i].Payoff = p
	}
}
