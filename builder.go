package filterparams

import (
	"fmt"

	"github.com/cbrand/go-filterparams/definition"
)


// QueryBuilder allows the configuration of a query which can represent
// query parameters.
type QueryBuilder struct {
	filters []*definition.Filter
}

// EnableFilter allows a filter to be registered against the query builder.
func (q *QueryBuilder) EnableFilter(filter *definition.Filter) {
	q.filters = append(q.filters, filter)
}

// RemoveFilters removes all filters.
func (q *QueryBuilder) RemoveFilters() {
	q.filters = []*definition.Filter{}
}

// HasFilter returns if the filter with the given Identification is
// registered
func (q *QueryBuilder) HasFilter(filterName string) bool {
	return q.filterIndexOf(filterName) != -1
}

// GetFilter returns the filter with the given name if it exists. Returns an error if none
// is present.
func (q *QueryBuilder) GetFilter(filterName string) (*definition.Filter, error) {
	index := q.filterIndexOf(filterName)
	if index == -1 {
		return nil, fmt.Errorf("Filter %s does not exist.", filterName)
	}
	return q.filters[index], nil
}

// filterIndexOf returns the index of the given filterName or -1 if none exists.
func (q *QueryBuilder) filterIndexOf(filterName string) int {
	for index, filter := range q.filters {
		if filter.Identification == filterName {
			return index
		}
	}
	return -1
}

// CreateQuery initializes a new Query and returns it.
func (q *QueryBuilder) CreateQuery() *Query {
	return newQuery(q.filters)
}

// NewBuilder initializes a new QueryBuilder and returns it.
// The builder can then be used to create query parsers.
func NewBuilder() *QueryBuilder {
	queryBuilder := &QueryBuilder{
		filters: []*definition.Filter{},
	}
	return queryBuilder
}
