package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("freezing-window")
	w2 := myApp.NewWindow("other-window")

	var tab *widget.Table
	tab = widget.NewTable(
		//length
		func() (rows int, cols int) {
			return 100, 100
		},
		//create
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		//update
		func(id widget.TableCellID, o fyne.CanvasObject) {
			l := o.(*widget.Label)
			l.SetText(fmt.Sprintf("(%d,%d)", id.Col, id.Row))
			if l.Size().Width < l.MinSize().Width {
				tab.SetColumnWidth(id.Col, l.MinSize().Width)
			}
		},
	)
	var d dialog.Dialog
	tab.OnSelected = func(id widget.TableCellID) {
		// We make this dialog part of w rather than w2, since making it part of w2 makes the bug disappear.
		if d != nil {
			d.Hide()
		}
		d = dialog.NewInformation("clicked", fmt.Sprintf("SELECTED %v\n", id), w2)
		d.Show()
	}

	l1 := widget.NewLabel("Do not interact with the table window, including mousing over it while the timer counts down.")
	l1.Wrapping = fyne.TextWrapWord
	l2 := widget.NewLabel("")
	l2.Wrapping = fyne.TextWrapWord
	t := time.Second * 120
	go func() {
		for t >= 0 {
			l2.SetText(fmt.Sprintf("%v", t))
			l2.Refresh()
			t = t - 1*time.Second
			time.Sleep(1 * time.Second)
		}
		l2.SetText(fmt.Sprintf("Now, mouse over the table window, try to scroll, etc. It is frozen. This can be corrected by resizing the window. Clicks still appear to work."))
		for {
			l2.Refresh()
			time.Sleep(1 * time.Second)
		}

	}()

	w.SetContent(tab)
	w.Resize(fyne.NewSize(400, 400))
	w2.SetContent(container.NewVBox(
		l1,
		l2,
		// This does nothing except provide correct width so the container sets min height correctly.
		widget.NewButton("I do nothing.", nil),
	))
	w2.Resize(fyne.NewSize(200, 200))
	w.Show()
	w2.Show()
	myApp.Run()
}
