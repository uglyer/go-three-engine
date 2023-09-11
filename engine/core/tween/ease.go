package tween

// TweenFunc 提供了用于缓动方程的接口。
// 可以使用提供的缓动函数之一或提供您自己的缓动函数。 t = 当前时间，b = 开始值，c = 自开始以来的变化，d = 持续时间

type TweenFunc func(t, b, c, d float32) float32
