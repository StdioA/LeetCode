package main

func destCity(paths [][]string) string {
	var pathMap = make(map[string]string)
	for _, path := range paths {
		pathMap[path[0]] = path[1]
	}
	loc := paths[0][1]
	_, ok := pathMap[loc]
	for ok {
		loc = pathMap[loc]
		_, ok = pathMap[loc]
	}
	return loc
}
