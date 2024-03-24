package formcontroller

import (
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

func RouteSubmitGet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("views/form.html"))
		err := tmpl.ExecuteTemplate(w, "form.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func RouteSubmitPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		name := r.FormValue("name")
		name = capitalize(name)
		address := r.FormValue("address")
		address = capitalize(address)

		//PPT
		gradePointPpt := r.FormValue("ppt")
		pointPpt, err := strconv.Atoi(gradePointPpt)
		gradePpt := grading(pointPpt)

		//Web Prog
		gradePointWebprog := r.FormValue("webprog")
		pointWebprog, err := strconv.Atoi(gradePointWebprog)
		gradeWebprog := grading(pointWebprog)

		//Software Engineering
		gradePointSofteng := r.FormValue("softeng")
		pointSofteng, err := strconv.Atoi(gradePointSofteng)
		gradeSofteng := grading(pointSofteng)

		data := map[string]any{
			"name":    name,
			"address": address,

			//PPT
			"gradePointPpt": gradePointPpt,
			"gradePpt":      gradePpt,

			//Web Prog
			"gradePointWebprog": gradePointWebprog,
			"gradeWebprog":      gradeWebprog,

			//Soft Eng
			"gradePointSofteng": gradePointSofteng,
			"gradeSofteng":      gradeSofteng,
		}

		tmpl := template.Must(template.ParseFiles("views/result.html"))
		err = tmpl.ExecuteTemplate(w, "result.html", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func grading(x int) string {
	if x <= 100 && x >= 90 {
		return "A+"
	} else if x < 90 && x >= 80 {
		return "A"
	} else if x < 80 && x >= 70 {
		return "B+"
	} else if x < 70 && x >= 60 {
		return "B"
	} else if x < 60 && x >= 55 {
		return "C+"
	} else if x < 55 && x >= 45 {
		return "C"
	} else if x < 45 && x >= 40 {
		return "D"
	} else if x < 40 && x >= 0 {
		return "F"
	} else {
		return "Invalid Grade"
	}
}

func capitalize(s string) string {
	// Split string menjadi slice kata-kata
	words := strings.Fields(s)

	// Lakukan loop pada setiap kata
	for i, word := range words {
		// Konversi huruf pertama dari setiap kata menjadi kapital
		words[i] = strings.ToUpper(word[:1]) + word[1:]
	}

	// Gabungkan kembali kata-kata menjadi satu string
	return strings.Join(words, " ")
}
