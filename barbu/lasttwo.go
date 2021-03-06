package barbu

import (
  "github.com/runningwild/barbu/util"
)

type LastTwo struct {
  StandardDoubling
  StandardTrickTaking
}

func makeLastTwo(players []Player, hands [][]string) BarbuGame {
  var lt LastTwo
  lt.StandardDoubling.Players = players[:]
  lt.StandardTrickTaking.Players = players[:]
  for _, hand := range hands {
    lt.StandardTrickTaking.Hands = append(lt.StandardTrickTaking.Hands, util.Hand(hand))
  }
  return &lt
}
func init() {
  RegisterBarbuGame("lasttwo", makeLastTwo)
}

func (lt *LastTwo) Scores() [4]int {
  var ret [4]int
  tricks := lt.StandardTrickTaking.Tricks
  ret[tricks[len(tricks)-2].Took] -= 10
  ret[tricks[len(tricks)-1].Took] -= 20
  return ret
}
