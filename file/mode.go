func NewMode(attributes string) (os.FileMode, error) {
	if len(attributes) != 10 {
		return result, fmt.Errorf("invalid attribute length %v %v", attributes, len(attributes))
	var fileModePosition = strings.Index(fileType, string(attributes[0]))
		if c == rune(attributes[i+1]) {