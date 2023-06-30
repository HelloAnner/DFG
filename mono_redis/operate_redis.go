package mono_redis

// redis 的一些基础操作

// KVOperate  基础的KV操作
// Set - 设置一个key的值
// Get - 查询key的值
// GetSet - 设置一个key的值，并返回这个key的旧值
// SetNX - 如果key不存在，则设置这个key的值
// MGet - 批量查询key的值
// MSet - 批量设置key的值
// Incr,IncrBy,IncrByFloat - 针对一个key的数值进行递增操作
// Decr,DecrBy - 针对一个key的数值进行递减操作
// Del - 删除key操作，可以批量删除
// Expire - 设置key的过期时间
type KVOperate interface {
	Set(k, v string) error
}
