package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func ReadLineWithJobName(line string, k int) string {
	if strings.Contains(line, "JOBNAME=") {
		hclRe := regexp.MustCompile(`JOBNAME="([^"]+)`)
		jobName := hclRe.FindString(line)
		jobNameIndex := strings.Index(jobName, `"`)
		jobNameVal := jobName[jobNameIndex+1 : len(jobName)]
		if strings.Contains(jobNameVal, " ") == false {
			return line
		}
		jobNameValNew := strings.Replace(jobNameVal, " ", "_", -1)
		tmp := fmt.Sprintf(`JOBNAME="%s"`, jobNameVal)
		tmpNew := fmt.Sprintf(`JOBNAME="%s"`, jobNameValNew)
		fmt.Printf("%s, %s\n", tmp, tmpNew)

		// fmt.Println(jobNameVal)
		// matches := hclRe.MatchString(line)
		line = strings.Replace(line, tmp, tmpNew, -1)
		fmt.Printf("%d: [%v, %v]  \n", k, jobNameVal, jobNameValNew)
	}
	return line
}
func ReadXml() {
	c, err := ioutil.ReadFile("2.xml")
	if err != nil {
		fmt.Println(err)
	}
	// lines := bytes.Split(c, []byte("\n"))
	lines := strings.Split(string(c), "\n")

	for k, line := range lines {
		if len(line) == 0 {
			continue
		}
		// fmt.Println(string(line))
		lines[k] = ReadLineWithJobName(line, k)
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile("2_auto.xml", []byte(output), 0644)
	if err != nil {
		fmt.Println("write error=", err)
	}

	/* fmt.Println("lines:", len(lines)-1)
	fmt.Println("last", string(lines[len(lines)-2])) */

}

func main() {
	ReadXml()
}
