package main

import (
    "time"
	"fmt"
    "math/rand"
    "bufio"
    "os"
)

type Card struct{
	name string
	suit string
	value int
    faceup bool
    
}

type Deck struct{
	deck []Card
}

func waitOnInput(){
	fmt.Print("Press 'Enter' to continue...")
    bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func (d *Deck) build(num_deck int){
	if(num_deck < 1 || num_deck > 8){
		fmt.Println("Invalid number of decks. Creating single deck game.")
		num_deck = 1
		}
    suits := []string{"Hearts","Diamonds","Spades","Clubs"}
	card_typ := []string{"Ace","2","3","4","5","6","7","8","9","10","Jack","Queen","King"}
	card_val := []int{11,2,3,4,5,6,7,8,9,10,10,10,10}
	num_card := 52 * num_deck
//	card_counter := 4 * num_deck
	card_create := 13
	d.deck = make([]Card, num_card)
	for(num_card > 0){
        for card_counter:= (4 * num_deck); card_counter > 0; card_counter--{
			d.deck[num_card-1] = Card{value: card_val[card_create-1], suit: suits[card_counter%4], name: card_typ[card_create-1], faceup: false}
//			d.deck[num_card-1] = tempCard
			num_card--		
		}
		card_create--
	}
}

func (d *Deck) fy_shuffle(){
    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
     
    for i:=len(d.deck) - 1; i > 0; i--{
        // Pick a random index from 0 to i
        j := r1.Intn(200000) % (i+1);
 
        // Swap arr[i] with the element at random index
        d.deck[i], d.deck[j] = d.deck[j], d.deck[i]
    }
        
}

func (d *Deck) draw() (Card){
    tempCard := d.deck[len(d.deck)-1]
    d.deck = d.deck[:len(d.deck)-1]
    
    return tempCard
}

type Hand struct{
	hand []Card
	score int
	numAce int
}

func (h *Hand) clear(){
    var newHand []Card
    newHand = make([]Card,0)
    h.hand = newHand
    h.numAce = 0
}

func (h *Hand) scoreUp(){
    sum:=0
    for i:=0;i<len(h.hand);i++{
    	if(h.hand[i].name == "Ace"){
    		h.numAce++
    	}
        sum = sum + h.hand[i].value
    }
    for (sum > 21 && h.numAce > 0) {
    	sum = sum - 10
    	h.numAce--
    }
    h.score = sum
}

func (h *Hand) display(){
    fmt.Println("Hand Contains:")
    for i:=0;i<len(h.hand);i++{
        fmt.Println(h.hand[i].name, "of", h.hand[i].suit)
    }
    fmt.Println("")
}

type Dealer struct{
    hand Hand
    deck Deck
}

type Player struct{
	hand Hand
//	money int
}

func (GM *Dealer) giveCard(hand *Hand){
    tempCard := GM.deck.draw()
//    fmt.Println("Card drawn is, ",tempCard)
    hand.hand = append(hand.hand, tempCard)
//    fmt.Println("Hand now contains: ",hand.hand)
}

func (GM *Dealer) newRound(p1 *Player){
    GM.hand.clear()
    p1.hand.clear()
    GM.giveCard(&p1.hand)
    GM.giveCard(&GM.hand)
    GM.giveCard(&p1.hand)
    GM.giveCard(&GM.hand)
    GM.hand.hand[len(GM.hand.hand)-1].faceup = true
    fmt.Println("Dealer shows: ",GM.hand.hand[1].name, "of", GM.hand.hand[1].suit)
}

func (p *Player) hit(GM *Dealer) (int){
    GM.giveCard(&p.hand)
    p.hand.scoreUp()
    if p.hand.score > 21{
        fmt.Println("You went Bust. Dealer wins round.")
        return -1
    }
    return p.hand.score
}

func (GM *Dealer) playHand(p1 *Player){
        roundFlag := 0	//set if player goes bust
        GM.newRound(p1)
        p1.hand.scoreUp()
        GM.hand.scoreUp()
        if(p1.hand.score == 21){
        	if(GM.hand.score < 21){
				fmt.Println("BLACK JACK! PLAYER WINS!!!!")
				fmt.Println("")
				fmt.Println("******************")
				fmt.Println("End of round")
				fmt.Println("******************")
				fmt.Println("")
				return
			}else{
				fmt.Println("Double Black Jack, dealer wins draw.")
				fmt.Println("")
				fmt.Println("******************")
				fmt.Println("End of round")
				fmt.Println("******************")
				fmt.Println("")
				return
			}
        }
        if(GM.hand.score == 21){
        	fmt.Println("Dealer has Black Jack, dealer wins.")
        	fmt.Println("")
			fmt.Println("******************")
			fmt.Println("End of round")
			fmt.Println("******************")
			fmt.Println("")
        	return
        }
        fmt.Println("Players score is: ",p1.hand.score)
        p1.hand.display()
        //fmt.Println("Dealers score is: ",dealer.hand.score)
        stand := 1 
        for stand == 1{
            fmt.Println("Would you like to hit? enter 1 to hit or 0 to stand")
            fmt.Scanln(&stand)
            if(stand == 1){
                round_score := p1.hit(GM)
                if(round_score == -1){
                    fmt.Println("Player total is: ", p1.hand.score)
                    p1.hand.display()
                    roundFlag = 1 
                    fmt.Print("Dealer's ")
                    GM.hand.display()
                    fmt.Println("Dealer has won by default.")                    
                    break
                }
                p1.hand.scoreUp()
                fmt.Println("Players hand now has score of: ", p1.hand.score)
                p1.hand.display()
                if(p1.hand.score == 21){
                	break
                }
                if(len(p1.hand.hand) == 5){
                	fmt.Println("Player has won by default with 5 cards!")
                	roundFlag = 1
                	fmt.Print("Dealer's ")
                    GM.hand.display()
                	break
                }
            }else if(stand != 0){
                fmt.Println("A wrong value was entered")
                stand = 1
            }else{
                fmt.Println("Player stands with: ", p1.hand.score)
                
            }
            
        }
    if(roundFlag == 1){
		fmt.Println("******************")
		fmt.Println("End of round")
		fmt.Println("******************")
		fmt.Println("")
	}else{
		fmt.Println("")
		fmt.Println("******************")
		fmt.Println("End of player round")
		fmt.Println("******************")
		fmt.Println("")
		waitOnInput()
		fmt.Println("Dealer's hand is: ", GM.hand.score)
		GM.hand.display()
		for GM.hand.score < 15 {
			fmt.Println("Dealer will hit")
			GM.giveCard(&GM.hand)
			GM.hand.scoreUp()
			fmt.Println("Dealer score is now: ",GM.hand.score)
			GM.hand.display()
			waitOnInput()
		}
		if(GM.hand.score > 21){
			fmt.Println("Dealer has gone bust with score: ", GM.hand.score)
			GM.hand.display()
			fmt.Println("You have won by default.")
		}else{
			if(p1.hand.score > GM.hand.score){
				fmt.Println("Dealers hand is: ")
				GM.hand.display()
				fmt.Println("You have won ", p1.hand.score,"to", GM.hand.score,"! ...this round")
			}else if(p1.hand.score == GM.hand.score){
				fmt.Println("Dealers hand is: ")
				GM.hand.display()
				fmt.Println("Dealer wins ties!")
			}else{
				fmt.Println("Dealers hand is: ")
				GM.hand.display()
				fmt.Println("The Dealer has outscored you! ", GM.hand.score, " to your ", p1.hand.score)
			}
		}
		fmt.Println("")
		fmt.Println("******************")
		fmt.Println("End of round")
		fmt.Println("******************")
		fmt.Println("")
	}
}

func main() {
	var numDeck int
	fmt.Println("Welcome to the Black Jack table!")
	fmt.Println("Please enter a number of decks between 1 and 8")
	fmt.Scanln(&numDeck)
	var deck Deck
	deck.build(numDeck)
//	fmt.Println("deck is:", deck)
    deck.fy_shuffle()
//  fmt.Println("shuffled deck is:", deck)
    var p1 Player
    var dealer Dealer
    dealer = Dealer{deck: deck}
    fmt.Println("")
    fmt.Println("The game is prepared! Lets Play")
    fmt.Println("")
    for len(dealer.deck.deck) > 10{
        dealer.playHand(&p1)
        duration := time.Second
        time.Sleep(duration)
        fmt.Println("Would you like to Continue? enter quit to stop, or anything else to continue")
        var cont string
        fmt.Scanln(&cont)
        if cont == "quit"{
        	break
        }
    }
    fmt.Println("Thank you for playing Blackjack!")
    
}
