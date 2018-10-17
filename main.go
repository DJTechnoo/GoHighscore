package main

import (
	//"fmt"
	"net/http"
	"encoding/json"
	"google.golang.org/appengine"
	"strconv"
	"sort"

)


//	High score and 4 recent scores

type Score struct {
	Highscore int 	`json:"highscore"`	// Store the highest score
	Scores []int 	`json:"scores"`		// Store other scores
}




func scoreHandler(w http.ResponseWriter, r * http.Request){

	switch r.Method {
		case http.MethodGet:						// Requesting the scores
			json.NewEncoder(w).Encode(scoreData)
		case http.MethodPost:						// Adding a score
			if err := r.ParseForm(); err != nil {
				return
			}

			
			//	add newest score
			freshScore, _ := strconv.Atoi(r.FormValue("score"))	
			addScore(freshScore)
					
	}

}




/*	Description: adds newest score and sorts accordingly
 *	@args "s" is the newest score
 	@return None
*/
func addScore(s int){
	scoreData.Scores = append(scoreData.Scores, s)
			
			// TODO: check if there is a score to check with
			sort.Slice(scoreData.Scores, func(i, j int) bool { return scoreData.Scores[i] > scoreData.Scores[j] })
			if scoreData.Scores[0] > scoreData.Highscore {
				tmp := scoreData.Highscore
				scoreData.Highscore = scoreData.Scores[0]
				scoreData.Scores[0] = tmp
			}
			
}


func p(w http.ResponseWriter, r * http.Request){
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
					return
				}

			
		//	add newest score
		freshScore, _ := strconv.Atoi(r.FormValue("score"))	
		addScore(freshScore)
	
	}else {
		http.ServeFile(w, r, "form.html")
	}
}



var scoreData = Score{		// Test data
					3000,
					[]int{22,55,80}}



func main() {
	http.HandleFunc("/api", scoreHandler)
	http.HandleFunc("/p", p)

	appengine.Main()

}
