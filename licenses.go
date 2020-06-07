// Generated with by licrep version v0.3.0
// https://github.com/kopoli/licrep
// Called with: licrep -o licenses.go

package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io"
	"strings"
)

// License is a representation of an embedded license.
type License struct {
	// Name of the license
	Name string

	// The text of the license
	Text string
}

// GetLicenses gets a map of Licenses where the keys are
// the package names.
func GetLicenses() (map[string]License, error) {
	type EncodedLicense struct {
		Name string
		Text string
	}
	data := map[string]EncodedLicense{

		"github.com/casimir/xdg-go": EncodedLicense{
			Name: "MIT",
			Text: `
H4sIAAAAAAAC/1xRzY7rJhTe8xSfZjUjWdMfqZvuGJvEqDZEmNw0S2KTMZUDliEd5e0rnMy9nbuyzDnf
79GjRcs1GtdbHy2eW65fCCnDfFvc+5jw3L/g919/+wOtWZLzKEfjvCVkZ5eLi9EFDxcx2sWebnhfjE92
KHBerEU4ox/N8m4LpADjb5jtEoNHOCXjvPPvMOjDfCPhjDS6iBjO6cMsFsYPMDGG3plkBwyhv16sTyZl
vbObbMRzGi2eugfi6WUVGayZiPPIs88RPlwawzVhsTEtrs8cBZzvp+uQPXyOJ3dxD4UMX+NHkgKu0Rar
zwKXMLhz/to11nw9TS6OBQaXqU/XZAvE/Li2WeQcv4QF0U4T6cPsbMSa9Ye7dSdbn3Oh6VFRzC8fY7h8
TeIiOV8X7+JoV8wQEMOq+I/tU37J6+cwTeEjR+uDH1xOFP8kJJ/anMK/ds1yv64PyfX3utcDzD+u+hjF
0UwTTvZRmB3gPMz/4ixZPibjkzMT5rCsej/HfCVE1wyd3OgDVQy8w07Jb7xiFZ5oB949FThwXcu9xoEq
RYU+Qm5AxRF/cVEVYH/vFOs6SEV4u2s4qwpwUTb7iost3vYaQmo0vOWaVdASWfBBxVmXyVqmypoKTd94
w/WxIBuuRebcSAWKHVWal/uGKuz2aic7BioqCCm42CgutqxlQr+CCwgJ9o0Jja6mTZOlCN3rWqrsD6Xc
HRXf1hq1bCqmOrwxNJy+NewuJY4oG8rbAhVt6ZatKKlrpkheu7vDoWb5KetRAVpqLkWOUUqhFS11AS2V
/g498I4VoIp3uZCNkm1Bcp1yk1e4yDjB7iy5any5iFTr/75j3wlRMdpwse0yOEf8XH4l/wUAAP//kAzo
9TgEAAA=`,
		},
		"github.com/jawher/mow.cli": EncodedLicense{
			Name: "MIT",
			Text: `
H4sIAAAAAAAC/3xRvY7jNhjs+RSDre4AYvODVOm4Em0xkUiDos9xSUv0ioEsGiIVw28fUPbeYVNEjcDv
Z+abGfz/ZwaHRhjUvnNTdIQU4Xqf/fuQ8KX7il9//uU3ij/sbXAzmrDEaAnZufniY/Rhgo8Y3OxOd7zP
dkqupzjPziGc0Q12fncUKcBOd1zdHMOEcErWT356h0UXrncSzkiDj4jhnG52drBTDxtj6LxNrkcfuuXi
pmRT5jv70UV8SYPDS/vcePm6kvTOjsRPyL2PFm4+DWFJmF1Ms+8yBoWfunHp8w0f7dFf/JMhr6/6I0kB
S3R0vZPiEnp/zn+3yroup9HHgaL3Gfq0JEcRc3E1kmYdP4UZ0Y0j6cLVu4hV64/r1pl8+jUbmp4WxVy5
DeHyWYmP5LzMk4+DW3f6gBhWxr9dl3Ilj5/DOIZbltaFqfdZUfydkJyyPYV/3KrlEe8Uku8edq8BXH+k
+mzFwY4jTu5pmOvhJ5JLH3LmTB+TnZK3I65hXvn+K/OVEFNxtGpjDkxziBY7rb6Jkpd4YS1E+0JxEKZS
e4MD05pJc4TagMkj/hSypOB/7TRvWyhNRLOrBS8phCzqfSnkFm97A6kMatEIw0sYhUz4hBK8zWAN10XF
pGFvohbmSMlGGJkxN0qDYce0EcW+Zhq7vd6ploPJElJJITdayC1vuDSvEBJSgX/j0qCtWF1nKsL2plI6
34dC7Y5abCuDStUl1y3eOGrB3mr+oJJHFDUTDUXJGrbl65YyFdckjz2uw6HiuZT5mAQrjFAyyyiUNJoV
hsIobb6vHkTLKZgWbTZko1VDSbZTbfKIkHlP8gdKthqfElF6fe9b/h0QJWe1kNsWQn6K75WQfwMAAP//
eVpV+1MEAAA=`,
		},
		"github.com/jawher/mow.cli/internal/container": EncodedLicense{
			Name: "MIT",
			Text: `
H4sIAAAAAAAC/3xRvY7jNhjs+RSDre4AYvODVOm4Em0xkUiDos9xSUv0ioEsGiIVw28fUPbeYVNEjcDv
Z+abGfz/ZwaHRhjUvnNTdIQU4Xqf/fuQ8KX7il9//uU3ij/sbXAzmrDEaAnZufniY/Rhgo8Y3OxOd7zP
dkqupzjPziGc0Q12fncUKcBOd1zdHMOEcErWT356h0UXrncSzkiDj4jhnG52drBTDxtj6LxNrkcfuuXi
pmRT5jv70UV8SYPDS/vcePm6kvTOjsRPyL2PFm4+DWFJmF1Ms+8yBoWfunHp8w0f7dFf/JMhr6/6I0kB
S3R0vZPiEnp/zn+3yroup9HHgaL3Gfq0JEcRc3E1kmYdP4UZ0Y0j6cLVu4hV64/r1pl8+jUbmp4WxVy5
DeHyWYmP5LzMk4+DW3f6gBhWxr9dl3Ilj5/DOIZbltaFqfdZUfydkJyyPYV/3KrlEe8Uku8edq8BXH+k
+mzFwY4jTu5pmOvhJ5JLH3LmTB+TnZK3I65hXvn+K/OVEFNxtGpjDkxziBY7rb6Jkpd4YS1E+0JxEKZS
e4MD05pJc4TagMkj/hSypOB/7TRvWyhNRLOrBS8phCzqfSnkFm97A6kMatEIw0sYhUz4hBK8zWAN10XF
pGFvohbmSMlGGJkxN0qDYce0EcW+Zhq7vd6ploPJElJJITdayC1vuDSvEBJSgX/j0qCtWF1nKsL2plI6
34dC7Y5abCuDStUl1y3eOGrB3mr+oJJHFDUTDUXJGrbl65YyFdckjz2uw6HiuZT5mAQrjFAyyyiUNJoV
hsIobb6vHkTLKZgWbTZko1VDSbZTbfKIkHlP8gdKthqfElF6fe9b/h0QJWe1kNsWQn6K75WQfwMAAP//
eVpV+1MEAAA=`,
		},
		"github.com/jawher/mow.cli/internal/flow": EncodedLicense{
			Name: "MIT",
			Text: `
H4sIAAAAAAAC/3xRvY7jNhjs+RSDre4AYvODVOm4Em0xkUiDos9xSUv0ioEsGiIVw28fUPbeYVNEjcDv
Z+abGfz/ZwaHRhjUvnNTdIQU4Xqf/fuQ8KX7il9//uU3ij/sbXAzmrDEaAnZufniY/Rhgo8Y3OxOd7zP
dkqupzjPziGc0Q12fncUKcBOd1zdHMOEcErWT356h0UXrncSzkiDj4jhnG52drBTDxtj6LxNrkcfuuXi
pmRT5jv70UV8SYPDS/vcePm6kvTOjsRPyL2PFm4+DWFJmF1Ms+8yBoWfunHp8w0f7dFf/JMhr6/6I0kB
S3R0vZPiEnp/zn+3yroup9HHgaL3Gfq0JEcRc3E1kmYdP4UZ0Y0j6cLVu4hV64/r1pl8+jUbmp4WxVy5
DeHyWYmP5LzMk4+DW3f6gBhWxr9dl3Ilj5/DOIZbltaFqfdZUfydkJyyPYV/3KrlEe8Uku8edq8BXH+k
+mzFwY4jTu5pmOvhJ5JLH3LmTB+TnZK3I65hXvn+K/OVEFNxtGpjDkxziBY7rb6Jkpd4YS1E+0JxEKZS
e4MD05pJc4TagMkj/hSypOB/7TRvWyhNRLOrBS8phCzqfSnkFm97A6kMatEIw0sYhUz4hBK8zWAN10XF
pGFvohbmSMlGGJkxN0qDYce0EcW+Zhq7vd6ploPJElJJITdayC1vuDSvEBJSgX/j0qCtWF1nKsL2plI6
34dC7Y5abCuDStUl1y3eOGrB3mr+oJJHFDUTDUXJGrbl65YyFdckjz2uw6HiuZT5mAQrjFAyyyiUNJoV
hsIobb6vHkTLKZgWbTZko1VDSbZTbfKIkHlP8gdKthqfElF6fe9b/h0QJWe1kNsWQn6K75WQfwMAAP//
eVpV+1MEAAA=`,
		},
		"github.com/jawher/mow.cli/internal/fsm": EncodedLicense{
			Name: "MIT",
			Text: `
H4sIAAAAAAAC/3xRvY7jNhjs+RSDre4AYvODVOm4Em0xkUiDos9xSUv0ioEsGiIVw28fUPbeYVNEjcDv
Z+abGfz/ZwaHRhjUvnNTdIQU4Xqf/fuQ8KX7il9//uU3ij/sbXAzmrDEaAnZufniY/Rhgo8Y3OxOd7zP
dkqupzjPziGc0Q12fncUKcBOd1zdHMOEcErWT356h0UXrncSzkiDj4jhnG52drBTDxtj6LxNrkcfuuXi
pmRT5jv70UV8SYPDS/vcePm6kvTOjsRPyL2PFm4+DWFJmF1Ms+8yBoWfunHp8w0f7dFf/JMhr6/6I0kB
S3R0vZPiEnp/zn+3yroup9HHgaL3Gfq0JEcRc3E1kmYdP4UZ0Y0j6cLVu4hV64/r1pl8+jUbmp4WxVy5
DeHyWYmP5LzMk4+DW3f6gBhWxr9dl3Ilj5/DOIZbltaFqfdZUfydkJyyPYV/3KrlEe8Uku8edq8BXH+k
+mzFwY4jTu5pmOvhJ5JLH3LmTB+TnZK3I65hXvn+K/OVEFNxtGpjDkxziBY7rb6Jkpd4YS1E+0JxEKZS
e4MD05pJc4TagMkj/hSypOB/7TRvWyhNRLOrBS8phCzqfSnkFm97A6kMatEIw0sYhUz4hBK8zWAN10XF
pGFvohbmSMlGGJkxN0qDYce0EcW+Zhq7vd6ploPJElJJITdayC1vuDSvEBJSgX/j0qCtWF1nKsL2plI6
34dC7Y5abCuDStUl1y3eOGrB3mr+oJJHFDUTDUXJGrbl65YyFdckjz2uw6HiuZT5mAQrjFAyyyiUNJoV
hsIobb6vHkTLKZgWbTZko1VDSbZTbfKIkHlP8gdKthqfElF6fe9b/h0QJWe1kNsWQn6K75WQfwMAAP//
eVpV+1MEAAA=`,
		},
		"github.com/jawher/mow.cli/internal/lexer": EncodedLicense{
			Name: "MIT",
			Text: `
H4sIAAAAAAAC/3xRvY7jNhjs+RSDre4AYvODVOm4Em0xkUiDos9xSUv0ioEsGiIVw28fUPbeYVNEjcDv
Z+abGfz/ZwaHRhjUvnNTdIQU4Xqf/fuQ8KX7il9//uU3ij/sbXAzmrDEaAnZufniY/Rhgo8Y3OxOd7zP
dkqupzjPziGc0Q12fncUKcBOd1zdHMOEcErWT356h0UXrncSzkiDj4jhnG52drBTDxtj6LxNrkcfuuXi
pmRT5jv70UV8SYPDS/vcePm6kvTOjsRPyL2PFm4+DWFJmF1Ms+8yBoWfunHp8w0f7dFf/JMhr6/6I0kB
S3R0vZPiEnp/zn+3yroup9HHgaL3Gfq0JEcRc3E1kmYdP4UZ0Y0j6cLVu4hV64/r1pl8+jUbmp4WxVy5
DeHyWYmP5LzMk4+DW3f6gBhWxr9dl3Ilj5/DOIZbltaFqfdZUfydkJyyPYV/3KrlEe8Uku8edq8BXH+k
+mzFwY4jTu5pmOvhJ5JLH3LmTB+TnZK3I65hXvn+K/OVEFNxtGpjDkxziBY7rb6Jkpd4YS1E+0JxEKZS
e4MD05pJc4TagMkj/hSypOB/7TRvWyhNRLOrBS8phCzqfSnkFm97A6kMatEIw0sYhUz4hBK8zWAN10XF
pGFvohbmSMlGGJkxN0qDYce0EcW+Zhq7vd6ploPJElJJITdayC1vuDSvEBJSgX/j0qCtWF1nKsL2plI6
34dC7Y5abCuDStUl1y3eOGrB3mr+oJJHFDUTDUXJGrbl65YyFdckjz2uw6HiuZT5mAQrjFAyyyiUNJoV
hsIobb6vHkTLKZgWbTZko1VDSbZTbfKIkHlP8gdKthqfElF6fe9b/h0QJWe1kNsWQn6K75WQfwMAAP//
eVpV+1MEAAA=`,
		},
		"github.com/jawher/mow.cli/internal/matcher": EncodedLicense{
			Name: "MIT",
			Text: `
H4sIAAAAAAAC/3xRvY7jNhjs+RSDre4AYvODVOm4Em0xkUiDos9xSUv0ioEsGiIVw28fUPbeYVNEjcDv
Z+abGfz/ZwaHRhjUvnNTdIQU4Xqf/fuQ8KX7il9//uU3ij/sbXAzmrDEaAnZufniY/Rhgo8Y3OxOd7zP
dkqupzjPziGc0Q12fncUKcBOd1zdHMOEcErWT356h0UXrncSzkiDj4jhnG52drBTDxtj6LxNrkcfuuXi
pmRT5jv70UV8SYPDS/vcePm6kvTOjsRPyL2PFm4+DWFJmF1Ms+8yBoWfunHp8w0f7dFf/JMhr6/6I0kB
S3R0vZPiEnp/zn+3yroup9HHgaL3Gfq0JEcRc3E1kmYdP4UZ0Y0j6cLVu4hV64/r1pl8+jUbmp4WxVy5
DeHyWYmP5LzMk4+DW3f6gBhWxr9dl3Ilj5/DOIZbltaFqfdZUfydkJyyPYV/3KrlEe8Uku8edq8BXH+k
+mzFwY4jTu5pmOvhJ5JLH3LmTB+TnZK3I65hXvn+K/OVEFNxtGpjDkxziBY7rb6Jkpd4YS1E+0JxEKZS
e4MD05pJc4TagMkj/hSypOB/7TRvWyhNRLOrBS8phCzqfSnkFm97A6kMatEIw0sYhUz4hBK8zWAN10XF
pGFvohbmSMlGGJkxN0qDYce0EcW+Zhq7vd6ploPJElJJITdayC1vuDSvEBJSgX/j0qCtWF1nKsL2plI6
34dC7Y5abCuDStUl1y3eOGrB3mr+oJJHFDUTDUXJGrbl65YyFdckjz2uw6HiuZT5mAQrjFAyyyiUNJoV
hsIobb6vHkTLKZgWbTZko1VDSbZTbfKIkHlP8gdKthqfElF6fe9b/h0QJWe1kNsWQn6K75WQfwMAAP//
eVpV+1MEAAA=`,
		},
		"github.com/jawher/mow.cli/internal/parser": EncodedLicense{
			Name: "MIT",
			Text: `
H4sIAAAAAAAC/3xRvY7jNhjs+RSDre4AYvODVOm4Em0xkUiDos9xSUv0ioEsGiIVw28fUPbeYVNEjcDv
Z+abGfz/ZwaHRhjUvnNTdIQU4Xqf/fuQ8KX7il9//uU3ij/sbXAzmrDEaAnZufniY/Rhgo8Y3OxOd7zP
dkqupzjPziGc0Q12fncUKcBOd1zdHMOEcErWT356h0UXrncSzkiDj4jhnG52drBTDxtj6LxNrkcfuuXi
pmRT5jv70UV8SYPDS/vcePm6kvTOjsRPyL2PFm4+DWFJmF1Ms+8yBoWfunHp8w0f7dFf/JMhr6/6I0kB
S3R0vZPiEnp/zn+3yroup9HHgaL3Gfq0JEcRc3E1kmYdP4UZ0Y0j6cLVu4hV64/r1pl8+jUbmp4WxVy5
DeHyWYmP5LzMk4+DW3f6gBhWxr9dl3Ilj5/DOIZbltaFqfdZUfydkJyyPYV/3KrlEe8Uku8edq8BXH+k
+mzFwY4jTu5pmOvhJ5JLH3LmTB+TnZK3I65hXvn+K/OVEFNxtGpjDkxziBY7rb6Jkpd4YS1E+0JxEKZS
e4MD05pJc4TagMkj/hSypOB/7TRvWyhNRLOrBS8phCzqfSnkFm97A6kMatEIw0sYhUz4hBK8zWAN10XF
pGFvohbmSMlGGJkxN0qDYce0EcW+Zhq7vd6ploPJElJJITdayC1vuDSvEBJSgX/j0qCtWF1nKsL2plI6
34dC7Y5abCuDStUl1y3eOGrB3mr+oJJHFDUTDUXJGrbl65YyFdckjz2uw6HiuZT5mAQrjFAyyyiUNJoV
hsIobb6vHkTLKZgWbTZko1VDSbZTbfKIkHlP8gdKthqfElF6fe9b/h0QJWe1kNsWQn6K75WQfwMAAP//
eVpV+1MEAAA=`,
		},
		"github.com/jawher/mow.cli/internal/values": EncodedLicense{
			Name: "MIT",
			Text: `
H4sIAAAAAAAC/3xRvY7jNhjs+RSDre4AYvODVOm4Em0xkUiDos9xSUv0ioEsGiIVw28fUPbeYVNEjcDv
Z+abGfz/ZwaHRhjUvnNTdIQU4Xqf/fuQ8KX7il9//uU3ij/sbXAzmrDEaAnZufniY/Rhgo8Y3OxOd7zP
dkqupzjPziGc0Q12fncUKcBOd1zdHMOEcErWT356h0UXrncSzkiDj4jhnG52drBTDxtj6LxNrkcfuuXi
pmRT5jv70UV8SYPDS/vcePm6kvTOjsRPyL2PFm4+DWFJmF1Ms+8yBoWfunHp8w0f7dFf/JMhr6/6I0kB
S3R0vZPiEnp/zn+3yroup9HHgaL3Gfq0JEcRc3E1kmYdP4UZ0Y0j6cLVu4hV64/r1pl8+jUbmp4WxVy5
DeHyWYmP5LzMk4+DW3f6gBhWxr9dl3Ilj5/DOIZbltaFqfdZUfydkJyyPYV/3KrlEe8Uku8edq8BXH+k
+mzFwY4jTu5pmOvhJ5JLH3LmTB+TnZK3I65hXvn+K/OVEFNxtGpjDkxziBY7rb6Jkpd4YS1E+0JxEKZS
e4MD05pJc4TagMkj/hSypOB/7TRvWyhNRLOrBS8phCzqfSnkFm97A6kMatEIw0sYhUz4hBK8zWAN10XF
pGFvohbmSMlGGJkxN0qDYce0EcW+Zhq7vd6ploPJElJJITdayC1vuDSvEBJSgX/j0qCtWF1nKsL2plI6
34dC7Y5abCuDStUl1y3eOGrB3mr+oJJHFDUTDUXJGrbl65YyFdckjz2uw6HiuZT5mAQrjFAyyyiUNJoV
hsIobb6vHkTLKZgWbTZko1VDSbZTbfKIkHlP8gdKthqfElF6fe9b/h0QJWe1kNsWQn6K75WQfwMAAP//
eVpV+1MEAAA=`,
		},
		"github.com/kopoli/appkit": EncodedLicense{
			Name: "MIT",
			Text: `
H4sIAAAAAAAC/1xRzW7jNhC+8yk+5JQAQrrYY2+MRVtEJNKg6HV9pCU6YiuLhkg3yNsXIzu7zZ4Eceb7
HTt4NNKiDp2fksdjI+0TY6t4+ZjD25Dx2D3h+7fv3/DqxtHj1U3/uNkztvXzOaQU4oSQMPjZHz/wNrsp
+77AafYe8YRucPObL5Aj3PSBi59TnBCP2YUpTG9w6OLlg8UT8hASUjzldzd7uKmHSyl2wWXfo4/d9eyn
7DLpncLoEx7z4PHQ3hEPT4tI793IwgSafY7wHvIQrxmzT3kOHXEUCFM3Xnvy8DkewzncFQi+xE8sR1yT
LxafBc6xDyf6+iXW5XocQxoK9IGoj9fsCyR6XNosKMcfcUby48i6eAk+Ycn6y92yQ9YvVGi+V5To5X2I
569JQmKn6zyFNPgF00ekuCj+7btML7R+iuMY3ylaF6c+UKL0J2N0aneM//oly+26U8yhu9W9HODy66r3
URrcOOLo74X5HmGC+1+cmeRTdlMObsQlzove7zGfGbOVQKvXds+NgGyxNfqHLEWJB95Ctg8F9tJWemex
58ZwZQ/Qa3B1wKtUZQHx19aItoU2TDbbWoqygFSreldKtcHLzkJpi1o20ooSVoME71RStETWCLOquLL8
RdbSHgq2llYR51obcGy5sXK1q7nBdme2uhXgqoTSSqq1kWojGqHsM6SC0hA/hLJoK17XJMX4zlbakD+s
9PZg5KayqHRdCtPiRaCW/KUWNyl1wKrmsilQ8oZvxILSthKG0drNHfaVoCfS4wp8ZaVWFGOllTV8ZQtY
bexP6F62ogA3sqVC1kY3BaM69ZpWpCKcEjcWqhpfLqLN8r9rxU9ClILXUm1aAlPEz+Vn9l8AAAD//7MD
VDw4BAAA`,
		},
		"github.com/kopoli/thelm/lib": EncodedLicense{
			Name: "MIT",
			Text: `
H4sIAAAAAAAC/1xRzW7jNhC+8yk+5JQAQvpz6KE3xqItIhJpUPS6PtISHbGVRUOkG+Tti5Gd3WZPgjjz
/Y4dPBppUYfOT8njsZH2ibFVvHzM4W3IeOye8Puvv/2BVzeOHq9u+sfNnrGtn88hpRAnhITBz/74gbfZ
Tdn3BU6z94gndIOb33yBHOGmD1z8nOKEeMwuTGF6g0MXLx8snpCHkJDiKb+72cNNPVxKsQsu+x597K5n
P2WXSe8URp/wmAePh/aOeHhaRHrvRhYm0OxzhPeQh3jNmH3Kc+iIo0CYuvHak4fP8RjO4a5A8CV+Yjni
mnyx+Cxwjn040dcvsS7X4xjSUKAPRH28Zl8g0ePSZkE5fokzkh9H1sVL8AlL1h/ulh2yfqFC872iRC/v
Qzx/TRISO13nKaTBL5g+IsVF8W/fZXqh9VMcx/hO0bo49YESpT8Zo1O7Y/zXL1lu151iDt2t7uUAlx9X
vY/S4MYRR38vzPcIE9z/4swkn7KbcnAjLnFe9H6O+cyYrQRavbZ7bgRki63R32QpSjzwFrJ9KLCXttI7
iz03hit7gF6DqwNepSoLiL+2RrQttGGy2dZSlAWkWtW7UqoNXnYWSlvUspFWlLAaJHinkqIlskaYVcWV
5S+ylvZQsLW0ijjX2oBjy42Vq13NDbY7s9WtAFcllFZSrY1UG9EIZZ8hFZSG+CaURVvxuiYpxne20ob8
YaW3ByM3lUWl61KYFi8CteQvtbhJqQNWNZdNgZI3fCMWlLaVMIzWbu6wrwQ9kR5X4CsrtaIYK62s4Stb
wGpjv0P3shUFuJEtFbI2uikY1anXtCIV4ZS4sVDV+HIRbZb/XSu+E6IUvJZq0xKYIn4uP/8XAAD//6vn
VXQ3BAAA`,
		},
		"github.com/mattn/go-runewidth": EncodedLicense{
			Name: "MIT",
			Text: `
H4sIAAAAAAAC/1xRTW/jNhO+81c8yCkBhLxve+ihN8aiLaISaVD0uj7SEh2xkEVDpBrk3xcjO7vdngRx
5vkcO3g00qIOnZ+Sx3Mj7Qtjm3j7nMP7kPHcveDX///yG04uLUOYIxqX03KNOTK29/M1pBTihJAw+Nmf
P/E+uyn7vsBl9h7xgm5w87svkCPc9Imbn1OcEM/ZhSlM73Do4u2TxQvyEBJSvOQPN3u4qYdLKXbBZd+j
j91y9VN2mfQuYfQJz3nweGofiKeXVaT3bmRhAs2+RvgIeYhLxuxTnkNHHAXC1I1LTx6+xmO4hocCwdcO
EssRS/LF6rPANfbhQl+/xrot5zGkoUAfiPq8ZF8g0eNaaUE5/hdnJD+OrIu34BPWrD/crTtk/UaF5kdF
iV4+hnj9OUlI7LLMU0iDXzF9RIqr4l++y/RC65c4jvGDonVx6gMlSr8zRvd25/i3X7PcTzzFHLp73esB
bj+u+hilwY0jzv5RmO8RJrh/xZlJPmU35eBG3OK86v035itjthJo9dYeuRGQLfZGf5OlKPHEW8j2qcBR
2kofLI7cGK7sCXoLrk74Q6qygPhzb0TbQhsmm30tRVlAqk19KKXa4e1gobRFLRtpRQmrQYIPKilaImuE
2VRcWf4ma2lPBdtKq4hzqw049txYuTnU3GB/MHvdCnBVQmkl1dZItRONUPYVUkFpiG9CWbQVr2uSYvxg
K23IHzZ6fzJyV1lUui6FafEmUEv+Vou7lDphU3PZFCh5w3diRWlbCcNo7e4Ox0rQE+lxBb6xUiuKsdHK
Gr6xBaw29jv0KFtRgBvZUiFbo5uCUZ16SytSEU6JOwtVjZ8uos36f2jFd0KUgtdS7VoCU8Sv5Vf2TwAA
AP//SpF7+z0EAAA=`,
		},
		"github.com/nsf/termbox-go": EncodedLicense{
			Name: "MIT",
			Text: `
H4sIAAAAAAAC/1xRT4/jJhS/8yl+mtOu5E7bPfbG2CRGdSDCZNMciU1iKgciIJ3m21c4mV1NT5Yf7/f3
1eF6j+48ZXypv+Lbb79/Q7bxcgz//nIOMLc8hZgI2dp4cSm54OESJhvt8Y5zND7bscIpWotwwjCZeLYV
coDxd1xtTMEjHLNx3vkzDIZwvZNwQp5cQgqn/G6ihfEjTEphcCbbEWMYbhfrs8lF7+Rmm/AlTxYv/RPx
8nURGa2ZifMobx9PeHd5CreMaFOObigcFZwf5ttYPHw8z+7ingoFvnSQSA64JVstPitcwuhO5WuXWNfb
cXZpqjC6Qn28ZVshleFgfUEZP/4aIpKdZzKEq7MJS9af7padYv1aCs3PilKZvE/h8jmJS+R0i96lyS6Y
MSCFRfFvO+QyKeunMM/hvUQbgh9dSZT+IERPFuYY/rFLlseJfchueNS9HOD686rPpzSZecbRPguzI5wn
ZfQRJxb5lI3Pzsy4hrjo/T/mKyG6ZejlSu+pYuA9tkp+5w1r8EJ78P6lwp7rVu409lQpKvQBcgUqDviT
i6YC+2urWN9DKsI3246zpgIXdbdruFjjbachpEbHN1yzBlqiCD6pOOsL2YapuqVC0zfecX2oyIprUThX
UoFiS5Xm9a6jCtud2sqegYoGQgouVoqLNdswoV/BBYQE+86ERt/SritShO50K1Xxh1puD4qvW41Wdg1T
Pd4YOk7fOvaQEgfUHeWbCg3d0DVbUFK3TJGy9nCHfcvKqOhRAVprLkWJUUuhFa11BS2V/gHd855VoIr3
pZCVkpuKlDrlqqxwUXCCPVhK1fh0EamW/13PfhCiYbTjYt2Di0/neyX/BQAA//9NO/UaJgQAAA==`,
		},
	}

	decode := func(input string) (string, error) {
		data := &bytes.Buffer{}
		br := base64.NewDecoder(base64.StdEncoding, strings.NewReader(input))

		r, err := gzip.NewReader(br)
		if err != nil {
			return "", err
		}

		_, err = io.Copy(data, r)
		if err != nil {
			return "", err
		}

		// Make sure the gzip is decoded successfully
		err = r.Close()
		if err != nil {
			return "", err
		}
		return data.String(), nil
	}

	ret := make(map[string]License)

	for k := range data {
		text, err := decode(data[k].Text)
		if err != nil {
			return nil, err
		}
		ret[k] = License{
			Name: data[k].Name,
			Text: text,
		}
	}

	return ret, nil
}
