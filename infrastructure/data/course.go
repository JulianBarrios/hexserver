package mooc

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type CourseID struct {
	value string
}

var ErrInvalidCourseID = errors.New("Invalid Course ID")

func NewCourseID(s string) (CourseID, error) {
	v, err := uuid.Parse(s)
	if err != nil {
		return CourseID{}, fmt.Errorf("%w %s", ErrInvalidCourseID, err)
	}
	return CourseID{
		value: v.String(),
	}, nil
}

func (id CourseID) String() string {
	return id.value
}

var ErrEmptyCourseName = errors.New("The Course Name cannot be empty")

type CourseName struct {
	value string
}

func NewCourseName(s string) (CourseName, error) {
	if s == "" {
		return CourseName{}, ErrEmptyCourseName
	}
	return CourseName{
		value: s,
	}, nil
}

func (name CourseName) String() string {
	return name.value
}

var ErrEmptyCourseDuration = errors.New("The Course Duration cannot be empty")

type CourseDuration struct {
	value string
}

func NewCourseDuration(s string) (CourseDuration, error) {
	if s == "" {
		return CourseDuration{}, ErrEmptyCourseDuration
	}
	return CourseDuration{
		value: s,
	}, nil
}

func (duration CourseDuration) String() string {
	return duration.value
}

type CourseRepository interface {
	Save(ctx context.Context, course Course) error
}

type Course struct {
	id       CourseID
	name     CourseName
	duration CourseDuration
}

func NewCourse(id, name, duration string) (Course, error) {

	idVO, err := NewCourseID(id)
	if err != nil {
		return Course{}, err
	}
	nameVO, err := NewCourseName(name)
	if err != nil {
		return Course{}, err
	}

	durationVO, err := NewCourseDuration(duration)
	if err != nil {
		return Course{}, err
	}

	return Course{
		id:       idVO,
		name:     nameVO,
		duration: durationVO,
	}, nil
}

func (c *Course) ID() CourseID {
	return c.id
}

func (c *Course) Name() CourseName {
	return c.name
}

func (c *Course) Duration() CourseDuration {
	return c.duration
}
