package initialize

import (
	"catering/global"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GormMysql() *gorm.DB {
	//获取配置文件的配置
	cfg := global.Config.Mysql
	//检查配置
	if err := cfg.Check(); err != nil {
		global.Log.Error(err.Error())
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       cfg.Dsn(), // DSN data source name
		DefaultStringSize:         255,       // string 类型字段的默认长度
		DisableDatetimePrecision:  true,      // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,      // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,      // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,     // 根据当前 MySQL 版本自动配置
	}
	//连接数据库
	db, err := gorm.Open(mysql.New(mysqlConfig), getGormConfig())
	if err != nil {
		return nil
	}
	sqlDB, _ := db.DB()
	if err := sqlDB.Ping(); err != nil {
		global.Log.Error(err.Error())
		return nil
	}
	// 设置默认值
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	Options(sqlDB, WithMaxIdelConns(cfg.MaxIdleConns), WithMaxOpenConns(cfg.MaxOpenConns), WithMaxLifeTime(cfg.MaxLifeTime))
	return db
}

func getGormConfig() *gorm.Config {
	//禁用外键约束
	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}

	//NewWriter 对log.New函数的再次封装，从而实现是否通过zap打印日志
	_default := logger.New(NewWriter(log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})

	//设置logger的日志输出等级
	switch global.Config.Mysql.LogMode {
	case "silent", "Silent":
		config.Logger = _default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = _default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = _default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = _default.LogMode(logger.Info)
	default:
		config.Logger = _default.LogMode(logger.Info)
	}
	return config
}

type writer struct {
	logger.Writer
}

// NewWriter writer 构造函数
func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

// Printf 格式化打印日志
func (w *writer) Printf(message string, data ...interface{}) {
	var logZap bool
	switch global.Config.System.DbType {
	case "mysql":
		logZap = global.Config.Mysql.LogZap
	}
	//通过zap打印日志，或者其他
	if logZap {
		global.Log.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}

//Option设计模式封装mysql的额外配置
type Option func(m *sql.DB)

func WithMaxIdelConns(idle int) Option {
	return func(m *sql.DB) {
		if idle == 0 {
			return
		}
		m.SetMaxIdleConns(idle)
	}
}

func WithMaxOpenConns(open int) Option {
	return func(m *sql.DB) {
		if open == 0 {
			return
		}
		m.SetMaxOpenConns(open)
	}
}

func WithMaxLifeTime(t int) Option {
	return func(m *sql.DB) {
		if t == 0 {
			return
		}
		m.SetConnMaxLifetime(time.Duration(t) * time.Second)
	}
}
func Options(m *sql.DB, opts ...Option) {
	for _, opt := range opts {
		opt(m)
	}
}
