package tween

type (
	// Tween 补间动画
	Tween struct {
		duration float32
		time     float32
		begin    float32
		end      float32
		change   float32
		Overflow float32
		easing   TweenFunc
		reverse  bool
	}
)

// New 构建一个新的过渡动画对象
func New(begin, end, duration float32, easing TweenFunc) *Tween {
	return &Tween{
		begin:    begin,
		end:      end,
		change:   end - begin,
		duration: duration,
		easing:   easing,
		Overflow: 0,
		reverse:  false,
	}
}

// Set 设置时间到指定值
func (tween *Tween) Set(time float32) (current float32, isFinished bool) {
	switch {
	case time <= 0:
		tween.Overflow = time
		tween.time = 0
		current = tween.begin
	case time >= tween.duration:
		tween.Overflow = time - tween.duration
		tween.time = tween.duration
		current = tween.end
	default:
		tween.Overflow = 0
		tween.time = time
		current = tween.easing(tween.time, tween.begin, tween.change, tween.duration)
	}

	if tween.reverse {
		return current, tween.time <= 0
	}
	return current, tween.time >= tween.duration
}

// Reset 会将 Tween 设置为两个值的开头。
func (tween *Tween) Reset() {
	if tween.reverse {
		tween.Set(tween.duration)
	} else {
		tween.Set(0)
	}
}

// Update 更新时间
func (tween *Tween) Update(dt float32) (current float32, isFinished bool) {
	if tween.reverse {
		return tween.Set(tween.time - dt)
	}
	return tween.Set(tween.time + dt)
}
