package launcher

import (
	"NovaPlus/module/mmcll"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"image/png"
	"path/filepath"
	runtime2 "runtime"
)

type AccountType struct {
	AType        string  `json:"type"`
	Name         string  `json:"name"`
	UUID         string  `json:"uuid"`
	AccessToken  *string `json:"access_token,omitempty"`
	RefreshToken *string `json:"refresh_token,omitempty"`
	ClientToken  *string `json:"client_token,omitempty"`
	Server       *string `json:"server,omitempty"`
	BaseCode     *string `json:"base_code,omitempty"`
	HeadSkin     string  `json:"head_skin"`
}
type AccountList struct {
	Accounts []AccountType `json:"accounts"`
}
type AccountMethod struct {
	Ctx context.Context
}

func CutHeadSkin(data []byte, isInner bool) ([]byte, error) {
	img, _, err := image.Decode(bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	rgba := img.(*image.NRGBA)
	subHead := rgba.SubImage(image.Rect(mmcll.If(isInner, 8, 40).(int), 8, mmcll.If(isInner, 16, 48).(int), 16)).(*image.NRGBA)
	var buf bytes.Buffer
	err = png.Encode(&buf, subHead)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (m *AccountMethod) SetAccountConfig(accounts AccountList) ExceptionHandler[any] {
	exception := func(err error) ExceptionHandler[any] {
		return NewExceptionHandler[any](500, false, "File System Error: "+err.Error(), nil)
	}
	realJson, err := json.MarshalIndent(accounts, "", "\t")
	if err != nil {
		return exception(err)
	}
	str := string(realJson)
	home, err := GetOtherDir()
	if err != nil {
		home = GetCurrentDir()
	}
	d := mmcll.If(runtime2.GOOS == "windows", []string{"PCL.Nova"}, mmcll.If(runtime2.GOOS == "darwin", []string{"PCL.Nova", "config"}, []string{".PCL.Nova"})).([]string)
	d = append([]string{home}, d...)
	d = append(d, "AccountJSON.json")
	sd := filepath.Join(d...)
	err = mmcll.SetFile(filepath.Join(sd), str)
	if err != nil {
		return exception(err)
	}
	return NewExceptionHandler[any](201, true, "Created!", nil)
}

func (m *AccountMethod) GetAccountConfig() ExceptionHandler[AccountList] {
	exception := func(err error) ExceptionHandler[AccountList] {
		return NewExceptionHandler(500, false, "File System Error: "+err.Error(), AccountList{Accounts: []AccountType{}})
	}
	home, err := GetOtherDir()
	if err != nil {
		home = GetCurrentDir()
	}
	d := mmcll.If(runtime2.GOOS == "windows", []string{"PCL.Nova"}, mmcll.If(runtime2.GOOS == "darwin", []string{"PCL.Nova", "config"}, []string{".PCL.Nova"})).([]string)
	d = append([]string{home}, d...)
	d = append(d, "AccountJSON.json")
	cur := filepath.Join(d...)
	if content, _ := mmcll.GetFile(cur); content == "" {
		err := mmcll.SetFile(cur, "{\n\t\"accounts\": []\n}")
		if err != nil {
			return exception(err)
		}
		return NewExceptionHandler(201, true, "Created!", AccountList{Accounts: []AccountType{}})
	} else {
		var arr AccountList
		err = json.Unmarshal([]byte(content), &arr)
		if err != nil {
			return exception(err)
		}
		return NewExceptionHandler(200, true, "OK!", arr)
	}
}

// GetUserCode 微软登录，但是此处用到了 MS_KEY，这个是在 privacy.go 里面被忽略了的，请手动新建并添加一个~
func (m *AccountMethod) GetUserCode() ExceptionHandler[[]string] {
	exception := func(err error) ExceptionHandler[[]string] {
		return NewExceptionHandler(500, false, "Get User Code has Exception: "+err.Error(), []string{})
	}
	url := mmcll.NewAccountLogin(MS_KEY)
	userCode, deviceCode, err := url.GetUserCode()
	if err != nil {
		return exception(err)
	}
	return NewExceptionHandler(200, true, "OK!", []string{userCode, deviceCode})
}

// GetAccessToken 获取到 AccessToken！
func (m *AccountMethod) GetAccessToken(deviceCode string) ExceptionHandler[AccountType] {
	exception := func(code int, err error) ExceptionHandler[AccountType] {
		return NewExceptionHandler(code, false, "Get Access Token has Exception: "+err.Error(), AccountType{})
	}
	url := mmcll.NewAccountLogin(MS_KEY)
	// 如果 err 的 code == -6 的话，说明没有完成登录，需要登录一次才能继续~
	accountLogin, err := url.LoginMicrosoft(deviceCode)
	if err != nil {
		var err2 mmcll.ErrorMMCLL
		errors.As(err, &err2)
		return exception(mmcll.If(err2.Code() != -6, 500, 401).(int), err)
	}
	var accessToken = accountLogin.AccessToken()
	var refreshToken = accountLogin.RefreshToken()
	var accountType = AccountType{
		AType:        "Microsoft",
		Name:         accountLogin.Name(),
		UUID:         accountLogin.Uuid(),
		AccessToken:  &accessToken,
		RefreshToken: &refreshToken,
	}
	return NewExceptionHandler(200, true, "OK!", accountType)
}

// GetUserHeadSkin 通过 用户 UUID 获取到用户的 头像！（Base64）
func (m *AccountMethod) GetUserHeadSkin(uuid string) ExceptionHandler[string] {
	exception := func(err error) ExceptionHandler[string] {
		return NewExceptionHandler(500, false, "Get Head Skin has Exception: "+err.Error(), "")
	}
	url := mmcll.NewUrlMethod(fmt.Sprintf("https://sessionserver.mojang.com/session/minecraft/profile/%s", uuid))
	get, err := url.GetDefault()
	if err != nil {
		return exception(err)
	}
	var profile map[string]any
	err = json.Unmarshal(get, &profile)
	if err != nil {
		return exception(err)
	}
	profileValue := mmcll.Safe(profile, "", "properties", 0, "value").(string)
	if profileValue == "" {
		return exception(errors.New("cannot get head skin profile json"))
	}
	decoded, err := base64.StdEncoding.DecodeString(profileValue)
	if err != nil {
		return exception(err)
	}
	var skin map[string]any
	err = json.Unmarshal(decoded, &skin)
	if err != nil {
		return exception(errors.New("cannot parse head skin base metadata"))
	}
	skinUrl := mmcll.Safe(skin, "", "textures", "SKIN", "url").(string)
	if skinUrl == "" {
		return exception(errors.New("cannot get head skin base64 decode"))
	}
	url2 := mmcll.NewUrlMethod(skinUrl)
	skinValue, err := url2.GetDefault()
	if err != nil {
		return exception(err)
	}
	// 以下是经过裁剪的！如果不需要裁剪，请删掉以下片段，并且将上述 skinValue 直接包装成 base64！
	resImgInner, err := CutHeadSkin(skinValue, true)
	if err != nil {
		return exception(err)
	}
	resImgOuter, err := CutHeadSkin(skinValue, false)
	if err != nil {
		return exception(err)
	}
	result := base64.StdEncoding.EncodeToString(resImgInner) + "**" + base64.StdEncoding.EncodeToString(resImgOuter)
	return NewExceptionHandler(200, true, "OK!", result)
}

func (m *AccountMethod) GetUserCodeThirdOAuth(server string) ExceptionHandler[[]string] {
	exception := func(err error) ExceptionHandler[[]string] {
		return NewExceptionHandler(500, false, "Get User Code has Exception: "+err.Error(), []string{})
	}
	url := mmcll.NewAccountLogin(LS_KEY)
	userCode, deviceCode, err := url.GetUserCodeThirdOAuth(server)
	if err != nil {
		return exception(err)
	}
	return NewExceptionHandler(200, true, "OK!", []string{userCode, deviceCode})
}

// 此处用到了 LittleSkin OAuth Code，如果你们需要，可以自行申请一个，随后放到 privacy.go 里面~
func (m *AccountMethod) GetThirdOAuthAccessToken(server, deviceCode string) ExceptionHandler[AccountType] {
	exception := func(code int, err error) ExceptionHandler[AccountType] {
		return NewExceptionHandler(code, false, "Get Access Token has Exception: "+err.Error(), AccountType{})
	}
	url := mmcll.NewAccountLogin(LS_KEY)
	// 如果 err 的 code == -106 的话，说明没有完成登录，需要登录一次才能继续~
	accountLogin, err := url.LoginThirdPartyOAuth(server, deviceCode)
	if err != nil {
		var err2 mmcll.ErrorMMCLL
		errors.As(err, &err2)
		return exception(mmcll.If(err2.Code() != -106, 500, 401).(int), err)
	}
	var accessToken = accountLogin.AccessToken()
	var refreshToken = accountLogin.RefreshToken()
	var accountType = AccountType{
		AType:        "ThirdPartyOAuth",
		Name:         accountLogin.Name(),
		UUID:         accountLogin.Uuid(),
		AccessToken:  &accessToken,
		RefreshToken: &refreshToken,
	}
	return NewExceptionHandler(200, true, "OK!", accountType)
}

// GetThirdOAuthAccessToken 通过 第三方用户 UUID 获取到用户的 头像！（Base64）
func (m *AccountMethod) GetThirdHeadSkin(server, uuid string) ExceptionHandler[string] {
	exception := func(err error) ExceptionHandler[string] {
		return NewExceptionHandler(500, false, "Get Head Skin has Exception: "+err.Error(), "")
	}
	url := mmcll.NewUrlMethod(fmt.Sprintf("%s/sessionserver/session/minecraft/profile/%s", server, uuid))
	get, err := url.GetDefault()
	if err != nil {
		return exception(err)
	}
	var profile map[string]any
	err = json.Unmarshal(get, &profile)
	if err != nil {
		return exception(err)
	}
	profileValue := mmcll.Safe(profile, "", "properties", 0, "value").(string)
	if profileValue == "" {
		return exception(fmt.Errorf("cannot get head skin profile json"))
	}
	sj2, err := base64.StdEncoding.DecodeString(profileValue)
	if err != nil {
		return exception(err)
	}
	var textureJSON map[string]any
	err = json.Unmarshal(sj2, &textureJSON)
	if err != nil {
		return exception(err)
	}
	surl := mmcll.Safe(textureJSON, "", "textures", "SKIN", "url").(string)
	pic := mmcll.NewUrlMethod(surl)
	skinValue, err := pic.GetDefault()
	if err != nil {
		return exception(err)
	}
	resImgInner, err := CutHeadSkin(skinValue, true)
	if err != nil {
		return exception(err)
	}
	resImgOuter, err := CutHeadSkin(skinValue, false)
	if err != nil {
		return exception(err)
	}
	result := base64.StdEncoding.EncodeToString(resImgInner) + "**" + base64.StdEncoding.EncodeToString(resImgOuter)
	return NewExceptionHandler(200, true, "OK!", result)
}
func (m *AccountMethod) GetThirdAPAccessToken(server, username, password string) ExceptionHandler[[]AccountType] {
	exception := func(code int, err error) ExceptionHandler[[]AccountType] {
		return NewExceptionHandler(code, false, "Get Access Token has Exception: "+err.Error(), []AccountType{})
	}
	url := mmcll.NewAccountLogin(server)
	accountLogin, err := url.LoginThirdParty(username, password)
	if err != nil {
		var err2 mmcll.ErrorMMCLL
		errors.As(err, &err2)
		return exception(mmcll.If(err2.Code() != 6, 500, 401).(int), err)
	}
	//var accessToken = accountLogin.
	var tokens []AccountType
	for _, v := range accountLogin {
		accessToken := v.AccessToken()
		tokens = append(tokens, AccountType{
			AType:       "ThirdParty",
			Name:        v.Name(),
			UUID:        v.Uuid(),
			AccessToken: &accessToken,
		})
	}
	return NewExceptionHandler(200, true, "OK!", tokens)
}
