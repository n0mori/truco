package truco

//PlayerState groups the informations about the player
type PlayerState struct {
	ID     int
	Tentos int
	Maos   int
	Cards  Deck
	Active bool
}

//GameState is the current state of the game
type GameState struct {
	PlayerStates []PlayerState
	TableCards   Deck
	Round        int
}

//NewPlayerState returns a pointer to a newly created PlayerState
func NewPlayerState(id int) PlayerState {
	return PlayerState{
		ID:     id,
		Tentos: 0,
		Maos:   0,
		Cards:  make(Deck, 0),
		Active: false}
}

//Deal the cards to the you and enemy players
func Deal(gs *GameState) {
	deck := NewDeck()
	gs.PlayerStates[0].Cards = deck[:5]
	gs.PlayerStates[1].Cards = deck[6:10]
	print(len(gs.PlayerStates[1].Cards), len(gs.PlayerStates[1].Cards))
}

//StartGame sets the initial state of the game.
func StartGame() GameState {
	you, enemy := NewPlayerState(0), NewPlayerState(1)
	you.Active = true
	enemy.Active = false
	gs := GameState{
		PlayerStates: []PlayerState{you, enemy},
		TableCards:   make(Deck, 2),
		Round:        0}
	Deal(&gs)

	return gs

}
