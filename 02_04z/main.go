package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// the amount of bidders we have at our auction
const bidderCount = 10

// initial wallet value for all bidders
const walletAmount = 250

// items is the map of auction items
var items = []string{
	"The \"Best Gopher\" trophy",
	"The \"Learn Go with Adelina\" experience",
	"Two tickets to a Go conference",
	"Signed copy of \"Beautiful Go code\"",
	"Vintage Gopher plushie",
}

// bid is a type that pairs the bidder id and the amount they want to bid
type bid struct {
	bidderID string
	amount   int
}

// auctioneer receives bids and announces winners
type auctioneer struct {
	bidders map[string]*bidder
}

// runAuction and manages the auction for all the items to be sold
// Change the signature of this function as required
// ==== recebe bids e open channels
func (a *auctioneer) runAuction(bids chan bid, open chan struct{}) {
	for _, item := range items {
		log.Printf("Opening bids for %s!\n", item)
		// chamar o método do auctioneer para abrir bid (escrever esse método auxiliar)
		a.openBids(open)
		// chamar o método do auctioneer para processar o vencedor (escrever esse método auxiliar)
		a.processWinner(item, bids)
	}
}

// openBids signals to all bidders that the auction for the item is open.
func (a *auctioneer) openBids(open chan struct{}) {
	//== loop em bidderCount, sinalizando no canal "open" que é recebido via parâmetro
	for i := 0; i < bidderCount; i++ {
		open <- struct{}{}
	}
}

// processWinner reads all the bids and finds the winner of the auction
// == como parâmetros são recebidos item e bids (via canal)
func (a *auctioneer) processWinner(item string, bids chan bid) {
	//== fazer o loop em bidderCount para buscar o winner
	var winner bid
	for i := 0; i < bidderCount; i++ {
		currentBid := <-bids
		log.Println("processWinner", currentBid.bidderID)
		if currentBid.amount > winner.amount {
			winner = currentBid
		}
	}
	//== enviar a mensagem informando: "item" is sold to "bitter" for "amount"
	log.Printf("%s is sold to %s for %d\n",
		item, winner.bidderID, winner.amount)
	//== dentro do objeto auctioneer temos a lista de bidders, portanto chamar o método payBid do com o id do winner
	a.bidders[winner.bidderID].payBid(winner.amount)
}

// bidder is a type that holds the bidder id and wallet
type bidder struct {
	id     string
	wallet int
}

// placeBid generates a random amount and places it on the bids channels
// Change the signature of this function as required
// == receber canais bid para output e open (sinalizador)
func (b *bidder) placeBid(output chan<- bid, open <-chan struct{}) {
	// loop na quantidade de items
	for i := 0; i < len(items); i++ {
		//== aguardar o sinal que a bid está open - canal open -- struct vazia
		log.Printf("I am %s. Waiting to place my bid.", b.id)
		<-open
		log.Printf("I am %s. Got the signal.", b.id)
		//== inicializar currentBid com o id
		currentBid := bid{
			bidderID: b.id,
		}
		//== se o valor que o bidder tiver na waller for maior que zero, pegar o valor random sendo o máximo que ele tem
		if b.wallet > 0 {
			currentBid.amount = getRandomAmount(b.wallet)
		}
		// colocar o currentBid no canal de saída
		output <- currentBid
	}
}

// payBid subtracts the bid amount from the wallet of the auction winner
func (b *bidder) payBid(amount int) {
	b.wallet -= amount
}

func main() {
	rand.Seed(time.Now().UnixNano())
	log.Println("Welcome to the LinkedIn Learning auction.")
	bidders := make(map[string]*bidder, bidderCount)
	// canal "bids" com tamanho de bidderCount (tipo bid) - buffered - não precisa esperar
	bids := make(chan bid, bidderCount)
	// canal "openBids" somente para sinalização
	openBids := make(chan struct{})
	for i := 0; i < bidderCount; i++ {
		id := fmt.Sprint("Bidder ", i)
		b := bidder{
			id:     id,
			wallet: walletAmount,
		}
		bidders[id] = &b
		// passar os dois canais para placeBid(), sendo esses canais bids e openBids
		go b.placeBid(bids, openBids)
	}
	a := auctioneer{
		bidders: bidders,
	}
	// passar os dois canais para placeBid(), sendo esses canais bids e openBids
	a.runAuction(bids, openBids)
	log.Println("The LinkedIn Learning auction has finished!")
}

// getRandomAmount generates a random integer amount up to max
func getRandomAmount(max int) int {
	return rand.Intn(int(max))
}
