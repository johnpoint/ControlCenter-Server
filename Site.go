package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func getSiteInfo(c echo.Context) error {
	user := checkAuth(c)
	if user != nil {
		site := Site{}
		if err := c.Bind(&site); err != nil {
			panic(err)
		}
		return c.JSON(http.StatusOK, getSite(site))
	}
	return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "ERROR"})
}

func addSiteInfo(c echo.Context) error {
	site := Site{}
	if err := c.Bind(&site); err != nil {
		panic(err)
	}
	if !addSite(site) {
		return c.JSON(http.StatusBadGateway, Callback{Code: 0, Info: "ERROR"})
	}
	return c.JSON(http.StatusOK, Callback{Code: 200, Info: "OK"})
}

func deleteSiteInfo(c echo.Context) error {
	site := Site{}
	if err := c.Bind(&site); err != nil {
		panic(err)
	}
	if !delSite(site) {
		return c.JSON(http.StatusBadGateway, Callback{Code: 0, Info: "ERROR"})
	}
	return c.JSON(http.StatusOK, Callback{Code: 200, Info: "OK"})
}

func linkServerSite(c echo.Context) error {
	sid := c.Param("ServerID")
	cid := c.Param("SiteID")
	Isid, _ := strconv.ParseInt(sid, 10, 64)
	Icid, _ := strconv.ParseInt(cid, 10, 64)
	payload := ServerLink{ServerID: Isid, SiteID: Icid}
	data := getLinkSite(payload)
	if len(data) == 0 {
		if (LinkServer(ServerLink{ServerID: Isid, SiteID: Icid})) {
			return c.JSON(http.StatusOK, Callback{Code: 200, Info: "OK"})
		}
		return c.JSON(http.StatusOK, Callback{Code: 0, Info: "ERROR"})
	} else {
		return c.JSON(http.StatusOK, Callback{Code: 0, Info: "already linked"})
	}
}

func unLinkServerSite(c echo.Context) error {
	sid := c.Param("ServerID")
	cid := c.Param("SiteID")
	Isid, _ := strconv.ParseInt(sid, 10, 64)
	Icid, _ := strconv.ParseInt(cid, 10, 64)
	payload := ServerLink{ServerID: Isid, SiteID: Icid}
	data := UnLinkServer(payload)
	if data {
		if len(getLinkCer(payload)) == 0 {
			return c.JSON(http.StatusOK, Callback{Code: 200, Info: "OK"})
		}
	}
	return c.JSON(http.StatusOK, Callback{Code: 0, Info: "ERROR"})
}

func getCertificatesInfo(c echo.Context) error {
	sid := c.Param("SiteID")
	Isid, _ := strconv.ParseInt(sid, 10, 64)
	SiteInfo := getSite(Site{ID: Isid})
	if len(SiteInfo) != 0 {
		SiteName := SiteInfo[0].Name
		resp, _ := http.Get("https://" + SiteName)
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return c.JSON(http.StatusOK, Callback{Code: 200, Info: "SSL error"})
		} else {
			return c.JSON(http.StatusOK, resp.TLS.PeerCertificates[0])
		}
	}
	return c.JSON(http.StatusOK, Callback{Code: 200, Info: "ERROR"})
}
