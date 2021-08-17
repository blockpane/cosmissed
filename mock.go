package missed

// no longer needed DELETEME
//func mkPex() []byte {
//	lines := make([]line3d, 0)
//	hub := point{}
//	mult1 := 1
//	if rand.Intn(3) % 2 == 0 {
//		mult1 = -1
//	}
//	hub[0] = float32(rand.Intn(100) * mult1)
//	hub[1] = float32(rand.Intn(100))
//	for i := 0; i < 200; i++{
//		next := line3d{}
//		mult2 := 1
//		if rand.Intn(3) % 2 == 0 {
//			mult2 = -1
//		}
//		next[0][0] = hub[0]
//		next[0][1] = hub[1]
//		if i % 2 == 0 {
//			next[0][0] = float32(rand.Intn(100) * mult2)
//			next[0][1] = float32(rand.Intn(50)+20)
//			next[1][0] = hub[0]
//			next[1][1] = hub[1]
//		} else {
//			next[1][0] = float32(rand.Intn(120) * mult2 + 30)
//			next[1][1] = float32(rand.Intn(50)+20)
//		}
//		lines = append(lines, next)
//		if i % 13 == 0 {
//			mult1 = 1
//			if rand.Intn(3) % 2 == 0 {
//				mult1 = -1
//			}
//			hub[0] = float32(rand.Intn(100) * mult1)
//			hub[1] = float32(rand.Intn(50))
//		}
//	}
//	j, _ := json.Marshal(lines)
//	return j
//}

