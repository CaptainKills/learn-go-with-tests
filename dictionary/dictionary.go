package dictionary

const (
	ErrorNotFound         = DictionaryError("could not find the key you were looking for")
	ErrorWordExists       = DictionaryError("key already exists")
	ErrorWordDoesNotExist = DictionaryError("key does not exist yet")
)

type (
	Dictionary      map[string]string
	DictionaryError string
)

func (e DictionaryError) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]

	if !ok {
		return "", ErrorNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrorNotFound:
		d[key] = value
	case nil:
		return ErrorWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrorNotFound:
		return ErrorWordDoesNotExist
	case nil:
		d[key] = value
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)

	switch err {
	case ErrorNotFound:
		return ErrorWordDoesNotExist
	case nil:
		delete(d, key)
	default:
		return err
	}

	return nil
}
