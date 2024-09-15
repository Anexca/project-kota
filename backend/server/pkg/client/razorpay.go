package client

import (
	"server/pkg/config"

	"github.com/razorpay/razorpay-go"
)

func NewRazorpayClient() (*razorpay.Client, error) {
	env, err := config.LoadEnvironment()
	if err != nil {
		return nil, err
	}

	return razorpay.NewClient(env.RazorpayKey, env.RazorpaySecret), nil
}
