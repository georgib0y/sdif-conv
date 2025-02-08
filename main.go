package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Code string
type USPS string

type SdifEncoder struct {
	w io.Writer
}

type SdifTag struct {
	start, len int
	code       string
}

func parseSdifTagInt(kvs map[string]string, key string) (int, error) {
	str, ok := kvs[key]
	if !ok {
		log.Printf("%s key is missing, using zero value")
		return 0, nil
	}
	n, err := strconv.Atoi(str)
	if err != nil {
		return 0, fmt.Errorf("could not parse int %s, %w", str, err)
	}
	return n, nil
}

func NewSdifTag(s string) (SdifTag, error) {
	if s == "" {
		return SdifTag{}, errors.New("Sdif tag string is empty")
	}

	kvs := map[string]string{}

	for _, kvStr := range strings.Split(s, ",") {
		kvPair := strings.Split(kvStr, "=")
		if len(kvPair) != 2 {
			return SdifTag{}, fmt.Errorf("invalid key value in tag string %s: %s", s, kvStr)
		}
		kvs[kvPair[0]] = kvPair[1]
	}

	start, err := parseSdifTagInt(kvs, "start")
	if err != nil {
		return SdifTag{}, fmt.Errorf("could not parse 'start': %w", err)
	}
	l, err := parseSdifTagInt(kvs, "len")
	if err != nil {
		return SdifTag{}, fmt.Errorf("could not parse 'len': %w", err)
	}

	sdifTag := SdifTag{
		start: start,
		len:   l,
		code:  kvs["code"],
	}

	return sdifTag, nil
}

type Record interface {
	ConstTag() string
}

func encodeRecord(w io.Writer, rec Record) error {
	if reflect.TypeOf(rec).Kind() != reflect.Struct {
		return errors.New("rec is not a struct")
	}

	bTag := []byte(rec.ConstTag())
	if len(bTag) != 2 {
		return fmt.Errorf("Tag %s is wrong byte len (%d)", rec.ConstTag(), len(bTag))
	}

	b := make([]byte, 162)
	// fill the whole row with "blanks" - spec is not super clear on what a 'blank' is
	for i := range b {
		b[i] = ' '
	}

	// write the tag
	b[0] = bTag[0]
	b[1] = bTag[1]

	for i := range reflect.ValueOf(rec).NumField() {
		t := reflect.TypeOf(rec).Field(i).Tag.Get("sdif")
		st, err := NewSdifTag(t)
		if err != nil {
			return err
		}

		// sdif spec is 1 indexed
		sl := b[st.start-1 : st.start-1+st.len]

		switch v := reflect.ValueOf(rec).Field(i).Interface().(type) {
		case string:
			err := encodeString(sl, v)
			if err != nil {
				return err
			}
		case int:
			err := encodeInt(sl, v)
			if err != nil {
				return err
			}
		case float32:
			err := encodeFloat(sl, v)
			if err != nil {
				return err
			}
		case USPS:
			err := encodeString(sl, string(v))
			if err != nil {
				return err
			}
		case Code:
			err := encodeString(sl, string(v))
			if err != nil {
				return err
			}
		case time.Time:
			err := encodeDate(sl, v)
			if err != nil {
				return err
			}
		}
	}

	// eugh
	b[160] = 'r'
	b[161] = 'n'

	n, err := w.Write(b)

	if err != nil {
		return err
	}

	if n != 162 {
		return fmt.Errorf("did not write all 162 bytes, wrote: %d bytes", n)
	}

	return nil
}

func encodeString(dst []byte, s string) error {
	if len([]byte(s)) > len(dst) {
		return fmt.Errorf("string '%s' too long len: %d exp %d", s, len([]byte(s)), len(dst))
	}

	// spec states that if alpha data type is numeric than right align
	if n, err := strconv.ParseFloat(s, 32); err == nil {
		return encodeFloat(dst, float32(n))
	}

	copy(dst, s)
	return nil
}

func encodeInt(dst []byte, n int) error {
	s := fmt.Sprintf("%d", n)
	b := []byte(s)

	if len(b) > len(dst) {
		return fmt.Errorf("int %s too long, len: %d exp %d", s, len(b), len(dst))
	}

	copy(dst, b)
	return nil
}

func encodeFloat(dst []byte, f float32) error {
	s := fmt.Sprintf("%f", f)
	b := []byte(s)

	if len(b) > len(dst) {
		return fmt.Errorf("float %s too long, len: %d exp %d", s, len(b), len(dst))
	}

	copy(dst, b)
	return nil
}

func encodeDate(dst []byte, d time.Time) error {
	s := d.Format("01022006") //MMDDYYYY
	copy(dst, s)
	return nil
}

type MeetPyramid struct {
	fileDesc FileDescriptionRecord
}

func main() {
	fd := FileDescriptionRecord{
		OrgCode:      "1",
		VersionNum:   "Ver. 3.0",
		FileCode:     "01",
		SoftwareName: "Custom",
		SoftwareVer:  "0.0.1",
		ContactName:  "George",
		PhoneNumber:  "somenum",
		LastModified: time.Now(),
		LscSubmitted: "na",
	}

	mr := MeetRecord{
		OrgCode:     "1",
		MeetName:    "My Meet",
		AddrLineOne: "1234 Layne Lane",
		AddrLineTwo: "",
		City:        "Cityzville",
		State:       "",
		PostCode:    "2575",
		Country:     "AUS",
		MeetCode:    "M",
		Start:       time.Date(2025, time.February, 1, 10, 0, 0, 0, time.Local),
		End:         time.Date(2025, time.February, 2, 10, 0, 0, 0, time.Local),
		Altitude:    700,
		Course:      "S",
	}

	var sb strings.Builder
	err := encodeRecord(&sb, fd)
	if err != nil {
		log.Panicln(err)
	}
	sb.WriteRune('\n')

	err = encodeRecord(&sb, mr)
	if err != nil {
		log.Panicln(err)
	}

	fmt.Printf("%s\n", sb.String())

	for i := range 16 {
		v := 'a'
		v += rune(i)
		fmt.Printf("%c123456789", v)
	}

	fmt.Println()
}
