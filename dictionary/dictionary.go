package dictionary

// Dictionary ...
type Dictionary map[string]string

//--
const (
	ErrNotFound         = ExceptionInfo("could not find the word you were looking for")
	ErrWordExists       = ExceptionInfo("cannot add word because it already exists")
	ErrWordDoesNotExist = ExceptionInfo("cannot update word because it does not exist")
)

// ExceptionInfo ...
type ExceptionInfo string

func (e ExceptionInfo) Error() string {
	return string(e)
}

// Search ...
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add ...
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

// Update ...
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

// Delete ...
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
