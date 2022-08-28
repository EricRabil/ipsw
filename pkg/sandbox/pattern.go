package sandbox

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/apex/log"
	"github.com/blacktop/ipsw/internal/utils"
)

const (
	STATE_UNKNOWN               = 0
	STATE_TOKEN_BYTE_READ       = 1
	STATE_CONCAT_BYTE_READ      = 2
	STATE_CONCAT_SAVE_BYTE_READ = 3
	STATE_END_BYTE_READ         = 4
	STATE_SPLIT_BYTE_READ       = 5
	STATE_TOKEN_READ            = 6
	STATE_RANGE_BYTE_READ       = 7
	STATE_CONSTANT_READ         = 8
	STATE_SINGLE_BYTE_READ      = 9
	STATE_PLUS_READ             = 10
	STATE_RESET_STRING          = 11
)

type SandboxString struct {
	Base          string
	Token         string
	OutputStrings []string
	BaseStack     []string
	TokenStack    []byte
	Pos           int
	Len           int
}

type Stack[T any] struct {
	vals []T
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.vals) == 0
}

func (s *Stack[T]) Push(val ...T) {
	s.vals = append(s.vals, val...)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}
	top := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return top, true
}

func (s *Stack[T]) Prev() T {
	if len(s.vals) == 0 {
		var zero T
		return zero
	}
	return s.vals[len(s.vals)-1]
}

// FIXME: this is NOT done and has bugs (I think this is func ___define_filter_block_invoke in libsandbox.1.dylib)
func ParseRSS(dat []byte, globals []string) ([]string, error) {
	var out []string
	var base string
	var token string
	var ss Stack[byte]
	var bs Stack[string]

	r := bytes.NewReader(dat)

	for {
		state, err := r.ReadByte()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("failed to read state byte: %v", err)
		}

		switch {
		case state == 0 || state == 7: // STATE_UNKNOWN
			ss.Push(STATE_UNKNOWN)
			continue
		case state == 10: // STATE_END_BYTE_READ
			ss.Push(STATE_END_BYTE_READ)
			out = append(out, base+token)
			token = ""
			if val, ok := bs.Pop(); ok {
				base = val
			}
		case state == 15: // STATE_CONCAT_BYTE_READ
			if ss.Prev() == STATE_TOKEN_READ ||
				ss.Prev() == STATE_CONSTANT_READ ||
				ss.Prev() == STATE_RANGE_BYTE_READ ||
				ss.Prev() == STATE_SINGLE_BYTE_READ ||
				ss.Prev() == STATE_PLUS_READ {
				base += token
				token = ""
			}
			ss.Push(STATE_CONCAT_BYTE_READ)
		case state >= 128: // STATE_SPLIT_BYTE_READ
			ss.Push(STATE_SPLIT_BYTE_READ)
			len := int(state) - 127
			dat := make([]byte, len)
			if err := binary.Read(r, binary.LittleEndian, &dat); err != nil {
				return nil, fmt.Errorf("failed to read state byte: %v", err)
			}
			subtokens, err := ParseRSS(dat, globals)
			if err != nil {
				return nil, fmt.Errorf("failed to parse subtokens: %v", err)
			}
			for _, st := range subtokens {
				out = append(out, base+token+st)
			}
			token = ""
		case state == 5 || state == 6: // STATE_RESET_STRING
			ss.Push(STATE_RESET_STRING)
			out = append(out, base+token)
			bs = Stack[string]{}
			token = ""
			base = ""
		case state == 8: // STATE_CONCAT_SAVE_BYTE_READ
			// XXX: Read two bytes. I don't know what they do.
			if _, err := r.ReadByte(); err != nil {
				return nil, fmt.Errorf("failed to read state byte: %v", err)
			}
			if _, err := r.ReadByte(); err != nil {
				return nil, fmt.Errorf("failed to read state byte: %v", err)
			}
			if ss.Prev() == STATE_TOKEN_READ ||
				ss.Prev() == STATE_CONSTANT_READ ||
				ss.Prev() == STATE_RANGE_BYTE_READ ||
				ss.Prev() == STATE_SINGLE_BYTE_READ ||
				ss.Prev() == STATE_PLUS_READ {
				bs.Push(base)
				base += token
				token = ""
			}
			ss.Push(STATE_CONCAT_SAVE_BYTE_READ)
		case state >= 16 && state < 63: // STATE_CONSTANT_READ
			ss.Push(STATE_CONSTANT_READ)
			if len(globals) <= int(state-16) {
				return nil, fmt.Errorf("failed to parse state 'STATE_CONSTANT_READ': %d > len(globals)", state-16) // FIXME: why is this happening
			}
			token = "${" + globals[state-16] + "}"
		case state == 11: // STATE_RANGE_BYTE_READ
			bs.Push(base)
			base += token
			token = ""
			count, err := r.ReadByte()
			if err != nil {
				return nil, fmt.Errorf("failed to read state byte: %v", err)
			}
			var barray [][]byte
			ascii := true
			token = ""
			for i := 0; i < int(count+1); i++ {
				dat := make([]byte, 2)
				if err := binary.Read(r, binary.LittleEndian, &dat); err != nil {
					return nil, fmt.Errorf("failed to read state byte: %v", err)
				}
				ascii = utils.IsASCII(string(dat))
				barray = append(barray, dat)
			}
			if ascii {
				token = "["
				for _, tuple := range barray {
					token += fmt.Sprintf("%c-%c", tuple[0], tuple[1])
				}
				token += "]"
			} else {
				if len(barray) >= 2 {
					tuple1 := barray[0]
					tuple2 := barray[1]
					if tuple1[1] == 0xff && tuple2[0] == 0x00 {
						if tuple1[0]-1 == tuple2[1]+1 {
							token = fmt.Sprintf("[^%c]", tuple1[0]-1) // range exclude
						} else {
							token = fmt.Sprintf("[^%c-%c]", tuple2[1]+1, tuple1[0]-1) // range exclude
						}
					} else {
						token = "[TODO]"
					}
				} else {
					token = "["
					for _, tuple := range barray {
						token += fmt.Sprintf("\\x%02x-\\x%02x", tuple[0], tuple[1])
					}
					token += "]"
					// return nil, fmt.Errorf("b-array should have length 2: got %d", len(barray))
				}
			}
			ss.Push(STATE_RANGE_BYTE_READ)
		case state == 2: // STATE_PLUS_READ
			if ss.Prev() == STATE_CONCAT_BYTE_READ {
				base += "+"
				token = ""
			}
			next, err := r.ReadByte()
			if err != nil {
				return nil, fmt.Errorf("failed to read state byte: %v", err)
			}
			token = string(next)
			ss.Push(STATE_PLUS_READ)
		default: // STATE_TOKEN_BYTE_READ
			if ss.Prev() != STATE_TOKEN_READ {
				var len int
				switch state {
				case 3:
					/*
					 *  STRING DATA:
					 *     0000000000000000:  03 2f 0f 0b 01 30 ff 00  2e 0f 02 2f 0f 47 53 43  |./...0...../.GSC|
					 *     0000000000000010:  5f 49 6e 66 6f 2f 0f 0a                           |_Info/..|
					 *
					 * Corresponding line in com.apple.WebKit.WebContent.sb.in:
					 *     (regex #"/[^/]+/SC_Info/")))
					 *
					 * The first byte is 0x03 and then there is a single char before the next "state" byte
					 * so I think the length should be 1. (TODO: I'd love to confirm via RE in libsandbox or the sandbox kext)
					 */
					// fmt.Println(utils.HexDump(dat, 0))
					len = 1
				case 4:
					l, err := r.ReadByte()
					if err != nil {
						return nil, fmt.Errorf("failed to read state byte: %v", err)
					}
					len = int(l) + 65
				default:
					len = int(state) - 63
				}
				dat := make([]byte, len)
				if err := binary.Read(r, binary.LittleEndian, &dat); err != nil {
					return nil, fmt.Errorf("failed to read state byte: %v", err)
				}
				token = string(dat)
				ss.Push(STATE_TOKEN_READ)
				continue
			} else {
				log.Warn("read token byte from token state")
			}
			ss.Push(STATE_TOKEN_BYTE_READ)
		}
	}

	return out, nil
}
