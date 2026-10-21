package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

// TitleMaxLength es el número máximo de caracteres permitido en el título.
const TitleMaxLength = 200

// DueDateLayout es el formato de una fecha límite: el de <input type="date"> y el del
// atributo datetime de <time>.
const DueDateLayout = "2006-01-02"

// Task representa una tarea de la lista.
type Task struct {
	ID        uuid.UUID  `json:"id" db:"id" form:"-"`
	Title     string     `json:"title" db:"title" form:"title"`
	Completed bool       `json:"completed" db:"completed" form:"-"`
	DueDate   *time.Time `json:"due_date" db:"due_date" form:"-"`
	CreatedAt time.Time  `json:"created_at" db:"created_at" form:"-"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at" form:"-"`

	// dueDateInvalid recuerda que la última entrada de SetDueDate no era una fecha.
	dueDateInvalid bool `db:"-"`
}

// Tasks es una colección de Task.
type Tasks []Task

// Validate se ejecuta en cada llamada a los métodos "pop.Validate*".
func (t *Task) Validate(tx *pop.Connection) (*validate.Errors, error) {
	verrs := validate.NewErrors()

	title := strings.TrimSpace(t.Title)
	if title == "" {
		verrs.Add("title", "El título es obligatorio.")
	} else if len([]rune(title)) > TitleMaxLength {
		verrs.Add("title", fmt.Sprintf("El título no puede superar los %d caracteres.", TitleMaxLength))
	}

	if t.dueDateInvalid {
		verrs.Add("due_date", "La fecha límite no es válida.")
	}

	return verrs, nil
}

// SetDueDate interpreta el texto de un formulario como fecha límite. Un texto vacío
// quita la fecha; uno que no es una fecha válida la quita y hace que Validate falle.
func (t *Task) SetDueDate(raw string) {
	raw = strings.TrimSpace(raw)
	t.dueDateInvalid = false
	if raw == "" {
		t.DueDate = nil
		return
	}

	d, err := time.Parse(DueDateLayout, raw)
	if err != nil || d.Year() < 1 {
		t.DueDate = nil
		t.dueDateInvalid = true
		return
	}
	t.DueDate = &d
}

// DueDateValue devuelve la fecha límite como aaaa-mm-dd, o "" si no tiene.
func (t Task) DueDateValue() string {
	if t.DueDate == nil {
		return ""
	}
	return t.DueDate.Format(DueDateLayout)
}

// DueDateLabel devuelve la fecha límite como dd/mm/aaaa, o "" si no tiene.
func (t Task) DueDateLabel() string {
	if t.DueDate == nil {
		return ""
	}
	return t.DueDate.Format("02/01/2006")
}
