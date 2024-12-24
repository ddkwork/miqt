package main

import (
	"os"

	"github.com/mappu/miqt/qt"
)

func createIntroPage() *qt.QWizardPage {
	page := qt.NewQWizardPage2()
	page.SetTitle("Introduction")

	label := qt.NewQLabel6("This wizard will help you register your copy "+
		"of Super Product Two.", page.QWidget, 0)
	label.SetWordWrap(true)

	layout := qt.NewQVBoxLayout2()
	layout.AddWidget3(label.QWidget, 0, 0)
	page.QWidget.SetLayout(layout.QBoxLayout.QLayout)

	return page
}

func createRegistrationPage() *qt.QWizardPage {
	page := qt.NewQWizardPage2()
	page.SetTitle("Registration")
	page.SetSubTitle("Please fill both fields.")

	nameLabel := qt.NewQLabel6("Name:", page.QWidget, 0)
	nameLineEdit := qt.NewQLineEdit(page.QWidget)

	emailLabel := qt.NewQLabel6("Email address:", page.QWidget, 0)
	emailLineEdit := qt.NewQLineEdit(page.QWidget)

	layout := qt.NewQGridLayout(page.QWidget)
	layout.AddWidget2(nameLabel.QFrame.QWidget, 0, 0)
	layout.AddWidget2(nameLineEdit.QWidget, 0, 1)
	layout.AddWidget2(emailLabel.QFrame.QWidget, 1, 0)
	layout.AddWidget2(emailLineEdit.QWidget, 1, 1)
	page.SetLayout(layout.QLayout)

	return page
}

func createConclusionPage() *qt.QWizardPage {
	page := qt.NewQWizardPage2()
	page.SetTitle("Conclusion")

	label := qt.NewQLabel6("You are now successfully registered. Have a "+
		"nice day!", page.QWidget, 0)
	label.SetWordWrap(true)

	layout := qt.NewQVBoxLayout2()
	layout.AddWidget3(label.QWidget, 0, 0)
	page.SetLayout(layout.QBoxLayout.QLayout)

	return page
}

func main() {
	qt.NewQApplication(os.Args)

	wizard := qt.NewQWizard2()
	wizard.AddPage(createIntroPage())
	wizard.AddPage(createRegistrationPage())
	wizard.AddPage(createConclusionPage())

	wizard.QWidget.SetWindowTitle("Trivial Wizard")
	wizard.QWidget.Show()

	qt.QApplication_Exec()
}
