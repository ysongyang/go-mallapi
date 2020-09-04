package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
)

// Setup Initialize the util
func Setup() {
	//jwtSecret = []byte(setting.AppSetting.JwtSecret)
}


func GetPassWord(password, salt string) string {
	md5Hash := md5.New()
	md5Hash.Write([]byte(password))                // 需要加密的字符串
	md5Pwd := hex.EncodeToString(md5Hash.Sum(nil)) //1f83a8abbacdcf8f17061a91a83c8fbc
	md5Pwd = fmt.Sprintf("%s%s", md5Pwd, salt)

	md5Hash = md5.New()
	md5Hash.Write([]byte(md5Pwd)) // 需要加密的字符串
	md5Pwd = hex.EncodeToString(md5Hash.Sum(nil))

	return md5Pwd
}

//转化为弧度(rad)
func rad(d float64) (r float64) {
	r = d * math.Pi / 180.0
	return
}

//获取2个经纬度之间的距离
func Distance(lon1, lat1, lon2, lat2 float64) (distance float64) {
	//赤道半径(单位m)
	const EARTH_RADIUS = 6378137
	rad_lat1 := rad(lat1)
	rad_lon1 := rad(lon1)
	rad_lat2 := rad(lat2)
	rad_lon2 := rad(lon2)
	if rad_lat1 < 0 {
		rad_lat1 = math.Pi/2 + math.Abs(rad_lat1)
	}
	if rad_lat1 > 0 {
		rad_lat1 = math.Pi/2 - math.Abs(rad_lat1)
	}
	if rad_lon1 < 0 {
		rad_lon1 = math.Pi*2 - math.Abs(rad_lon1)
	}
	if rad_lat2 < 0 {
		rad_lat2 = math.Pi/2 + math.Abs(rad_lat2)
	}
	if rad_lat2 > 0 {
		rad_lat2 = math.Pi/2 - math.Abs(rad_lat2)
	}
	if rad_lon2 < 0 {
		rad_lon2 = math.Pi*2 - math.Abs(rad_lon2)
	}
	x1 := EARTH_RADIUS * math.Cos(rad_lon1) * math.Sin(rad_lat1)
	y1 := EARTH_RADIUS * math.Sin(rad_lon1) * math.Sin(rad_lat1)
	z1 := EARTH_RADIUS * math.Cos(rad_lat1)

	x2 := EARTH_RADIUS * math.Cos(rad_lon2) * math.Sin(rad_lat2)
	y2 := EARTH_RADIUS * math.Sin(rad_lon2) * math.Sin(rad_lat2)
	z2 := EARTH_RADIUS * math.Cos(rad_lat2)
	d := math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2) + (z1-z2)*(z1-z2))
	theta := math.Acos((EARTH_RADIUS*EARTH_RADIUS + EARTH_RADIUS*EARTH_RADIUS - d*d) / (2 * EARTH_RADIUS * EARTH_RADIUS))
	distance = theta * EARTH_RADIUS
	return
}

func Substr(str string, start, length int) string {
	if length == 0 {
		return ""
	}
	rune_str := []rune(str)
	len_str := len(rune_str)

	if start < 0 {
		start = len_str + start
	}
	if start > len_str {
		start = len_str
	}
	end := start + length
	if end > len_str {
		end = len_str
	}
	if length < 0 {
		end = len_str + length
	}
	if start > end {
		start, end = end, start
	}
	return string(rune_str[start:end])
}

// @title截取小数位数
func FloatRound(f float64, n int) float64 {
	format := "%." + strconv.Itoa(n) + "f"
	res, _ := strconv.ParseFloat(fmt.Sprintf(format, f), 64)
	return res
}
