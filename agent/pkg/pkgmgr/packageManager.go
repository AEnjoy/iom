//go:build !windows
// +build !windows

package pkgmgr

import (
	"bufio"
	pb "IOM/agent/proto"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func getPackageManager(pkg string) string {
	switch pkg {
	case "debian":
		return "apt-get"
	case "ubuntu":
		return "apt-get"
	case "centos":
		return "yum"
	case "fedora":
		return "yum"
	case "opensuse":
		return "zypper"
	case "sles":
		return "zypper"
	case "arch":
		return "pacman"
	case "alpine":
		return "apk"
	default:
		return "apt-get"
	}
}

func yumGetInstalledPackages() []rpmPackage {
	var packages []rpmPackage
	cmd := exec.Command("yum", "list", "installed")
	out, err := cmd.CombinedOutput()
	os.WriteFile("out", out, 0666)
	if err == nil {
		f, _ := os.Open("out")
		r := bufio.NewReader(f)
		r.ReadString('\n')
		for {
			s, err := r.ReadString('\n')
			if err == io.EOF {
				break
			} else {
				var pack rpmPackage
				r := regexp.MustCompile("[^\\s]+")
				t := r.FindAllString(s, -1)
				t1 := strings.Split(t[0], ".")
				pack.name = t1[0]
				pack.arch = t1[1]
				pack.version = t[1]
				pack.repository = t[2]
				packages = append(packages, pack)
			}
		}
		f.Close()
		os.Remove("out")
		return packages
	}
	return nil
}
func debGetInstalledPackages() []debPackage {
	var packages []debPackage
	cmd := exec.Command("dpkg-query", "-W", "--showformat=${Package} ${Version} ${Architecture}\\n")
	out, err := cmd.CombinedOutput()
	os.WriteFile("out", out, 0666)
	if err == nil {
		f, _ := os.Open("out")
		r := bufio.NewReader(f)
		for {
			s, err := r.ReadString('\n')
			if err == io.EOF {
				break
			} else {
				var pack debPackage
				t := strings.Split(s, " ")
				pack.name = t[0]
				pack.version = t[1]
				pack.arch = t[2]
				//pack.description = t[3]
				packages = append(packages, pack)
			}
		}
		f.Close()
		os.Remove("out")
		return packages
	}
	return nil
}
func apkGetInstalledPackages() []apkPackage {
	var packages []apkPackage
	cmd := exec.Command("apk", "list", "-I")
	out, err := cmd.CombinedOutput()
	os.WriteFile("out", out, 0666)
	if err == nil {
		f, _ := os.Open("out")
		r := bufio.NewReader(f)
		for {
			s, err := r.ReadString('\n')
			if err == io.EOF {
				break
			} else {
				var pack apkPackage
				t := strings.Split(s, " ")
				pack.name = strings.Replace(t[2], "{", "", -1)
				pack.name = strings.Replace(pack.name, "}", "", -1)
				pack.version = t[0]
				pack.arch = t[1]
				packages = append(packages, pack)
			}
		}
		f.Close()
		os.Remove("out")
		return packages
	}
	return nil
}

func TimeToGetPackage() *pb.Packages {
	var packages pb.Packages
	switch getPkgManager() {
	case "apt-get":
		p := debGetInstalledPackages()
		for _, d := range p {
			var pack pb.Package
			pack.Name = d.name
			pack.Arch = d.arch
			pack.Description = d.description
			pack.Version = d.version
			packages.Packages = append(packages.Packages, &pack)
		}
		packages.Type = 0
	case "yum":
		p := yumGetInstalledPackages()
		for _, d := range p {
			var pack pb.Package
			pack.Name = d.name
			pack.Arch = d.arch
			pack.Description = d.description
			pack.Version = d.version
			packages.Packages = append(packages.Packages, &pack)
		}
		packages.Type = 1
	case "apk":
		p := apkGetInstalledPackages()
		for _, d := range p {
			var pack pb.Package
			pack.Name = d.name
			pack.Arch = d.arch
			pack.Description = d.description
			pack.Version = d.version
			packages.Packages = append(packages.Packages, &pack)
		}
		packages.Type = 2
	default:
		packages.Type = -1
	}
	return &packages
}

func getPkgManager() string {
	pkgManager := ""
	cmd := exec.Command("which", "apt-get")
	if cmd.Run() == nil {
		pkgManager = "apt-get"
	}
	cmd = exec.Command("which", "yum")
	if cmd.Run() == nil {
		pkgManager = "yum"
	}
	cmd = exec.Command("which", "apk")
	if cmd.Run() == nil {
		pkgManager = "apk"
	}
	return pkgManager
}
