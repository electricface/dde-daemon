package dock

import (
	"encoding/base64"
	"io/ioutil"
	"path/filepath"
	"pkg.linuxdeepin.com/lib/gio-2.0"
	"strings"
)

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func isEntryNameValid(name string) bool {
	if !strings.HasPrefix(name, entryDestPrefix) {
		return false
	}
	return true
}

func getEntryId(name string) (string, bool) {
	a := strings.SplitN(name, entryDestPrefix, 2)
	if len(a) >= 1 {
		return a[len(a)-1], true
	}
	return "", false
}

func guess_desktop_id(oldId string) string {
	allApp := gio.AppInfoGetAll()
	for _, app := range allApp {
		baseName := filepath.Base(gio.ToDesktopAppInfo(app).GetFilename())
		lowerBaseName := strings.ToLower(baseName)
		if oldId == lowerBaseName || strings.Replace(oldId, "-", "_", -1) == lowerBaseName {
			return baseName
		}
	}

	return ""
}

func getAppIcon(core *gio.DesktopAppInfo) string {
	gioIcon := core.GetIcon()
	if gioIcon == nil {
		logger.Warning("get icon from appinfo failed")
		return ""
	}

	logger.Debug("GetIcon:", gioIcon.ToString())
	icon := get_theme_icon(gioIcon.ToString(), 48)
	if icon == "" {
		logger.Warning("get icon from theme failed")
		return ""
	}

	logger.Debug("get_theme_icon:", icon)
	// the filepath.Ext return ".xxx"
	ext := filepath.Ext(icon)[1:]
	logger.Debug("ext:", ext)
	if strings.EqualFold(ext, "xpm") {
		logger.Debug("change xpm to data uri")
		return xpm_to_dataurl(icon)
	}

	return icon
}

func dataUriToFile(dataUri, path string) (string, error) {
	commaIndex := strings.Index(dataUri, ",")
	img, err := base64.StdEncoding.DecodeString(dataUri[commaIndex+1:])
	if err != nil {
		return path, err
	}

	return path, ioutil.WriteFile(path, img, 0744)
}