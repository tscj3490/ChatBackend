package config

import "time"

// EmailServer is email server address
var EmailServer = "bsemailmarketing.smtp.com" //smtp.live.com

// EmailPort is 465 by default 587
var EmailPort = 25025 // 587

// EmailUsername is username
var EmailUsername = "platben" // hantig1986@outlook.com

// EmailPassword is password
var EmailPassword = "ZdpDs95R" // tscj3490han919

// Constants for Email service
const (
	EmailTimeout = 10 * time.Second
)
