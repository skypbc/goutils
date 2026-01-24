package internal

// findFirstNonTrimChar ищет позицию первого символа, который не является служебным символом (./\).
func findFirstNonTrimChar(path string) int {
	for pos, r := range path {
		if r == '.' || r == '/' || r == '\\' {
			continue
		}
		return pos
	}
	return -1
}

// findLastNonTrimChar ищет позицию последнего символа, который не является служебным символом (./\).
func findLastNonTrimChar(path string) int {
	end := len(path) - 1
	for end >= 0 {
		if path[end] == '.' || path[end] == '/' || path[end] == '\\' {
			end--
		} else {
			break
		}
	}
	return end
}

// CleanPath удаляет лишние служебные символы ./\ в начале и конце пути.
func CleanPath(path string) string {
	n := len(path)
	if n == 0 {
		return path
	}

	start := findFirstNonTrimChar(path)
	if start < 0 {
		return ""
	}
	end := findLastNonTrimChar(path)
	if end < 0 || end < start {
		return ""
	}

	// Если перед и после пути есть "точки" их нужно восстановить
	for {
		if start > 0 && path[start-1] == '.' {
			start--
			continue
		}
		if end < n-1 && path[end+1] == '.' {
			end++
			continue
		}
		break
	}

	return path[start : end+1]
}
