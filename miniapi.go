package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// Display current time
func rootHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		hour, minute, _:= time.Now().Clock()
		fmt.Printf("%02dh%02d\n", hour, minute)
		fmt.Fprintf(writer, "%02dh%02d", hour, minute)
	
	default:
		fmt.Printf("Unsupported %s request received for route '/'\n", request.Method)
		fmt.Fprintf(writer, "Method %s is not supported", request.Method)
	}
}

// Roll a dice of 1000 sides
func diceHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		// Roll dice
		dice := rand.Intn(1000) + 1

		fmt.Printf("%04d\n", dice)
		fmt.Fprintf(writer, "%04d", dice)
	
	default:
		fmt.Printf("Unsupported %s request received for route '/dice'\n", request.Method)
		fmt.Fprintf(writer, "Method %s is not supported", request.Method)
	}

}

// Roll 15 dices of a given type in URL parameter
func dicesHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		// Define dice types
		diceTypes := map[string]int{
			"d2": 2,
			"d4": 4,
			"d6": 6,
			"d8": 8,
			"d10": 10,
			"d12": 12,
			"d20": 20,
			"d100": 100,
		}
		var choosedDiceTypes []int

		if request.URL.Query().Get("type") != "" {
			// If type is specified in URL parameter, get it
			choosedDiceTypes = append(choosedDiceTypes, diceTypes[strings.ToLower(request.URL.Query().Get("type"))])
		} else {
			// Otherwise, use all dice types
			for _, diceType := range diceTypes {
				choosedDiceTypes = append(choosedDiceTypes, diceType)
			}
		}

		// Roll dices
		for i := 0; i < 15; i++ {
			diceType := choosedDiceTypes[rand.Intn(len(choosedDiceTypes))]
			dice := rand.Intn(diceType) + 1

			// Print dice with 0 padding, depending on dice type
			fmt.Printf("%0*d ", len(fmt.Sprint(diceType)), dice)
			fmt.Fprintf(writer, "%0*d ", len(fmt.Sprint(diceType)), dice)
		}
		fmt.Println()

	default:
		fmt.Printf("Unsupported %s request received for route '/dices'\n", request.Method)
		fmt.Fprintf(writer, "Method %s is not supported", request.Method)
	}
}

// Shuffles words in a sentence
func randomizeWordsHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodPost:
		// Get sentence from payload
		sentence := request.FormValue("words")

		// Check if sentence is empty
		if sentence == "" {
			fmt.Println("Sentence is empty")
			fmt.Fprint(writer, "Sentence is empty")
			return
		}
		
		// Split sentence into words
		words := strings.Split(sentence, " ")

		// Shuffle words
		rand.Shuffle(len(words), func(i, j int) {
			words[i], words[j] = words[j], words[i]
		})

		// Join words into sentence
		sentence = strings.Join(words, " ")

		fmt.Println(sentence)
		fmt.Fprint(writer, sentence)

	default:
		fmt.Printf("Unsupported %s request received for route '/randomize-words'\n", request.Method)
		fmt.Fprintf(writer, "Method %s is not supported", request.Method)
	}
}

// Capitalize one letter out of two in a sentence
func semiCapitalizeSentenceHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodPost:
		// Get sentence from payload
		sentence := request.FormValue("sentence")

		// Check if sentence is empty
		if sentence == "" {
			fmt.Println("Sentence is empty")
			fmt.Fprint(writer, "Sentence is empty")
			return
		}

		// Capitalize one letter out of two in each word
		for i, letter := range sentence {
			if i % 2 == 0 {
				sentence = sentence[:i] + strings.ToUpper(string(letter)) + sentence[i+1:]
			} else {
				sentence = sentence[:i] + strings.ToLower(string(letter)) + sentence[i+1:]
			}
		}

		fmt.Println(sentence)
		fmt.Fprint(writer, sentence)

	default:
		fmt.Printf("Unsupported %s request received for route '/semi-capitalize-sentence'\n", request.Method)
		fmt.Fprintf(writer, "Method %s is not supported", request.Method)
	}
}

func main() {
	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	// Define routes
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/dice", diceHandler)
	http.HandleFunc("/dices", dicesHandler)
	http.HandleFunc("/randomize-words", randomizeWordsHandler)
	http.HandleFunc("/semi-capitalize-sentence", semiCapitalizeSentenceHandler)

	// Start server
	http.ListenAndServe(":4567", nil)
}
