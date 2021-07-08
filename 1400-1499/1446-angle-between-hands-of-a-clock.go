func angleClock(hour int, minutes int) float64 {
	var (
		minAngle  float64 = float64(minutes) * (360.0 / 60.0)
		hourAngle float64 = float64(hour%12)*(360.0/12.0) + float64(minutes)*(30.0/60.0)
	)
	delta := hourAngle - minAngle
	if delta < 0 {
		delta += 360
	}
	if delta > 180 {
		delta = 360 - delta
	}
	return delta
}
