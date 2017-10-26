package useragent

import "net/url"

// Keep them sorted
var libraries = map[string]*url.URL{
	"curl":      u("https://curl.haxx.se/"),
	"HTTPie":    u("https://github.com/jakubroztocil/httpie"),
	"PhantomJS": u("http://phantomjs.org/"),
}

// Keep them sorted
var crawlers = map[string]*url.URL{
	"Google AdsBot":    u("https://support.google.com/webmasters/answer/1061943"),
	"Google AdSense":   u("https://support.google.com/webmasters/answer/1061943"),
	"Googlebot":        u("http://www.google.com/bot.html"),
	"Googlebot Images": u("https://support.google.com/webmasters/answer/1061943"),
	"Googlebot News":   u("https://support.google.com/news/publisher/answer/93977"),
	"Googlebot Video":  u("https://support.google.com/webmasters/answer/1061943"),
}

// Keep them sorted
var browsers = map[string]*url.URL{
	"Chrome":    u("http://www.chromium.org/"),
	"Dillo":     u("http://www.dillo.org/"),
	"Edge":      u("https://www.microsoft.com/en-us/windows/microsoft-edge"),
	"Firefox":   u("https://www.mozilla.org/en-US/firefox"),
	"IceCat":    u("https://www.gnu.org/software/gnuzilla/"),
	"Iceweasel": u("https://wiki.debian.org/Iceweasel"),
	"NetSurf":   u("http://www.netsurf-browser.org/"),
	"Opera":     u("http://www.opera.com/"),
	"Silk":      u("http://aws.amazon.com/documentation/silk/"),
	"WebView":   u("http://developer.android.com/guide/webapps/webview.html"),
}

func parseGeneric(l *lex) *UserAgent {
	ua := new()
	if !parseNameVersion(l, ua) {
		return nil
	}

	if url, ok := libraries[ua.Name]; ok {
		ua.Type = Library
		ua.URL = url
		return ua
	}

	if url, ok := browsers[ua.Name]; ok {
		ua.Type = Browser
		ua.URL = url
		return ua
	}

	if url, ok := crawlers[ua.Name]; ok {
		ua.Type = Crawler
		ua.URL = url
		return ua
	}

	return nil
}
