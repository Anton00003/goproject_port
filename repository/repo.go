package repository

import (
	"errors"
	"fmt"
	"goproject_port/datastruct"
	"math/rand"
	"sync"
)

type repo struct {
	In  []*datastruct.Port
	Out []*datastruct.Port
	mu  sync.Mutex
}

func New(numIn, numOut int) *repo {
	in := make([]*datastruct.Port, numIn)
	out := make([]*datastruct.Port, numOut)

	for i := range in {
		in[i] = &datastruct.Port{Value: rand.Intn(2), Ch: make(chan datastruct.Reqest)}
	}
	for i := range out {
		out[i] = &datastruct.Port{Ch: make(chan datastruct.Reqest)}
	}

	return &repo{In: in, Out: out}
}

func (r *repo) GetIn(nIn int) (int, error) {
	if nIn >= len(r.In) || nIn < 0 {
		err := errors.New("invalid number of In port")
		return 0, err
	}
	return r.In[nIn].Value, nil
}

func (r *repo) GetChIn(nIn int) (chan datastruct.Reqest, error) {
	if nIn >= len(r.In) || nIn < 0 {
		err := errors.New("invalid number of In port")
		return nil, err
	}
	return r.In[nIn].Ch, nil
}

func (r *repo) PostOut(nOut int, newValue int) error {
	if nOut >= len(r.Out) || nOut < 0 {
		err := errors.New("invalid number of Out port")
		return err
	}
	r.mu.Lock()
	r.Out[nOut].Value = newValue
	r.mu.Unlock()
	return nil
}

func (r *repo) PostChOut(nOut int) (chan datastruct.Reqest, error) {
	if nOut >= len(r.Out) || nOut < 0 {
		err := errors.New("invalid number of Out port")
		return nil, err
	}
	return r.Out[nOut].Ch, nil
}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func (r *repo) GetAll() {
	in := make([]datastruct.Port, len(r.In))
	out := make([]datastruct.Port, len(r.Out))

	for i, port := range r.In {
		in[i] = *port
	}
	for i, port := range r.Out {
		out[i] = *port
	}

	fmt.Println(in)
	fmt.Println(out)
}
