package dictionary

const (
	ErrorNotFound         = DictionaryErr("could not find the word you are looking for")
	ErrorWordExists       = DictionaryErr("cannot add because word already exists")
	ErrorWordDoesNotExist = DictionaryErr("cannot perform operation as word does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrorNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrorNotFound:
		d[word] = definition
	case nil:
		return ErrorWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrorNotFound:
		return ErrorWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case ErrorNotFound:
		return ErrorWordDoesNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}
	return nil
}
