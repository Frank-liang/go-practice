package grade

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/51reboot/golang-01-homework/lesson6/jungle85gopy/stuServer/class"
)

// Grade map from class name to it's data
type Grade struct {
	// allClasses map from class name to data
	allClasses map[string]*class.Class `json:"classes"`
	// curClass is the current class name for operation
	curClass string `json:"current"`
}

// Grade methods

func (g *Grade) searchClass(name string) bool {
	if len(name) == 0 {
		return false
	}
	for curName := range g.allClasses {
		if name == curName {
			return true
		}
	}
	return false
}

// Init to init the allClasses map
func (g *Grade) init() {
	g.allClasses = make(map[string]*class.Class)
}

// PrintGrade print the grade info of all class
func (g *Grade) PrintGrade() {
	fmt.Println("grade info:\nclass\tInfo:")
	for cName, cMap := range g.allClasses {
		fmt.Printf("%s\t%v\n", cName, cMap)
	}
	// buf, _ := json.Marshal(g)
	// fmt.Printf("\n---- marshal ---\n%v", buf)
}

// GetcurClass return the current class name
func (g *Grade) GetcurClass() string {
	return g.curClass
}

// Create to creat a new Class
func (g *Grade) Create(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("args not enougth")
	}
	name := args[0]
	if existed := g.searchClass(name); existed {
		return fmt.Errorf("class name: %s existed", name)
	}
	if len(g.allClasses) == 0 {
		g.init()
	}
	g.allClasses[name] = new(class.Class)
	g.allClasses[name].SetName(name)
	g.curClass = name
	g.PrintGrade()
	return nil
}

// Change changes the current class
func (g *Grade) Change(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("args not enougth")
	}
	name := args[0]
	if existed := g.searchClass(name); !existed {
		return g.Create(args)
	}
	g.curClass = name
	// for debug
	// g.PrintGrade()
	return nil
}

// Save save all class info
func (g *Grade) Save(args []string) error {
	// for debug
	// g.PrintGrade()
	var err error
	var fd *os.File
	if len(args) < 1 {
		return fmt.Errorf("no save file")
	}
	name := args[0]
	// file existed and override
	if checkFileExist(name) {
		if !checkYes(fmt.Sprintf("override file %s", name)) {
			return fmt.Errorf("cancel save to file %s", name)
		}
		if fd, err = os.OpenFile(name, os.O_RDWR|os.O_TRUNC, 0644); err != nil {
			return fmt.Errorf("open file error of %s", name)
		}
		defer fd.Close()
	}
	// save to new file
	if fd, err = os.Create(name); err != nil {
		return fmt.Errorf("open new file error of %s", name)
	}
	defer fd.Close()

	// saving
	buf, err := json.Marshal(g)
	if err != nil {
		return fmt.Errorf("marshal stu info error")
	}
	fmt.Println("- start to marshal:", buf)
	if _, err := fd.Write(buf); err != nil {
		return fmt.Errorf("saving error")
	}
	return nil
}

// Load load grade info from the given file
func (g *Grade) Load(args []string) error {
	var err error
	var buf []byte

	if len(args) < 1 {
		return fmt.Errorf("args not enougth")
	}
	name := args[0]
	if !checkFileExist(name) {
		return fmt.Errorf("%s not existed", name)
	}
	if !checkYes("clear up stu info in mem for load") {
		return fmt.Errorf("give up load")
	}
	if buf, err = ioutil.ReadFile(name); err != nil {
		return fmt.Errorf("read from file error")
	}
	if err := json.Unmarshal(buf, g); err != nil {
		return fmt.Errorf("unmarshal stu error")
	}
	PrintInfo("load success")
	return nil
}

// MarshalJSON my own marshal.
// cls must be a struct copy. not pointer
func (g Grade) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		AllClasses map[string]*class.Class `json:"classes"`
		CurClass   string                  `json:"current"`
	}{
		AllClasses: g.allClasses,
		CurClass:   g.curClass,
	})
}

// UnmarshalJSON un
// cls must be a pointer of Student struct
func (g *Grade) UnmarshalJSON(data []byte) error {
	var tmp struct {
		AllClasses map[string]*class.Class `json:"classes"`
		CurClass   string                  `json:"current"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		fmt.Println("err of UN,", err)
		return err
	}
	g.allClasses = tmp.AllClasses
	g.curClass = tmp.CurClass
	return nil
}

// ------- implement all Method for Class -------

// Add add student info to current class
func (g *Grade) Add(args []string) error {
	cur := g.curClass
	if cur == "" {
		return fmt.Errorf("current class is blank")
	}
	return g.allClasses[cur].Add(args)
}

// List list student info of current class
func (g *Grade) List(args []string) error {
	cur := g.GetcurClass()
	if cur == "" {
		return fmt.Errorf("current class is blank")
	}
	fmt.Println(" - current class:", cur)
	return g.allClasses[cur].List(args)
}

// Update update student info of current class
func (g *Grade) Update(args []string) error {
	cur := g.curClass
	if cur == "" {
		return fmt.Errorf("current class is blank")
	}
	return g.allClasses[cur].Update(args)
}

// Delete delete student from current class
func (g *Grade) Delete(args []string) error {
	cur := g.curClass
	if cur == "" {
		return fmt.Errorf("current class is blank")
	}
	return g.allClasses[cur].Delete(args)
}

// check file f exist or not
func checkFileExist(f string) bool {
	if _, err := os.Stat(f); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	}
	return false
}

// check yes or no for the give question
func checkYes(f string) bool {
	var s string
	for {
		PrintInfo(fmt.Sprintf("%s, y or n?", f))
		fmt.Scanf("%s", &s)
		if string(s[0]) == "y" || string(s[0]) == "Y" {
			return true
		} else if string(s[0]) == "n" || string(s[0]) == "N" {
			return false
		}
	}
}

// PrintInfo print the give string
func PrintInfo(outs string) {
	fmt.Printf("  -- [%s] %s\n", time.Now().String(), outs)
}
