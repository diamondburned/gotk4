package pen

import (
	"fmt"
	"strings"
)

// Joints is a string builder that joins using a separator.
type Joints struct {
	sep  string
	strs []string
}

// NewJoints creates a new Joints instance.
func NewJoints(sep string, cap int) Joints {
	return Joints{
		sep:  sep,
		strs: make([]string, 0, cap),
	}
}

// Add adds a new joint.
func (j *Joints) Add(str string) { j.strs = append(j.strs, str) }

// Addf adds a new joint with Sprintf.
func (j *Joints) Addf(f string, v ...interface{}) {
	j.Add(fmt.Sprintf(f, v...))
}

// Len returns the length of joints
func (j *Joints) Len() int { return len(j.strs) }

// Join joins the joints.
func (j *Joints) Join() string {
	if j == nil {
		return ""
	}

	return strings.Join(j.strs, j.sep)
}

// Joints returns the list of joints.
func (j *Joints) Joints() []string {
	return j.strs
}
