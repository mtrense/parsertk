package lexer

type BufferedRuneReader interface {
	Mark() int
	Offset() int
	Read() rune
	Peek() rune
	Rewind()
	Error() error
	EOF() bool
}

type stringReader struct {
	input  []rune
	offset int
	marks  []int
}

func (s *stringReader) Mark() int {
	s.marks = append(s.marks, s.offset)
	return s.offset
}

func (s *stringReader) Offset() int {
	return s.offset
}

func (s *stringReader) Read() rune {
	if !s.EOF() {
		r := s.input[s.offset]
		s.offset++
		return r
	}
	return '\uFFFD'
}

func (s *stringReader) Peek() rune {
	if !s.EOF() {
		return s.input[s.offset]
	}
	return '\uFFFD'
}

func (s *stringReader) Rewind() {
	if len(s.marks) > 0 {
		lastMark := s.marks[len(s.marks)-1]
		s.marks = s.marks[0 : len(s.marks)-1]
		s.offset = lastMark
	}
}

func (s *stringReader) Error() error {
	return nil
}

func (s *stringReader) EOF() bool {
	return s.offset >= len(s.input)
}

func StringReader(input string) BufferedRuneReader {
	return &stringReader{
		input:  []rune(input),
		offset: 0,
		marks:  make([]int, 0),
	}
}
