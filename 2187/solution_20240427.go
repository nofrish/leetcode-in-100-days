package _187

func minimumTime(time []int, totalTrips int) int64 {

	low, high := 1, time[0]*totalTrips
	for _, t := range time {
		high = min(high, t*totalTrips)
	}

	totalTime := 0
	for low <= high {
		mid := (low + high) >> 1
		if f(time, mid) < totalTrips {
			low = mid + 1
		} else {
			totalTime = mid
			high = mid - 1
		}
	}
	return int64(totalTime)
}

// f 是单调递增函数：总时间越多，趟数越多
func f(times []int, totalTime int) int {
	trips := 0
	for _, time := range times {
		trips += totalTime / time
	}
	return trips
}
