package main

func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}

type Dictionary map[string]string

const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	defination, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return defination, nil
}

func (d Dictionary) Add(word, defination string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = defination
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}
