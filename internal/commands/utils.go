package commands

import (
	"errors"
	"fmt"

	"github.com/tcoyne1729/todo/internal/storage"
)

func checkUniqueness(ids []string, prefixLen int, startIndex int) int {
	if len(ids) == 0 {
		return 0
	}
	seen := make(map[byte]bool, len(ids)-startIndex)
	for k, elt := range ids {
		if k < startIndex {
			continue
		}
		shortenedId := elt[prefixLen]
		if seen[shortenedId] {
			return k
		} else {
			seen[shortenedId] = true
		}
	}
	return len(ids)
}

type Bijection struct {
	ShortToLong map[string]string
	LongToShort map[string]string
}

func NewBijection() *Bijection {
	return &Bijection{
		ShortToLong: make(map[string]string),
		LongToShort: make(map[string]string),
	}
}

func (b *Bijection) Insert(shortId string, longId string) error {
	if _, ok := b.LongToShort[longId]; ok {
		return fmt.Errorf("long id %s already exists and cant be inserted", longId)
	}
	b.LongToShort[longId] = shortId
	if _, ok := b.ShortToLong[shortId]; ok {
		return fmt.Errorf("short id %s already exists and cant be inserted", shortId)
	}
	b.ShortToLong[shortId] = longId
	return nil
}

func (b *Bijection) GetLongToShort(longId string) (string, bool) {
	val, ok := b.LongToShort[longId]
	return val, ok
}

func (b *Bijection) GetShortToLong(shortId string) (string, bool) {
	val, ok := b.ShortToLong[shortId]
	return val, ok
}

func ShortIds(fullIds []string) (*Bijection, error) {
	// strip out any ids which are empty
	strippedIds := make([]string, 0)
	for _, elt := range fullIds {
		if elt != "" {
			strippedIds = append(strippedIds, elt)
		}
	}
	initialIndex := 0
	k := 0
	for k < len(strippedIds) {
		k = checkUniqueness(strippedIds, initialIndex, k)
		initialIndex++
	}
	// create the shortened output
	bij := NewBijection()
	for _, elt := range strippedIds {
		err := bij.Insert(elt[:initialIndex], elt)
		if err != nil {
			return &Bijection{}, err
		}
	}
	return bij, nil
}

func ShortToLongId(shortId string, store *storage.Store) (string, error) {
	listCmd := ListCmd{}
	allIds, err := listCmd.Run(store, false)
	if err != nil {
		return "", err
	}
	shortIds, err := ShortIds(allIds)
	if err != nil {
		return "", err
	}
	id, ok := shortIds.GetShortToLong(shortId)
	if !ok {
		return "", errors.New("error retrieving short id")
	}
	return id, nil
}

func PointString(taskID string, currentTask string) string {
	if taskID == currentTask {
		return "-->"
	}
	return "   " // 3 spaces
}
