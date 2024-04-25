package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

const path = "songs.json"

// Song stores all the song related information
type Song struct {
	Name      string `json:"name"`
	Album     string `json:"album"`
	PlayCount int64  `json:"play_count"`
	// contadores de Album e Song
}

// An PlaylistHeap is a max-heap of PlaylistEntries.

// Implementar as interfaces necessárias para o package
// container/heap
// interessante é que ela usou ponteiro apenas para o
// Push e Pop

// makePlaylist makes the merged sorted list of songs
func makePlaylist(albums [][]Song) []Song {
	// criar o valor playlist
	// ponteiro pHeap
	// verificar se o album tem algo

	// initialize the heap and add first of each album, since they are the max

	// enquanto tiver algo em pHeap
	// take max elem from the list
	// the next song after the max is a good candidate to look at

	// retornar a playlist

	panic("NOT IMPLEMENTED")
}

func main() {
	albums := importData()
	printTable(makePlaylist(albums))
}

// printTable prints merged playlist as a table
func printTable(songs []Song) {
	w := tabwriter.NewWriter(os.Stdout, 3, 3, 3, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "####\tSong\tAlbum\tPlay count")
	for i, s := range songs {
		fmt.Fprintf(w, "[%d]:\t%s\t%s\t%d\n", i+1, s.Name, s.Album, s.PlayCount)
	}
	w.Flush()

}

// importData reads the input data from file and creates the friends map
func importData() [][]Song {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var data [][]Song
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
