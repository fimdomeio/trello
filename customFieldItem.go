// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package trello

//CustomFieldItem for card custom fields
type CustomFieldItem struct {
	ID            string           `json:"id"`
	Value         CustomFieldValue `json:"value"`
	IDCustomField string           `json:"idCustomField"`
	IDModel       string           `json:"idModel"`
	ModelType     string           `json:"modelType"`
}

//CustomFieldValue for custom field value
type CustomFieldValue struct {
	Text   string  `json:"text,omitempty"`
	Number float32 `json:"number,omitempty"`
	Date   string  `json:"checked,omitempty"`
}
