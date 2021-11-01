package etc

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

const (
	issue = "/etc/os-release"
)

func GetUbuntuRelease() (version string, err error) {
	file, err := os.Open(issue)
	if err != nil {
		return "", err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	re := regexp.MustCompile(`VERSION_ID=".*"$`)
	for scanner.Scan() {
		line := scanner.Text()
		match := re.Find([]byte(line))
		if match != nil {
			version = string(match)
			version = strings.TrimSuffix(version, `"`)
			version = strings.TrimPrefix(version, `VERSION_ID="`)
			return version, nil
		}
	}
	return "", nil
}
