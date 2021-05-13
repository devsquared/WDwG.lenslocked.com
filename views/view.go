package views

import "html/template"

/*
 NewView is the constructor to be used for the View struct.
*/
func NewView(files ...string) *View {
	files = append(files, "views/layouts/footer.gohtml") //TODO: hardcoded for now

	t, err := template.ParseFiles(files...)
	if err != nil {
		// panic is fine assuming this is used on setup and not runtime
		panic(err)
	}

	return &View{
		Template: t,
	}
}

type View struct {
	Template *template.Template
}
