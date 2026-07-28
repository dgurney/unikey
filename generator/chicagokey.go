package generator

import (
	"fmt"
	"io"
	"math/rand/v2"

	"golang.org/x/crypto/md4"
)

// ChicagoCredentials contains a Chicago site ID, password, and build.
type ChicagoCredentials struct {
	Site string
	// Password may contain the four-character password prefix before generation.
	Password string
	Build    string
}

func getText(build string) (string, error) {
	switch build {
	case "73f":
		return "Microsoft Chicago PDK Release, November 1993", nil
	case "73g":
		return "Microsoft Chicago PDK2 Release, December 1993", nil
	case "81":
		return "Chicago Preliminary PDK Release, January 1994", nil
	case "99":
		return "Chicago Preliminary Beta 1 Release, May 1994", nil
	case "122":
		return "Chicago Beta 1 Release, May 1994", nil
	case "216":
		return "Windows 95 Beta 2 Release, October 1994", nil
	case "ie4july":
		return "Microsoft Internet Explorer 4.0 alpha 2 July 1996 release", nil
	case "ie4sept":
		return "Microsoft Internet Explorer 4.0 Beta - Sept. 1996 release", nil
	default:
		return "", fmt.Errorf("invalid build %q", build)
	}
}

func (c ChicagoCredentials) String() string {
	return c.Site + "/" + c.Password
}

// Generate populates credentials for c.Build.
func (c *ChicagoCredentials) Generate() error {
	site := c.Site
	if c.Site == "" {
		site = fmt.Sprintf("%06d", rand.IntN(1_000_000))
	}

	pass := c.Password
	if c.Password == "" {
		pass = fmt.Sprintf("%04x", rand.IntN(1<<16))
	}

	hash := md4.New()
	text, err := getText(c.Build)
	if err != nil {
		return err
	}

	io.WriteString(hash, site+pass+text)
	sum := hash.Sum(nil)

	last := fmt.Sprintf("%x%x", sum[1:2], sum[0:1])

	middle := 0
	for i := range site {
		middle += int(site[i])
	}
	for i := range pass {
		middle += int(pass[i])
	}
	for i := range last {
		middle += int(last[i])
	}

	c.Site = site
	c.Password = fmt.Sprintf("%s%d%s", pass, middle%9, last)
	return nil
}
