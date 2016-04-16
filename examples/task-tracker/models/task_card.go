package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*TaskCard a card for a task

A task card is a minimalistic representation of a task. Useful for display in list views, like a card list.


swagger:model TaskCard
*/
type TaskCard struct {

	/* assigned to
	 */
	AssignedTo *UserCard `json:"assignedTo,omitempty"`

	/* The description of the task.

	The task description is a longer, more detailed description of the issue.
	Perhaps it even mentions steps to reproduce.

	*/
	Description string `json:"description,omitempty"`

	/* the level of effort required to get this task completed

	Maximum: 27
	Multiple Of: 3
	*/
	Effort int32 `json:"effort,omitempty"`

	/* The id of the task.

	A unique identifier for the task. These are created in ascending order.

	Read Only: true
	*/
	ID int64 `json:"id,omitempty"`

	/* the karma donated to this item.

	Karma is a lot like voting.  Users can donate a certain amount or karma to an issue.
	This is used to determine the weight users place on an issue. Not that +1 comments aren't great.


	Minimum: > 0
	Multiple Of: 0.5
	*/
	Karma *float64 `json:"karma,omitempty"`

	/* milestone
	 */
	Milestone *Milestone `json:"milestone,omitempty"`

	/* The time at which this issue was reported.

	This field is read-only, so it's only sent as part of the response.


	Read Only: true
	*/
	ReportedAt strfmt.DateTime `json:"reportedAt,omitempty"`

	/* severity

	Maximum: 5
	Minimum: 1
	*/
	Severity int32 `json:"severity,omitempty"`

	/* the status of the issue

	There are 4 possible values for a status.
	Ignored means as much as accepted but not now, perhaps later.


	Required: true
	*/
	Status *string `json:"status"`

	/* task tags.

	a task can be tagged with text blurbs.

	Max Items: 5
	Unique: true
	*/
	Tags []string `json:"tags,omitempty"`

	/* The title of the task.

	The title for a task, this needs to be at least 5 chars long.
	Titles don't allow any formatting, besides emoji.


	Required: true
	Max Length: 150
	Min Length: 5
	*/
	Title *string `json:"title"`
}

// Validate validates this task card
func (m *TaskCard) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEffort(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateKarma(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSeverity(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskCard) validateEffort(formats strfmt.Registry) error {

	if swag.IsZero(m.Effort) { // not required
		return nil
	}

	if err := validate.MaximumInt("effort", "body", int64(m.Effort), 27, false); err != nil {
		return err
	}

	if err := validate.MultipleOf("effort", "body", float64(m.Effort), 3); err != nil {
		return err
	}

	return nil
}

func (m *TaskCard) validateKarma(formats strfmt.Registry) error {

	if swag.IsZero(m.Karma) { // not required
		return nil
	}

	if err := validate.Minimum("karma", "body", float64(*m.Karma), 0, true); err != nil {
		return err
	}

	if err := validate.MultipleOf("karma", "body", float64(*m.Karma), 0.5); err != nil {
		return err
	}

	return nil
}

func (m *TaskCard) validateSeverity(formats strfmt.Registry) error {

	if swag.IsZero(m.Severity) { // not required
		return nil
	}

	if err := validate.MinimumInt("severity", "body", int64(m.Severity), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("severity", "body", int64(m.Severity), 5, false); err != nil {
		return err
	}

	return nil
}

var taskCardTypeStatusPropEnum []interface{}

// prop value enum
func (m *TaskCard) validateStatusEnum(path, location string, value string) error {
	if taskCardTypeStatusPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["open","closed","ignored","rejected"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			taskCardTypeStatusPropEnum = append(taskCardTypeStatusPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, taskCardTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TaskCard) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *TaskCard) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	iTagsSize := int64(len(m.Tags))

	if err := validate.MaxItems("tags", "body", iTagsSize, 5); err != nil {
		return err
	}

	if err := validate.UniqueItems("tags", "body", m.Tags); err != nil {
		return err
	}

	for i := 0; i < len(m.Tags); i++ {

		if err := validate.MinLength("tags"+"."+strconv.Itoa(i), "body", string(m.Tags[i]), 3); err != nil {
			return err
		}

		if err := validate.Pattern("tags"+"."+strconv.Itoa(i), "body", string(m.Tags[i]), `\w[\w- ]+`); err != nil {
			return err
		}

	}

	return nil
}

func (m *TaskCard) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	if err := validate.MinLength("title", "body", string(*m.Title), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("title", "body", string(*m.Title), 150); err != nil {
		return err
	}

	return nil
}
