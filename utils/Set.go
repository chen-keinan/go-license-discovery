package utils

//Set object
type Set struct {
	setMap map[string]SetInterface
}

//NewSet instantiate new Set
func NewSet() *Set {
	return &Set{setMap: make(map[string]SetInterface)}
}

//Add add new item to set
func (set Set) Add(value SetInterface) {
	if set.setMap[value.ToString()] == nil {
		set.setMap[value.ToString()] = value
	}
}

//GetValue get value from set
func (set Set) GetValue(key string) string {
	v, ok := set.setMap[key]
	if !ok {
		return ""
	}
	return v.ToString()
}

//AddString add string value to set
func (set Set) AddString(value string) {
	if set.setMap[value] == nil {
		set.setMap[value] = SetString(value)
	}
}

//AddValues add multi values to set
func (set Set) AddValues(values []SetInterface) {
	for _, value := range values {
		set.Add(value)
	}
}

//AddStringValues add multi string values to set
func (set Set) AddStringValues(values []string) {
	for _, value := range values {
		set.AddString(value)
	}
}

//Update update set value
func (set Set) Update(value SetInterface) {
	set.setMap[value.ToString()] = value
}

//Remove remove value from set
func (set Set) Remove(value SetInterface) {
	if set.setMap[value.ToString()] != nil {
		delete(set.setMap, value.ToString())
	}
}

//Size set size
func (set Set) Size() int {
	return len(set.setMap)
}

//Values return set values
func (set Set) Values() []SetInterface {
	var values []SetInterface = make([]SetInterface, 0)
	for _, value := range set.setMap {
		values = append(values, value)
	}
	return values
}

//StringValues return set string values
func (set Set) StringValues() []string {
	values := []string{}
	for _, value := range set.setMap {
		values = append(values, value.ToString())
	}
	return values
}

//SetInterface interface
type SetInterface interface {
	ToString() string
}

//SetString set string type
type SetString string

//ToString return value as string
func (s SetString) ToString() string {
	return string(s)
}

//SetID set ID object
type SetID struct {
	ID string
}

//ToString print set id object
func (id SetID) ToString() string {
	return id.ID
}

//RemoveDuplicatedObjectIds remove duplicate from set
func RemoveDuplicatedObjectIds(ids []string) []string {
	set := NewSet()
	for _, id := range ids {
		set.Add(SetID{id})
	}

	res := make([]string, set.Size())
	for i, v := range set.Values() {
		res[i] = v.(SetID).ToString()
	}
	return res
}
