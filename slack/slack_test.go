package slack

import (
	"testing"

	"github.com/pkg/errors"
)

func TestSend(t *testing.T) {
	type in struct {
		msg Msg
	}
	type out struct {
		result string
		err    error
	}
	tests := []struct {
		in  in
		out out
	}{
		// all arguments is empty
		{
			in: in{
				msg: Msg{},
			},
			out: out{
				result: "",
				err:    errors.Wrap(errors.Errorf("SlackMsg: %#v", Msg{}), "Illegal argument."),
			},
		},
		// Text is empty
		{
			in: in{
				msg: Msg{
					Text:    "",
					Channel: "@kitahara_yuki",
				},
			},
			out: out{
				result: "",
				err:    errors.Wrap(errors.Errorf("SlackMsg: %#v", Msg{Channel: "@kitahara_yuki"}), "Illegal argument."),
			},
		},
		// Channel is empty
		{
			in: in{
				msg: Msg{
					Text:    "Test!",
					Channel: "",
				},
			},
			out: out{
				result: "",
				err:    errors.Wrap(errors.Errorf("SlackMsg: %#v", Msg{Text: "Test!"}), "Illegal argument."),
			},
		},
		// Not exists channel
		{
			in: in{
				msg: Msg{
					Text:    "Test!",
					Channel: "test",
				},
			},
			out: out{
				result: "channel_not_found",
				err:    nil,
			},
		},
	}

	s := &Slack{}

	for _, test := range tests {
		res, err := s.Send(test.in.msg)
		if res != test.out.result {
			t.Errorf("res='%v', want='%v'", res, test.out.result)
		}
		if err == nil || test.out.err == nil {
			if err != test.out.err {
				t.Errorf("err='%v', want='%v'", err, test.out.err)
			}
		} else {
			if err.Error() != test.out.err.Error() {
				t.Errorf("err='%v', want='%v'", err.Error(), test.out.err.Error())
			}
		}
	}
}
