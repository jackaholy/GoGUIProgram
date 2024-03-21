package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Create a new app and window
	chordFinder := app.New()
	window := chordFinder.NewWindow("Ukulele Chord Diagrams")

	// Reset the size of the window
	window.Resize(fyne.NewSize(800, 400))

	// Define the chords
	chords := []string{"C", "G", "F", "D", "Am", "A", "Dm", "G7", "Em", "E7", "A7", "C7"}

	// Load the default chord diagram image
	defaultChordImage := loadChordImage(chords[0])
	defaultChordImage.FillMode = canvas.ImageFillContain
	defaultChordImage.SetMinSize(fyne.NewSize(400, 300))

	// Create a container for the image
	imageContainer := container.New(layout.NewCenterLayout(), defaultChordImage)

	// Create a dropdown widget for chords
	chordDropdown := widget.NewSelect(chords, func(selectedChord string) {
		// Load and display the corresponding chord diagram image
		chordImage := loadChordImage(selectedChord)
		chordImage.FillMode = canvas.ImageFillContain
		chordImage.SetMinSize(fyne.NewSize(400, 300)) // Adjust size as needed
		imageContainer.Objects = []fyne.CanvasObject{chordImage}
		imageContainer.Refresh()
	})

	// Set the default chord selection
	chordDropdown.SetSelected(chords[0])

	// Create a container for the dropdown
	dropdownContainer := container.New(layout.NewVBoxLayout(), chordDropdown)
	dropdownContainer.Resize(fyne.NewSize(200, 400)) // Set size as needed

	// Create a main container with horizontal layout
	mainContainer := container.New(layout.NewHBoxLayout(), dropdownContainer, imageContainer)

	// Set the initial content of the window
	window.SetContent(mainContainer)

	// Show the window
	window.ShowAndRun()
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
