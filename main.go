package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Create a new app and window
	chordFinder := app.New()
	myWindow := chordFinder.NewWindow("Ukulele Chord Diagrams")

	// Define the chords
	chords := []string{"C", "G", "F", "D", "Am", "A", "Dm", "G7", "Em", "E7", "A7", "C7"}

	// Create a dropdown widget for chords
	chordDropdown := widget.NewSelect(chords, func(selectedChord string) {
		// Load and display the corresponding chord diagram image
		chordImage := loadChordImage(selectedChord)
		myWindow.SetContent(container.NewVBox(
			chordImage,
		))
	})

	// Set the default chord selection
	chordDropdown.SetSelected(chords[0])

	// Load the default chord diagram image
	defaultChordImage := loadChordImage(chords[0])

	// Set the initial content of the window
	myWindow.SetContent(container.NewVBox(
		chordDropdown,
		defaultChordImage,
	))

	// Show the window
	myWindow.ShowAndRun()
}

// Load chord diagram image based on the selected chord
func loadChordImage(chord string) *canvas.Image {
	var imagePath string
	switch chord {
	case "C":
		imagePath = "chord_images/c_chord.png"
	case "G":
		imagePath = "chord_images/g_chord.png"
	case "F":
		imagePath = "chord_images/f_chord.png"
	case "D":
		imagePath = "chord_images/d_chord.png"
	case "Am":
		imagePath = "chord_images/am_chord.png"
	case "A":
		imagePath = "chord_images/a_chord.png"
	case "Dm":
		imagePath = "chord_images/dm_chord.png"
	case "G7":
		imagePath = "chord_images/g7_chord.png"
	case "Em":
		imagePath = "chord_images/em_chord.png"
	case "E7":
		imagePath = "chord_images/e7_chord.png"
	case "A7":
		imagePath = "chord_images/a7_chord.png"
	case "C7":
		imagePath = "chord_images/c7_chord.png"
	default:
		// Default image
		imagePath = "chord_images/treble_clef.png"
	}

	// Load the chord image from the file
	chordImage := canvas.NewImageFromFile(imagePath)
	return chordImage
}
