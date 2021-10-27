package setting

import (
	"time"
)

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
}

type DatabaseSettingS struct {
	DBType       string
	Username     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

func (s *Setting) ReadSection(k string, v interface{}) error {

	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	//if reflect.ValueOf(v).IsNil() {
	//	return errors.New(fmt.Sprintf("% config not found", reflect.TypeOf(v)))
	//}
	return nil

}
