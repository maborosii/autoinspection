package email

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestNewMail(t *testing.T) {
	m := NewMail("yo", 123)
	m.Login("username", "password")
	mockCtl := gomock.NewController(t)
	MockEmailDialer := NewMockEmailDialer(mockCtl)

	t.Run("ok", func(t *testing.T) {
		MockEmailDialer.EXPECT().DialAndSend(gomock.Any()).Return(nil)
		err := m.Send(
			"subject",
			m.BuildMessage("content"),
			[]string{"to1m@email.com", "to2@email.com"},
			WithMailSendDialer(func(host string, port int, username, passwd string) EmailDialer {
				return MockEmailDialer
			}),
		)
		require.NoError(t, err)
	})
	t.Run("err", func(t *testing.T) {
		errWant := errors.New("haha")
		MockEmailDialer.EXPECT().DialAndSend(gomock.Any()).Return(errWant)
		err := m.Send(
			"subject",
			m.BuildMessage("content"),
			[]string{"to1m@email.com", "to2@email.com"},
			WithMailSendDialer(func(host string, port int, username, passwd string) EmailDialer {
				return MockEmailDialer
			}),
		)
		require.True(t, errors.Is(err, errWant))
	})
}
