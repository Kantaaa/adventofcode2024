package handlers

import (
	"adventofcode2024/solutions"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Templates for rendering HTML pages
var (
	mainTmpl = template.Must(template.ParseFiles("templates/main.html"))
	dayTmpl  = template.Must(template.ParseFiles("templates/day.html"))
)

// MainHandler serves the main page.
func MainHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title string
		Days  []int
	}{
		Title: "Advent of Code 2024",
		Days:  generateDays(1, 25), // Generate days from 1 to 25
	}
	err := mainTmpl.Execute(w, data)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// DayHandler returns an http.HandlerFunc that handles requests for specific days.
// It now dynamically executes the solution for each day.
func DayHandler(day int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		solution := getSolutionForDay(day)

		data := struct {
			Title       string
			Description string
			Solution    string
		}{
			Title:       fmt.Sprintf("Day %d: Advent of Code", day),
			Description: fmt.Sprintf("This is the summary of the Day %d challenge...", day),
			Solution:    solution,
		}
		err := dayTmpl.Execute(w, data)
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// getSolutionForDay returns the solution for a given day.
// It calls the appropriate function from the solutions package.
func getSolutionForDay(day int) string {
	switch day {
	case 1:
		return solutions.SolutionDay1()
	// Add more cases for other days:
	// case 2:
	//     return solutions.SolutionDay2()
	default:
		return "Solution for this day is not available yet."
	}
}

// generateDays creates a slice of integers from start to end.
func generateDays(start, end int) []int {
	days := make([]int, end-start+1)
	for i := range days {
		days[i] = start + i
	}
	return days
}
