package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

// GameState holds the score and remaining time
type GameState struct {
	Score        int
	TimeRemaining int
}

// This will serve the main game page
func gamePage(w http.ResponseWriter, r *http.Request) {
	// Set the initial game state (score 0, 30 seconds remaining)
	state := GameState{Score: 0, TimeRemaining: 30}
	tmpl, err := template.New("game").Parse(gameHTML)
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}

	// Render the template with the current state
	tmpl.Execute(w, state)
}

// This will handle the score update when the button is clicked
func clickHandler(w http.ResponseWriter, r *http.Request) {
	// Decrease time remaining with each click
	state := GameState{Score: 0, TimeRemaining: 30}
	state.Score = state.Score + 1 // Increment score for each click

	// Send the updated score and time back to the frontend
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"score": %d, "timeRemaining": %d}`, state.Score, state.TimeRemaining)
}

// HTML template for the game
var gameHTML = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Click Counter Game</title>
	<style>
		body { font-family: Arial, sans-serif; text-align: center; }
		#game-over { color: red; display: none; }
	</style>
</head>
<body>
	<h1>Click Counter Game</h1>
	<p>Click the button as many times as you can before time runs out!</p>
	<p>Time Remaining: <span id="time">{{.TimeRemaining}}</span> seconds</p>
	<p>Your Score: <span id="score">{{.Score}}</span></p>
	<button id="clickButton">Click Me!</button>
	<p id="game-over">Game Over! Your final score is: <span id="final-score"></span></p>

	<script>
		// Setup JavaScript to interact with Go backend
		let score = {{.Score}};
		let timeRemaining = {{.TimeRemaining}};
		let gameOver = false;

		document.getElementById('clickButton').onclick = function() {
			if (gameOver) return; // Disable button after game over
			score++;
			timeRemaining--;

			// Update the UI with the new score and time
			document.getElementById('score').innerText = score;
			document.getElementById('time').innerText = timeRemaining;

			// If time runs out, show game over
			if (timeRemaining <= 0) {
				gameOver = true;
				document.getElementById('game-over').style.display = 'block';
				document.getElementById('final-score').innerText = score;
				document.getElementById('clickButton').disabled = true;
			}
		};

		// Start the countdown timer
		setInterval(function() {
			if (!gameOver && timeRemaining > 0) {
				timeRemaining--;
				document.getElementById('time').innerText = timeRemaining;
			}
		}, 1000);
	</script>
</body>
</html>
`

func main() {
	// Serve the game page at the root URL
	http.HandleFunc("/", gamePage)
	http.HandleFunc("/click", clickHandler)

	// Start the web server
	fmt.Println("Starting game on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
