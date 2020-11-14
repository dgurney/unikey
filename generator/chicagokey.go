package generator

import (
	"fmt"
	"io"
	"math/rand"

	"golang.org/x/crypto/md4"
)

// ChicagoCredentials is exactly what it says on the tin
type ChicagoCredentials struct {
	Site     string
	Password string
	Build    string
}

// Used for md4 hash generation
func getText(build string) string {
	switch {
	case build == "73f":
		return "Microsoft Chicago PDK Release, November 1993"
	case build == "73g":
		return "Microsoft Chicago PDK2 Release, December 1993"
	case build == "81":
		return "Chicago Preliminary PDK Release, January 1994"
	case build == "99":
		// up to 116
		return "Chicago Preliminary Beta 1 Release, May 1994"
	case build == "122":
		// up to 189
		return "Chicago Beta 1 Release, May 1994"
	case build == "216":
		// Up to 302
		return "Windows 95 Beta 2 Release, October 1994"
		// Internet Explorer 4
	case build == "ie4july":
		// 4.70.1169
		return "Microsoft Internet Explorer 4.0 alpha 2 July 1996 release"
	case build == "ie4sept":
		// 4.71.0225
		return "Microsoft Internet Explorer 4.0 Beta - Sept. 1996 release"
	}
	panic("invalid build")
}

func (c ChicagoCredentials) String() string {
	return c.Build + ": " + c.Site + "/" + c.Password
}

// Generate generates credentials for the build specified in c.Build
func (c ChicagoCredentials) Generate(k chan KeyGenerator) {
	site := fmt.Sprintf("%06d", rand.Intn(999999))

	// First part of the password
	pass := fmt.Sprintf("%04x", rand.Intn(65535))

	// Highly secure hashing right here
	hash := md4.New()
	text := getText(c.Build)

	io.WriteString(hash, site+pass+text)
	sum := hash.Sum(nil)

	last := fmt.Sprintf("%x%x", sum[1:2], sum[0:1])

	// Sum all ascii character codes together
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
	// Middle must be mod9'd
	c.Password = fmt.Sprintf("%s%d%s", pass, middle%9, last)
	k <- c
}
