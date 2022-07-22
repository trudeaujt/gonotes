package maps

type Dictionary map[string]string

var (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add this word because it is already in the dictionary")
	//we could reuse ErrNotFound for this, but having specific errors gives us more information about what went wrong.
	//For isntance, we can redirect the user when ErrNotFound is encountered but display an error message when ErrDoesNotExist is encountered.
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

//We are implementing a DictionaryErr type which implements the error interface, to make the errors more reusable and immutable.
//Please read https://dave.cheney.net/2016/04/07/constant-errors for more information.
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	//the second value returned from a map lookup is a boolean which indicates if the key was found successfully.
	definition, found := d[word]
	if !found {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word string, definition string) error {
	_, err := d.Search(word)

	//Using a switch statement here provides an extra safety net in case Search returns an error other than ErrNotFound.
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

func (d Dictionary) Update(word string, definition string) error {
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

func (d Dictionary) Delete(word string) {
	//the delete function returns nothing, so we are basing our delete method on the same concept and returning nothing.
	//Since deleting a value that is not there has no effect, unlike the add and update methods, we don't need to complicate our API with errors.
	delete(d, word)
}
