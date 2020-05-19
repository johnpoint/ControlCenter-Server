package apis

import (
	"main/src/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func getSiteInfo(c echo.Context) error {
	user := checkAuth(c)
	if user != nil {
		site := model.Site{}
		if err := c.Bind(&site); err != nil {
			panic(err)
		}
		return c.JSON(http.StatusOK, getSite(site))
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "ERROR"})
}

func addSiteInfo(c echo.Context) error {
	site := model.Site{}
	if err := c.Bind(&site); err != nil {
		panic(err)
	}
	if !addSite(site) {
		return c.JSON(http.StatusBadGateway, model.Callback{Code: 0, Info: "ERROR"})
	}
	return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "OK"})
}

func deleteSiteInfo(c echo.Context) error {
	site := model.Site{}
	if err := c.Bind(&site); err != nil {
		panic(err)
	}
	if !delSite(site) {
		return c.JSON(http.StatusBadGateway, model.Callback{Code: 0, Info: "ERROR"})
	}
	return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "OK"})
}

func linkServerSite(c echo.Context) error {
	sid := c.Param("ServerID")
	cid := c.Param("SiteID")
	Isid, _ := strconv.ParseInt(sid, 10, 64)
	Icid, _ := strconv.ParseInt(cid, 10, 64)
	payload := model.ServerLink{ServerID: Isid, SiteID: Icid}
	data := getLinkSite(payload)
	if len(data) == 0 {
		if (LinkServer(model.ServerLink{ServerID: Isid, SiteID: Icid})) {
			return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "OK"})
		}
		return c.JSON(http.StatusOK, model.Callback{Code: 0, Info: "ERROR"})
	} else {
		return c.JSON(http.StatusOK, model.Callback{Code: 0, Info: "already linked"})
	}
}

func unLinkServerSite(c echo.Context) error {
	sid := c.Param("ServerID")
	cid := c.Param("SiteID")
	Isid, _ := strconv.ParseInt(sid, 10, 64)
	Icid, _ := strconv.ParseInt(cid, 10, 64)
	payload := model.ServerLink{ServerID: Isid, SiteID: Icid}
	data := UnLinkServer(payload)
	if data {
		if len(getLinkCer(payload)) == 0 {
			return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "OK"})
		}
	}
	return c.JSON(http.StatusOK, model.Callback{Code: 0, Info: "ERROR"})
}

func getCertificatesInfo(c echo.Context) error {
	sid := c.Param("SiteID")
	Isid, _ := strconv.ParseInt(sid, 10, 64)
	SiteInfo := getSite(model.Site{ID: Isid})
	if len(SiteInfo) != 0 {
		SiteName := SiteInfo[0].Name
		resp, _ := http.Get("https://" + SiteName)
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "SSL error"})
		} else {
			return c.JSON(http.StatusOK, resp.TLS.PeerCertificates[0])
		}
	}
	return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "ERROR"})
}
