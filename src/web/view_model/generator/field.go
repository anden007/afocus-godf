package generator

import "github.com/anden007/afocus-godf/src/web/view_model"

type Field struct {
	view_model.BaseViewModel
	Field           string  `json:"field"`
	Name            string  `json:"name"`
	Level           string  `json:"level"`
	TableShow       bool    `json:"tableShow"`
	SortOrder       float32 `json:"sortOrder"`
	Searchable      bool    `json:"searchable"`
	Editable        bool    `json:"editable"`
	Type            string  `json:"type"`
	Validate        bool    `json:"validate"`
	SearchType      string  `json:"searchType"`
	SearchLevel     string  `json:"searchLevel"`
	Sortable        bool    `json:"sortable"`
	DefaultSort     bool    `json:"defaultSort"`
	DefaultSortType string  `json:"defaultSortType"`
}
