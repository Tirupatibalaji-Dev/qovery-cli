package io

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func CheckAuthenticationOrQuitWithMessage() {
	authorizationToken := strings.TrimSpace(GetAuthorizationToken())
	refreshToken := strings.TrimSpace(GetRefreshToken())
	accountId := strings.TrimSpace(GetAccountId())
	expiration, err := GetAuthorizationTokenExpiration()

	if err == nil && expiration.Before(time.Now()) && refreshToken != "" {
		refreshTokenOrQuitWithMessage()
	}

	if authorizationToken == "" || accountId == "" {
		if refreshToken != "" {
			refreshTokenOrQuitWithMessage()
		} else {
			fmt.Println("Are you authenticated? Consider doing 'qovery auth' to authenticate yourself")
			os.Exit(1)
		}
	}
}

func refreshTokenOrQuitWithMessage() {
	err := RefreshAccessToken()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func CheckHTTPResponse(resp *http.Response) error {
	if resp == nil {
		return errors.New("Qovery is in maintenance. Try again later or contact #support on https://discord.qovery.com")
	}

	if resp.StatusCode == http.StatusUnauthorized {
		err := RefreshAccessToken()
		if err != nil {
			return errors.New("Your authentication has expired. Please re-authenticate yourself with 'qovery auth'")
		}
		return errors.New("Your authentication token has expired. Refreshed session. Please, re-run the command. ")
	} else if resp.StatusCode == http.StatusForbidden {
		return errors.New("Your account must be approved by an administrator to get access to this resource. " +
			"Please join #support on https://discord.qovery.com")
	} else if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return errors.New("Something goes wrong while requesting the Qovery API. Please try again later or " +
			"contact the #support on https://discord.qovery.com")
	}

	return nil
}
