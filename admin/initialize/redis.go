package initialize

func Redis() {
	// redisCfg := global.Config.Redis
	// client := redis.NewClient(&redis.Options{
	// 	Addr:     redisCfg.Addr,
	// 	Password: redisCfg.Password, // no password set
	// 	DB:       redisCfg.DB,       // use default DB
	// })
	// pong, err := client.Ping(context.Background()).Result()
	// if err != nil {
	// 	global.Log.Error("redis connect ping failed, err:", zap.Error(err))
	// } else {
	// 	global.Log.Info("redis connect ping response:", zap.String("pong", pong))
	// 	global.GVA_REDIS = client
	// }
}
