package sse

import (
	"encoding/json"
	"fmt"
)

// StreamEvent wraps a single event of a typed stream, mirroring the
// StreamEvent shape of the JavaScript and Python clients.
type StreamEvent[T any] struct {
	// Data is the deserialized payload of the event, or nil when Err is set.
	Data *T
	// Raw is the original, unparsed event.
	Raw Event
	// Err is the error that occurred while deserializing the event payload,
	// if any. It only concerns this event: the stream continues past it.
	Err error
}

// Stream deserializes the payload of each event of a Server-Sent Events
// stream into a value of type T.
type Stream[T any] struct {
	decoder Decoder
	cur     StreamEvent[T]
	err     error
}

// NewStream wraps a [Decoder] into a typed Stream. The err argument allows
// callers to forward the error of the call that produced the decoder in a
// single expression: when non-nil, the stream is created in a failed state
// and Err reports it.
func NewStream[T any](decoder Decoder, err error) *Stream[T] {
	return &Stream[T]{decoder: decoder, err: err}
}

// Next advances the stream to the next event and deserializes its payload
// into the value returned by Current. It returns false when the stream is
// exhausted or a transport error occurred, in which case Err reports it.
//
// A payload that fails to unmarshal is not terminal: the event is delivered
// with Err set and a nil Data, and the stream continues, consistent with the
// JavaScript and Python clients.
func (s *Stream[T]) Next() bool {
	if s.err != nil || s.decoder == nil {
		return false
	}

	if !s.decoder.Next() {
		s.err = s.decoder.Err()

		return false
	}

	evt := s.decoder.Event()

	var cur T

	err := json.Unmarshal(evt.Data, &cur)
	if err != nil {
		s.cur = StreamEvent[T]{Raw: evt, Err: fmt.Errorf("cannot unmarshal event data: %w", err)}

		return true
	}

	s.cur = StreamEvent[T]{Data: &cur, Raw: evt}

	return true
}

// Current returns the event read by the last successful call to Next.
func (s *Stream[T]) Current() StreamEvent[T] {
	return s.cur
}

// Err returns the terminal error of the stream, or nil if the stream simply
// ended.
func (s *Stream[T]) Err() error {
	return s.err
}

// Close closes the underlying reader.
func (s *Stream[T]) Close() error {
	if s.decoder == nil {
		return nil
	}

	err := s.decoder.Close()
	if err != nil {
		return fmt.Errorf("cannot close stream: %w", err)
	}

	return nil
}
